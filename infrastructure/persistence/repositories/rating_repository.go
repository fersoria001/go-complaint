package repositories

import (
	"context"
	"go-complaint/domain/model/common"
	"go-complaint/domain/model/complaint"
	"go-complaint/domain/model/recipient"
	"go-complaint/infrastructure/persistence/datasource"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type RatingRepository struct {
	schema datasource.Schema
}

func NewRatingRepository(schema datasource.Schema) RatingRepository {
	return RatingRepository{
		schema: schema,
	}
}

func (r RatingRepository) Remove(ctx context.Context, id uuid.UUID) error {
	conn, err := r.schema.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()
	deleteCommand := string(`DELETE FROM RATINGS WHERE ID=$1;`)
	_, err = conn.Exec(ctx, deleteCommand, &id)
	if err != nil {
		return err
	}
	return nil
}

func (r RatingRepository) Save(ctx context.Context, rating complaint.Rating) error {
	conn, err := r.schema.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()
	insertCommand := string(`INSERT INTO RATINGS(
	ID,
	CREATED_AT,
	LAST_UPDATE
	) VALUES ($1, $2, $3);`)
	var (
		id         uuid.UUID = rating.Id()
		createdAt  string    = common.StringDate(rating.CreatedAt())
		lastUpdate string    = common.StringDate(rating.LastUpdate())
	)
	_, err = conn.Exec(
		ctx,
		insertCommand,
		&id,
		&createdAt,
		&lastUpdate,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r RatingRepository) Update(ctx context.Context, rating complaint.Rating) error {
	conn, err := r.schema.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()
	updateCommand := string(`
	UPDATE RATINGS SET 
		LAST_UPDATE=$2,
		SENT_TO_REVIEW_BY_ID=$3,
		RATING=$4,
		COMMENT=$5,
		RATED_BY_ID=$6
	WHERE ID=$1`)
	var (
		id               uuid.UUID = rating.Id()
		lastUpdate       string    = common.StringDate(rating.LastUpdate())
		sentToReviewById uuid.UUID = rating.SentToReviewBy().Id()
		rate             int       = rating.Rate()
		comment          string    = rating.Comment()
		ratedById        uuid.UUID = rating.RatedBy().Id()
	)
	_, err = conn.Exec(ctx, updateCommand, &id, &lastUpdate, &sentToReviewById,
		&rate, &comment, &ratedById)
	if err != nil {
		return err
	}
	return nil
}

func (r RatingRepository) Get(ctx context.Context, id uuid.UUID) (*complaint.Rating, error) {
	conn, err := r.schema.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	selectQuery := string(`SELECT 
	ID,
	CREATED_AT,
	LAST_UPDATE,
	SENT_TO_REVIEW_BY_ID,
	RATING,
	COMMENT,
	RATED_BY_ID
	FROM RATINGS
	WHERE ID = $1
	`)
	row := conn.QueryRow(ctx, selectQuery, &id)
	return r.load(ctx, row)
}

func (r RatingRepository) load(ctx context.Context, row pgx.Row) (*complaint.Rating, error) {
	var (
		id               uuid.UUID
		createdAt        string
		lastUpdate       string
		sentToReviewById uuid.UUID
		rate             int
		comment          string
		ratedById        uuid.UUID
		sentToReviewBy   *recipient.Recipient = &recipient.Recipient{}
		ratedBy          *recipient.Recipient = &recipient.Recipient{}
	)
	err := row.Scan(&id, &createdAt, &lastUpdate, &sentToReviewById,
		&rate, &comment, &ratedById)
	if err != nil {
		return nil, err
	}
	reg := MapperRegistryInstance()
	recipientRepository, ok := reg.Get("Recipient").(RecipientRepository)
	if !ok {
		return nil, ErrWrongTypeAssertion
	}
	if sentToReviewById != uuid.Nil {
		sentToReviewBy, err = recipientRepository.Get(ctx, sentToReviewById)
		if err != nil {
			return nil, err
		}
	}
	if ratedById != uuid.Nil {
		ratedBy, err = recipientRepository.Get(ctx, ratedById)
		if err != nil {
			return nil, err
		}
	}
	createdAtDate, err := common.NewDateFromString(createdAt)
	if err != nil {
		return nil, err
	}
	lastUpdateDate, err := common.NewDateFromString(lastUpdate)
	if err != nil {
		return nil, err
	}
	return complaint.NewRating(
		id,
		*sentToReviewBy,
		*ratedBy,
		rate,
		comment,
		createdAtDate.Date(),
		lastUpdateDate.Date(),
	), nil
}
