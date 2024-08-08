package commands

import (
	"context"
	"fmt"
	"go-complaint/domain"
	"go-complaint/domain/model/enterprise"
	"go-complaint/domain/model/identity"
	"go-complaint/infrastructure/persistence/finders/find_user"
	"go-complaint/infrastructure/persistence/repositories"
	"reflect"

	"github.com/google/uuid"
)

type PromoteEmployeeCommand struct {
	EmployeeId   string `json:"employeeId"`
	PromoteTo    string `json:"promoteTo"`
	PromotedById string `json:"promotedById"`
}

func NewPromoteEmployeeCommand(employeeId, promoteTo, promoteById string) *PromoteEmployeeCommand {
	return &PromoteEmployeeCommand{
		EmployeeId:   employeeId,
		PromoteTo:    promoteTo,
		PromotedById: promoteById,
	}
}

func (c PromoteEmployeeCommand) Execute(ctx context.Context) error {
	employeeId, err := uuid.Parse(c.EmployeeId)
	if err != nil {
		return err
	}
	promotedById, err := uuid.Parse(c.PromotedById)
	if err != nil {
		return err
	}
	promoteTo := enterprise.ParsePosition(c.PromoteTo)
	if promoteTo < 0 {
		return fmt.Errorf("this enterprise position doesn't exists")
	}
	reg := repositories.MapperRegistryInstance()
	enterpriseRepository, ok := reg.Get("Enterprise").(repositories.EnterpriseRepository)
	if !ok {
		return ErrWrongTypeAssertion
	}
	employeeRepository, ok := reg.Get("Employee").(repositories.EmployeeRepository)
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
			if e, ok := event.(*enterprise.EmployeePromoted); ok {
				reg := repositories.MapperRegistryInstance()
				recipientRepository, ok := reg.Get("Recipient").(repositories.RecipientRepository)
				if !ok {
					return ErrWrongTypeAssertion
				}
				userRepository, ok := reg.Get("User").(repositories.UserRepository)
				if !ok {
					return ErrWrongTypeAssertion
				}
				user, err := userRepository.Find(ctx, find_user.ByUsername(e.EmployeeUsername()))
				if err != nil {
					return err
				}
				enterpriseRecipient, err := recipientRepository.Get(ctx, e.EnterpriseId())
				if err != nil {
					return err
				}
				c := NewSendNotificationCommand(
					user.Id().String(),
					e.EnterpriseId().String(),
					fmt.Sprintf("You have been promoted in %s", enterpriseRecipient.SubjectName()),
					fmt.Sprintf("You have been promoted to %s at %s", e.NewPosition().String(), enterpriseRecipient.SubjectName()),
					"/",
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
			return reflect.TypeOf(&enterprise.EmployeePromoted{})
		},
	})
	domain.DomainEventPublisherInstance().Subscribe(domain.DomainEventSubscriber{
		HandleEvent: func(event domain.DomainEvent) error {
			if e, ok := event.(*enterprise.EmployeePromoted); ok {
				reg := repositories.MapperRegistryInstance()
				employeeRepository, ok := reg.Get("Employee").(repositories.EmployeeRepository)
				if !ok {
					return ErrWrongTypeAssertion
				}
				userRepository, ok := reg.Get("User").(repositories.UserRepository)
				if !ok {
					return ErrWrongTypeAssertion
				}
				employee, err := employeeRepository.Get(ctx, employeeId)
				if err != nil {
					return err
				}
				user, err := userRepository.Get(ctx, employee.Id())
				if err != nil {
					return err
				}
				prevRole, err := identity.ParseRole(e.PrevPosition().String())
				if err != nil {
					return err
				}
				err = user.RemoveUserRole(ctx, prevRole, e.EnterpriseId())
				if err != nil {
					return err
				}
				newRole, err := identity.ParseRole(e.NewPosition().String())
				if err != nil {
					return err
				}
				err = user.AddRole(ctx, newRole, e.EnterpriseId(), dbEnterprise.Name())
				if err != nil {
					return err
				}
				err = userRepository.Update(ctx, user)
				if err != nil {
					return err
				}
				return nil
			}
			return nil
		},
		SubscribedToEventType: func() reflect.Type {
			return reflect.TypeOf(&enterprise.EmployeePromoted{})
		},
	})
	err = dbEnterprise.PromoteEmployee(ctx, promotedById, employee.UserName(), promoteTo)
	if err != nil {
		return err
	}
	err = enterpriseRepository.Update(ctx, dbEnterprise)
	if err != nil {
		return err
	}
	return nil
}
