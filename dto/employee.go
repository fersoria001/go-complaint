package dto

import (
	"go-complaint/domain/model/employee"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/google/uuid"
)

// EmployeeWaitingForApproval is a struct type
type Employee struct {
	ID                  uuid.UUID `json:"id"`
	EnterpriseID        string    `json:"enterprise_id"`
	UserID              string    `json:"user_id"`
	ProfileIMG          string    `json:"profile_img"`
	FirstName           string    `json:"first_name"`
	LastName            string    `json:"last_name"`
	Age                 int       `json:"age"`
	Email               string    `json:"email"`
	Phone               string    `json:"phone"`
	HiringDate          string    `json:"hiring_date"`
	ApprovedHiring      bool      `json:"approved_hiring"`
	ApprovedHiringAt    string    `json:"approved_hiring_at"`
	Position            string    `json:"position"`
	ComplaintsSolved    int       `json:"complaint_solved"`
	ComplaintsSolvedIDs []string  `json:"complaints_solved"`
}

func NewEmployee(domainObj employee.Employee) Employee {
	return Employee{
		ID:                  domainObj.ID(),
		ProfileIMG:          domainObj.ProfileIMG(),
		FirstName:           domainObj.FirstName(),
		LastName:            domainObj.LastName(),
		Age:                 domainObj.Age(),
		Email:               domainObj.Email(),
		Phone:               domainObj.Phone(),
		HiringDate:          domainObj.HiringDate().StringRepresentation(),
		ApprovedHiring:      domainObj.ApprovedHiring(),
		ApprovedHiringAt:    domainObj.ApprovedHiringAt().StringRepresentation(),
		Position:            domainObj.Position().String(),
		ComplaintsSolved:    0,
		ComplaintsSolvedIDs: []string{},
	}
}

func NewEmployeeList(domainObjs mapset.Set[employee.Employee]) []Employee {
	employees := []Employee{}
	for domainObj := range domainObjs.Iter() {
		employee := NewEmployee(domainObj)
		employees = append(employees, employee)
	}
	return employees
}
