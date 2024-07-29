package repositories

import (
	"context"
	"go-complaint/domain/model/common"
	"go-complaint/domain/model/complaint"
	"go-complaint/infrastructure/persistence/datasource"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type ComplaintRepliesRepository struct {
	complaintSchema datasource.Schema
}

func NewComplaintRepliesRepository(complaintSchema datasource.Schema) ComplaintRepliesRepository {
	return ComplaintRepliesRepository{
		complaintSchema: complaintSchema,
	}
}

func (rr ComplaintRepliesRepository) Get(
	ctx context.Context,
	replyID uuid.UUID,
) (*complaint.Reply, error) {
	conn, err := rr.complaintSchema.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	selectQuery := string(`
	SELECT
	id,
	complaint_id,
	author_id,
	body,
	is_read,
	read_at,
	created_at,
	updated_at
	FROM complaint_replies WHERE id = $1
	`)
	row := conn.QueryRow(ctx, selectQuery, replyID)
	reply, err := rr.load(ctx, row)
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	return reply, nil
}

func (rr ComplaintRepliesRepository) FindAll(
	ctx context.Context,
	source StatementSource,
) (mapset.Set[*complaint.Reply], error) {
	conn, err := rr.complaintSchema.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	rows, err := conn.Query(ctx, source.Query(), source.Args()...)
	if err != nil {
		return nil, err
	}
	replies, err := rr.loadAll(ctx, rows)
	if err != nil {
		return nil, err
	}
	defer func() {
		rows.Close()
		conn.Release()
	}()
	return replies, nil
}
func (rr ComplaintRepliesRepository) loadAll(
	ctx context.Context,
	rows pgx.Rows,
) (mapset.Set[*complaint.Reply], error) {
	replies := mapset.NewSet[*complaint.Reply]()
	for rows.Next() {
		reply, err := rr.load(ctx, rows)
		if err != nil {
			return nil, err
		}
		replies.Add(reply)
	}
	return replies, nil
}

func (rr ComplaintRepliesRepository) DeleteAll(
	ctx context.Context,
	id uuid.UUID,
) error {
	conn, err := rr.complaintSchema.Acquire(ctx)
	defer conn.Release()
	if err != nil {
		return err
	}
	deleteCommand := string(`DELETE FROM complaint_replies WHERE complaint_id = $1;`)
	_, err = conn.Exec(ctx, deleteCommand, &id)
	if err != nil {
		return err
	}
	return nil
}

func (rr ComplaintRepliesRepository) Remove(
	ctx context.Context,
	id uuid.UUID,
) error {
	conn, err := rr.complaintSchema.Acquire(ctx)
	defer conn.Release()
	if err != nil {
		return err
	}
	deleteCommand := string(`DELETE FROM complaint_replies WHERE id = $1;`)
	_, err = conn.Exec(ctx, deleteCommand, id)
	if err != nil {
		return err
	}
	return nil
}

func (crr ComplaintRepliesRepository) Save(ctx context.Context, reply complaint.Reply) error {
	conn, err := crr.complaintSchema.Acquire(ctx)
	defer conn.Release()
	if err != nil {
		return err
	}
	insertCommand := string(`
	INSERT INTO complaint_replies (
		id,
		complaint_id,
		author_id,
		body,
		is_read,
		read_at,
		created_at,
		updated_at
	) VALUES ($1,$2, $3, $4, $5, $6, $7, $8);`)
	var (
		id          uuid.UUID = reply.ID()
		complaintId uuid.UUID = reply.ComplaintId()
		authorId    uuid.UUID = reply.Sender().Id()
		body        string    = reply.Body()
		isRead      bool      = reply.Read()
		readAt      string    = reply.ReadAt().StringRepresentation()
		createdAt   string    = reply.CreatedAt().StringRepresentation()
		updatedAt   string    = reply.UpdatedAt().StringRepresentation()
	)
	_, err = conn.Exec(
		ctx,
		insertCommand,
		&id,
		&complaintId,
		&authorId,
		&body,
		&isRead,
		&readAt,
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		return err
	}
	return nil
}

func (rr ComplaintRepliesRepository) SaveAll(
	ctx context.Context,
	replies mapset.Set[complaint.Reply],
) error {
	conn, err := rr.complaintSchema.Acquire(ctx)
	defer conn.Release()
	if err != nil {
		return err
	}
	if replies.Cardinality() == 0 {
		return nil
	}
	insertCommand := string(`
	INSERT INTO complaint_replies (
		id,
		complaint_id,
		author_id,
		body,
		is_read,
		read_at,
		created_at,
		updated_at
	) VALUES ($1,$2, $3, $4, $5, $6, $7, $8);`)
	tx, err := conn.Begin(ctx)
	if err != nil {
		return err
	}
	for reply := range replies.Iter() {
		var (
			id          uuid.UUID = reply.ID()
			complaintId uuid.UUID = reply.ComplaintId()
			authorId    uuid.UUID = reply.Sender().Id()
			body        string    = reply.Body()
			isRead      bool      = reply.Read()
			readAt      string    = reply.ReadAt().StringRepresentation()
			createdAt   string    = reply.CreatedAt().StringRepresentation()
			updatedAt   string    = reply.UpdatedAt().StringRepresentation()
		)
		_, err = tx.Exec(
			ctx,
			insertCommand,
			&id,
			&complaintId,
			&authorId,
			&body,
			&isRead,
			&readAt,
			&createdAt,
			&updatedAt,
		)
		if err != nil {
			return err
		}
	}
	err = tx.Commit(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (rr ComplaintRepliesRepository) UpdateAll(
	ctx context.Context,
	replies mapset.Set[complaint.Reply],
) error {
	conn, err := rr.complaintSchema.Acquire(ctx)
	if err != nil {
		return err
	}
	insertCommand := string(`
	UPDATE complaint_replies
	SET
		body=$2,
		is_read=$3,
		read_at=$4,
		created_at=$5,
		updated_at=$6,
	WHERE id = $1;`)
	tx, err := conn.Begin(ctx)
	if err != nil {
		return err
	}
	for reply := range replies.Iter() {
		var (
			id         uuid.UUID = reply.ID()
			body       string    = reply.Body()
			readStatus bool      = reply.Read()
			readAt     string    = reply.ReadAt().StringRepresentation()
			createdAt  string    = reply.CreatedAt().StringRepresentation()
			updatedAt  string    = reply.UpdatedAt().StringRepresentation()
		)
		_, err = tx.Exec(
			ctx,
			insertCommand,
			&id,
			&body,
			&readStatus,
			&readAt,
			&createdAt,
			&updatedAt,
		)
		if err != nil {
			return err
		}
	}
	err = tx.Commit(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()
	return nil
}

func (rr ComplaintRepliesRepository) load(ctx context.Context, row pgx.Row) (*complaint.Reply, error) {
	var (
		id          uuid.UUID
		complaintId uuid.UUID
		authorId    uuid.UUID
		body        string
		isRead      bool
		readAt      string
		createdAt   string
		updatedAt   string
	)

	err := row.Scan(
		&id,
		&complaintId,
		&authorId,
		&body,
		&isRead,
		&readAt,
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		return nil, err
	}
	reg := MapperRegistryInstance()
	recipientRepository, ok := reg.Get("Recipient").(RecipientRepository)
	if !ok {
		return nil, ErrWrongTypeAssertion
	}
	author, err := recipientRepository.Get(ctx, authorId)
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
	commonReadAt, err := common.NewDateFromString(readAt)
	if err != nil {
		return nil, err
	}

	reply, err := complaint.NewReply(
		id,
		complaintId,
		*author,
		body,
		isRead,
		commonCreatedAt,
		commonReadAt,
		commonUpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return reply, nil
}
