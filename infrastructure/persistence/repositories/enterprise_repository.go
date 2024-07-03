package repositories

import (
	"context"
	"go-complaint/domain/model/common"
	"go-complaint/domain/model/employee"
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
	castedEmployees := mapset.NewSet[employee.Employee]()
	for emp := range updatedEnterprise.Employees().Iter() {
		cast, ok := emp.(*employee.Employee)
		if !ok {
			return ErrWrongTypeAssertion
		}
		castedEmployees.Add(*cast)
	}
	err = employeeRepository.DeleteAll(ctx, updatedEnterprise.Name())
	if err != nil {
		return err
	}
	err = employeeRepository.SaveAll(ctx, castedEmployees)
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
		name           string = updatedEnterprise.Name()
		owner          string = updatedEnterprise.Owner()
		logoIMG        string = updatedEnterprise.LogoIMG()
		bannerIMG      string = updatedEnterprise.BannerIMG()
		website        string = updatedEnterprise.Website()
		email          string = updatedEnterprise.Email()
		phone          string = updatedEnterprise.Phone()
		industryID     int    = updatedEnterprise.Industry().ID()
		createdAt      string = updatedEnterprise.CreatedAt().StringRepresentation()
		updatedAt      string = updatedEnterprise.UpdatedAt().StringRepresentation()
		foundationDate string = updatedEnterprise.FoundationDate().StringRepresentation()
	)
	_, err = conn.Exec(
		ctx,
		updateCommand,
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

func (er EnterpriseRepository) Save(
	ctx context.Context,
	enterprise *enterprise.Enterprise,
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
	err = addressRepository.Save(ctx, enterprise.Address())
	if err != nil {
		return err
	}
	insertCommand := string(
		`INSERT INTO enterprise(
		enterprise_id,
		owner_user_id,
		logo_img,
		banner_img,
		website,
		email,
		phone,
		address_id,
		industry_id,
		created_at,
		updated_at,
		foundation_date
		) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11,$12)`,
	)
	var (
		name           string    = enterprise.Name()
		owner          string    = enterprise.Owner()
		logoIMG        string    = enterprise.LogoIMG()
		bannerIMG      string    = enterprise.BannerIMG()
		website        string    = enterprise.Website()
		email          string    = enterprise.Email()
		phone          string    = enterprise.Phone()
		addressID      uuid.UUID = enterprise.Address().ID()
		industryID     int       = enterprise.Industry().ID()
		createdAt      string    = enterprise.CreatedAt().StringRepresentation()
		updatedAt      string    = enterprise.UpdatedAt().StringRepresentation()
		foundationDate string    = enterprise.FoundationDate().StringRepresentation()
	)
	_, err = conn.Exec(
		ctx,
		insertCommand,
		&name,
		&owner,
		&logoIMG,
		&bannerIMG,
		&website,
		&email,
		&phone,
		&addressID,
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

func (er EnterpriseRepository) Get(
	ctx context.Context,
	enterpriseID string,
) (*enterprise.Enterprise, error) {
	conn, err := er.schema.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	selectQuery := string(
		`SELECT
		enterprise_id,
		owner_user_id,
		logo_img,
		banner_img,
		website,
		email,
		phone,
		address_id,
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
		enterpriseId string
		ownerUserID  string
		logoImg      string
		bannerImg    string
		website      string
		email        string
		phone        string
		addressID    uuid.UUID
		industryID   int
		createdAt    string
		updatedAt    string
		foundation   string
	)
	err := row.Scan(
		&enterpriseId,
		&ownerUserID,
		&logoImg,
		&bannerImg,
		&website,
		&email,
		&phone,
		&addressID,
		&industryID,
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
		employeefindall.NewByEnterpriseID(enterpriseId),
	)
	if err != nil {

		return nil, err
	}
	castToEmployeeInterface := mapset.NewSet[enterprise.Employee]()
	for emp := range employees.Iter() {
		castToEmployeeInterface.Add(emp)
	}
	address, err := addressRepository.Get(ctx, addressID)
	if err != nil {

		return nil, err
	}
	industry, err := industryRepository.Get(ctx, industryID)
	if err != nil {
		return nil, err
	}
	return enterprise.NewEnterprise(
		ownerUserID,
		enterpriseId,
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
		castToEmployeeInterface,
	)
}
