package models

import (
	"go-complaint/domain/model/enterprise"

	mapset "github.com/deckarep/golang-set/v2"
)

type Employee struct {
	ID               string
	ProfileIMG       string
	FirstName        string
	LastName         string
	Age              int
	Email            string
	Phone            string
	HiringDate       string
	ApprovedHiring   bool
	ApprovedHiringAt string
	Position         string
}

func NewEmployee(emp *enterprise.Employee) *Employee {
	return &Employee{
		ID:               emp.ID(),
		ProfileIMG:       emp.ProfileIMG(),
		FirstName:        emp.FirstName(),
		LastName:         emp.LastName(),
		Age:              emp.Age(),
		Email:            emp.Email(),
		Phone:            emp.Phone(),
		HiringDate:       emp.HiringDate().StringRepresentation(),
		ApprovedHiring:   emp.ApprovedHiring(),
		ApprovedHiringAt: emp.ApprovedHiringAt().StringRepresentation(),
		Position:         emp.Position().String(),
	}
}

func NewEmployees(domains map[string]*enterprise.Employee) mapset.Set[*Employee] {
	var employees mapset.Set[*Employee] = mapset.NewSet[*Employee]()
	for _, emp := range domains {
		employees.Add(NewEmployee(emp))
	}
	return employees
}

func (e *Employee) Table() string {
	return "employees"
}

func (e *Employee) Columns() Columns {
	return Columns{
		"id",
		"profile_img",
		"first_name",
		"last_name",
		"age",
		"email",
		"phone",
		"hiring_date",
		"approved_hiring",
		"approved_hiring_at",
		"position",
	}
}

func (e *Employee) Values() []interface{} {
	return []interface{}{
		&e.ID,
		&e.ProfileIMG,
		&e.FirstName,
		&e.LastName,
		&e.Age,
		&e.Email,
		&e.Phone,
		&e.HiringDate,
		&e.ApprovedHiring,
		&e.ApprovedHiringAt,
		&e.Position,
	}
}
