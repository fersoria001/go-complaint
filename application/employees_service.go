package application

import (
	"context"
	"fmt"
	"go-complaint/domain"
	"go-complaint/domain/model/enterprise"
	"go-complaint/dto"
	"go-complaint/erros"
	"go-complaint/infrastructure/persistence/repositories"
	"log"
	"reflect"
	"time"

	"github.com/google/uuid"
)

type EmployeesService struct {
	repository      *repositories.EmployeeRepository
	identityService *IdentityService
}

func NewEmployeesService(repository *repositories.EmployeeRepository, identityService *IdentityService) *EmployeesService {
	return &EmployeesService{
		repository:      repository,
		identityService: identityService,
	}
}

func (employeeService *EmployeesService) Employee(ctx context.Context, employeeID string) (*dto.Employee, error) {
	employee, err := employeeService.repository.Get(ctx, employeeID)
	if err != nil {
		return nil, err
	}
	return dto.NewEmployee(*employee), nil
}

func (employeeService *EmployeesService) Employees(ctx context.Context, enterpriseID string) ([]*dto.Employee, error) {
	employees, err := employeeService.repository.FindByEnterpriseID(ctx, enterpriseID)
	if err != nil {
		return nil, err
	}
	var employeesDTO []*dto.Employee
	for employee := range employees.Iter() {
		employeesDTO = append(employeesDTO, dto.NewEmployee(*employee))
	}
	return employeesDTO, nil
}

// func (employeeService *EmployeesService) EnterpriseEmployees(ctx context.Context, enterpriseName string, callerEmail string) ([]*dto.Employee, error)

// If invitation not ignored, then start the hiring proccess
// This is called by the notify service based on
// an event restored from the event store
// This method should be named as "StartHiringProcess" or "StartHiring" or "AcceptHiringInvitation"
func (employeeService *EmployeesService) AcceptJobSelectionInvitation(ctx context.Context,
	invitationID string,
	enterpriseName, managerID,
	profileIMG,
	firstName, lastName, email, phone string, age int) error {
	var (
		manager             *enterprise.Employee
		unconfirmedEmployee *enterprise.Employee
		err                 error
	)
	domain.DomainEventPublisherInstance().Subscribe(domain.DomainEventSubscriber{
		HandleEvent: func(event domain.DomainEvent) error {
			if _, ok := event.(*enterprise.EmployeeHired); ok {
				fmt.Println("Employee Hired Handler")
				return nil
			}
			return &erros.ValueNotFoundError{}
		},
		SubscribedToEventType: func() reflect.Type {
			return reflect.TypeOf(&enterprise.EmployeeHired{})
		},
	})
	manager, err = employeeService.repository.Get(ctx, managerID)
	if err != nil {
		log.Println("MANAGER NOT FOUND", managerID)
		return err
	}
	//its a 'pre-hire'
	parsedInvitationID, err := uuid.Parse(invitationID)
	if err != nil {
		return err
	}
	unconfirmedEmployee, err = manager.HireEmployee(ctx,
		parsedInvitationID,
		employeeService.NewEmployeeID(
			enterpriseName, manager.Email(), email),
		profileIMG,
		firstName, lastName, email, phone, age)
	if err != nil {
		return err
	}
	err = employeeService.repository.Save(ctx, unconfirmedEmployee)
	if err != nil {
		return err
	}
	return nil
}

// Notify the manager about the hiring invitation

// SentHiringInvitation(ctx context.Context,managerID, userID string) error
func (employeeService *EmployeesService) SendHiringInvitation(
	ctx context.Context,
	managerID,
	userID string,
) error {
	var (
		manager *enterprise.Employee
		err     error
	)
	domain.DomainEventPublisherInstance().Subscribe(domain.DomainEventSubscriber{
		HandleEvent: func(event domain.DomainEvent) error {
			if _, ok := event.(*enterprise.HiringInvitationSent); ok {
				fmt.Println("Hiring Invitation Sent Handler")
				return nil
			}
			return &erros.ValueNotFoundError{}
		},
		SubscribedToEventType: func() reflect.Type {
			return reflect.TypeOf(&enterprise.HiringInvitationSent{})
		},
	})

	manager, err = employeeService.repository.Get(ctx, managerID)
	if err != nil {
		log.Println("MANAGER NOT FOUND")
		return err
	}

	user, err := employeeService.identityService.User(ctx, userID)
	if err != nil {
		log.Println("USER NOT FOUND")
		return err
	}
	age := time.Now().Year() - user.Person().BirthDate().Date().Year()
	return manager.SentHiringInvitation(
		ctx,
		user.ProfileIMG(),
		user.Person().FirstName(),
		user.Person().LastName(),
		user.Email(),
		user.Person().Phone(),
		age)
}

func (employeeService *EmployeesService) NewEmployeeID(enterpriseName, managerEmail, userID string) string {
	return enterpriseName + "-" + managerEmail + "-" + time.Now().Format("02/01/2006") + "-" + userID
}
