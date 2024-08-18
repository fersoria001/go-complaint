package repositories

import (
	"context"
	"errors"
	"go-complaint/domain/model/common"
	"go-complaint/domain/model/complaint"
	"go-complaint/infrastructure/persistence/datasource"
	"go-complaint/infrastructure/persistence/finders/find_all_complaint_replies"

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

func (pr ComplaintRepository) Find(
	ctx context.Context,
	src StatementSource,
) (*complaint.Complaint, error) {
	conn, err := pr.schema.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	row := conn.QueryRow(ctx, src.Query(), src.Args()...)
	complaint, err := pr.load(ctx, row)
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	return complaint, nil
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
	id,
	author_id,
	receiver_id,
	status,
	title,
	description,
	created_at,
	updated_at
	FROM complaint WHERE id = $1
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
) ([]*complaint.Complaint, error) {
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
) ([]*complaint.Complaint, error) {
	complaints := make([]*complaint.Complaint, 0)
	for rows.Next() {
		complaint, err := pr.load(ctx, rows)
		if err != nil {
			return nil, err
		}
		complaints = append(complaints, complaint)
	}
	return complaints, nil
}

func (pr ComplaintRepository) load(
	ctx context.Context,
	row pgx.Row,
) (*complaint.Complaint, error) {
	var (
		id          uuid.UUID
		authorId    uuid.UUID
		receiverId  uuid.UUID
		status      string
		title       string
		description string
		createdAt   string
		updatedAt   string
	)
	err := row.Scan(
		&id,
		&authorId,
		&receiverId,
		&status,
		&title,
		&description,
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		return nil, err
	}
	reg := MapperRegistryInstance()
	complaintRepliesRepository, ok := reg.Get("Reply").(ComplaintRepliesRepository)
	if !ok {
		return nil, ErrWrongTypeAssertion
	}
	recipientRepository, ok := reg.Get("Recipient").(RecipientRepository)
	if !ok {
		return nil, ErrWrongTypeAssertion
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
	replies, err := complaintRepliesRepository.FindAll(
		ctx,
		find_all_complaint_replies.ByComplaintID(id),
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			replies = mapset.NewSet[*complaint.Reply]()
		} else {
			return nil, err
		}
	}
	author, err := recipientRepository.Get(ctx, authorId)
	if err != nil {
		return nil, err
	}
	receiver, err := recipientRepository.Get(ctx, receiverId)
	if err != nil {
		return nil, err
	}
	ratingRepository, ok := reg.Get("Rating").(RatingRepository)
	if !ok {
		return nil, ErrWrongTypeAssertion
	}
	dbR, err := ratingRepository.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return complaint.NewComplaint(
		id,
		*author,
		*receiver,
		parsedStatus,
		title,
		description,
		commonCreatedAt,
		commonUpdatedAt,
		dbR,
		replies,
	)
}

func (pr ComplaintRepository) Remove(ctx context.Context, id uuid.UUID) error {
	conn, err := pr.schema.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()
	deleteCommand := string(`DELETE FROM complaint WHERE ID = $1`)
	_, err = conn.Exec(
		ctx,
		deleteCommand,
		&id,
	)
	if err != nil {
		return err
	}
	reg := MapperRegistryInstance()
	complaintRepliesRepository, ok := reg.Get("Reply").(ComplaintRepliesRepository)
	if !ok {
		return ErrWrongTypeAssertion
	}
	err = complaintRepliesRepository.DeleteAll(ctx, id)
	if err != nil {
		return err
	}
	ratingRepository, ok := reg.Get("Rating").(RatingRepository)
	if !ok {
		return ErrWrongTypeAssertion
	}
	err = ratingRepository.Remove(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (pr ComplaintRepository) Save(
	ctx context.Context,
	complaint *complaint.Complaint,
) error {
	conn, err := pr.schema.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()
	insertCommand := string(`
		INSERT INTO complaint
			(
				id,
				author_id,
				receiver_id,
				status,
				title,
				description,
				created_at,
				updated_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`)
	_, err = conn.Exec(
		ctx,
		insertCommand,
		complaint.Id(),
		complaint.Author().Id(),
		complaint.Receiver().Id(),
		complaint.Status().String(),
		complaint.Title(),
		complaint.Description(),
		complaint.CreatedAt().StringRepresentation(),
		complaint.UpdatedAt().StringRepresentation(),
	)
	if err != nil {
		return err
	}
	reg := MapperRegistryInstance()
	ratingRepository, ok := reg.Get("Rating").(RatingRepository)
	if !ok {
		return ErrWrongTypeAssertion
	}
	err = ratingRepository.Save(ctx, complaint.Rating())
	if err != nil {
		return err
	}
	complaintRepliesRepository, ok := reg.Get("Reply").(ComplaintRepliesRepository)
	if !ok {
		return ErrWrongTypeAssertion
	}
	err = complaintRepliesRepository.SaveAll(ctx, complaint.Replies())
	if err != nil {
		return err
	}
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
	reg := MapperRegistryInstance()
	complaintRepliesRepository, ok := reg.Get("Reply").(ComplaintRepliesRepository)
	if !ok {
		return ErrWrongTypeAssertion
	}
	err = complaintRepliesRepository.DeleteAll(ctx, updatedComplaint.Id())
	if err != nil {
		return err
	}
	err = complaintRepliesRepository.SaveAll(ctx, updatedComplaint.Replies())
	if err != nil {
		return err
	}
	ratingRepository, ok := reg.Get("Rating").(RatingRepository)
	if !ok {
		return ErrWrongTypeAssertion
	}
	err = ratingRepository.Update(ctx, updatedComplaint.Rating())
	if err != nil {
		return err
	}
	var (
		id          uuid.UUID = updatedComplaint.Id()
		authorId    uuid.UUID = updatedComplaint.Author().Id()
		receiverId  uuid.UUID = updatedComplaint.Receiver().Id()
		status      string    = updatedComplaint.Status().String()
		title       string    = updatedComplaint.Title()
		description string    = updatedComplaint.Description()
		createdAt   string    = updatedComplaint.CreatedAt().StringRepresentation()
		updatedAt   string    = updatedComplaint.UpdatedAt().StringRepresentation()
	)
	updateCommand := string(`
	UPDATE complaint
		SET
		author_id=$2,
		receiver_id=$3,
		status=$4,
		title=$5,
		description=$6,
		created_at=$7,
		updated_at=$8
	WHERE id = $1;`,
	)
	_, err = conn.Exec(
		ctx,
		updateCommand,
		&id,
		&authorId,
		&receiverId,
		&status,
		&title,
		&description,
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		return err
	}
	defer conn.Release()
	return nil
}
