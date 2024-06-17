package repositories

import (
	"context"
	"go-complaint/domain/model/common"
	"go-complaint/domain/model/feedback"
	"go-complaint/infrastructure/persistence/datasource"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type FeedbackReviewRepository struct {
	schema datasource.Schema
}

func NewFeedbackReviewRepository(
	feedbackSchema datasource.Schema,
) FeedbackReviewRepository {
	return FeedbackReviewRepository{schema: feedbackSchema}
}

func (fr FeedbackReviewRepository) Save(
	ctx context.Context,
	review feedback.Review,
) error {
	conn, err := fr.schema.Acquire(ctx)
	if err != nil {
		return nil
	}
	insertCommand := string(`
		INSERT INTO feedback_reviews
		(
			id,
			reviewer_id,
			reviewed_at,
			review_comment
		)
		VALUES ($1, $2, $3, $4)`)
	var (
		reviewID   = review.ReplyReviewID()
		email      = review.ReviewerID()
		reviewedAt = review.ReviewedAt().StringRepresentation()
		comment    = review.Comment()
	)
	_, err = conn.Exec(
		ctx,
		insertCommand,
		&reviewID,
		&email,
		&reviewedAt,
		&comment,
	)
	if err != nil {
		return err
	}
	defer conn.Release()
	return nil
}

func (fr FeedbackReviewRepository) Find(
	ctx context.Context,
	source StatementSource,
) (*feedback.Review, error) {
	conn, err := fr.schema.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	row := conn.QueryRow(ctx, source.Query(), source.Args()...)
	review, err := fr.load(ctx, row)
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	return review, nil
}

func (fr FeedbackReviewRepository) FindAll(
	ctx context.Context,
	source StatementSource,
) (mapset.Set[*feedback.Review], error) {
	conn, err := fr.schema.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	rows, err := conn.Query(ctx, source.Query(), source.Args()...)
	if err != nil {
		return nil, err
	}
	reviews, err := fr.loadAll(ctx, rows)
	if err != nil {
		return nil, err
	}
	defer func() {
		rows.Close()
		conn.Release()
	}()
	return reviews, nil
}

func (fr FeedbackReviewRepository) loadAll(
	ctx context.Context,
	rows pgx.Rows,
) (mapset.Set[*feedback.Review], error) {
	reviews := mapset.NewSet[*feedback.Review]()
	for rows.Next() {
		review, err := fr.load(ctx, rows)
		if err != nil {
			return nil, err
		}
		reviews.Add(review)
	}
	return reviews, nil
}

func (fr FeedbackReviewRepository) load(
	_ context.Context,
	row pgx.Row,
) (*feedback.Review, error) {
	mapper := MapperRegistryInstance().Get("Employee")
	if mapper == nil {
		return nil, ErrMapperNotRegistered
	}
	var (
		ID            uuid.UUID
		reviewerID    string
		reviewedAt    string
		reviewComment string
	)
	err := row.Scan(&ID, &reviewerID, &reviewedAt, &reviewComment)
	if err != nil {
		return nil, err
	}
	reviewedAtDate, err := common.NewDateFromString(reviewedAt)
	if err != nil {
		return nil, err
	}
	return feedback.NewReview(
		ID,
		reviewerID,
		reviewedAtDate,
		reviewComment,
	)
}
