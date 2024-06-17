package repositories

import (
	"context"
	"errors"
	"go-complaint/domain/model/feedback"
	"go-complaint/infrastructure/persistence/datasource"
	"go-complaint/infrastructure/persistence/finders/find_all_feedback_answer"
	"go-complaint/infrastructure/persistence/finders/find_all_feedback_reply_review"
	"log"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type FeedbackRepository struct {
	schema datasource.Schema
}

func NewFeedbackRepository(feedbackSchema datasource.Schema) FeedbackRepository {
	return FeedbackRepository{schema: feedbackSchema}
}

func (fr FeedbackRepository) Update(
	ctx context.Context,
	feedback *feedback.Feedback,
) error {
	conn, err := fr.schema.Acquire(ctx)
	if err != nil {
		return err
	}
	feedbackReplyReviewRepository := NewFeedbackReplyReviewRepository(fr.schema)
	feedbackAnswerRepository := NewFeedbackAnswerRepository(fr.schema)
	err = feedbackReplyReviewRepository.DeleteAll(ctx, feedback.ComplaintID())
	if err != nil {
		return err
	}
	err = feedbackReplyReviewRepository.SaveAll(
		ctx,
		feedback.ComplaintID(),
		feedback.ReplyReview(),
	)
	if err != nil {
		return err
	}
	err = feedbackAnswerRepository.DeleteAll(ctx, feedback.ComplaintID())
	if err != nil {
		return err
	}
	err = feedbackAnswerRepository.Save(ctx, feedback.FeedbackAnswers())
	if err != nil {
		return err
	}
	defer conn.Release()
	return nil
}

func (fr FeedbackRepository) Save(
	ctx context.Context,
	feedback *feedback.Feedback,
) error {
	conn, err := fr.schema.Acquire(ctx)
	if err != nil {
		return err
	}
	feedbackReplyReviewRepository := NewFeedbackReplyReviewRepository(fr.schema)
	feedbackAnswerRepository := NewFeedbackAnswerRepository(fr.schema)
	err = feedbackReplyReviewRepository.SaveAll(
		ctx,
		feedback.ComplaintID(),
		feedback.ReplyReview(),
	)
	if err != nil {
		return err
	}
	err = feedbackAnswerRepository.Save(ctx, feedback.FeedbackAnswers())
	if err != nil {
		return err
	}
	insertCommand := string(`
		INSERT INTO feedback 
			(
			ID,
			complaint_id,
			reviewed_id
			)
		VALUES ($1,$2, $3)`)
	var (
		id          = feedback.ID()
		complaintID = feedback.ComplaintID()
		reviewedID  = feedback.ReviewedID()
	)
	_, err = conn.Exec(
		ctx,
		insertCommand,
		&id,
		&complaintID,
		&reviewedID,
	)
	if err != nil {
		return err
	}
	defer conn.Release()
	return nil
}

func (fr FeedbackRepository) Get(
	ctx context.Context,
	feedbackID uuid.UUID,
) (*feedback.Feedback, error) {
	conn, err := fr.schema.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	selectQuery := string(`
		SELECT 
			id,
			complaint_id,
			reviewed_id
		FROM feedback
		WHERE id = $1`)
	row := conn.QueryRow(ctx, selectQuery, feedbackID)
	feedback, err := fr.load(ctx, row)
	if err != nil {
		log.Println("err at load")
		return nil, err
	}
	defer conn.Release()
	return feedback, nil
}

func (fr FeedbackRepository) FindAll(
	ctx context.Context,
	source StatementSource,
) (mapset.Set[*feedback.Feedback], error) {
	conn, err := fr.schema.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	rows, err := conn.Query(ctx, source.Query(), source.Args()...)
	if err != nil {
		return nil, err
	}
	feedbacks, err := fr.loadAll(ctx, rows)
	if err != nil {
		return nil, err
	}
	defer func() {
		rows.Close()
		conn.Release()
	}()
	return feedbacks, nil
}

func (fr FeedbackRepository) loadAll(
	ctx context.Context,
	rows pgx.Rows,
) (mapset.Set[*feedback.Feedback], error) {
	feedbacks := mapset.NewSet[*feedback.Feedback]()
	for rows.Next() {
		feedback, err := fr.load(ctx, rows)
		if err != nil {
			return nil, err
		}
		feedbacks.Add(feedback)
	}
	return feedbacks, nil
}

func (fr FeedbackRepository) load(ctx context.Context, row pgx.Row) (*feedback.Feedback, error) {
	var (
		feedbackID  uuid.UUID
		complaintID uuid.UUID
		reviewedID  string
	)
	err := row.Scan(&feedbackID, &complaintID, &reviewedID)
	if err != nil {
		log.Println("err at scan")
		return nil, err
	}
	reply_reviews, err := MapperRegistryInstance().Get("ReplyReview").(FeedbackReplyReviewRepository).FindAll(
		ctx,
		find_all_feedback_reply_review.ByFeedbackID(feedbackID),
	)
	if err != nil {
		log.Println("err at replyreview load")
		return nil, err
	}
	answers, err := MapperRegistryInstance().Get("Answer").(FeedbackAnswerRepository).FindAll(
		ctx,
		find_all_feedback_answer.ByFeedbackID(feedbackID),
	)
	if err != nil {
		log.Println("err at answer load")
		if errors.Is(err, pgx.ErrNoRows) {
			answers = mapset.NewSet[*feedback.Answer]()
		} else {
			return nil, err
		}
	}
	return feedback.NewFeedback(
		feedbackID,
		complaintID,
		reviewedID,
		reply_reviews,
		answers,
	)
}
