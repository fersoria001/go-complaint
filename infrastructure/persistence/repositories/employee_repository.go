package repositories

import (
	"context"
	"go-complaint/domain/model/common"
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
	id uuid.UUID,
) error {
	conn, err := er.schema.Acquire(ctx)
	if err != nil {
		return nil
	}
	deleteCommand := string(`
	DELETE FROM employee
	WHERE enterprise_id = $1
	`)

	_, err = conn.Exec(
		ctx,
		deleteCommand,
		&id,
	)
	if err != nil {

		return err
	}

	defer conn.Release()
	return nil
}

func (er EmployeeRepository) UpdateAll(
	ctx context.Context,
	employees mapset.Set[enterprise.Employee],
) error {
	conn, err := er.schema.Acquire(ctx)
	if err != nil {
		return nil
	}
	updateCommand := string(
		`
		UPDATE employee
		SET
		approved_hiring = $2,
		approved_hiring_at = $3,
		job_position = $4
		WHERE enterprise.employee_id = $1
		`,
	)
	tx, err := conn.Begin(ctx)
	if err != nil {
		return err
	}
	for employee := range employees.Iter() {
		var (
			employeeId       uuid.UUID = employee.ID()
			approvedHiring   bool      = employee.ApprovedHiring()
			approvedHiringAt string    = employee.ApprovedHiringAt().StringRepresentation()
			jobPosition      string    = employee.Position().String()
		)
		_, err := tx.Exec(
			ctx,
			updateCommand,
			&employeeId,
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

func (er EmployeeRepository) Remove(ctx context.Context, id uuid.UUID) error {
	conn, err := er.schema.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()
	deleteCommand := string(`
	DELETE FROM employee WHERE employee_id = $1`)
	_, err = conn.Exec(ctx, deleteCommand, &id)
	if err != nil {
		return err
	}
	return nil
}
func (er EmployeeRepository) Find(
	ctx context.Context,
	src StatementSource,
) (
	*enterprise.Employee, error) {
	conn, err := er.schema.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	row := conn.QueryRow(ctx, src.Query(), src.Args()...)
	employee, err := er.load(ctx, row)
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	return employee, nil
}
func (er EmployeeRepository) Get(
	ctx context.Context,
	employeeID uuid.UUID,
) (
	*enterprise.Employee, error) {
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
	employee *enterprise.Employee,
) error {
	conn, err := er.schema.Acquire(ctx)
	if err != nil {
		return err
	}
	var (
		employeeId       uuid.UUID = employee.ID()
		enterpriseId     uuid.UUID = employee.EnterpriseId()
		userId           uuid.UUID = employee.User.Id()
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
		employeeId,
		enterpriseId,
		userId,
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
	employees mapset.Set[enterprise.Employee],
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
	for emp := range employees.Iter() {
		var (
			employeeId       uuid.UUID = emp.ID()
			enterpriseId     uuid.UUID = emp.EnterpriseId()
			userId           uuid.UUID = emp.User.Id()
			hiringDate       string    = emp.HiringDate().StringRepresentation()
			approvedHiring   bool      = emp.ApprovedHiring()
			approvedHiringAt string    = emp.ApprovedHiringAt().StringRepresentation()
			jobPosition      string    = emp.Position().String()
		)
		_, err = tx.Exec(
			ctx,
			insertCommand,
			employeeId,
			enterpriseId,
			userId,
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
) (mapset.Set[*enterprise.Employee], error) {
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
) (mapset.Set[*enterprise.Employee], error) {
	employees := mapset.NewSet[*enterprise.Employee]()
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
	*enterprise.Employee, error) {
	var (
		employeeId       uuid.UUID
		enterpriseId     uuid.UUID
		userId           uuid.UUID
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
		&enterpriseId,
		&userId,
		&hiringDate,
		&approvedHiring,
		&approvedHiringAt,
		&jobPosition,
	)
	if err != nil {
		return nil, err
	}
	parsedJobPosition := enterprise.ParsePosition(jobPosition)
	if parsedJobPosition < 0 {
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
	user, err := userRepository.Get(ctx, userId)
	if err != nil {
		return nil, err
	}
	return enterprise.NewEmployee(
		employeeId,
		enterpriseId,
		user,
		parsedJobPosition,
		commonHiringDate,
		approvedHiring,
		commonApprovedHiringAt,
	)
}
