package repositories

import (
	"context"
	"go-complaint/domain/model/common"

	"go-complaint/domain/model/enterprise"
	"go-complaint/infrastructure/persistence/datasource"
	employeefindall "go-complaint/infrastructure/persistence/finders/employee_findall"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type EnterpriseRepository struct {
	schema datasource.Schema
}

func NewEnterpriseRepository(schema datasource.Schema) EnterpriseRepository {
	return EnterpriseRepository{
		schema: schema,
	}
}

func (er EnterpriseRepository) Update(
	ctx context.Context,
	updatedEnterprise *enterprise.Enterprise,
) error {
	conn, err := er.schema.Acquire(ctx)
	if err != nil {
		return err
	}
	mapper := MapperRegistryInstance().Get("Address")
	if mapper == nil {
		return ErrMapperNotRegistered
	}
	addressRepository, ok := mapper.(AddressRepository)
	if !ok {
		return ErrWrongTypeAssertion
	}
	mapper = MapperRegistryInstance().Get("Employee")
	if mapper == nil {
		return ErrMapperNotRegistered
	}
	employeeRepository, ok := mapper.(EmployeeRepository)
	if !ok {
		return ErrWrongTypeAssertion
	}
	err = addressRepository.Update(
		ctx,
		updatedEnterprise.Address(),
	)
	if err != nil {
		return err
	}
	err = employeeRepository.DeleteAll(ctx, updatedEnterprise.Id())
	if err != nil {
		return err
	}
	err = employeeRepository.SaveAll(ctx, updatedEnterprise.Employees())
	if err != nil {
		return err
	}
	updateCommand := string(
		`
		UPDATE enterprise
		SET 
			owner_user_id = $2,
			logo_img = $3,
			banner_img = $4,
			website = $5,
			email = $6,
			phone = $7,
			industry_id = $8,
			created_at = $9,
			updated_at = $10,
			foundation_date = $11
		WHERE enterprise.enterprise_id = $1
		`,
	)
	var (
		id             uuid.UUID = updatedEnterprise.Id()
		name           string    = updatedEnterprise.Name()
		owner          uuid.UUID = updatedEnterprise.OwnerId()
		logoIMG        string    = updatedEnterprise.LogoIMG()
		bannerIMG      string    = updatedEnterprise.BannerIMG()
		website        string    = updatedEnterprise.Website()
		email          string    = updatedEnterprise.Email()
		phone          string    = updatedEnterprise.Phone()
		industryID     int       = updatedEnterprise.Industry().ID()
		createdAt      string    = updatedEnterprise.CreatedAt().StringRepresentation()
		updatedAt      string    = updatedEnterprise.UpdatedAt().StringRepresentation()
		foundationDate string    = updatedEnterprise.FoundationDate().StringRepresentation()
	)
	_, err = conn.Exec(
		ctx,
		updateCommand,
		&id,
		&name,
		&owner,
		&logoIMG,
		&bannerIMG,
		&website,
		&email,
		&phone,
		&industryID,
		&createdAt,
		&updatedAt,
		&foundationDate,
	)
	if err != nil {
		return err
	}
	defer conn.Release()
	return nil
}

func (er EnterpriseRepository) Remove(ctx context.Context, id uuid.UUID) error {
	conn, err := er.schema.Acquire(ctx)
	defer conn.Release()
	if err != nil {
		return err
	}
	addressRepository, ok := MapperRegistryInstance().Get("Address").(AddressRepository)
	if !ok {
		return ErrWrongTypeAssertion
	}
	err = addressRepository.Remove(ctx, id)
	if err != nil {
		return err
	}
	employeesRepository, ok := MapperRegistryInstance().Get("Employee").(EmployeeRepository)
	if !ok {
		return ErrWrongTypeAssertion
	}
	err = employeesRepository.DeleteAll(ctx, id)
	if err != nil {
		return err
	}
	_, err = conn.Exec(ctx, "DELETE FROM ENTERPRISE WHERE ENTERPRISE_ID=$1", &id)
	if err != nil {
		return err
	}
	return nil
}

func (er EnterpriseRepository) Save(
	ctx context.Context,
	enterprise *enterprise.Enterprise,
) error {
	conn, err := er.schema.Acquire(ctx)
	defer conn.Release()
	if err != nil {
		return err
	}
	mapper := MapperRegistryInstance().Get("Address")
	if mapper == nil {
		return ErrMapperNotRegistered
	}
	addressRepository, ok := mapper.(AddressRepository)
	if !ok {
		return ErrWrongTypeAssertion
	}
	err = addressRepository.Save(ctx, enterprise.Address())
	if err != nil {
		return err
	}
	insertCommand := string(
		`INSERT INTO enterprise(
		enterprise_id,
		enterprise_name,
		owner_user_id,
		logo_img,
		banner_img,
		website,
		email,
		phone,
		industry_id,
		created_at,
		updated_at,
		foundation_date
		) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11,$12)`,
	)
	var (
		id             uuid.UUID = enterprise.Id()
		name           string    = enterprise.Name()
		ownerId        uuid.UUID = enterprise.OwnerId()
		logoImg        string    = enterprise.LogoIMG()
		bannerImg      string    = enterprise.BannerIMG()
		website        string    = enterprise.Website()
		email          string    = enterprise.Email()
		phone          string    = enterprise.Phone()
		industryId     int       = enterprise.Industry().ID()
		createdAt      string    = enterprise.CreatedAt().StringRepresentation()
		updatedAt      string    = enterprise.UpdatedAt().StringRepresentation()
		foundationDate string    = enterprise.FoundationDate().StringRepresentation()
	)
	_, err = conn.Exec(
		ctx,
		insertCommand,
		&id,
		&name,
		&ownerId,
		&logoImg,
		&bannerImg,
		&website,
		&email,
		&phone,
		&industryId,
		&createdAt,
		&updatedAt,
		&foundationDate,
	)
	if err != nil {
		return err
	}
	return nil
}

func (er EnterpriseRepository) Find(ctx context.Context, src StatementSource) (*enterprise.Enterprise, error) {
	conn, err := er.schema.Acquire(ctx)
	defer conn.Release()
	if err != nil {
		return nil, err
	}
	row := conn.QueryRow(ctx, src.Query(), src.Args()...)
	return er.load(ctx, row)
}

func (er EnterpriseRepository) Get(
	ctx context.Context,
	enterpriseID uuid.UUID,
) (*enterprise.Enterprise, error) {
	conn, err := er.schema.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	selectQuery := string(
		`SELECT
		enterprise_id,
		enterprise_name,
		owner_user_id,
		logo_img,
		banner_img,
		website,
		email,
		phone,
		industry_id,
		created_at,
		updated_at,
		foundation_date
		FROM enterprise
		WHERE enterprise_id = $1`,
	)
	row := conn.QueryRow(ctx, selectQuery, enterpriseID)
	enterprise, err := er.load(ctx, row)
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	return enterprise, nil
}

func (er EnterpriseRepository) FindAll(
	ctx context.Context,
	source StatementSource,
) (mapset.Set[enterprise.Enterprise], error) {
	conn, err := er.schema.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	rows, err := conn.Query(ctx, source.Query(), source.Args()...)
	if err != nil {
		return nil, err
	}
	enterprises, err := er.loadAll(ctx, rows)
	if err != nil {
		return nil, err
	}
	defer func() {
		rows.Close()
		conn.Release()
	}()
	return enterprises, nil
}

func (er EnterpriseRepository) loadAll(
	ctx context.Context,
	rows pgx.Rows,
) (mapset.Set[enterprise.Enterprise], error) {
	enterprises := mapset.NewSet[enterprise.Enterprise]()
	for rows.Next() {
		enterprise, err := er.load(ctx, rows)
		if err != nil {

			return nil, err
		}
		enterprises.Add(*enterprise)
	}
	return enterprises, nil
}
func (er EnterpriseRepository) load(
	ctx context.Context,
	row pgx.Row,
) (*enterprise.Enterprise, error) {
	mapper := MapperRegistryInstance().Get("Address")
	if mapper == nil {
		return nil, ErrMapperNotRegistered
	}
	addressRepository, ok := mapper.(AddressRepository)
	if !ok {
		return nil, ErrWrongTypeAssertion
	}
	mapper = MapperRegistryInstance().Get("Employee")
	if mapper == nil {
		return nil, ErrMapperNotRegistered
	}
	employeeRepository, ok := mapper.(EmployeeRepository)
	if !ok {
		return nil, ErrWrongTypeAssertion
	}
	mapper = MapperRegistryInstance().Get("Industry")
	if mapper == nil {
		return nil, ErrMapperNotRegistered
	}
	industryRepository, ok := mapper.(IndustryRepository)
	if !ok {
		return nil, ErrWrongTypeAssertion
	}
	var (
		id          uuid.UUID
		name        string
		ownerUserId uuid.UUID
		logoImg     string
		bannerImg   string
		website     string
		email       string
		phone       string
		industryId  int
		createdAt   string
		updatedAt   string
		foundation  string
	)
	err := row.Scan(
		&id,
		&name,
		&ownerUserId,
		&logoImg,
		&bannerImg,
		&website,
		&email,
		&phone,
		&industryId,
		&createdAt,
		&updatedAt,
		&foundation,
	)
	if err != nil {

		return nil, err
	}
	commonCreatedAt, err := common.NewDateFromString(createdAt)
	if err != nil {
		return nil, err
	}
	commonUpdatedAt, err := common.NewDateFromString(updatedAt)
	if err != nil {
		return nil, err
	}
	commonFoundation, err := common.NewDateFromString(foundation)
	if err != nil {
		return nil, err
	}
	employees, err := employeeRepository.FindAll(
		ctx,
		employeefindall.ByEnterpriseId(id),
	)
	if err != nil {

		return nil, err
	}
	address, err := addressRepository.Get(ctx, id)
	if err != nil {

		return nil, err
	}
	industry, err := industryRepository.Get(ctx, industryId)
	if err != nil {
		return nil, err
	}
	return enterprise.NewEnterprise(
		id,
		ownerUserId,
		name,
		logoImg,
		bannerImg,
		website,
		email,
		phone,
		address,
		industry,
		commonCreatedAt,
		commonUpdatedAt,
		commonFoundation,
		employees,
	)
}
