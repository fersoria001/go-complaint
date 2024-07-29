package repositories

import (
	"context"
	"go-complaint/dto"
	"go-complaint/infrastructure/persistence/datasource"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type EventRepository struct {
	schema datasource.Schema
}

func NewEventRepository(schema datasource.Schema) EventRepository {
	return EventRepository{
		schema: schema,
	}
}

func (r EventRepository) Save(ctx context.Context, e *dto.StoredEvent) error {
	var conn, err = r.schema.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()
	insertCommand := string(`
	INSERT INTO events (
		event_id,
		event_body,
		occurred_on,
		type_name
		)
	VALUES ($1, $2, $3, $4)
	`)
	_, err = conn.Exec(
		ctx,
		insertCommand,
		&e.EventId,
		&e.EventBody,
		&e.OccurredOn,
		&e.TypeName,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r EventRepository) Get(ctx context.Context, id uuid.UUID) (*dto.StoredEvent, error) {
	var conn, err = r.schema.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	selectCommand := string(`
	SELECT
		event_id,
		event_body,
		occurred_on,
		type_name
	FROM events
	WHERE event_id = $1
	`)
	row := conn.QueryRow(ctx, selectCommand, id)
	return r.load(ctx, row)
}

func (r EventRepository) FindAll(ctx context.Context, source StatementSource) ([]*dto.StoredEvent, error) {
	var conn, err = r.schema.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	rows, err := conn.Query(ctx, source.Query(), source.Args()...)
	if err != nil {
		return nil, err
	}
	return r.loadAll(ctx, rows)
}

func (r EventRepository) load(
	_ context.Context,
	row pgx.Row,
) (*dto.StoredEvent, error) {
	var e = dto.StoredEvent{}
	err := row.Scan(
		&e.EventId,
		&e.EventBody,
		&e.OccurredOn,
		&e.TypeName,
	)
	if err != nil {
		return nil, err
	}
	return &e, nil
}
func (r EventRepository) loadAll(
	_ context.Context,
	rows pgx.Rows,
) ([]*dto.StoredEvent, error) {
	var events = make([]*dto.StoredEvent, 0)
	for rows.Next() {
		e, err := r.load(context.Background(), rows)
		if err != nil {
			return nil, err
		}
		events = append(events, e)
	}
	return events, nil
}
