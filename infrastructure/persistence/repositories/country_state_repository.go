package repositories

import (
	"context"
	"go-complaint/domain/model/common"
	"go-complaint/infrastructure/persistence/datasource"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/jackc/pgx/v5"
)

type CountryStateRepository struct {
	schema datasource.Schema
}

func NewCountryStateRepository(schema datasource.Schema) CountryStateRepository {
	return CountryStateRepository{schema: schema}
}

func (cr CountryStateRepository) Get(
	ctx context.Context,
	countryStateID int,
) (common.CountryState, error) {
	conn, err := cr.schema.Acquire(ctx)
	if err != nil {
		return common.CountryState{}, err
	}
	selectQuery := string(`
		SELECT
		id,
		name
		FROM country_states
		WHERE id = $1
		`)
	row := conn.QueryRow(
		ctx,
		selectQuery,
		countryStateID,
	)
	countryState, err := cr.load(ctx, row)
	if err != nil {
		return common.CountryState{}, err
	}
	defer conn.Release()
	return countryState, nil
}

func (cr CountryStateRepository) Find(
	ctx context.Context,
	source StatementSource,
) (common.CountryState, error) {
	conn, err := cr.schema.Acquire(ctx)
	if err != nil {
		return common.CountryState{}, err
	}
	row := conn.QueryRow(
		ctx,
		source.Query(),
		source.Args()...,
	)
	countryState, err := cr.load(ctx, row)
	if err != nil {
		return common.CountryState{}, err
	}
	defer conn.Release()
	return countryState, nil
}

func (cr CountryStateRepository) FindAll(
	ctx context.Context,
	source StatementSource,
) (mapset.Set[common.CountryState], error) {
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
	countryStates, err := cr.loadAll(ctx, rows)
	if err != nil {
		return nil, err
	}
	defer func() {
		rows.Close()
		conn.Release()
	}()
	return countryStates, nil
}

func (cr CountryStateRepository) loadAll(
	ctx context.Context,
	rows pgx.Rows,
) (mapset.Set[common.CountryState], error) {
	countryStates := mapset.NewSet[common.CountryState]()
	for rows.Next() {
		countryState, err := cr.load(ctx, rows)
		if err != nil {
			return mapset.NewSet[common.CountryState](), err
		}
		countryStates.Add(countryState)
	}
	return countryStates, nil
}

func (cr CountryStateRepository) load(
	_ context.Context,
	row pgx.Row,
) (common.CountryState, error) {
	var (
		id   int
		name string
	)
	err := row.Scan(
		&id,
		&name,
	)
	if err != nil {
		return common.CountryState{}, err
	}
	return common.NewCountryState(id, name), nil
}
