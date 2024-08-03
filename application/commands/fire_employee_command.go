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

type FireEmployeeCommmand struct {
	EmployeeId    string `json:"employeeId"`
	TriggeredById string `json:"triggeredById"`
	FireReason    string `json:"fireReason"`
}

func NewFireEmployeeCommand(employeeId, triggeredById, fireReason string) *FireEmployeeCommmand {
	return &FireEmployeeCommmand{
		EmployeeId:    employeeId,
		TriggeredById: triggeredById,
		FireReason:    fireReason,
	}
}

func (c FireEmployeeCommmand) Execute(ctx context.Context) error {
	employeeId, err := uuid.Parse(c.EmployeeId)
	if err != nil {
		return err
	}
	triggeredById, err := uuid.Parse(c.TriggeredById)
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
	employeeRepository, ok := reg.Get("Employee").(repositories.EmployeeRepository)
	if !ok {
		return ErrWrongTypeAssertion
	}
	employee, err := employeeRepository.Get(ctx, employeeId)
	if err != nil {
		return err
	}
	hiringProccess, err := hiringProccessRepository.Find(ctx, find_hiring_proccess.ByUserIdAndEnterpriseId(employee.GetUser().Id(), employee.EnterpriseId()))
	if err != nil {
		return err
	}
	triggeredBy, err := recipientRepository.Get(ctx, triggeredById)
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
					fmt.Sprintf("Your work has end at %s", hiringProccess.Enterprise().SubjectName()),
					fmt.Sprintf("You no longer hold the position %s at %s ", hiringProccess.Role(), hiringProccess.Enterprise().SubjectName()),
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
			return reflect.TypeOf(&enterprise.HiringProccessStatusChanged{})
		},
	})
	domain.DomainEventPublisherInstance().Subscribe(domain.DomainEventSubscriber{
		HandleEvent: func(event domain.DomainEvent) error {
			if e, ok := event.(*enterprise.HiringProccessStatusChanged); ok {
				if enterprise.ParseHiringProccessStatus(e.NewStatus()) != enterprise.FIRED {
					return fmt.Errorf("unexpected status in HiringProccessStatusChanged event handler at FireEmployeeCommand")
				}
				reg := repositories.MapperRegistryInstance()
				userRepository, ok := reg.Get("User").(repositories.UserRepository)
				if !ok {
					return ErrWrongTypeAssertion
				}
				hiringProccessRepository, ok := reg.Get("HiringProccess").(repositories.HiringProccessRepository)
				if !ok {
					return ErrWrongTypeAssertion
				}
				hiringProccess, err := hiringProccessRepository.Get(ctx, e.HiringProccessId())
				if err != nil {
					return err
				}
				user, err := userRepository.Get(ctx, hiringProccess.User().Id())
				if err != nil {
					return err
				}
				role, err := identity.ParseRole(employee.Position().String())
				if err != nil {
					return err
				}
				err = user.RemoveUserRole(ctx, role, hiringProccess.Enterprise().Id())
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
				if enterprise.ParseHiringProccessStatus(e.NewStatus()) != enterprise.FIRED {
					return fmt.Errorf("unexpected status in HiringProccessStatusChanged event handler at FireEmployeeCommand")
				}
				reg := repositories.MapperRegistryInstance()
				hiringProccessRepository, ok := reg.Get("HiringProccess").(repositories.HiringProccessRepository)
				if !ok {
					return ErrWrongTypeAssertion
				}
				enterpriseRepository, ok := reg.Get("Enterprise").(repositories.EnterpriseRepository)
				if !ok {
					return ErrWrongTypeAssertion
				}
				hiringProccess, err := hiringProccessRepository.Get(ctx, e.HiringProccessId())
				if err != nil {
					return err
				}
				dbE, err := enterpriseRepository.Get(ctx, hiringProccess.Enterprise().Id())
				if err != nil {
					return err
				}
				err = dbE.FireEmployee(ctx, hiringProccess.UpdatedBy().Id(), hiringProccess.User().Id())
				if err != nil {
					return err
				}
				err = enterpriseRepository.Update(ctx, dbE)
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
	err = hiringProccess.ChangeStatus(ctx, enterprise.FIRED, *triggeredBy)
	if err != nil {
		return err
	}
	hiringProccess.WriteAReason(c.FireReason, *triggeredBy)
	err = hiringProccessRepository.Update(ctx, *hiringProccess)
	if err != nil {
		return err
	}
	return nil
}
