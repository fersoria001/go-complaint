package employee

import (
	"go-complaint/domain/model/common"
	"go-complaint/domain/model/enterprise"
	"go-complaint/domain/model/identity"
	"go-complaint/erros"

	"github.com/google/uuid"
)

/*
Id = enterpriseName + - + emailOFManagerWhoHiredThis + dd/yy//mm + employeeID
*/
type Employee struct {
	id           uuid.UUID
	enterpriseID string
	*identity.User
	hiringDate       common.Date
	approvedHiring   bool
	approvedHiringAt common.Date
	position         enterprise.Position
}

func NewEmployee(
	id uuid.UUID,
	enterpriseID string,
	user *identity.User,
	position enterprise.Position,
	hiringDate common.Date,
	approvedHiring bool,
	approvedHiringAt common.Date) (*Employee, error) {
	var e *Employee = new(Employee)
	err := e.setID(id)
	if err != nil {
		return nil, err
	}
	err = e.setEnterpriseID(enterpriseID)
	if err != nil {
		return nil, err
	}
	err = e.setUser(user)
	if err != nil {
		return nil, err
	}
	err = e.SetPosition(position)
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

func (e *Employee) setID(id uuid.UUID) error {
	if id == (uuid.UUID{}) {
		return &erros.NullValueError{}
	}
	e.id = id
	return nil
}

func (e *Employee) setUser(user *identity.User) error {
	if user == nil {
		return &erros.NullValueError{}
	}
	e.User = user
	return nil
}

func (e *Employee) setApprovedHiringAt(approvedHiringAt common.Date) error {
	if approvedHiringAt == (common.Date{}) {
		return &erros.NullValueError{}
	}
	e.approvedHiringAt = approvedHiringAt
	return nil
}

func (e *Employee) setEnterpriseID(enterpriseID string) error {
	if enterpriseID == "" {
		return &erros.NullValueError{}
	}
	e.enterpriseID = enterpriseID
	return nil
}

func (e *Employee) setHiringDate(hiringDate common.Date) error {
	if hiringDate == (common.Date{}) {
		return &erros.NullValueError{}
	}
	e.hiringDate = hiringDate
	return nil

}

func (e *Employee) SetPosition(position enterprise.Position) error {
	if position < 0 {
		return &erros.ValidationError{Expected: "position needs to be greater than 0"}
	}
	if position > 3 {
		return &erros.ValidationError{Expected: "position needs to be less than 3"}
	}
	e.position = position
	return nil
}

func (e Employee) GetUser() *identity.User {
	return e.User
}

func (e Employee) Email() string {
	return e.User.Email()
}

func (e Employee) EnterpriseID() string {
	return e.enterpriseID
}

func (e Employee) HiringDate() common.Date {
	return e.hiringDate
}

func (e Employee) ApprovedHiring() bool {
	return e.approvedHiring
}

func (e Employee) ApprovedHiringAt() common.Date {
	return e.approvedHiringAt
}

func (e Employee) Position() enterprise.Position {
	return e.position
}

func (e Employee) ID() uuid.UUID {
	return e.id
}
