package repositories

import (
	"context"
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

func (fr FeedbackReviewRepository) DeleteAll(
	ctx context.Context,
	feedbackID uuid.UUID,
) error {
	conn, err := fr.schema.Acquire(ctx)
	if err != nil {
		return err
	}
	deleteCommand := string(`DELETE FROM feedback_reviews WHERE feedback_id = $1`)
	_, err = conn.Exec(ctx, deleteCommand, feedbackID)
	if err != nil {
		return err
	}
	defer conn.Release()
	return nil
}

func (fr FeedbackReviewRepository) Update(
	ctx context.Context,
	review feedback.Review,
) error {
	con, err := fr.schema.Acquire(ctx)
	if err != nil {
		return err
	}
	updateCommand := string(`
		UPDATE feedback_reviews
		SET review_comment = $1
		WHERE id = $2`)
	var (
		reviewID = review.ReplyReviewID()
		comment  = review.Comment()
	)
	_, err = con.Exec(
		ctx,
		updateCommand,
		&comment,
		&reviewID,
	)
	if err != nil {
		return err
	}
	defer con.Release()
	return nil
}

func (fr FeedbackReviewRepository) Save(
	ctx context.Context,
	review feedback.Review,
) error {
	conn, err := fr.schema.Acquire(ctx)
	if err != nil {
		return err
	}
	insertCommand := string(`
		INSERT INTO feedback_reviews
		(
			id,
			feedback_id,
			review_comment
		)
		VALUES ($1, $2,$3)`)
	var (
		reviewID   = review.ReplyReviewID()
		feedbackID = review.FeedbackID()
		comment    = review.Comment()
	)
	_, err = conn.Exec(
		ctx,
		insertCommand,
		&reviewID,
		&feedbackID,
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
	ctx context.Context,
	row pgx.Row,
) (*feedback.Review, error) {
	var (
		ID            uuid.UUID
		reviewComment string
		feedbackID    uuid.UUID
	)
	err := row.Scan(&ID, &feedbackID, &reviewComment)
	if err != nil {
		return nil, err
	}
	return feedback.NewReview(
		feedbackID,
		ID,
		reviewComment,
	), nil
}
