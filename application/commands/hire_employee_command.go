package commands

import (
	"context"
	"fmt"
	"go-complaint/domain"
	"go-complaint/domain/model/common"
	"go-complaint/domain/model/enterprise"
	"go-complaint/domain/model/identity"
	"go-complaint/infrastructure/persistence/repositories"
	"reflect"
	"time"

	"github.com/google/uuid"
)

type HireEmployeeCommand struct {
	HiringProccessId string `json:"hiringProccessId"`
	EmployeeId       string `json:"employeeId"`
}

func NewHireEmployeeCommand(hiringProccessId, employeeId string) *HireEmployeeCommand {
	return &HireEmployeeCommand{
		HiringProccessId: hiringProccessId,
		EmployeeId:       employeeId,
	}
}

func (c HireEmployeeCommand) Execute(ctx context.Context) error {
	employeeUserId, err := uuid.Parse(c.EmployeeId)
	if err != nil {
		return err
	}
	hiringProccessId, err := uuid.Parse(c.HiringProccessId)
	if err != nil {
		return err
	}
	domain.DomainEventPublisherInstance().Subscribe(domain.DomainEventSubscriber{
		HandleEvent: func(event domain.DomainEvent) error {
			if e, ok := event.(*enterprise.HiringProccessStatusChanged); ok {
				reg := repositories.MapperRegistryInstance()
				hiringProccessRepository, ok := reg.Get("HiringProccess").(repositories.HiringProccessRepository)
				if !ok {
					return ErrWrongTypeAssertion
				}
				hiringProccess, err := hiringProccessRepository.Get(ctx, e.HiringProccessId())
				if err != nil {
					return err
				}
				c := NewSendNotificationCommand(
					hiringProccess.User().Id().String(),
					hiringProccess.Enterprise().Id().String(),
					fmt.Sprintf("You have been hired  to be part of %s", hiringProccess.Enterprise().SubjectName()),
					fmt.Sprintf("You have been hired by %s as %s in %s", hiringProccess.UpdatedBy().SubjectName(),
						hiringProccess.Enterprise().SubjectName(), hiringProccess.Role().String()),
					fmt.Sprintf("/%s", hiringProccess.Enterprise().SubjectName()),
				)
				err = c.Execute(ctx)
				if err != nil {
					return err
				}
				return nil
			}
			return ErrWrongTypeAssertion
		},
		SubscribedToEventType: func() reflect.Type {
			return reflect.TypeOf(&enterprise.HiringProccessStatusChanged{})
		},
	})
	domain.DomainEventPublisherInstance().Subscribe(domain.DomainEventSubscriber{
		HandleEvent: func(event domain.DomainEvent) error {
			if e, ok := event.(*enterprise.HiringProccessStatusChanged); ok {
				reg := repositories.MapperRegistryInstance()
				hiringProccessRepository, ok := reg.Get("HiringProccess").(repositories.HiringProccessRepository)
				if !ok {
					return ErrWrongTypeAssertion
				}
				userRepository, ok := reg.Get("User").(repositories.UserRepository)
				if !ok {
					return ErrWrongTypeAssertion
				}
				hiringProccess, err := hiringProccessRepository.Get(ctx, e.HiringProccessId())
				if err != nil {
					return err
				}
				toRole, err := identity.ParseRole(hiringProccess.Role().String())
				if err != nil {
					return err
				}
				user, err := userRepository.Get(ctx, hiringProccess.User().Id())
				if err != nil {
					return err
				}
				err = user.AddRole(ctx, toRole, hiringProccess.Enterprise().Id())
				if err != nil {
					return err
				}
				err = userRepository.Update(ctx, user)
				if err != nil {
					return err
				}
				return nil
			}
			return ErrWrongTypeAssertion
		},
		SubscribedToEventType: func() reflect.Type {
			return reflect.TypeOf(&enterprise.HiringProccessStatusChanged{})
		},
	})
	domain.DomainEventPublisherInstance().Subscribe(domain.DomainEventSubscriber{
		HandleEvent: func(event domain.DomainEvent) error {
			if e, ok := event.(*enterprise.HiringProccessStatusChanged); ok {
				if enterprise.ParseHiringProccessStatus(e.NewStatus()) != enterprise.HIRED {
					return fmt.Errorf("unexpected status at hire employee HiringProccessStatusChanged event handler")
				}
				reg := repositories.MapperRegistryInstance()
				enterpriseRepository, ok := reg.Get("Enterprise").(repositories.EnterpriseRepository)
				if !ok {
					return ErrWrongTypeAssertion
				}
				hiringProccessRepository, ok := reg.Get("HiringProccess").(repositories.HiringProccessRepository)
				if !ok {
					return ErrWrongTypeAssertion
				}
				userRepository, ok := reg.Get("User").(repositories.UserRepository)
				if !ok {
					return ErrWrongTypeAssertion
				}
				hiringProccess, err := hiringProccessRepository.Get(ctx, hiringProccessId)
				if err != nil {
					return err
				}
				dbEnterprise, err := enterpriseRepository.Get(ctx, hiringProccess.Enterprise().Id())
				if err != nil {
					return err
				}
				user, err := userRepository.Get(ctx, hiringProccess.User().Id())
				if err != nil {
					return err
				}
				newId := uuid.New()
				newEmployee, err := enterprise.NewEmployee(
					newId,
					dbEnterprise.Id(),
					user,
					hiringProccess.Role(),
					common.NewDate(time.Now()),
					true,
					common.NewDate(time.Now()),
				)
				if err != nil {
					return err
				}
				err = dbEnterprise.HireEmployee(ctx, newId, newEmployee)
				if err != nil {
					return err
				}
				return enterpriseRepository.Update(ctx, dbEnterprise)
			}
			return ErrWrongTypeAssertion
		},
		SubscribedToEventType: func() reflect.Type {
			return reflect.TypeOf(&enterprise.HiringProccessStatusChanged{})
		},
	})
	reg := repositories.MapperRegistryInstance()
	recipientRepository, ok := reg.Get("Recipient").(repositories.RecipientRepository)
	if !ok {
		return ErrWrongTypeAssertion
	}
	hiringProccessRepository, ok := reg.Get("HiringProccess").(repositories.HiringProccessRepository)
	if !ok {
		return ErrWrongTypeAssertion
	}
	employeeRecipient, err := recipientRepository.Get(ctx, employeeUserId)
	if err != nil {
		return err
	}
	hiringProccess, err := hiringProccessRepository.Get(ctx, hiringProccessId)
	if err != nil {
		return err
	}
	err = hiringProccess.ChangeStatus(ctx, enterprise.HIRED, *employeeRecipient)
	if err != nil {
		return err
	}
	return hiringProccessRepository.Update(ctx, *hiringProccess)
}
