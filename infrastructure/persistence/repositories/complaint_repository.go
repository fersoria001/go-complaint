package repositories

import (
	"context"
	"fmt"
	"go-complaint/domain/model/common"
	"go-complaint/domain/model/complaint"
	"go-complaint/erros"
	"go-complaint/infrastructure/persistence/datasource"
	"go-complaint/infrastructure/persistence/models"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ComplaintRepository struct {
	schema *datasource.Schema
}

func NewComplaintRepository(schema *datasource.Schema) *ComplaintRepository {
	return &ComplaintRepository{
		schema: schema,
	}
}

func (cr *ComplaintRepository) Save(ctx context.Context, complaint *complaint.Complaint) error {
	var (
		conn            *pgxpool.Conn
		err             error
		model           models.Complaint = models.NewComplaint(complaint)
		insertComplaint string           = fmt.Sprintf(`INSERT INTO %s (%s) VALUES (%s)`,
			model.Table(), models.StringColumns(model.Columns()), models.Args(model.Columns()))
	)
	conn, err = cr.schema.Pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()
	_, err = conn.Exec(ctx, insertComplaint, model.Values()...)
	if err != nil {

		return err
	}
	return nil
}

func (cr *ComplaintRepository) Get(ctx context.Context, id string) (*complaint.Complaint, error) {
	var (
		conn            *pgxpool.Conn
		err             error
		row             pgx.Row
		status          complaint.Status
		message         complaint.Message
		createdAt       common.Date
		updatedAt       common.Date
		rating          *complaint.Rating
		newComplaint    *complaint.Complaint
		parsedID        uuid.UUID
		model           = &models.Complaint{}
		selectComplaint = fmt.Sprintf(`SELECT %s FROM %s WHERE id = $1`,
			models.StringColumns(model.Columns()), model.Table())
	)
	parsedID, err = uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	conn, err = cr.schema.Pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	row = conn.QueryRow(ctx, selectComplaint, parsedID)
	err = row.Scan(model.Values()...)
	if err != nil {
		return nil, err
	}

	status, err = complaint.ParseStatus(model.Status)
	if err != nil {
		return nil, err
	}
	message, err = complaint.NewMessage(model.Title, model.Description, model.Content)
	if err != nil {
		return nil, err
	}
	createdAt, err = common.NewDateFromString(model.CreatedAt)
	if err != nil {
		return nil, err
	}
	updatedAt, err = common.NewDateFromString(model.UpdatedAt)
	if err != nil {
		return nil, err
	}
	rating, err = complaint.NewRating(model.Rate, model.Comment)
	if err != nil {
		if rating == nil {
			rating = &complaint.Rating{}
		}
	}
	newComplaint, err = complaint.NewComplaint(
		model.ID,
		model.AuthorID,
		model.ReceiverID,
		status,
		message,
		createdAt,
		updatedAt,
		*rating)
	if err != nil {
		return nil, err
	}
	return newComplaint, nil
}

func (cr *ComplaintRepository) GetAll(ctx context.Context) (mapset.Set[*complaint.Complaint], error) {
	var (
		conn            *pgxpool.Conn
		err             error
		model           *models.Complaint = &models.Complaint{}
		status          complaint.Status
		message         complaint.Message
		createdAt       common.Date
		updatedAt       common.Date
		rating          *complaint.Rating
		newComplaint    *complaint.Complaint
		selectComplaint string = fmt.Sprintf(`SELECT %s FROM %s`,
			models.StringColumns(model.Columns()), model.Table())
	)
	conn, err = cr.schema.Pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	rows, err := conn.Query(ctx, selectComplaint)
	if err != nil {
		return nil, err
	}
	complaints := mapset.NewSet[*complaint.Complaint]()
	for rows.Next() {
		err = rows.Scan(model.Values()...)
		if err != nil {
			return nil, err
		}
		status, err = complaint.ParseStatus(model.Status)
		if err != nil {
			return nil, err
		}
		message, err = complaint.NewMessage(model.Title, model.Description, model.Content)
		if err != nil {
			return nil, err
		}
		createdAt, err = common.NewDateFromString(model.CreatedAt)
		if err != nil {
			return nil, err
		}
		updatedAt, err = common.NewDateFromString(model.UpdatedAt)
		if err != nil {
			return nil, err
		}
		rating, err = complaint.NewRating(model.Rate, model.Comment)
		if err != nil {
			return nil, err
		}
		newComplaint, err = complaint.NewComplaint(
			model.ID,
			model.AuthorID,
			model.ReceiverID,
			status,
			message,
			createdAt,
			updatedAt,
			*rating)
		if err != nil {
			return nil, err
		}
		complaints.Add(newComplaint)
	}
	return complaints, nil

}

func (cr *ComplaintRepository) Update(ctx context.Context, complaint *complaint.Complaint) error {
	var (
		conn            *pgxpool.Conn
		err             error
		model           models.Complaint = models.NewComplaint(complaint)
		updateComplaint string           = fmt.Sprintf(`UPDATE %s SET %s WHERE id = $1`,
			model.Table(), models.StringKeyArgs(model.Columns()))
	)
	conn, err = cr.schema.Pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()
	_, err = conn.Exec(ctx, updateComplaint, model.Values()...)
	if err != nil {
		return err
	}
	return nil
}

func (cr *ComplaintRepository) Remove(ctx context.Context, id string) error {

	var (
		conn            *pgxpool.Conn
		err             error
		model           *models.Complaint = &models.Complaint{}
		deleteComplaint string            = fmt.Sprintf(`DELETE FROM %s WHERE id = $1`,
			model.Table())
		parsedID uuid.UUID
	)
	parsedID, err = uuid.Parse(id)
	if err != nil {
		return err
	}
	conn, err = cr.schema.Pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()
	_, err = conn.Exec(ctx, deleteComplaint, parsedID)
	if err != nil {
		return err
	}

	return nil
}

//Additional behaviour

func (cr *ComplaintRepository) FindByReceiver(ctx context.Context, id string, limit, offset int) (mapset.Set[*complaint.Complaint], int, error) {
	var (
		conn            *pgxpool.Conn
		err             error
		model           *models.Complaint = &models.Complaint{}
		status          complaint.Status
		message         complaint.Message
		createdAt       common.Date
		updatedAt       common.Date
		rating          *complaint.Rating
		newComplaint    *complaint.Complaint
		count           int
		countQuery      string = fmt.Sprintf(`SELECT COUNT(*) FROM %s where receiver_id = $1`, model.Table())
		selectComplaint string = fmt.Sprintf(`SELECT %s FROM %s where receiver_id = $1 LIMIT %d OFFSET %d`,
			models.StringColumns(model.Columns()), model.Table(), limit, offset)
	)
	conn, err = cr.schema.Pool.Acquire(ctx)
	if err != nil {
		return nil, 0, err
	}
	defer conn.Release()
	complaints := mapset.NewSet[*complaint.Complaint]()
	err = conn.QueryRow(ctx, countQuery, id).Scan(&count)
	if err != nil {
		return nil, 0, err
	}
	if count == 0 {
		return complaints, 0, &erros.ValueNotFoundError{}
	}
	rows, err := conn.Query(ctx, selectComplaint, id)
	if err != nil {
		return nil, 0, err
	}
	for rows.Next() {
		err = rows.Scan(model.Values()...)
		if err != nil {
			return nil, 0, err
		}
		status, err = complaint.ParseStatus(model.Status)
		if err != nil {
			return nil, 0, err
		}
		message, err = complaint.NewMessage(model.Title, model.Description, model.Content)
		if err != nil {
			return nil, 0, err
		}
		createdAt, err = common.NewDateFromString(model.CreatedAt)
		if err != nil {
			return nil, 0, err
		}
		updatedAt, err = common.NewDateFromString(model.UpdatedAt)
		if err != nil {
			return nil, 0, err
		}
		rating, err = complaint.NewRating(model.Rate, model.Comment)
		if err != nil {
			return nil, 0, err
		}
		newComplaint, err = complaint.NewComplaint(
			model.ID,
			model.AuthorID,
			model.ReceiverID,
			status,
			message,
			createdAt,
			updatedAt,
			*rating)
		if err != nil {
			return nil, 0, err
		}
		complaints.Add(newComplaint)
	}
	return complaints, count, nil
}

func (cr *ComplaintRepository) FindByAuthor(ctx context.Context, id string, limit, offset int) (mapset.Set[*complaint.Complaint], int, error) {
	var (
		conn            *pgxpool.Conn
		err             error
		model           *models.Complaint = &models.Complaint{}
		status          complaint.Status
		message         complaint.Message
		createdAt       common.Date
		updatedAt       common.Date
		rating          *complaint.Rating
		newComplaint    *complaint.Complaint
		count           int
		countQuery      string = fmt.Sprintf(`SELECT COUNT(*) FROM %s where author_id = $1`, model.Table())
		selectComplaint string = fmt.Sprintf(`SELECT %s FROM %s where author_id = $1 LIMIT %d OFFSET %d`,
			models.StringColumns(model.Columns()), model.Table(), limit, offset)
	)
	conn, err = cr.schema.Pool.Acquire(ctx)
	if err != nil {
		return nil, 0, err
	}
	defer conn.Release()
	complaints := mapset.NewSet[*complaint.Complaint]()
	err = conn.QueryRow(ctx, countQuery, id).Scan(&count)
	if err != nil {
		return nil, 0, err
	}
	if count == 0 {
		return complaints, 0, &erros.ValueNotFoundError{}
	}
	rows, err := conn.Query(ctx, selectComplaint, id)
	if err != nil {
		return nil, 0, err
	}
	for rows.Next() {
		err = rows.Scan(model.Values()...)
		if err != nil {
			return nil, 0, err
		}
		status, err = complaint.ParseStatus(model.Status)
		if err != nil {
			return nil, 0, err
		}
		message, err = complaint.NewMessage(model.Title, model.Description, model.Content)
		if err != nil {
			return nil, 0, err
		}
		createdAt, err = common.NewDateFromString(model.CreatedAt)
		if err != nil {
			return nil, 0, err
		}
		updatedAt, err = common.NewDateFromString(model.UpdatedAt)
		if err != nil {
			return nil, 0, err
		}
		rating, err = complaint.NewRating(model.Rate, model.Comment)
		if err != nil {
			return nil, 0, err
		}
		newComplaint, err = complaint.NewComplaint(
			model.ID,
			model.AuthorID,
			model.ReceiverID,
			status,
			message,
			createdAt,
			updatedAt,
			*rating)
		if err != nil {
			return nil, 0, err
		}
		complaints.Add(newComplaint)
	}
	return complaints, count, nil
}
