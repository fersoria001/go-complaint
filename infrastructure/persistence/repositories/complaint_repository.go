package repositories

import (
	"context"
	"errors"
	"go-complaint/domain/model/common"
	"go-complaint/domain/model/complaint"
	"go-complaint/infrastructure/persistence/datasource"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type ComplaintRepository struct {
	schema datasource.Schema
}

func NewComplaintRepository(schema datasource.Schema) ComplaintRepository {
	return ComplaintRepository{
		schema: schema,
	}
}

func (pr ComplaintRepository) Get(
	ctx context.Context,
	complaintID uuid.UUID,
) (*complaint.Complaint, error) {
	conn, err := pr.schema.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	selectQuery := string(`
	SELECT
	"complaint".id,
	"complaint".author_id,
	"complaint".receiver_id,
	"complaint".complaint_status,
	"complaint".title,
	"complaint".complaint_description,
	"complaint".body,
	"complaint".rating_rate,
	"complaint".rating_comment,
	"complaint".created_at,
	"complaint".updated_at
	FROM 
	public."complaint"
	WHERE "complaint".id = $1
	`)
	row := conn.QueryRow(ctx, selectQuery, complaintID)
	complaint, err := pr.load(ctx, row)
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	return complaint, nil
}

func (pr ComplaintRepository) Count(
	ctx context.Context,
	source StatementSource,
) (int, error) {
	conn, err := pr.schema.Acquire(ctx)
	if err != nil {
		return 0, err
	}
	row := conn.QueryRow(ctx, source.Query(), source.Args()...)
	var count int
	err = row.Scan(&count)
	if err != nil {
		return 0, err
	}
	defer conn.Release()
	return count, nil
}

func (pr ComplaintRepository) FindAll(
	ctx context.Context,
	source StatementSource,
) (mapset.Set[complaint.Complaint], error) {
	conn, err := pr.schema.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	rows, err := conn.Query(ctx, source.Query(), source.Args()...)
	if err != nil {
		return nil, err
	}
	complaints, err := pr.loadAll(ctx, rows)
	if err != nil {
		return nil, err
	}
	defer func() {
		rows.Close()
		conn.Release()
	}()
	return complaints, nil
}

func (pr ComplaintRepository) loadAll(
	ctx context.Context,
	rows pgx.Rows,
) (mapset.Set[complaint.Complaint], error) {
	complaints := mapset.NewSet[complaint.Complaint]()
	for rows.Next() {
		complaint, err := pr.load(ctx, rows)
		if err != nil {
			return nil, err
		}
		complaints.Add(*complaint)
	}
	return complaints, nil
}

func (pr ComplaintRepository) load(
	ctx context.Context,
	row pgx.Row,
) (*complaint.Complaint, error) {
	var (
		id            uuid.UUID
		authorID      string
		receiverID    string
		status        string
		title         string
		description   string
		body          string
		ratingRate    int
		ratingComment string
		createdAt     string
		updatedAt     string
	)
	err := row.Scan(
		&id,
		&authorID,
		&receiverID,
		&status,
		&title,
		&description,
		&body,
		&ratingRate,
		&ratingComment,
		&createdAt,
		&updatedAt,
	)

	mapper := MapperRegistryInstance().Get("Reply")
	if mapper == nil {
		return nil, ErrMapperNotRegistered
	}
	complaintRepliesRepository, ok := mapper.(ComplaintRepliesRepository)
	if !ok {
		return nil, ErrWrongTypeAssertion
	}
	if err != nil {
		return nil, err
	}
	parsedStatus, err := complaint.ParseStatus(status)
	if err != nil {
		return nil, err
	}

	commonCreatedAt, err := common.NewDateFromString(createdAt)
	if err != nil {
		return nil, err
	}
	commonUpdatedAt, err := common.NewDateFromString(updatedAt)
	if err != nil {
		return nil, err
	}
	message, err := complaint.NewMessage(
		title,
		description,
		body,
	)
	if err != nil {
		return nil, err
	}
	rating, err := complaint.NewRating(
		ratingRate,
		ratingComment,
	)
	if err != nil {
		return nil, err
	}

	replies, err := complaintRepliesRepository.FindAllByComplaintID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			replies = mapset.NewSet[*complaint.Reply]()
		} else {
			return nil, err
		}
	}
	return complaint.NewComplaint(
		id,
		authorID,
		receiverID,
		parsedStatus,
		message,
		commonCreatedAt,
		commonUpdatedAt,
		rating,
		replies,
	)
}

func (pr ComplaintRepository) Save(
	ctx context.Context,
	complaint *complaint.Complaint,
) error {
	conn, err := pr.schema.Acquire(ctx)
	if err != nil {
		return err
	}
	insertCommand := string(`
		INSERT INTO complaint
			(
				id,
				author_id,
				receiver_id,
				complaint_status,
				title,
				complaint_description,
				body,
				rating_rate,
				rating_comment,
				created_at,
				updated_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`)
	_, err = conn.Exec(
		ctx,
		insertCommand,
		complaint.ID(),
		complaint.AuthorID(),
		complaint.ReceiverID(),
		complaint.Status().String(),
		complaint.Message().Title(),
		complaint.Message().Description(),
		complaint.Message().Body(),
		complaint.Rating().Rate(),
		complaint.Rating().Comment(),
		complaint.CreatedAt().StringRepresentation(),
		complaint.UpdatedAt().StringRepresentation(),
	)
	if err != nil {
		return err
	}
	mapper := MapperRegistryInstance().Get("Reply")
	if mapper == nil {
		return ErrMapperNotRegistered
	}
	complaintRepliesRepository, ok := mapper.(ComplaintRepliesRepository)
	if !ok {
		return ErrWrongTypeAssertion
	}
	err = complaintRepliesRepository.SaveAll(ctx, complaint.Replies())
	if err != nil {
		return err
	}
	defer conn.Release()
	return nil
}

func (pr ComplaintRepository) Update(
	ctx context.Context,
	updatedComplaint *complaint.Complaint,
) error {
	conn, err := pr.schema.Acquire(ctx)
	if err != nil {
		return err
	}
	mapper := MapperRegistryInstance().Get("Reply")
	if mapper == nil {
		return ErrMapperNotRegistered
	}
	complaintRepliesRepository, ok := mapper.(ComplaintRepliesRepository)
	if !ok {
		return ErrWrongTypeAssertion
	}
	err = complaintRepliesRepository.DeleteAll(ctx, updatedComplaint.Replies())
	if err != nil {
		return err
	}
	err = complaintRepliesRepository.SaveAll(ctx, updatedComplaint.Replies())
	if err != nil {
		return err
	}
	var (
		id            uuid.UUID = updatedComplaint.ID()
		authorID      string    = updatedComplaint.AuthorID()
		receiverID    string    = updatedComplaint.ReceiverID()
		status        string    = updatedComplaint.Status().String()
		title         string    = updatedComplaint.Message().Title()
		description   string    = updatedComplaint.Message().Description()
		body          string    = updatedComplaint.Message().Body()
		ratingRate    int       = updatedComplaint.Rating().Rate()
		ratingComment string    = updatedComplaint.Rating().Comment()
		createdAt     string    = updatedComplaint.CreatedAt().StringRepresentation()
		updatedAt     string    = updatedComplaint.UpdatedAt().StringRepresentation()
	)
	updateCommand := string(`
	UPDATE public."complaint"
		SET
		id=$1,
		author_id=$2,
		receiver_id=$3,
		complaint_status=$4,
		title=$5,
		complaint_description=$6,
		body=$7,
		rating_rate=$8,
		rating_comment=$9,
		created_at=$10,
		updated_at=$11
	WHERE id = $1;`,
	)
	_, err = conn.Exec(
		ctx,
		updateCommand,
		&id,
		&authorID,
		&receiverID,
		&status,
		&title,
		&description,
		&body,
		&ratingRate,
		&ratingComment,
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		return err
	}
	defer conn.Release()
	return nil
}
