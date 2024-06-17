package repositories

import (
	"context"
	"go-complaint/domain/model/common"
	"go-complaint/infrastructure/persistence/datasource"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/jackc/pgx/v5"
)

type CountryRepository struct {
	schema datasource.Schema
}

func NewCountryRepository(schema datasource.Schema) CountryRepository {
	return CountryRepository{schema: schema}
}

func (cr CountryRepository) Get(
	ctx context.Context,
	countryID int,
) (common.Country, error) {
	conn, err := cr.schema.Acquire(ctx)
	if err != nil {
		return common.Country{}, err
	}
	selectQuery := string(`
		SELECT
		"country".id,
		"country".name,
		"country".phonecode
		FROM public."country"
		WHERE id = $1
		`)
	row := conn.QueryRow(
		ctx,
		selectQuery,
		countryID,
	)
	country, err := cr.load(ctx, row)
	if err != nil {
		return common.Country{}, err
	}
	defer conn.Release()
	return country, nil
}

func (cr CountryRepository) FindAll(
	ctx context.Context,
	source StatementSource,
) (mapset.Set[common.Country], error) {
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
	countries, err := cr.loadAll(ctx, rows)
	if err != nil {
		return nil, err
	}
	defer func() {
		rows.Close()
		conn.Release()
	}()
	return countries, nil
}

func (cr CountryRepository) loadAll(
	ctx context.Context,
	rows pgx.Rows,
) (mapset.Set[common.Country], error) {
	countries := mapset.NewSet[common.Country]()
	for rows.Next() {
		country, err := cr.load(ctx, rows)
		if err != nil {
			return nil, err
		}
		countries.Add(country)
	}
	return countries, nil
}

func (cr CountryRepository) load(
	_ context.Context,
	row pgx.Row,
) (common.Country, error) {
	var (
		id        int
		name      string
		phonecode string
	)
	err := row.Scan(
		&id,
		&name,
		&phonecode,
	)
	if err != nil {
		return common.Country{}, err
	}
	return common.NewCountry(id, name, phonecode), nil
}
