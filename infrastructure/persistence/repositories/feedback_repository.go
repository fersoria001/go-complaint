package repositories

import (
	"context"
	"errors"
	"go-complaint/domain/model/common"
	"go-complaint/domain/model/feedback"
	"go-complaint/infrastructure/persistence/datasource"
	"go-complaint/infrastructure/persistence/finders/find_all_feedback_answer"
	"go-complaint/infrastructure/persistence/finders/find_all_feedback_reply_review"
	"go-complaint/infrastructure/persistence/removers/remove_all_feedback_replies"

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
	f *feedback.Feedback,
) error {
	conn, err := fr.schema.Acquire(ctx)
	if err != nil {
		return err
	}
	feedbackReplyReviewRepository := NewFeedbackReplyReviewRepository(fr.schema)
	feedbackAnswerRepository := NewFeedbackAnswerRepository(fr.schema)
	if f.ReplyReviews() != nil {
		for rr := range f.ReplyReviews().Iter() {
			err = MapperRegistryInstance().Get("FeedbackReply").(FeedbackRepliesRepository).DeleteAll(
				ctx,
				remove_all_feedback_replies.WhereReplyReviewID(rr.ID()),
			)
			if err != nil {
				return err
			}
		}
		err = feedbackReplyReviewRepository.DeleteAll(ctx, f.ID())
		if err != nil {
			return err
		}
		err = feedbackReplyReviewRepository.SaveAll(
			ctx,
			f.ReplyReviews(),
		)
		if err != nil {
			return err
		}
	}
	if f.FeedbackAnswers() != nil {
		err = feedbackAnswerRepository.DeleteAll(ctx, f.ComplaintID())
		if err != nil {
			return err
		}
		err = feedbackAnswerRepository.Save(ctx, f.FeedbackAnswers())
		if err != nil {
			return err
		}
	}
	updateCommand := string(`
		UPDATE feedback
		SET
		updated_at = $2,
		is_done = $3
		WHERE id = $1`)
	var (
		id        = f.ID()
		updatedAt = common.StringDate(f.UpdatedAt())
		isDone    = f.IsDone()
	)
	_, err = conn.Exec(
		ctx,
		updateCommand,
		&id,
		&updatedAt,
		isDone,
	)
	if err != nil {
		return err
	}
	defer conn.Release()
	return nil
}

func (fr FeedbackRepository) Save(
	ctx context.Context,
	f *feedback.Feedback,
) error {
	conn, err := fr.schema.Acquire(ctx)
	if err != nil {
		return err
	}
	if f.ReplyReviews() != nil {
		feedbackReplyReviewRepository := NewFeedbackReplyReviewRepository(fr.schema)
		err = feedbackReplyReviewRepository.SaveAll(
			ctx,
			f.ReplyReviews(),
		)
		if err != nil {
			return err
		}
	}
	if f.FeedbackAnswers() != nil {
		feedbackAnswerRepository := NewFeedbackAnswerRepository(fr.schema)
		err = feedbackAnswerRepository.Save(ctx, f.FeedbackAnswers())
		if err != nil {
			return err
		}
	}
	insertCommand := string(`
		INSERT INTO feedback 
			(
			ID,
			enterprise_id,
			complaint_id,
			reviewed_at,
			updated_at,
			is_done
			)
		VALUES ($1,$2,$3,$4,$5,$6)`)
	var (
		id           = f.ID()
		complaintID  = f.ComplaintID()
		enterpriseID = f.EnterpriseID()
		reviewedAt   = common.StringDate(f.ReviewedAt())
		updatedAt    = common.StringDate(f.UpdatedAt())
		isDone       = f.IsDone()
	)
	_, err = conn.Exec(
		ctx,
		insertCommand,
		&id,
		&enterpriseID,
		&complaintID,
		&reviewedAt,
		&updatedAt,
		&isDone,
	)
	if err != nil {
		return err
	}
	defer conn.Release()
	return nil
}

func (fr FeedbackRepository) Count(
	ctx context.Context,
	source StatementSource,
) (int, error) {
	conn, err := fr.schema.Acquire(ctx)
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
			enterprise_id,
			reviewed_at,
			updated_at,
			is_done
		FROM feedback
		WHERE id = $1`)
	row := conn.QueryRow(ctx, selectQuery, feedbackID)
	feedback, err := fr.load(ctx, row)
	if err != nil {

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
		feedbackID   uuid.UUID
		complaintID  uuid.UUID
		enterpriseID string
		reviewedAt   string
		updatedAt    string
		isDone       bool
	)
	err := row.Scan(&feedbackID, &complaintID, &enterpriseID, &reviewedAt, &updatedAt, &isDone)
	if err != nil {

		return nil, err
	}
	reply_reviews, err := MapperRegistryInstance().Get("ReplyReview").(FeedbackReplyReviewRepository).FindAll(
		ctx,
		find_all_feedback_reply_review.ByFeedbackID(feedbackID),
	)
	if err != nil {

		return nil, err
	}
	answers, err := MapperRegistryInstance().Get("Answer").(FeedbackAnswerRepository).FindAll(
		ctx,
		find_all_feedback_answer.ByFeedbackID(feedbackID),
	)
	if err != nil {

		if errors.Is(err, pgx.ErrNoRows) {
			answers = mapset.NewSet[*feedback.Answer]()
		} else {
			return nil, err
		}
	}
	revAt, err := common.ParseDate(reviewedAt)
	if err != nil {
		return nil, err
	}
	updAt, err := common.ParseDate(updatedAt)
	if err != nil {
		return nil, err
	}
	return feedback.NewFeedback(
		feedbackID,
		complaintID,
		enterpriseID,
		reply_reviews,
		answers,
		revAt,
		updAt,
		isDone,
	)
}
