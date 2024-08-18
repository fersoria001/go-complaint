package queries

import (
	"context"
	"go-complaint/dto"
	"go-complaint/infrastructure/persistence/finders/find_hiring_proccess"
	"go-complaint/infrastructure/persistence/repositories"

	"github.com/google/uuid"
)

type HiringProccessByEmployeeIdQuery struct {
	EmployeeId string `json:"employeeId"`
}

func NewHiringProccessByEmployeeIdQuery(employeeId string) *HiringProccessByEmployeeIdQuery {
	return &HiringProccessByEmployeeIdQuery{
		EmployeeId: employeeId,
	}
}

func (q HiringProccessByEmployeeIdQuery) Execute(ctx context.Context) (*dto.HiringProccess, error) {
	employeeId, err := uuid.Parse(q.EmployeeId)
	if err != nil {
		return nil, err
	}
	reg := repositories.MapperRegistryInstance()
	r, ok := reg.Get("HiringProccess").(repositories.HiringProccessRepository)
	if !ok {
		return nil, ErrWrongTypeAssertion
	}
	employeeRepository, ok := reg.Get("Employee").(repositories.EmployeeRepository)
	if !ok {
		return nil, ErrWrongTypeAssertion
	}
	employee, err := employeeRepository.Get(ctx, employeeId)
	if err != nil {
		return nil, err
	}
	hiringProccess, err := r.Find(ctx, find_hiring_proccess.ByUserIdAndEnterpriseId(employee.GetUser().Id(), employee.EnterpriseId()))
	if err != nil {
		return nil, err
	}
	return dto.NewHiringProccess(*hiringProccess), nil
}
