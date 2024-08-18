package queries

import (
	"context"
	"go-complaint/dto"
	"go-complaint/infrastructure/persistence/repositories"

	"github.com/google/uuid"
)

type EmployeeByEmployeeIdQuery struct {
	EmployeeId string `json:"employeeId"`
}

func NewEmployeeByEmployeeIdQuery(employeeId string) *EmployeeByEmployeeIdQuery {
	return &EmployeeByEmployeeIdQuery{
		EmployeeId: employeeId,
	}
}

func (q EmployeeByEmployeeIdQuery) Execute(ctx context.Context) (*dto.Employee, error) {
	employeeId, err := uuid.Parse(q.EmployeeId)
	if err != nil {
		return nil, err
	}
	reg := repositories.MapperRegistryInstance()
	r, ok := reg.Get("Employee").(repositories.EmployeeRepository)
	if !ok {
		return nil, ErrWrongTypeAssertion
	}
	dbE, err := r.Get(ctx, employeeId)
	if err != nil {
		return nil, err
	}
	return dto.NewEmployee(dbE), nil
}
