package application_test

import (
	"context"
	"go-complaint/application"
	"go-complaint/domain/model/common"
	"go-complaint/domain/model/enterprise"
	"go-complaint/infrastructure/persistence/datasource"
	"go-complaint/infrastructure/persistence/repositories"
	"log"
	"net/mail"
	"strings"
	"sync"
	"testing"
	"time"
)

var employeeEventsMap map[string][]interface{}
var employeeServiceTest sync.Once

func EmployeeEventsMap() map[string][]interface{} {
	employeeServiceTest.Do(func() {
		employeeEventsMap = map[string][]interface{}{}
	})
	return employeeEventsMap
}

func TestEmployeeService(t *testing.T) {
	ctx := context.Background()
	employeeRepository := repositories.NewEmployeeRepository(datasource.EnterpriseSchema())
	userRepository := repositories.NewUserRepository(datasource.IdentityAndAccessSchema())
	identityService := application.NewIdentityService(userRepository)
	employeeService := application.NewEmployeesService(employeeRepository, identityService)
	for _, user := range userInfo {
		err := identityService.RegisterUser(
			ctx,
			user.ProfileIMG,
			user.Email,
			user.Password,
			user.FirstName,
			user.LastName,
			common.StringDate(user.BirthDate),
			user.Phone,
			user.Country,
			user.County,
			user.City)
		if err != nil {
			t.Errorf("Error registering user: %v", err)
		}
	}
	managerEmail := "manager@gmail.com"
	enterpriseName := "John Doe"
	owner, err := identityService.User(ctx, userInfo["user1"].Email)
	if err != nil {
		t.Errorf("Error getting user: %v", err)
	}
	CEOID := employeeService.NewEmployeeID(enterpriseName, owner.Email(), owner.Email())
	runtimeEmployeesIDS := map[string]string{}
	runtimeEmployeesIDS["CEO"] = CEOID

	age := time.Now().Year() - owner.Person().BirthDate().Date().Year()
	newEmployee, err := enterprise.NewEmployee(
		CEOID,
		"iamge.jpg",
		owner.Person().FirstName(),
		owner.Person().LastName(),
		age,
		owner.Email(),
		owner.Person().Phone(),
		enterprise.MANAGER,
		common.NewDate(time.Now()),
		true,
		common.NewDate(time.Now()),
	)
	if err != nil {
		t.Errorf("Error creating employee: %v", err)
	}
	err = employeeRepository.Save(ctx, newEmployee)
	if err != nil {
		t.Errorf("Error saving employee: %v", err)
	}
	t.Run("Creates a new employeeID", func(t *testing.T) {
		expectedSegments := 4
		employeeID := employeeService.NewEmployeeID(enterpriseName, managerEmail, "employeeEmail@gmail.com")
		if employeeID == "" {
			t.Errorf("EmployeeID is empty")
		}
		segments := strings.Split(employeeID, "-")
		log.Println(segments)
		if len(segments) != expectedSegments {
			t.Errorf("EmployeeID does not have the expected segments")
		}
		if segments[0] != enterpriseName {
			t.Errorf("EmployeeID does not contain the enterprise name")
		}
		if segments[1] != managerEmail {
			t.Errorf("EmployeeID does not contain the manager email")
		}
		_, err := time.Parse("02/01/2006", segments[2])
		if err != nil {
			t.Errorf("EmployeeID does not contain a valid timestamp")
		}
		if segments[3] == "" {
			t.Errorf("EmployeeID does not contain a random segment")
		}
		if _, err := mail.ParseAddress(segments[3]); err != nil {
			t.Errorf("EmployeeID does not contain a valid user ID in the last segment (an email address)")
		}
	})

	t.Run(`Start the hiring process for the user, based on the
	published event, a new event will be published`, func(t *testing.T) {
		err := employeeService.SendHiringInvitation(ctx,
			CEOID,
			userInfo["user2"].Email,
		)
		if err != nil {
			t.Errorf("Error hiring employee: %v", err)
		}

	})

	t.Cleanup(func() {
		for _, user := range userInfo {
			err := userRepository.Remove(ctx, user.Email)
			if err != nil {
				t.Errorf("Error deleting user: %v", err)
			}
		}
		for _, employeeID := range runtimeEmployeesIDS {
			err := employeeRepository.Remove(ctx, employeeID)
			if err != nil {
				t.Errorf("Error deleting employee: %v", err)
			}
		}
		eventsMap := EmployeeEventsMap()
		for key := range eventsMap {
			delete(eventsMap, key)
		}
	})
}
