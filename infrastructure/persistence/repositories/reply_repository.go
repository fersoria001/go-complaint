package repositories

import (
	"context"
	"fmt"
	"go-complaint/domain/model/common"
	"go-complaint/domain/model/complaint"
	"go-complaint/infrastructure/persistence/datasource"
	"go-complaint/infrastructure/persistence/models"

	"github.com/google/uuid"
)

type ReplyRepository struct {
	complaintSchema *datasource.Schema
}

func NewReplyRepository(complaintSchema *datasource.Schema) *ReplyRepository {
	return &ReplyRepository{
		complaintSchema: complaintSchema,
	}
}

func (replyRepository *ReplyRepository) Get(ctx context.Context, id string) (*complaint.Reply, error) {
	conn, err := replyRepository.complaintSchema.Pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	parsedID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	model := &models.Reply{}
	selectReply := fmt.Sprintf(
		"SELECT %s FROM %s WHERE id = $1",
		models.StringColumns(model.Columns()),
		model.Table(),
	)
	row := conn.QueryRow(
		ctx,
		selectReply,
		parsedID,
	)
	err = row.Scan(model.Values()...)
	if err != nil {
		return nil, err
	}
	createdAt, err := common.NewDateFromString(model.CreatedAt)
	if err != nil {
		return nil, err
	}
	updatedAt, err := common.NewDateFromString(model.UpdatedAt)
	if err != nil {
		return nil, err
	}
	readAt, err := common.NewDateFromString(model.ReadAt)
	if err != nil {
		return nil, err
	}
	reply, err := complaint.NewReply(
		model.ID,
		model.ComplaintID,
		model.SenderID,
		model.SenderName,
		model.SenderIMG,
		model.Body,
		model.Read,
		createdAt,
		readAt,
		updatedAt,
	)
	if err != nil {
		return nil, err
	}

	return reply, nil
}

func (replyRepository *ReplyRepository) Save(
	ctx context.Context,
	reply *complaint.Reply,
) error {
	conn, err := replyRepository.complaintSchema.Pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()
	model := models.NewReply(reply)
	insertReply := fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES (%s)",
		model.Table(),
		models.StringColumns(model.Columns()),
		model.Args(),
	)
	_, err = conn.Exec(
		ctx,
		insertReply,
		model.Values()...,
	)
	if err != nil {
		return err
	}
	return nil
}

func (replyRepository *ReplyRepository) Update(
	ctx context.Context,
	reply *complaint.Reply,
) error {
	conn, err := replyRepository.complaintSchema.Pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()
	model := models.NewReply(reply)
	updateReply := fmt.Sprintf(
		"UPDATE %s SET %s WHERE id = $1",
		model.Table(),
		models.StringKeyArgs(model.Columns()),
	)
	_, err = conn.Exec(
		ctx,
		updateReply,
		model.Values()...,
	)
	if err != nil {
		return err
	}
	return nil
}

func (replyRepository *ReplyRepository) Remove(
	ctx context.Context,
	id string,
) error {
	conn, err := replyRepository.complaintSchema.Pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()
	parsedID, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	model := &models.Reply{}
	deleteReply := fmt.Sprintf(
		"DELETE FROM %s WHERE id = $1",
		model.Table(),
	)
	_, err = conn.Exec(
		ctx,
		deleteReply,
		parsedID,
	)
	if err != nil {
		return err
	}
	return nil
}

// Additional behaviour
func (replyRepository *ReplyRepository) FindByComplaintID(
	ctx context.Context,
	complaintID string,
) ([]*complaint.Reply, error) {
	conn, err := replyRepository.complaintSchema.Pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	parsedID, err := uuid.Parse(complaintID)
	if err != nil {
		return nil, err
	}
	model := &models.Reply{}
	selectReply := fmt.Sprintf(
		"SELECT %s FROM %s WHERE complaint_id = $1",
		models.StringColumns(model.Columns()),
		model.Table(),
	)
	rows, err := conn.Query(
		ctx,
		selectReply,
		parsedID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	replies := []*complaint.Reply{}
	for rows.Next() {
		err = rows.Scan(model.Values()...)
		if err != nil {
			return nil, err
		}
		createdAt, err := common.NewDateFromString(model.CreatedAt)
		if err != nil {
			return nil, err
		}
		updatedAt, err := common.NewDateFromString(model.UpdatedAt)
		if err != nil {
			return nil, err
		}
		readAt, err := common.NewDateFromString(model.ReadAt)
		if err != nil {
			return nil, err
		}
		reply, err := complaint.NewReply(
			model.ID,
			model.ComplaintID,
			model.SenderID,
			model.SenderName,
			model.SenderIMG,
			model.Body,
			model.Read,
			createdAt,
			readAt,
			updatedAt,
		)
		if err != nil {
			return nil, err
		}
		replies = append(replies, reply)
	}
	return replies, nil
}

//additional behaviour

func (replyRepository *ReplyRepository) Count(ctx context.Context, complaintID string) (int, error) {
	conn, err := replyRepository.complaintSchema.Pool.Acquire(ctx)
	if err != nil {
		return 0, err
	}
	defer conn.Release()
	parsedID, err := uuid.Parse(complaintID)
	if err != nil {
		return 0, err
	}
	model := &models.Reply{}
	countReply := fmt.Sprintf(
		"SELECT COUNT(*) FROM %s WHERE complaint_id = $1",
		model.Table(),
	)
	row := conn.QueryRow(
		ctx,
		countReply,
		parsedID,
	)
	var count int
	err = row.Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}
