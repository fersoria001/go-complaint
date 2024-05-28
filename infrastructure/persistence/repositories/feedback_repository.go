package repositories

import (
	"context"
	"fmt"
	"go-complaint/domain/model/common"
	"go-complaint/domain/model/feedback"
	"go-complaint/erros"
	"go-complaint/infrastructure/persistence/datasource"
	"go-complaint/infrastructure/persistence/models"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/google/uuid"
)

type FeedbackRepository struct {
	schema *datasource.Schema
}

func NewFeedbackRepository(feedbackSchema *datasource.Schema) *FeedbackRepository {
	return &FeedbackRepository{schema: feedbackSchema}
}

func (feedbackRepository *FeedbackRepository) Save(
	ctx context.Context,
	entity interface{},
) error {
	aggregate, ok := entity.(*feedback.Feedback)
	if !ok {
		return &erros.InvalidTypeError{}
	}
	if aggregate.ReplyReview().Cardinality() == 0 {
		return &erros.NoElementError{}
	}
	emptyModel := models.Feedback{}
	insertcommand := fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES (%s)",
		emptyModel.Table(),
		models.StringColumns(emptyModel.Columns()),
		emptyModel.Args(),
	)
	conn, err := feedbackRepository.schema.Pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()

	tx, err := conn.Begin(ctx)
	if err != nil {
		return err
	}

	for item := range aggregate.ReplyReview().Iter() {
		feedbackModel := models.NewFeedback(
			aggregate.ComplaintID(),
			aggregate.ReviewerID(),
			aggregate.ReviewedID(),
			item,
		)
		_, err = tx.Exec(
			ctx,
			insertcommand,
			feedbackModel.Values()...,
		)
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

func (feedbackRepository *FeedbackRepository) Get(
	ctx context.Context,
	id string,
) (interface{}, error) {
	conn, err := feedbackRepository.schema.Pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	aggregateModel := models.Feedback{}
	parsedID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	selectQuery := fmt.Sprintf(
		"SELECT %s FROM %s WHERE feedback_id = $1",
		models.StringColumns(aggregateModel.Columns()),
		aggregateModel.Table(),
	)
	rows, err := conn.Query(ctx, selectQuery, parsedID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	repliesReviewsMapset := mapset.NewSet[*feedback.ReplyReview]()
	var rootID uuid.UUID
	var reviewerID string
	var reviewedID string
	for rows.Next() {
		var feedbackModel models.Feedback
		err = rows.Scan(feedbackModel.Values()...)
		if err != nil {
			return nil, err
		}
		rootID = feedbackModel.FeedbackID
		reviewerID = feedbackModel.ReviewerID
		reviewedID = feedbackModel.ReviewedID
		createdAt, err := common.NewDateFromString(feedbackModel.ReplyCreatedAt)
		if err != nil {
			return nil, err
		}
		readAt, err := common.NewDateFromString(feedbackModel.ReplyReadAt)
		if err != nil {
			return nil, err
		}
		updatedAt, err := common.NewDateFromString(feedbackModel.ReplyUpdatedAt)
		if err != nil {
			return nil, err
		}
		reviewedAt, err := common.NewDateFromString(feedbackModel.ReviewedAt)
		if err != nil {
			return nil, err
		}

		newReply, err := feedback.NewReply(
			feedbackModel.ReplySenderID,
			feedbackModel.ReplySenderIMG,
			feedbackModel.ReplySenderName,
			feedbackModel.ReplyBody,
			createdAt,
			feedbackModel.ReplyRead,
			readAt,
			updatedAt,
		)
		if err != nil {
			return nil, err
		}
		newReview, err := feedback.NewReview(
			feedbackModel.ReviewerID,
			feedbackModel.ReviewerIMG,
			feedbackModel.ReviewerName,
			reviewedAt,
			feedbackModel.ReviewComment,
		)
		if err != nil {
			return nil, err
		}
		newReplyReview, err := feedback.NewReplyReview(
			feedbackModel.ID,
			feedbackModel.FeedbackID,
			newReply,
			newReview,
		)
		if err != nil {
			return nil, err
		}
		repliesReviewsMapset.Add(newReplyReview)
	}

	if rootID == uuid.Nil || reviewerID == "" {
		return nil, &erros.NullValueError{}
	}
	feedbackAggregate, err := feedback.NewFeedback(
		rootID,
		reviewerID,
		reviewedID,
		repliesReviewsMapset,
	)
	if err != nil {
		return nil, err
	}
	return feedbackAggregate, err
}

func (feedbackRepository *FeedbackRepository) Remove(ctx context.Context, id string) error {
	conn, err := feedbackRepository.schema.Pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()
	aggregateModel := models.Feedback{}
	parsedID, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	deleteQuery := fmt.Sprintf(
		"DELETE FROM %s WHERE feedback_id = $1",
		aggregateModel.Table(),
	)
	_, err = conn.Exec(ctx, deleteQuery, parsedID)
	if err != nil {
		return err
	}
	return nil
}
