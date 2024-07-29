package repositories

import (
	"context"
	"go-complaint/domain/model/complaint"
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
	defer conn.Release()
	if err != nil {
		return err
	}
	deleteCommand := string(`DELETE FROM RATINGS WHERE ID=$1;`)
	_, err = conn.Exec(ctx, deleteCommand, &id)
	if err != nil {
		return err
	}
	return nil
}

func (r RatingRepository) Save(ctx context.Context, rating complaint.Rating) error {
	conn, err := r.schema.Acquire(ctx)
	defer conn.Release()
	if err != nil {
		return err
	}
	insertCommand := string(`INSERT INTO RATINGS
	(ID, RATING, COMMENT) VALUES ($1, $2, $3);`)
	var (
		id      uuid.UUID = rating.Id()
		rate    int       = rating.Rate()
		comment string    = rating.Comment()
	)
	_, err = conn.Exec(
		ctx,
		insertCommand,
		&id,
		&rate,
		&comment,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r RatingRepository) Get(ctx context.Context, id uuid.UUID) (*complaint.Rating, error) {
	conn, err := r.schema.Acquire(ctx)
	defer conn.Release()
	if err != nil {
		return nil, err
	}
	selectQuery := string(`SELECT ID, RATING, COMMENT FROM RATINGS WHERE ID=$1;`)
	row := conn.QueryRow(ctx, selectQuery, &id)
	return r.load(row)
}

func (r RatingRepository) load(row pgx.Row) (*complaint.Rating, error) {
	var (
		id      uuid.UUID
		rating  int
		comment string
	)
	err := row.Scan(&id, &rating, &comment)
	if err != nil {
		return nil, err
	}
	dbR, err := complaint.NewRating(id, rating, comment)
	if err != nil {
		return nil, err
	}
	return &dbR, nil
}
