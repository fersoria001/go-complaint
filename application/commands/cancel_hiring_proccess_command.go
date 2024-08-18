package commands

import (
	"context"
	"fmt"
	"go-complaint/domain"
	"go-complaint/domain/model/enterprise"
	"go-complaint/infrastructure/persistence/repositories"
	"reflect"

	"github.com/google/uuid"
)

type CancelHiringProccessCommand struct {
	HiringProccessId  string `json:"hiringProccessId"`
	EmployeeUserId    string `json:"employeeUserId"`
	CancelationReason string `json:"cancelationReason"`
}

func NewCancelHiringProccessCommand(hiringProccessId, employeeUserId, cancelationReason string) *CancelHiringProccessCommand {
	return &CancelHiringProccessCommand{
		HiringProccessId:  hiringProccessId,
		EmployeeUserId:    employeeUserId,
		CancelationReason: cancelationReason,
	}
}

func (c CancelHiringProccessCommand) Execute(ctx context.Context) error {
	hiringProccessId, err := uuid.Parse(c.HiringProccessId)
	if err != nil {
		return err
	}
	employeeUserId, err := uuid.Parse(c.EmployeeUserId)
	if err != nil {
		return err
	}
	reg := repositories.MapperRegistryInstance()
	hiringProccessRepository, ok := reg.Get("HiringProccess").(repositories.HiringProccessRepository)
	if !ok {
		return ErrWrongTypeAssertion
	}
	recipientRepository, ok := reg.Get("Recipient").(repositories.RecipientRepository)
	if !ok {
		return ErrWrongTypeAssertion
	}
	employeeUserRecipient, err := recipientRepository.Get(ctx, employeeUserId)
	if err != nil {
		return err
	}
	hiringProccess, err := hiringProccessRepository.Get(ctx, hiringProccessId)
	if err != nil {
		return err
	}
	domain.DomainEventPublisherInstance().Subscribe(domain.DomainEventSubscriber{
		HandleEvent: func(event domain.DomainEvent) error {
			if e, ok := event.(*enterprise.HiringProccessStatusChanged); ok {
				if enterprise.ParseHiringProccessStatus(e.NewStatus()) == enterprise.CANCELED {
					c := NewSendNotificationCommand(
						hiringProccess.User().Id().String(),
						hiringProccess.Enterprise().Id().String(),
						fmt.Sprintf("Your hiring process at %s has been canceled", hiringProccess.Enterprise().SubjectName()),
						fmt.Sprintf("Your hiring process at %s as %s has been canceled by %s",
							hiringProccess.Enterprise().SubjectName(), hiringProccess.Role(), employeeUserRecipient.SubjectName()),
						"/hiring",
					)
					return c.Execute(ctx)
				}
				return fmt.Errorf("incorrect status found at cancel hiring proccess event handler")
			}
			return ErrWrongTypeAssertion
		},
		SubscribedToEventType: func() reflect.Type {
			return reflect.TypeOf(&enterprise.HiringProccessStatusChanged{})
		},
	})
	err = hiringProccess.ChangeStatus(ctx, enterprise.CANCELED, *employeeUserRecipient)
	if err != nil {
		return err
	}
	hiringProccess.WriteAReason(c.CancelationReason, *employeeUserRecipient)
	err = hiringProccessRepository.Update(ctx, *hiringProccess)
	if err != nil {
		return err
	}
	return nil
}
