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

type Owner struct {
	id string
}

func NewOwner(id string) *Owner {
	return &Owner{id: id}
}
func (owner *Owner) ApproveHiring(ctx context.Context, employee *Employee, enterpriseID string) error {
	var (
		publisher = domain.DomainEventPublisherInstance()
		err       error
	)
	employee.SetApprovedHiring(true)
	err = employee.setApprovedHiringAt(common.NewDate(time.Now()))
	if err != nil {
		return err
	}
	err = employee.setPosition(ASSISTANT)
	if err != nil {
		return err
	}
	err = publisher.Publish(ctx, NewEmployeeHired(enterpriseID, employee.ID(), owner.id))
	if err != nil {
		return err
	}
	return nil
}
func (owner *Owner) CancelHiring(ctx context.Context, employee *Employee) error {

	var (
		publisher = domain.DomainEventPublisherInstance()
		err       error
	)

	err = publisher.Publish(ctx, NewHiringProccessCanceled(employee))
	if err != nil {
		return err
	}
	return nil
}

func (owner *Owner) FireEmployee(ctx context.Context, employee *Employee) error {
	if !employee.ApprovedHiring() {
		return &erros.ValidationError{Expected: "employee needs to be approved first"}
	}
	var (
		publisher = domain.DomainEventPublisherInstance()
		err       error
	)

	err = publisher.Publish(ctx, NewEmployeeFired(employee))
	if err != nil {
		return err
	}
	return nil
}

func (owner *Owner) PromoteEmployee(
	ctx context.Context,
	enterpriseName string,
	employee *Employee,
	position string) error {
	if enterpriseName != strings.Split(employee.ID(), "-")[0] {
		return &erros.ValidationError{Expected: "employee is not part of this enterprise"}
	}
	if !employee.ApprovedHiring() {
		return &erros.ValidationError{Expected: "employee needs to be approved first"}
	}
	var (
		err       error
		publisher = domain.DomainEventPublisherInstance()
		event     *EmployeePromoted
	)
	newPosition, err := ParsePosition(position)
	if err != nil {
		return err
	}
	err = employee.setPosition(newPosition)
	if err != nil {
		return err
	}
	event = NewEmployeePromoted(enterpriseName, owner.id, employee.Email(), newPosition)
	err = publisher.Publish(ctx, event)
	if err != nil {
		return err
	}
	return nil
}

func (owner *Owner) AcceptHiringInvitation(
	ctx context.Context,
	invitationID uuid.UUID,
	employeeID,
	profileIMG,
	firstName,
	lastName,
	email,
	phone string,
	age int,
	position Position,
) (*Employee, error) {
	var (
		publisher = domain.DomainEventPublisherInstance()
		employee  *Employee
		err       error
	)
	today := common.NewDate(time.Now())
	employee, err = NewEmployee(
		employeeID,
		profileIMG,
		firstName,
		lastName,
		age,
		email,
		phone,
		position,
		today,
		true,
		today)

	if err != nil {
		return nil, err
	}
	err = publisher.Publish(ctx, NewEmployeeWaitingForApproval(employee.ID(), owner.id, invitationID))
	if err != nil {
		return nil, err
	}
	return employee, nil
}

func (owner *Owner) SendHiringInvitation(
	ctx context.Context,
	enterpriseName string,
	newID string,
	position Position,
	profileIMG,
	firstName,
	lastName,
	email,
	phone string,
	age int,
) error {
	publisher := domain.DomainEventPublisherInstance()
	return publisher.Publish(
		ctx,
		NewHiringInvitationSent(
			enterpriseName,
			owner.id,
			profileIMG,
			firstName,
			lastName,
			email,
			phone,
			age,
			position),
	)
}
