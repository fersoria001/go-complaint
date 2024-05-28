package repositories

import (
	"context"
	"fmt"
	"go-complaint/domain/model/common"
	"go-complaint/domain/model/enterprise"
	"go-complaint/infrastructure/persistence/datasource"
	"go-complaint/infrastructure/persistence/models"

	mapset "github.com/deckarep/golang-set/v2"
)

type EmployeeRepository struct {
	schema *datasource.Schema
}

func NewEmployeeRepository(enterpriseSchema *datasource.Schema) *EmployeeRepository {
	return &EmployeeRepository{
		schema: enterpriseSchema,
	}
}

func (employeeRepository *EmployeeRepository) Save(ctx context.Context, employee *enterprise.Employee) error {
	conn, err := employeeRepository.schema.Pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()
	employeeModel := models.NewEmployee(employee)
	insertCommand := fmt.Sprintf(`INSERT INTO %s (%s) VALUES (%s)`, employeeModel.Table(),
		models.StringColumns(employeeModel.Columns()), models.Args(employeeModel.Columns()))
	_, err = conn.Exec(ctx, insertCommand, employeeModel.Values()...)
	if err != nil {
		return err
	}
	return nil
}

func (employeeRepository *EmployeeRepository) Get(ctx context.Context, id string) (*enterprise.Employee, error) {
	conn, err := employeeRepository.schema.Pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	employeeModel := &models.Employee{}
	selectQuery := fmt.Sprintf(`SELECT * FROM %s WHERE id = $1`, employeeModel.Table())
	row := conn.QueryRow(ctx, selectQuery, id)
	err = row.Scan(employeeModel.Values()...)
	if err != nil {
		return nil, err
	}
	hiringDate, err := common.NewDateFromString(employeeModel.HiringDate)
	if err != nil {
		return nil, err
	}
	approvedHiringAt, err := common.NewDateFromString(employeeModel.ApprovedHiringAt)
	if err != nil {
		return nil, err
	}
	parsedPosition, err := enterprise.ParsePosition(employeeModel.Position)
	if err != nil {
		return nil, err
	}
	employee, err := enterprise.NewEmployee(
		employeeModel.ID,
		employeeModel.ProfileIMG,
		employeeModel.FirstName,
		employeeModel.LastName,
		employeeModel.Age,
		employeeModel.Email,
		employeeModel.Phone,
		parsedPosition,
		hiringDate,
		employeeModel.ApprovedHiring,
		approvedHiringAt,
	)
	if err != nil {
		return nil, err
	}
	return employee, nil
}

func (employeeRepository *EmployeeRepository) Update(ctx context.Context, employee *enterprise.Employee) error {
	conn, err := employeeRepository.schema.Pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()
	employeeModel := models.NewEmployee(employee)
	updateCommand := fmt.Sprintf(`UPDATE %s SET %s WHERE id = $1`, employeeModel.Table(),
		models.StringKeyArgs(employeeModel.Columns()))
	_, err = conn.Exec(ctx, updateCommand, employeeModel.Values()...)
	if err != nil {
		return err
	}
	return nil
}

func (employeeRepository *EmployeeRepository) Remove(ctx context.Context, id string) error {
	conn, err := employeeRepository.schema.Pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()
	employeeModel := &models.Employee{}
	deleteCommand := fmt.Sprintf(`DELETE FROM %s WHERE id = $1`, employeeModel.Table())
	_, err = conn.Exec(ctx, deleteCommand, id)
	if err != nil {
		return err
	}
	return nil
}

func (employeeRepository *EmployeeRepository) GetAll(ctx context.Context) (mapset.Set[*enterprise.Employee], error) {
	conn, err := employeeRepository.schema.Pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	employeeModel := &models.Employee{}
	selectQuery := fmt.Sprintf(`SELECT * FROM %s`, employeeModel.Table())
	rows, err := conn.Query(ctx, selectQuery)
	if err != nil {
		return nil, err
	}
	var employees = mapset.NewSet[*enterprise.Employee]()
	for rows.Next() {
		err = rows.Scan(employeeModel.Values()...)
		if err != nil {
			return nil, err
		}
		hiringDate, err := common.NewDateFromString(employeeModel.HiringDate)
		if err != nil {
			return nil, err
		}
		approvedHiringAt, err := common.NewDateFromString(employeeModel.ApprovedHiringAt)
		if err != nil {
			return nil, err
		}
		parsedPosition, err := enterprise.ParsePosition(employeeModel.Position)
		if err != nil {
			return nil, err
		}
		employee, err := enterprise.NewEmployee(
			employeeModel.ID,
			employeeModel.ProfileIMG,
			employeeModel.FirstName,
			employeeModel.LastName,
			employeeModel.Age,
			employeeModel.Email,
			employeeModel.Phone,
			parsedPosition,
			hiringDate,
			employeeModel.ApprovedHiring,
			approvedHiringAt,
		)
		if err != nil {
			return nil, err
		}
		employees.Add(employee)
	}
	return employees, nil
}

func (employeeRepository *EmployeeRepository) FindByEmail(ctx context.Context, email string) (mapset.Set[*enterprise.Employee], error) {
	conn, err := employeeRepository.schema.Pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	employeeModel := &models.Employee{}
	selectQuery := fmt.Sprintf(`SELECT * FROM %s WHERE email = $1`, employeeModel.Table())
	rows, err := conn.Query(ctx, selectQuery, email)
	if err != nil {
		return nil, err
	}
	var employees = mapset.NewSet[*enterprise.Employee]()
	for rows.Next() {
		err = rows.Scan(employeeModel.Values()...)
		if err != nil {
			return nil, err
		}
		hiringDate, err := common.NewDateFromString(employeeModel.HiringDate)
		if err != nil {
			return nil, err
		}
		approvedHiringAt, err := common.NewDateFromString(employeeModel.ApprovedHiringAt)
		if err != nil {
			return nil, err
		}
		parsedPosition, err := enterprise.ParsePosition(employeeModel.Position)
		if err != nil {
			return nil, err
		}
		employee, err := enterprise.NewEmployee(
			employeeModel.ID,
			employeeModel.ProfileIMG,
			employeeModel.FirstName,
			employeeModel.LastName,
			employeeModel.Age,
			employeeModel.Email,
			employeeModel.Phone,
			parsedPosition,
			hiringDate,
			employeeModel.ApprovedHiring,
			approvedHiringAt,
		)
		if err != nil {
			return nil, err
		}
		employees.Add(employee)
	}
	return employees, nil
}

func (employeeRepository *EmployeeRepository) FindByEnterpriseID(
	ctx context.Context, enterpriseID string) (mapset.Set[*enterprise.Employee], error) {
	conn, err := employeeRepository.schema.Pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	employeeModel := &models.Employee{}
	selectQuery := fmt.Sprintf(`SELECT %s FROM %s WHERE SUBSTRING(id FROM 1 FOR POSITION('-' IN id) - 1) = $1 `,
		models.StringColumns(employeeModel.Columns()),
		employeeModel.Table())
	rows, err := conn.Query(ctx, selectQuery, enterpriseID)
	if err != nil {
		return nil, err
	}
	var employees = mapset.NewSet[*enterprise.Employee]()
	for rows.Next() {
		err = rows.Scan(employeeModel.Values()...)
		if err != nil {
			return nil, err
		}
		hiringDate, err := common.NewDateFromString(employeeModel.HiringDate)
		if err != nil {
			return nil, err
		}
		approvedHiringAt, err := common.NewDateFromString(employeeModel.ApprovedHiringAt)
		if err != nil {
			return nil, err
		}
		parsedPosition, err := enterprise.ParsePosition(employeeModel.Position)
		if err != nil {
			return nil, err
		}
		employee, err := enterprise.NewEmployee(
			employeeModel.ID,
			employeeModel.ProfileIMG,
			employeeModel.FirstName,
			employeeModel.LastName,
			employeeModel.Age,
			employeeModel.Email,
			employeeModel.Phone,
			parsedPosition,
			hiringDate,
			employeeModel.ApprovedHiring,
			approvedHiringAt,
		)
		if err != nil {
			return nil, err
		}
		employees.Add(employee)
	}
	return employees, nil
}
