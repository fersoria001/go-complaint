package enterprise

import (
	"encoding/json"
	"go-complaint/domain/model/common"
	"time"
)

// Package enterprise
// <<Domain event>> implements domain.DomainEvent
type EmployeeFired struct {
	id               string
	firstName        string
	lastName         string
	age              int
	email            string
	phone            string
	hiringDate       time.Time
	approvedHiring   bool
	approvedHiringAt time.Time
	position         Position
	occurredOn       time.Time
}

func NewEmployeeFired(employee *Employee) *EmployeeFired {
	return &EmployeeFired{
		id:               employee.id,
		firstName:        employee.firstName,
		lastName:         employee.lastName,
		age:              employee.age,
		email:            employee.email,
		phone:            employee.phone,
		hiringDate:       employee.hiringDate.Date(),
		approvedHiring:   employee.approvedHiring,
		approvedHiringAt: employee.approvedHiringAt.Date(),
		position:         employee.position,
		occurredOn:       time.Now(),
	}
}

func (ef *EmployeeFired) OccurredOn() time.Time {
	return ef.occurredOn
}

func (ef *EmployeeFired) MarshalJSON() ([]byte, error) {
	j, err := json.Marshal(struct {
		ID               string `json:"id"`
		FirstName        string `json:"first_name"`
		LastName         string `json:"last_name"`
		Age              int    `json:"age"`
		Email            string `json:"email"`
		Phone            string `json:"phone"`
		HiringDate       string `json:"hiring_date"`
		ApprovedHiring   bool   `json:"approved_hiring"`
		ApprovedHiringAt string `json:"approved_hiring_at"`
		Position         string `json:"position"`
		OccurredOn       string `json:"occurred_on"`
	}{
		ID:               ef.id,
		FirstName:        ef.firstName,
		LastName:         ef.lastName,
		Age:              ef.age,
		Email:            ef.email,
		Phone:            ef.phone,
		HiringDate:       common.StringDate(ef.hiringDate),
		ApprovedHiring:   ef.approvedHiring,
		ApprovedHiringAt: common.StringDate(ef.approvedHiringAt),
		Position:         ef.position.String(),
		OccurredOn:       common.StringDate(ef.occurredOn),
	})
	if err != nil {
		return nil, err
	}
	return j, nil
}

func (ef *EmployeeFired) UnmarshalJSON(data []byte) error {
	aux := struct {
		ID               string `json:"id"`
		FirstName        string `json:"first_name"`
		LastName         string `json:"last_name"`
		Age              int    `json:"age"`
		Email            string `json:"email"`
		Phone            string `json:"phone"`
		HiringDate       string `json:"hiring_date"`
		ApprovedHiring   bool   `json:"approved_hiring"`
		ApprovedHiringAt string `json:"approved_hiring_at"`
		Position         string `json:"position"`
		OccurredOn       string `json:"occurred_on"`
	}{}
	err := json.Unmarshal(data, &aux)
	if err != nil {
		return err
	}
	ef.id = aux.ID
	ef.firstName = aux.FirstName
	ef.lastName = aux.LastName
	ef.age = aux.Age
	ef.email = aux.Email
	ef.phone = aux.Phone
	ef.hiringDate, err = common.ParseDate(aux.HiringDate)
	if err != nil {
		return err
	}
	ef.approvedHiring = aux.ApprovedHiring
	ef.approvedHiringAt, err = common.ParseDate(aux.ApprovedHiringAt)
	if err != nil {
		return err
	}
	ef.position, err = ParsePosition(aux.Position)
	if err != nil {
		return err
	}
	ef.occurredOn, err = common.ParseDate(aux.OccurredOn)
	if err != nil {
		return err
	}
	return nil
}
