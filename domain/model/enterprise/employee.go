package enterprise

import (
	"context"
	"go-complaint/domain"
	"go-complaint/domain/model/common"
	"go-complaint/erros"
	"strings"
	"time"

	"github.com/google/uuid"
)

/*
Id = enterpriseName + - + emailOFManagerWhoHiredThis + dd/yy//mm + employeeID
*/
type Employee struct {
	id               string
	profileIMG       string
	firstName        string
	lastName         string
	age              int
	email            string
	phone            string
	hiringDate       common.Date
	approvedHiring   bool
	approvedHiringAt common.Date
	position         Position
}

// It's a pre-hire
func (manager *Employee) HireEmployee(
	ctx context.Context, invitationID uuid.UUID,
	newID, profileIMG, firstName, lastName, email, phone string, age int) (*Employee, error) {
	if manager.Position() < MANAGER {
		return nil, &erros.ValidationError{Expected: "only managers can hire employees"}
	}
	var (
		publisher = domain.DomainEventPublisherInstance()
		employee  *Employee
		err       error
	)
	today := common.NewDate(time.Now())
	employee, err = NewEmployee(newID,
		profileIMG,
		firstName, lastName, age, email, phone, NOT_ASSIGNED, today, false, today)
	if err != nil {
		return nil, err
	}
	err = publisher.Publish(ctx, NewEmployeeWaitingForApproval(employee.ID(), manager.ID(), invitationID))
	if err != nil {
		return nil, err
	}
	return employee, nil
}

func (manager *Employee) SentHiringInvitation(
	ctx context.Context,
	profileIMG,
	firstName,
	lastName,
	email,
	phone string,
	age int,
) error {
	enterpriseName := strings.Split(manager.ID(), "-")[0]
	publisher := domain.DomainEventPublisherInstance()
	return publisher.Publish(
		ctx,
		NewJobSelectionInvitationSent(
			enterpriseName,
			manager.ID(),
			profileIMG,
			firstName,
			lastName,
			email,
			phone,
			age),
	)
}

func (e *Employee) GetFullName() string {
	return e.firstName + " " + e.lastName
}

func NewEmployee(
	id string,
	profileIMG string,
	firstName string,
	lastName string,
	age int,
	email string,
	phone string,
	position Position,
	hiringDate common.Date,
	approvedHiring bool,
	approvedHiringAt common.Date) (*Employee, error) {
	var e *Employee = new(Employee)
	err := e.setID(id)
	if err != nil {
		return nil, err
	}
	e.profileIMG = profileIMG
	err = e.setFirstName(firstName)
	if err != nil {
		return nil, err
	}
	err = e.setLastName(lastName)
	if err != nil {
		return nil, err
	}
	err = e.setAge(age)
	if err != nil {
		return nil, err
	}
	err = e.setEmail(email)
	if err != nil {
		return nil, err
	}
	err = e.setPhone(phone)
	if err != nil {
		return nil, err
	}
	err = e.setPosition(position)
	if err != nil {

		return nil, err
	}
	err = e.setHiringDate(hiringDate)
	if err != nil {
		return nil, err
	}
	err = e.setApprovedHiringAt(approvedHiringAt)
	if err != nil {
		return nil, err
	}
	e.SetApprovedHiring(approvedHiring)
	return e, nil
}

func (e *Employee) SetApprovedHiring(approvedHiring bool) {
	e.approvedHiring = approvedHiring
}

func (e *Employee) setApprovedHiringAt(approvedHiringAt common.Date) error {
	if approvedHiringAt == (common.Date{}) {
		return &erros.NullValueError{}
	}
	e.approvedHiringAt = approvedHiringAt
	return nil
}

func (e *Employee) setID(id string) error {
	if id == "" {
		return &erros.NullValueError{}
	}
	e.id = id
	return nil
}

func (e *Employee) setFirstName(firstName string) error {
	if firstName == "" {
		return &erros.NullValueError{}
	}
	e.firstName = firstName
	return nil
}

func (e *Employee) setLastName(lastName string) error {
	if lastName == "" {
		return &erros.NullValueError{}
	}
	e.lastName = lastName
	return nil
}

func (e *Employee) setAge(age int) error {
	if age < 0 {
		return &erros.ValidationError{Expected: "age needs to be greater than 0"}
	}
	e.age = age
	return nil
}

func (e *Employee) setEmail(email string) error {
	if email == "" {
		return &erros.NullValueError{}
	}
	e.email = email
	return nil
}

func (e *Employee) setPhone(phone string) error {
	if phone == "" {
		return &erros.NullValueError{}
	}
	e.phone = phone
	return nil
}

func (e *Employee) setHiringDate(hiringDate common.Date) error {
	if hiringDate == (common.Date{}) {
		return &erros.NullValueError{}
	}
	e.hiringDate = hiringDate
	return nil

}

func (e *Employee) setPosition(position Position) error {
	if position < 0 {
		return &erros.ValidationError{Expected: "position needs to be greater than 0"}
	}
	if position > 3 {
		return &erros.ValidationError{Expected: "position needs to be less than 3"}
	}
	e.position = position
	return nil
}

func (e *Employee) ID() string {
	return e.id
}

func (e *Employee) FirstName() string {
	return e.firstName
}

func (e *Employee) LastName() string {
	return e.lastName
}

func (e *Employee) Age() int {
	return e.age
}

func (e *Employee) Email() string {
	return e.email
}

func (e *Employee) Phone() string {
	return e.phone
}

func (e *Employee) HiringDate() common.Date {
	return e.hiringDate
}

func (e *Employee) ApprovedHiring() bool {
	return e.approvedHiring
}

func (e *Employee) ApprovedHiringAt() common.Date {
	return e.approvedHiringAt
}

func (e *Employee) Position() Position {
	return e.position
}

func (e *Employee) ProfileIMG() string {
	return e.profileIMG
}
