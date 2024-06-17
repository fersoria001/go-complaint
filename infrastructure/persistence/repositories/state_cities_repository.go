package repositories

import (
	"context"
	"go-complaint/domain/model/common"
	"go-complaint/infrastructure/persistence/datasource"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/jackc/pgx/v5"
)

type StateCitiesRepository struct {
	schema datasource.Schema
}

func NewStateCitiesRepository(schema datasource.Schema) StateCitiesRepository {
	return StateCitiesRepository{schema: schema}
}

func (cr StateCitiesRepository) Get(
	ctx context.Context,
	countryID int,
) (common.City, error) {
	conn, err := cr.schema.Acquire(ctx)
	if err != nil {
		return common.City{}, err
	}
	selectQuery := string(`
		SELECT
		id,
		name,
		country_code,
	 	latitude,
	  	longitude
		FROM public."state_cities"
		WHERE id = $1
		`)
	row := conn.QueryRow(
		ctx,
		selectQuery,
		countryID,
	)
	countryState, err := cr.load(ctx, row)
	if err != nil {
		return common.City{}, err
	}
	defer conn.Release()
	return countryState, nil
}

func (cr StateCitiesRepository) FindAll(
	ctx context.Context,
	source StatementSource,
) (mapset.Set[common.City], error) {
	conn, err := cr.schema.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	rows, err := conn.Query(
		ctx,
		source.Query(),
		source.Args()...,
	)
	if err != nil {
		return nil, err
	}
	cities, err := cr.loadAll(ctx, rows)
	if err != nil {
		return nil, err
	}
	defer func() {
		rows.Close()
		conn.Release()
	}()
	return cities, nil
}

func (cr StateCitiesRepository) loadAll(
	ctx context.Context,
	rows pgx.Rows,
) (mapset.Set[common.City], error) {
	cities := mapset.NewSet[common.City]()
	for rows.Next() {
		city, err := cr.load(ctx, rows)
		if err != nil {
			return nil, err
		}
		cities.Add(city)
	}
	return cities, nil
}

func (cr StateCitiesRepository) load(
	_ context.Context,
	row pgx.Row,
) (common.City, error) {
	var (
		id          int
		name        string
		countryCode string
		latitude    float64
		longitude   float64
	)
	err := row.Scan(
		&id,
		&name,
		&countryCode,
		&latitude,
		&longitude,
	)
	if err != nil {
		return common.City{}, err
	}
	return common.NewCity(
		id,
		name,
		countryCode,
		latitude,
		longitude,
	), nil
}
