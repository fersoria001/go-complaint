package repositories

import (
	"context"
	"go-complaint/domain/model/complaint"
	"go-complaint/domain/model/feedback"
	"go-complaint/infrastructure/persistence/datasource"
	"go-complaint/infrastructure/persistence/finders/find_all_feedback_replies"
	"go-complaint/infrastructure/persistence/finders/find_all_feedback_review"

	"log"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type FeedbackReplyReviewRepository struct {
	schema datasource.Schema
}

func NewFeedbackReplyReviewRepository(
	feedbackSchema datasource.Schema,
) FeedbackReplyReviewRepository {
	return FeedbackReplyReviewRepository{schema: feedbackSchema}
}

func (fr FeedbackReplyReviewRepository) DeleteAll(
	ctx context.Context,
	feedbackID uuid.UUID,
) error {
	conn, err := fr.schema.Acquire(ctx)
	if err != nil {
		return err
	}
	deleteCommand := string(`
	DELETE FROM feedback_reply_review
	WHERE feedback_id = $1`)
	_, err = conn.Exec(ctx, deleteCommand, &feedbackID)
	if err != nil {
		return err
	}

	defer conn.Release()
	return nil
}

func (fr FeedbackReplyReviewRepository) SaveAll(
	ctx context.Context,
	feedbackID uuid.UUID,
	replyReviews mapset.Set[feedback.ReplyReview],
) error {
	conn, err := fr.schema.Acquire(ctx)
	if err != nil {
		return err
	}
	mapper := MapperRegistryInstance().Get("FeedbackReply")
	if mapper == nil {
		return ErrMapperNotRegistered
	}
	feedbackRepliesRepository, ok := mapper.(FeedbackRepliesRepository)
	if !ok {
		return ErrWrongTypeAssertion
	}
	mapper = MapperRegistryInstance().Get("Review")
	if mapper == nil {
		return ErrMapperNotRegistered
	}
	feedbackReviewRepository, ok := mapper.(FeedbackReviewRepository)
	if !ok {
		return ErrWrongTypeAssertion
	}
	insertCommand := string(`
		INSERT INTO feedback_reply_review
		(
			id,
			feedback_id,
			review_ID,
			color
		)
		VALUES ($1, $2, $3, $4)`)
	for replyReview := range replyReviews.Iter() {
		var (
			ID         = replyReview.ID()
			feedbackId = replyReview.FeedbackID()
			reviewID   = replyReview.Review().ReplyReviewID()
			color      = replyReview.Color()
		)
		replyReview := replyReview
		_, err = conn.Exec(
			ctx,
			insertCommand,
			&ID,
			&feedbackId,
			&reviewID,
			&color,
		)
		if err != nil {
			return err
		}
		err = feedbackRepliesRepository.Save(
			ctx,
			replyReview.ID(),
			replyReview.Replies())
		if err != nil {
			return err
		}
		err = feedbackReviewRepository.Save(ctx, replyReview.Review())
		if err != nil {
			return err
		}
	}
	defer conn.Release()
	return nil
}

func (fr FeedbackReplyReviewRepository) FindAll(
	ctx context.Context,
	statementSource StatementSource,
) (mapset.Set[*feedback.ReplyReview], error) {
	conn, err := fr.schema.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	rows, err := conn.Query(
		ctx,
		statementSource.Query(),
		statementSource.Args()...,
	)
	if err != nil {
		return nil, err
	}
	replyReviews, err := fr.loadAll(ctx, rows)
	if err != nil {
		return nil, err
	}
	defer func() {
		rows.Close()
		conn.Release()
	}()
	return replyReviews, nil
}

func (fr FeedbackReplyReviewRepository) loadAll(
	ctx context.Context,
	rows pgx.Rows,
) (mapset.Set[*feedback.ReplyReview], error) {
	var replyReviews = mapset.NewSet[*feedback.ReplyReview]()
	for rows.Next() {
		replyReview, err := fr.load(ctx, rows)
		if err != nil {
			return nil, err
		}
		replyReviews.Add(replyReview)
	}
	return replyReviews, nil
}

func (fr FeedbackReplyReviewRepository) load(
	ctx context.Context,
	row pgx.Row,
) (*feedback.ReplyReview, error) {
	mapper := MapperRegistryInstance().Get("FeedbackReply")
	if mapper == nil {
		return nil, ErrMapperNotRegistered
	}
	feedbackRepliesRepository, ok := mapper.(FeedbackRepliesRepository)
	if !ok {
		return nil, ErrWrongTypeAssertion
	}
	mapper = MapperRegistryInstance().Get("Review")
	if mapper == nil {
		return nil, ErrMapperNotRegistered
	}
	feedbackReviewRepository, ok := mapper.(FeedbackReviewRepository)
	if !ok {
		return nil, ErrWrongTypeAssertion
	}
	var (
		ID         uuid.UUID
		feedbackID uuid.UUID
		reviewID   string
		color      string
	)
	err := row.Scan(&ID, &feedbackID, &reviewID, &color)
	if err != nil {
		return nil, err
	}
	replies, err := feedbackRepliesRepository.FindAll(
		ctx,
		find_all_feedback_replies.ByReplyReviewID(ID),
	)
	if err != nil {
		return nil, err
	}
	review, err := feedbackReviewRepository.Find(
		ctx,
		find_all_feedback_review.ByReplyReviewID(ID),
	)
	if err != nil {
		log.Println("review find", err, ID)
		return nil, err
	}
	repliesValueCopy := mapset.NewSet[complaint.Reply]()
	for reply := range replies.Iter() {
		repliesValueCopy.Add(*reply)
	}
	replyReview, err := feedback.NewReplyReview(
		ID,
		feedbackID,
		repliesValueCopy,
		review,
		color,
	)
	if err != nil {
		return nil, err
	}
	return replyReview, nil
}
