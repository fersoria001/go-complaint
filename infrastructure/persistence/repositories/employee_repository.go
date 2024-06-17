package repositories

import (
	"context"
	"go-complaint/domain/model/common"
	"go-complaint/domain/model/employee"
	"go-complaint/domain/model/enterprise"
	"go-complaint/infrastructure/persistence/datasource"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type EmployeeRepository struct {
	schema datasource.Schema
}

func NewEmployeeRepository(enterpriseSchema datasource.Schema) EmployeeRepository {
	return EmployeeRepository{
		schema: enterpriseSchema,
	}
}

func (er EmployeeRepository) DeleteAll(
	ctx context.Context,
	employees mapset.Set[employee.Employee],
) error {
	conn, err := er.schema.Acquire(ctx)
	if err != nil {
		return nil
	}
	deleteCommand := string(`
	DELETE FROM employee
	WHERE employee.employee_id = $1
	`)
	tx, err := conn.Begin(ctx)
	if err != nil {
		return err
	}
	for employee := range employees.Iter() {
		var employeeID = employee.ID()
		_, err = tx.Exec(
			ctx,
			deleteCommand,
			employeeID,
		)
		if err != nil {
			tx.Rollback(ctx)
			return err
		}
	}
	err = tx.Commit(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()
	return nil
}

func (er EmployeeRepository) UpdateAll(
	ctx context.Context,
	employees mapset.Set[employee.Employee],
) error {
	conn, err := er.schema.Acquire(ctx)
	if err != nil {
		return nil
	}
	updateCommand := string(
		`
		UPDATE employee
		SET
		employee_id = $1,
		enterprise_id = $2,
		user_id = $3,
		hiring_date = $4,
		approved_hiring = $5,
		approved_hiring_at = $6,
		job_position = $7
		WHERE employee.employee_id = $1
		`,
	)
	tx, err := conn.Begin(ctx)
	if err != nil {
		return err
	}
	for employee := range employees.Iter() {
		var (
			employeeID       uuid.UUID = employee.ID()
			enterpriseID     string    = employee.EnterpriseID()
			userID           string    = employee.Email()
			hiringDate       string    = employee.HiringDate().StringRepresentation()
			approvedHiring   bool      = employee.ApprovedHiring()
			approvedHiringAt string    = employee.ApprovedHiringAt().StringRepresentation()
			jobPosition      string    = employee.Position().String()
		)
		_, err := tx.Exec(
			ctx,
			updateCommand,
			&employeeID,
			&enterpriseID,
			&userID,
			&hiringDate,
			&approvedHiring,
			&approvedHiringAt,
			&jobPosition,
		)
		if err != nil {
			tx.Rollback(ctx)
			return err
		}
	}
	err = tx.Commit(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()
	return nil
}

func (er EmployeeRepository) Get(
	ctx context.Context,
	employeeID uuid.UUID,
) (
	*employee.Employee, error) {
	conn, err := er.schema.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	selectQuery := string(
		`SELECT 
		employee_id,
		enterprise_id,
		user_id,
		hiring_date,
		approved_hiring,
		approved_hiring_at,
		job_position
		FROM employee
		WHERE employee_id = $1`,
	)
	row := conn.QueryRow(ctx, selectQuery, employeeID)
	employee, err := er.load(ctx, row)
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	return employee, nil
}
func (er EmployeeRepository) Save(
	ctx context.Context,
	employee *employee.Employee,
) error {
	conn, err := er.schema.Acquire(ctx)
	if err != nil {
		return err
	}
	var (
		employeeID       uuid.UUID = employee.ID()
		enterpriseID     string    = employee.EnterpriseID()
		userID           string    = employee.Email()
		hiringDate       string    = employee.HiringDate().StringRepresentation()
		approvedHiring   bool      = employee.ApprovedHiring()
		approvedHiringAt string    = employee.ApprovedHiringAt().StringRepresentation()
		jobPosition      string    = employee.Position().String()
	)
	insertCommand := string(
		`INSERT INTO employee (
		employee_id,
		enterprise_id,
		user_id,
		hiring_date,
		approved_hiring,
		approved_hiring_at,
		job_position
		) VALUES ($1, $2, $3, $4, $5, $6, $7)`)
	_, err = conn.Exec(
		ctx,
		insertCommand,
		employeeID,
		enterpriseID,
		userID,
		hiringDate,
		approvedHiring,
		approvedHiringAt,
		jobPosition,
	)
	if err != nil {
		return err
	}
	defer conn.Release()
	return nil
}

func (er EmployeeRepository) SaveAll(
	ctx context.Context,
	employees mapset.Set[employee.Employee],
) error {
	conn, err := er.schema.Acquire(ctx)
	if err != nil {
		return err
	}

	insertCommand := string(
		`INSERT INTO employee (
		employee_id,
		enterprise_id,
		user_id,
		hiring_date,
		approved_hiring,
		approved_hiring_at,
		job_position
		) VALUES ($1, $2, $3, $4, $5, $6, $7)`)
	tx, err := conn.Begin(ctx)
	if err != nil {
		return err
	}
	for employee := range employees.Iter() {
		var (
			employeeID       uuid.UUID = employee.ID()
			enterpriseID     string    = employee.EnterpriseID()
			userID           string    = employee.Email()
			hiringDate       string    = employee.HiringDate().StringRepresentation()
			approvedHiring   bool      = employee.ApprovedHiring()
			approvedHiringAt string    = employee.ApprovedHiringAt().StringRepresentation()
			jobPosition      string    = employee.Position().String()
		)
		_, err = tx.Exec(
			ctx,
			insertCommand,
			employeeID,
			enterpriseID,
			userID,
			hiringDate,
			approvedHiring,
			approvedHiringAt,
			jobPosition,
		)
		if err != nil {
			tx.Rollback(ctx)
			return err
		}
	}
	err = tx.Commit(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()
	return nil
}

func (er EmployeeRepository) FindAll(
	ctx context.Context,
	statementSource StatementSource,
) (mapset.Set[*employee.Employee], error) {
	conn, err := er.schema.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	rows, err := conn.Query(
		ctx,
		statementSource.Query(),
		statementSource.Args()...,
	)
	if err != nil {
		return nil, err
	}
	employees, err := er.loadAll(ctx, rows)
	if err != nil {
		return nil, err
	}
	defer func() {
		rows.Close()
		conn.Release()
	}()
	return employees, nil
}

func (er EmployeeRepository) loadAll(
	ctx context.Context,
	rows pgx.Rows,
) (mapset.Set[*employee.Employee], error) {
	employees := mapset.NewSet[*employee.Employee]()
	for rows.Next() {
		employee, err := er.load(ctx, rows)
		if err != nil {
			return nil, err
		}
		employees.Add(employee)
	}
	return employees, nil
}

func (er EmployeeRepository) load(ctx context.Context, row pgx.Row) (
	*employee.Employee, error) {
	var (
		employeeId       uuid.UUID
		enterpriseID     string
		userID           string
		hiringDate       string
		approvedHiring   bool
		approvedHiringAt string
		jobPosition      string
	)
	mapper := MapperRegistryInstance().Get("User")
	if mapper == nil {
		return nil, ErrMapperNotRegistered
	}
	userRepository, ok := mapper.(UserRepository)
	if !ok {
		return nil, ErrWrongTypeAssertion
	}

	err := row.Scan(
		&employeeId,
		&enterpriseID,
		&userID,
		&hiringDate,
		&approvedHiring,
		&approvedHiringAt,
		&jobPosition,
	)
	if err != nil {
		return nil, err
	}
	parsedJobPosition := enterprise.ParsePosition(jobPosition)
	if parsedJobPosition == enterprise.NOT_EXISTS {
		return nil, enterprise.ErrPositionNotExists
	}
	commonHiringDate, err := common.NewDateFromString(hiringDate)
	if err != nil {
		return nil, err
	}
	commonApprovedHiringAt, err := common.NewDateFromString(approvedHiringAt)
	if err != nil {
		return nil, err
	}
	user, err := userRepository.Get(ctx, userID)
	if err != nil {
		return nil, err
	}
	return employee.NewEmployee(
		employeeId,
		enterpriseID,
		user,
		parsedJobPosition,
		commonHiringDate,
		approvedHiring,
		commonApprovedHiringAt,
	)
}
