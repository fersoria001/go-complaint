package repositories

import (
	"context"
	"fmt"
	"go-complaint/dto"
	"go-complaint/infrastructure/persistence/datasource"
	"go-complaint/infrastructure/persistence/models"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/google/uuid"
)

type EventRepository struct {
	schema *datasource.Schema
}

func NewEventRepository(schema *datasource.Schema) *EventRepository {
	return &EventRepository{
		schema: schema,
	}
}

func (r *EventRepository) Save(ctx context.Context, e dto.StoredEvent) error {
	var conn, err = r.schema.Pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()
	model, err := models.NewEvent(e)
	if err != nil {
		return err
	}
	var cols = strings.Join(model.Columns(), ", ")
	var insertEvent = fmt.Sprintf(`INSERT INTO %s (%s) VALUES (%s)`, model.Table(), cols, model.Args())
	_, err = conn.Exec(ctx, insertEvent, model.Values()...)
	if err != nil {
		return err
	}
	return nil
}

func (r *EventRepository) SaveToLog(ctx context.Context, e dto.StoredEvent) error {
	var conn, err = r.schema.Pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()
	model, err := models.NewEvent(e)
	if err != nil {
		return err
	}
	var cols = strings.Join(model.Columns(), ", ")
	var insertEvent = fmt.Sprintf(`INSERT INTO %s_log (%s) VALUES (%s)`, model.Table(), cols, model.Args())
	_, err = conn.Exec(ctx, insertEvent, model.Values()...)
	if err != nil {
		return err
	}
	return nil
}

func (r *EventRepository) SaveAll(ctx context.Context, events mapset.Set[*dto.StoredEvent]) error {
	var conn, err = r.schema.Pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()
	emptyModel := &models.Event{}
	var insertEvent = fmt.Sprintf(
		`INSERT INTO %s (%s) VALUES (%s)`,
		emptyModel.Table(),
		models.StringColumns(emptyModel.Columns()),
		emptyModel.Args(),
	)
	tx, err := conn.Begin(ctx)
	if err != nil {
		return err
	}
	for event := range events.Iter() {
		model, err := models.NewEvent(*event)
		if err != nil {
			return err
		}
		_, err = tx.Exec(ctx, insertEvent, model.Values()...)
		if err != nil {
			tx.Rollback(ctx)
			return err
		}
	}
	err = tx.Commit(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (r *EventRepository) GetAll(ctx context.Context) (mapset.Set[*dto.StoredEvent], error) {
	var conn, err = r.schema.Pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	model := &models.Event{}
	var selectEvents = fmt.Sprintf(`SELECT %s FROM %s`, models.StringColumns(model.Columns()), model.Table())
	rows, err := conn.Query(ctx, selectEvents)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var events = mapset.NewSet[*dto.StoredEvent]()
	for rows.Next() {
		err = rows.Scan(model.Values()...)
		if err != nil {
			return nil, err
		}
		event := &dto.StoredEvent{
			EventId:    model.EventId,
			EventBody:  model.EventBody,
			OccurredOn: model.OccurredOn,
			TypeName:   model.TypeName,
		}
		events.Add(event)
	}
	return events, nil
}

func (r *EventRepository) GetAllFromLog(ctx context.Context) (mapset.Set[*dto.StoredEvent], error) {
	var conn, err = r.schema.Pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	model := &models.Event{}
	var selectEvents = fmt.Sprintf(`SELECT %s FROM %s_log`, models.StringColumns(model.Columns()), model.Table())
	rows, err := conn.Query(ctx, selectEvents)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var events = mapset.NewSet[*dto.StoredEvent]()
	for rows.Next() {
		err = rows.Scan(model.Values()...)
		if err != nil {
			return nil, err
		}
		event := &dto.StoredEvent{
			EventId:    model.EventId,
			EventBody:  model.EventBody,
			OccurredOn: model.OccurredOn,
			TypeName:   model.TypeName,
		}
		events.Add(event)
	}
	return events, nil
}

func (r *EventRepository) Get(ctx context.Context, eventID string) (*dto.StoredEvent, error) {
	var conn, err = r.schema.Pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	parsedID, err := uuid.Parse(eventID)
	if err != nil {
		return nil, err
	}
	var model = &models.Event{}
	var cols = strings.Join(model.Columns(), ", ")
	var select_event = fmt.Sprintf(`SELECT %s FROM %s WHERE event_id = $1`, cols, model.Table())
	row := conn.QueryRow(ctx, select_event, parsedID)
	err = row.Scan(model.Values()...)

	if err != nil {

		return nil, err
	}
	return &dto.StoredEvent{
		EventId:    model.EventId,
		EventBody:  model.EventBody,
		OccurredOn: model.OccurredOn,
		TypeName:   model.TypeName,
	}, nil
}

func (r *EventRepository) Remove(ctx context.Context, eventID string) error {
	var conn, err = r.schema.Pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()
	parsedID, err := uuid.Parse(eventID)
	if err != nil {
		return err
	}
	var model = &models.Event{}
	var deleteEvent = fmt.Sprintf(`DELETE FROM %s WHERE event_id = $1`, model.Table())
	_, err = conn.Exec(ctx, deleteEvent, parsedID)
	if err != nil {
		return err
	}
	return nil
}

func (r *EventRepository) RemoveAll(ctx context.Context, eventIDS []string) error {
	var conn, err = r.schema.Pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()
	var model = &models.Event{}
	var deleteEvent = fmt.Sprintf(`DELETE FROM %s WHERE event_id = $1`, model.Table())
	for _, eventID := range eventIDS {
		parsedID, err := uuid.Parse(eventID)
		if err != nil {
			return err
		}
		_, err = conn.Exec(ctx, deleteEvent, parsedID)
		if err != nil {
			return err
		}
	}
	return nil
}
func (r *EventRepository) FindByTypeName(
	ctx context.Context,
	typeName string,
) (mapset.Set[*dto.StoredEvent], error) {
	var conn, err = r.schema.Pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	model := &models.Event{}
	var selectEvents = fmt.Sprintf(`
	SELECT %s FROM %s WHERE type_name=$1`,
		models.StringColumns(model.Columns()), model.Table())
	rows, err := conn.Query(ctx, selectEvents, typeName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var events = mapset.NewSet[*dto.StoredEvent]()
	for rows.Next() {
		err = rows.Scan(model.Values()...)
		if err != nil {
			return nil, err
		}
		event := &dto.StoredEvent{
			EventId:    model.EventId,
			EventBody:  model.EventBody,
			OccurredOn: model.OccurredOn,
			TypeName:   model.TypeName,
		}
		events.Add(event)
	}
	return events, nil
}
