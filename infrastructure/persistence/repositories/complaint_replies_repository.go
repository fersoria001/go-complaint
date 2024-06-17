package repositories

import (
	"context"
	"go-complaint/domain/model/common"
	"go-complaint/domain/model/complaint"
	"go-complaint/infrastructure/persistence/datasource"
	"log"

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
	"complaint_replies".id,
	"complaint_replies".complaint_id,
	"complaint_replies".author_id,
	"complaint_replies".body,
	"complaint_replies".read_status,
	"complaint_replies".read_at,
	"complaint_replies".created_at,
	"complaint_replies".updated_at,
	"complaint_replies".is_enterprise,
	"complaint_replies".enterprise_id
	FROM
	public."complaint_replies"
	WHERE id = $1
	`)
	row := conn.QueryRow(ctx, selectQuery, replyID)
	reply, err := rr.load(ctx, row)
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	return reply, nil
}

func (rr ComplaintRepliesRepository) FindAllByComplaintID(
	ctx context.Context,
	complaintID uuid.UUID,
) (mapset.Set[*complaint.Reply], error) {
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
	read_status,
	read_at,
	created_at,
	updated_at,
	is_enterprise,
	enterprise_id
	FROM
	public."complaint_replies"
	WHERE complaint_id = $1
	`)
	rows, err := conn.Query(ctx, selectQuery, complaintID)
	if err != nil {
		return nil, err
	}
	replies := mapset.NewSet[*complaint.Reply]()
	for rows.Next() {
		reply, err := rr.load(ctx, rows)
		if err != nil {
			return nil, err
		}
		replies.Add(reply)
	}
	defer func() {
		rows.Close()
		conn.Release()
	}()
	return replies, nil
}

func (rr ComplaintRepliesRepository) DeleteAll(
	ctx context.Context,
	replies mapset.Set[complaint.Reply],
) error {
	conn, err := rr.complaintSchema.Acquire(ctx)
	if err != nil {
		return err
	}
	tx, err := conn.Begin(ctx)
	if err != nil {
		return err
	}
	deleteCommand := string(`
	DELETE FROM public."complaint_replies"
	WHERE id = $1;
	`)
	for reply := range replies.Iter() {
		var replyID = reply.ID()
		_, err = tx.Exec(ctx, deleteCommand, replyID)
		if err != nil {
			tx.Rollback(ctx)
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

func (rr ComplaintRepliesRepository) SaveAll(
	ctx context.Context,
	replies mapset.Set[complaint.Reply],
) error {
	conn, err := rr.complaintSchema.Acquire(ctx)
	if err != nil {
		return err
	}
	if replies.Cardinality() == 0 {
		return nil
	}
	insertCommand := string(`
	INSERT INTO public."complaint_replies" (
		id,
		complaint_id,
		author_id,
		body,
		read_status,
		read_at,
		created_at,
		updated_at,
		is_enterprise,
		enterprise_id
	) VALUES ($1,$2, $3, $4, $5, $6, $7, $8, $9, $10);`)
	tx, err := conn.Begin(ctx)
	if err != nil {
		return err
	}
	for reply := range replies.Iter() {
		var (
			id           uuid.UUID = reply.ID()
			complaintID  uuid.UUID = reply.ComplaintID()
			authorID     string    = reply.SenderID()
			body         string    = reply.Body()
			readStatus   bool      = reply.Read()
			readAt       string    = reply.ReadAt().StringRepresentation()
			createdAt    string    = reply.CreatedAt().StringRepresentation()
			updatedAt    string    = reply.UpdatedAt().StringRepresentation()
			isEnterprise bool      = reply.IsEnterprise()
			enterpriseID string    = reply.EnterpriseID()
		)
		_, err = tx.Exec(
			ctx,
			insertCommand,
			&id,
			&complaintID,
			&authorID,
			&body,
			&readStatus,
			&readAt,
			&createdAt,
			&updatedAt,
			&isEnterprise,
			&enterpriseID,
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

func (rr ComplaintRepliesRepository) UpdateAll(
	ctx context.Context,
	replies mapset.Set[complaint.Reply],
) error {
	conn, err := rr.complaintSchema.Acquire(ctx)
	if err != nil {
		return err
	}
	insertCommand := string(`
	UPDATE public."complaint_replies"
	SET
		id=$1,
		complaint_id=$2,
		author_id=$3,
		body=$4,
		read_status=$5,
		read_at=$6,
		created_at=$7,
		updated_at=$8,
		is_enterprise=$9,
		enterprise_id=$10
	WHERE id = $1;`)
	tx, err := conn.Begin(ctx)
	if err != nil {
		return err
	}
	for reply := range replies.Iter() {
		var (
			id           uuid.UUID = reply.ID()
			complaintID  uuid.UUID = reply.ComplaintID()
			authorID     string    = reply.SenderID()
			body         string    = reply.Body()
			readStatus   bool      = reply.Read()
			readAt       string    = reply.ReadAt().StringRepresentation()
			createdAt    string    = reply.CreatedAt().StringRepresentation()
			updatedAt    string    = reply.UpdatedAt().StringRepresentation()
			isEnterprise bool      = reply.IsEnterprise()
			enterpriseID string    = reply.EnterpriseID()
		)
		_, err = tx.Exec(
			ctx,
			insertCommand,
			&id,
			&complaintID,
			&authorID,
			&body,
			&readStatus,
			&readAt,
			&createdAt,
			&updatedAt,
			&isEnterprise,
			&enterpriseID,
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
	mapper := MapperRegistryInstance().Get("User")
	if mapper == nil {
		return nil, ErrMapperNotRegistered
	}
	userRepository, ok := mapper.(UserRepository)
	if !ok {
		return nil, ErrWrongTypeAssertion
	}
	mapper = MapperRegistryInstance().Get("Enterprise")
	if mapper == nil {
		return nil, ErrMapperNotRegistered
	}
	enterpriseRepository, ok := mapper.(EnterpriseRepository)
	if !ok {
		return nil, ErrWrongTypeAssertion
	}
	var (
		id           uuid.UUID
		complaintID  uuid.UUID
		authorID     string
		body         string
		readStatus   bool
		readAt       string
		createdAt    string
		updatedAt    string
		isEnterprise bool
		enterpriseID string
		authorIMG    string
		authorName   string
	)

	err := row.Scan(
		&id,
		&complaintID,
		&authorID,
		&body,
		&readStatus,
		&readAt,
		&createdAt,
		&updatedAt,
		&isEnterprise,
		&enterpriseID,
	)
	if err != nil {
		return nil, err
	}
	author, err := userRepository.Get(ctx, authorID)
	if err != nil {
		return nil, err
	}
	authorIMG = author.ProfileIMG()
	authorName = author.FullName()
	if isEnterprise {
		dbEnterprise, err := enterpriseRepository.Get(ctx, enterpriseID)
		if err != nil {
			log.Printf("Error: %v", err)
			return nil, err
		}
		authorIMG = dbEnterprise.LogoIMG()
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
		complaintID,
		authorID,
		authorIMG,
		authorName,
		body,
		readStatus,
		commonCreatedAt,
		commonReadAt,
		commonUpdatedAt,
		isEnterprise,
		enterpriseID,
	)
	if err != nil {
		return nil, err
	}
	return reply, nil
}
