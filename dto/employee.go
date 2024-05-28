package dto

import "go-complaint/domain/model/enterprise"

// EmployeeWaitingForApproval is a struct type
type Employee struct {
	ID               string `json:"id"`
	ProfileIMG       string `json:"profile_img"`
	FirstName        string `json:"first_name"`
	LastName         string `json:"last_name"`
	Age              int    `json:"age"`
	Email            string `json:"email"`
	Phone            string `json:"phone"`
	HiringDate       string `json:"hiring_date"`
	ApprovedHiring   bool   `json:"approved_hiring"`
	ApprovedHiringAt string `json:"approved_hiring_at"`
	Position         string `json:"position"`
}

func NewEmployee(domainObj enterprise.Employee) *Employee {
	return &Employee{
		ID:               domainObj.ID(),
		ProfileIMG:       domainObj.ProfileIMG(),
		FirstName:        domainObj.FirstName(),
		LastName:         domainObj.LastName(),
		Age:              domainObj.Age(),
		Email:            domainObj.Email(),
		Phone:            domainObj.Phone(),
		HiringDate:       domainObj.HiringDate().StringRepresentation(),
		ApprovedHiring:   domainObj.ApprovedHiring(),
		ApprovedHiringAt: domainObj.ApprovedHiringAt().StringRepresentation(),
		Position:         domainObj.Position().String(),
	}
}
