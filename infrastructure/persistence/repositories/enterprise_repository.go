package repositories

import (
	"context"
	"fmt"
	"go-complaint/domain/model/common"
	"go-complaint/domain/model/enterprise"
	"go-complaint/dto"
	"go-complaint/infrastructure/persistence/datasource"
	"go-complaint/infrastructure/persistence/models"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type EnterpriseRepository struct {
	schema *datasource.Schema
}

func NewEnterpriseRepository(schema *datasource.Schema) *EnterpriseRepository {
	return &EnterpriseRepository{
		schema: schema,
	}
}

func (r *EnterpriseRepository) Save(ctx context.Context, e *enterprise.Enterprise) error {
	var (
		conn             *pgxpool.Conn
		err              error
		enterprise       *models.Enterprise = models.NewEnterprise(e)
		insertEnterprise string             = fmt.Sprintf(`INSERT INTO %s (%s) VALUES (%s)`,
			enterprise.Table(), models.StringColumns(enterprise.Columns()), models.Args(enterprise.Columns()))
	)
	conn, err = r.schema.Pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()
	_, err = conn.Exec(ctx, insertEnterprise, enterprise.Values()...)
	if err != nil {
		return err
	}
	return nil
}

func (r *EnterpriseRepository) Get(ctx context.Context, id string) (*enterprise.Enterprise, error) {
	var (
		conn             *pgxpool.Conn
		row              pgx.Row
		err              error
		enterpriseModel  *models.Enterprise = &models.Enterprise{}
		industry         enterprise.Industry
		address          common.Address
		fDate            common.Date
		regAt            common.Date
		result           *enterprise.Enterprise
		selectEnterprise = fmt.Sprintf(`SELECT %s FROM %s WHERE id = $1`, models.StringColumns(enterpriseModel.Columns()),
			enterpriseModel.Table())
	)
	conn, err = r.schema.Pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	row = conn.QueryRow(ctx, selectEnterprise, id)
	err = row.Scan(enterpriseModel.Values()...)
	if err != nil {
		return nil, err
	}
	industry, err = enterprise.NewIndustry(0, enterpriseModel.Industry)
	if err != nil {
		return nil, err
	}
	address, err = common.NewAddress(enterpriseModel.Country, enterpriseModel.County, enterpriseModel.City)
	if err != nil {
		return nil, err
	}
	fDate, err = common.NewDateFromString(enterpriseModel.FoundationDate)
	if err != nil {
		return nil, err
	}
	regAt, err = common.NewDateFromString(enterpriseModel.RegisterAt)
	if err != nil {
		return nil, err
	}
	result, err = enterprise.NewEnterprise(
		enterpriseModel.OwnerID,
		enterpriseModel.Name,
		enterpriseModel.LogoIMG,
		enterpriseModel.BannerIMG,
		enterpriseModel.Website,
		enterpriseModel.Email,
		enterpriseModel.Phone,
		address,
		industry,
		regAt,
		fDate)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Designed only for tests id LIKE $1 can behave unexpectedly
func (r *EnterpriseRepository) Remove(ctx context.Context, id string) error {
	var (
		conn             *pgxpool.Conn
		err              error
		enterpriseModel  *models.Enterprise = &models.Enterprise{}
		employeeModel    *models.Employee   = &models.Employee{}
		deleteEnterprise                    = fmt.Sprintf(`DELETE FROM %s WHERE id = $1`, enterpriseModel.Table())
		deleteEmployees                     = fmt.Sprintf(`DELETE FROM %s WHERE id LIKE $1`, employeeModel.Table())
	)
	conn, err = r.schema.Pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()
	tx, err := conn.Begin(ctx)
	if err != nil {
		return err
	}
	_, err = tx.Exec(ctx, deleteEnterprise, id)
	if err != nil {
		tx.Rollback(ctx)
		return err
	}
	_, err = tx.Exec(ctx, deleteEmployees, id)
	if err != nil {
		tx.Rollback(ctx)
		return err
	}
	err = tx.Commit(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (r *EnterpriseRepository) GetAll(ctx context.Context) (mapset.Set[*enterprise.Enterprise], error) {
	var (
		conn             *pgxpool.Conn
		err              error
		enterpriseModel  *models.Enterprise = &models.Enterprise{}
		industry         enterprise.Industry
		address          common.Address
		fDate            common.Date
		regAt            common.Date
		result           *enterprise.Enterprise
		selectEnterprise = fmt.Sprintf(`SELECT %s FROM %s`,
			models.StringColumns(enterpriseModel.Columns()),
			enterpriseModel.Table())
	)
	conn, err = r.schema.Pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	rows, err := conn.Query(ctx, selectEnterprise)
	if err != nil {
		return nil, err
	}
	enterprises := mapset.NewSet[*enterprise.Enterprise]()
	for rows.Next() {
		err = rows.Scan(enterpriseModel.Values()...)
		if err != nil {

			return nil, err
		}

		industry, err = enterprise.NewIndustry(0, enterpriseModel.Industry)
		if err != nil {

			return nil, err
		}
		address, err = common.NewAddress(enterpriseModel.Country, enterpriseModel.County, enterpriseModel.City)
		if err != nil {

			return nil, err
		}
		fDate, err = common.NewDateFromString(enterpriseModel.FoundationDate)
		if err != nil {

			return nil, err
		}
		regAt, err = common.NewDateFromString(enterpriseModel.RegisterAt)
		if err != nil {

			return nil, err
		}
		result, err = enterprise.NewEnterprise(
			enterpriseModel.OwnerID,
			enterpriseModel.Name,
			enterpriseModel.LogoIMG,
			enterpriseModel.BannerIMG,
			enterpriseModel.Website,
			enterpriseModel.Email,
			enterpriseModel.Phone,
			address,
			industry,
			regAt,
			fDate)
		if err != nil {

			return nil, err
		}
		enterprises.Add(result)
	}

	return enterprises, nil
}

func (r *EnterpriseRepository) Update(ctx context.Context, e *enterprise.Enterprise) error {
	var (
		conn             *pgxpool.Conn
		err              error
		enterprise       *models.Enterprise = models.NewEnterprise(e)
		updateEnterprise string             = fmt.Sprintf(`UPDATE %s SET %s WHERE id = $2`,
			enterprise.Table(), models.StringKeyArgs(enterprise.Columns()))
	)
	conn, err = r.schema.Pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()

	_, err = conn.Exec(ctx, updateEnterprise, enterprise.Values()...)
	if err != nil {

		return err
	}

	return nil
}

//additional behaviour

func (r *EnterpriseRepository) GetIndustriesSlice(ctx context.Context) ([]dto.Industry, error) {
	var (
		conn       *pgxpool.Conn
		rows       pgx.Rows
		err        error
		industries []dto.Industry
	)
	conn, err = r.schema.Pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	rows, err = conn.Query(ctx, `SELECT id,name  FROM industries`)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var industry dto.Industry
		err = rows.Scan(&industry.ID, &industry.Name)
		if err != nil {
			return nil, err
		}
		industries = append(industries, industry)
	}
	return industries, nil
}

func (r *EnterpriseRepository) FindByName(ctx context.Context, term string) ([]*dto.Receiver, error) {
	conn, err := r.schema.Pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	wildCardTerm := "%" + strings.ToLower(term) + "%"
	rows, err := conn.Query(ctx, `SELECT id, logo_IMG FROM enterprises WHERE LOWER(id) LIKE $1`, wildCardTerm)
	if err != nil {
		return nil, err
	}
	receivers := make([]*dto.Receiver, 0)
	var id string
	var logoIMG string
	for rows.Next() {
		var receiver = &dto.Receiver{}
		err = rows.Scan(&id, &logoIMG)
		if err != nil {
			return nil, err
		}
		receiver.ID = id
		receiver.IMG = logoIMG
		receiver.FullName = id
		receivers = append(receivers, receiver)
	}
	return receivers, nil
}
func (r *EnterpriseRepository) FindByOwnerID(ctx context.Context, id string) (mapset.Set[*enterprise.Enterprise], error) {
	var (
		conn             *pgxpool.Conn
		err              error
		enterpriseModel  *models.Enterprise = &models.Enterprise{}
		industry         enterprise.Industry
		address          common.Address
		fDate            common.Date
		regAt            common.Date
		result           *enterprise.Enterprise
		selectEnterprise = fmt.Sprintf(`SELECT %s FROM %s where owner_id = $1`,
			models.StringColumns(enterpriseModel.Columns()),
			enterpriseModel.Table())
	)
	conn, err = r.schema.Pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	rows, err := conn.Query(ctx, selectEnterprise, id)
	if err != nil {
		return nil, err
	}
	enterprises := mapset.NewSet[*enterprise.Enterprise]()
	for rows.Next() {
		err = rows.Scan(enterpriseModel.Values()...)
		if err != nil {

			return nil, err
		}

		industry, err = enterprise.NewIndustry(0, enterpriseModel.Industry)
		if err != nil {

			return nil, err
		}
		address, err = common.NewAddress(enterpriseModel.Country, enterpriseModel.County, enterpriseModel.City)
		if err != nil {

			return nil, err
		}
		fDate, err = common.NewDateFromString(enterpriseModel.FoundationDate)
		if err != nil {

			return nil, err
		}
		regAt, err = common.NewDateFromString(enterpriseModel.RegisterAt)
		if err != nil {

			return nil, err
		}
		result, err = enterprise.NewEnterprise(
			enterpriseModel.OwnerID,
			enterpriseModel.Name,
			enterpriseModel.LogoIMG,
			enterpriseModel.BannerIMG,
			enterpriseModel.Website,
			enterpriseModel.Email,
			enterpriseModel.Phone,
			address,
			industry,
			regAt,
			fDate)
		if err != nil {

			return nil, err
		}
		enterprises.Add(result)
	}

	return enterprises, nil
}
