package commands

import (
	"context"
	"fmt"
	"go-complaint/domain"
	"go-complaint/domain/model/enterprise"
	"go-complaint/domain/model/identity"
	"go-complaint/infrastructure/persistence/finders/find_hiring_proccess"
	"go-complaint/infrastructure/persistence/repositories"
	"reflect"

	"github.com/google/uuid"
)

type LeaveEnterpriseCommand struct {
	EmployeeId string `json:"employeeId"`
}

func NewLeaveEnterpriseCommand(employeeId string) *LeaveEnterpriseCommand {
	return &LeaveEnterpriseCommand{
		EmployeeId: employeeId,
	}
}

func (c LeaveEnterpriseCommand) Execute(ctx context.Context) error {
	employeeId, err := uuid.Parse(c.EmployeeId)
	if err != nil {
		return err
	}
	reg := repositories.MapperRegistryInstance()
	employeeRepository, ok := reg.Get("Employee").(repositories.EmployeeRepository)
	if !ok {
		return ErrWrongTypeAssertion
	}
	enterpriseRepository, ok := reg.Get("Enterprise").(repositories.EnterpriseRepository)
	if !ok {
		return ErrWrongTypeAssertion
	}
	employee, err := employeeRepository.Get(ctx, employeeId)
	if err != nil {
		return err
	}
	dbEnterprise, err := enterpriseRepository.Get(ctx, employee.EnterpriseId())
	if err != nil {
		return err
	}
	domain.DomainEventPublisherInstance().Subscribe(domain.DomainEventSubscriber{
		HandleEvent: func(event domain.DomainEvent) error {
			if e, ok := event.(*enterprise.EmployeeLeaved); ok {
				reg := repositories.MapperRegistryInstance()
				userRepository, ok := reg.Get("User").(repositories.UserRepository)
				if !ok {
					return ErrWrongTypeAssertion
				}
				user, err := userRepository.Get(ctx, e.UserId())
				if err != nil {
					return err
				}
				c := NewSendNotificationCommand(
					e.EnterpriseId().String(),
					e.UserId().String(),
					fmt.Sprintf("%s has leaved %s", user.FullName(), dbEnterprise.Name()),
					fmt.Sprintf("%s is no longer part of %s", user.FullName(), dbEnterprise.Name()),
					fmt.Sprintf("/enterprises/%s/employees/hiring", dbEnterprise.Name()),
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
			return reflect.TypeOf(&enterprise.EmployeeLeaved{})
		},
	})
	domain.DomainEventPublisherInstance().Subscribe(domain.DomainEventSubscriber{
		HandleEvent: func(event domain.DomainEvent) error {
			if e, ok := event.(*enterprise.EmployeeLeaved); ok {
				reg := repositories.MapperRegistryInstance()
				userRepository, ok := reg.Get("User").(repositories.UserRepository)
				if !ok {
					return ErrWrongTypeAssertion
				}
				user, err := userRepository.Get(ctx, e.UserId())
				if err != nil {
					return err
				}
				role, err := identity.ParseRole(e.Position().String())
				if err != nil {
					return err
				}
				err = user.RemoveUserRole(ctx, role, e.EnterpriseId())
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
			return reflect.TypeOf(&enterprise.EmployeeLeaved{})
		},
	})
	domain.DomainEventPublisherInstance().Subscribe(domain.DomainEventSubscriber{
		HandleEvent: func(event domain.DomainEvent) error {
			if e, ok := event.(*enterprise.EmployeeLeaved); ok {
				reg := repositories.MapperRegistryInstance()
				hiringProccessRepository, ok := reg.Get("HiringProccess").(repositories.HiringProccessRepository)
				if !ok {
					return ErrWrongTypeAssertion
				}
				recipientRepository, ok := reg.Get("Recipient").(repositories.RecipientRepository)
				if !ok {
					return ErrWrongTypeAssertion
				}
				user, err := recipientRepository.Get(ctx, e.UserId())
				if err != nil {
					return err
				}
				hiringProccess, err := hiringProccessRepository.Find(ctx, find_hiring_proccess.ByUserIdAndEnterpriseId(
					e.UserId(), e.EnterpriseId(),
				))
				if err != nil {
					return err
				}
				err = hiringProccess.ChangeStatus(ctx, enterprise.LEAVED, *user)
				if err != nil {
					return err
				}
				err = hiringProccessRepository.Update(ctx, *hiringProccess)
				if err != nil {
					return err
				}
				return nil
			}
			return ErrWrongTypeAssertion
		},
		SubscribedToEventType: func() reflect.Type {
			return reflect.TypeOf(&enterprise.EmployeeLeaved{})
		},
	})
	err = dbEnterprise.EmployeeLeave(ctx, employeeId)
	if err != nil {
		return err
	}
	err = enterpriseRepository.Update(ctx, dbEnterprise)
	if err != nil {
		return err
	}
	return nil
}
