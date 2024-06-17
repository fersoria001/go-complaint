package repositories

import (
	"context"
	"go-complaint/domain/model/enterprise"
	"go-complaint/infrastructure/persistence/datasource"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/jackc/pgx/v5"
)

type IndustryRepository struct {
	schema datasource.Schema
}

func NewIndustryRepository(schema datasource.Schema) IndustryRepository {
	return IndustryRepository{
		schema: schema,
	}
}

func (ir IndustryRepository) Get(ctx context.Context, industryID int) (
	enterprise.Industry, error) {
	conn, err := ir.schema.Acquire(ctx)
	if err != nil {
		return enterprise.Industry{}, err
	}
	selectQuery := string(
		`SELECT 
		industry_id,
		name
		FROM industry
		WHERE industry_id = $1`,
	)
	row := conn.QueryRow(ctx, selectQuery, industryID)
	industry, err := ir.load(ctx, row)
	if err != nil {
		return enterprise.Industry{}, err
	}
	defer conn.Release()
	return industry, nil
}

func (ir IndustryRepository) FindAll(
	ctx context.Context,
	source StatementSource,
) (mapset.Set[enterprise.Industry], error) {
	conn, err := ir.schema.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	rows, err := conn.Query(ctx, source.Query(), source.Args()...)
	if err != nil {
		return nil, err
	}
	industries, err := ir.loadAll(ctx, rows)
	if err != nil {
		return nil, err
	}
	defer func() {
		rows.Close()
		conn.Release()
	}()
	return industries, nil
}

func (ir *IndustryRepository) loadAll(
	ctx context.Context,
	rows pgx.Rows,
) (mapset.Set[enterprise.Industry], error) {
	industries := mapset.NewSet[enterprise.Industry]()
	for rows.Next() {
		industry, err := ir.load(ctx, rows)
		if err != nil {
			return nil, err
		}
		industries.Add(industry)
	}
	return industries, nil
}

func (ir *IndustryRepository) load(_ context.Context, row pgx.Row) (
	enterprise.Industry, error) {
	var (
		industryID int
		name       string
	)
	err := row.Scan(&industryID, &name)
	if err != nil {
		return enterprise.Industry{}, err
	}
	industry, err := enterprise.NewIndustry(industryID, name)
	if err != nil {
		return enterprise.Industry{}, err
	}
	return industry, nil
}
