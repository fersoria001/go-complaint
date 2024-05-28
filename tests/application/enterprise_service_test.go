package application_test

import (
	"context"
	"encoding/json"
	"go-complaint/application"
	"go-complaint/domain"
	"go-complaint/domain/model/common"
	"go-complaint/domain/model/enterprise"
	"go-complaint/infrastructure/persistence/datasource"
	"go-complaint/infrastructure/persistence/repositories"
	"reflect"
	"testing"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
)

type HiringInviteNotification struct {
	EventID    string
	Invitation *enterprise.HiringInvitationSent
}

func TestEnterpriseService(t *testing.T) {
	ctx := context.Background()
	userRepository := repositories.NewUserRepository(datasource.IdentityAndAccessSchema())
	enterpriseRepository := repositories.NewEnterpriseRepository(datasource.EnterpriseSchema())
	employeesRepository := repositories.NewEmployeeRepository(datasource.EnterpriseSchema())
	eventStore := repositories.NewEventRepository(datasource.EventSchema())
	notificationService := application.NewNotificationService(eventStore)
	identityService := application.NewIdentityService(userRepository)
	enterpriseService := application.NewEnterpriseService(
		enterpriseRepository,
		identityService,
		employeesRepository,
		notificationService,
	)

	for _, userInfo := range usersInfo {
		err := identityService.RegisterUser(
			ctx,
			userInfo.profileImage,
			userInfo.email,
			userInfo.password,
			userInfo.firstName,
			userInfo.lastName,
			userInfo.birthDate,
			userInfo.phone,
			userInfo.country,
			userInfo.county,
			userInfo.city)
		if err != nil {
			t.Error(err)
		}
	}
	fixedEmployee, err := enterprise.NewEmployee(
		enterpriseService.NewEmployeeID("DocuSign", usersInfo["owner"].email, usersInfo["employee"].email),
		"image.jpg",
		usersInfo["employee"].firstName,
		usersInfo["employee"].lastName,
		25,
		usersInfo["employee"].email,
		usersInfo["employee"].phone,
		enterprise.MANAGER,
		common.NewDate(time.Now()),
		true,
		common.NewDate(time.Now()))
	if err != nil {
		t.Error(err)
	}
	err = employeesRepository.Save(ctx, fixedEmployee)
	if err != nil {
		t.Error(err)
	}
	var enterpriseID = "DocuSign"
	var stringDate = common.StringDate(time.Now())

	t.Run(`Create a new enterprise`, func(t *testing.T) {

		err := enterpriseService.CreateEnterprise(
			ctx,
			usersInfo["owner"].email,
			enterpriseID,
			"docusign.com",
			"docusign@live.com",
			"01234567890",
			"Country",
			"County",
			"City",
			"Tech",
			stringDate)
		if err != nil {
			t.Error(err)
		}
		_, err = enterpriseRepository.Get(ctx, "DocuSign")
		if err != nil {
			t.Error(err)
		}
	})

	t.Run(`Update an enterprise`, func(t *testing.T) {

		err := enterpriseService.UpdateEnterprise(
			ctx,
			enterpriseID,
			"www.docusign.com",
			"",
			"",
			"",
			"",
			"",
			"")
		if err != nil {
			t.Error(err)
		}
		_, err = enterpriseRepository.Get(ctx, "DocuSign")
		if err != nil {
			t.Error(err)
		}

	})

	t.Run(`Obtain the enterprises of an owner`, func(t *testing.T) {
		results, err := enterpriseService.ProvideOwnerEnterprises(ctx, usersInfo["owner"].email)
		if err != nil {
			t.Error(err)
		}
		if results.Cardinality() == 0 {
			t.Error("expected at least one enterprise")
		}
		if results.Cardinality() > 1 {
			t.Error("expected only one enterprise")
		}
		for _, result := range results.ToSlice() {
			if result.OwnerID != usersInfo["owner"].email {
				t.Error("expected the owner id to be the same")
			}
		}
	})

	t.Run(`Send hiring invitation to an user and check for an error`,
		func(t *testing.T) {
			err := enterpriseService.InviteToProject(
				ctx,
				enterpriseID,
				usersInfo["owner"].email,
				usersInfo["user"].email,
				"MANAGER",
			)
			if err != nil {
				t.Error(err)
			}
		})

	t.Run(`Send hiring invitation to an user and check for an error, then
	use the event to hire the user`,
		func(t *testing.T) {
			domain.DomainEventPublisherInstance().Subscribe(
				application.NewEventProcessor().Subscriber(),
			)
			err = enterpriseService.InviteToProject(
				ctx,
				enterpriseID,
				usersInfo["owner"].email,
				usersInfo["user"].email,
				"MANAGER",
			)
			if err != nil {
				t.Error(err)
			}
			restoredEvents, err := eventStore.GetAll(ctx)
			if err != nil {
				t.Error(err)
			}
			castedEvents := mapset.NewSet[HiringInviteNotification]()
			for event := range restoredEvents.Iter() {
				if event.TypeName != reflect.TypeOf(
					&enterprise.HiringInvitationSent{},
				).String() {
					restoredEvents.Remove(event)
				}
			}
			if restoredEvents.Cardinality() == 0 {
				t.Error("expected at least one event")
			}
			for restoredEvent := range restoredEvents.Iter() {
				var castedEvent *enterprise.HiringInvitationSent
				err := json.Unmarshal(restoredEvent.EventBody, &castedEvent)
				if err != nil {
					t.Error(err)
				}
				castedEvents.Add(HiringInviteNotification{
					EventID:    restoredEvent.EventId.String(),
					Invitation: castedEvent,
				})
			}
			if castedEvents.Cardinality() == 0 {
				t.Error("expected at least one event")
			}
			for castedEvent := range castedEvents.Iter() {
				if castedEvent.Invitation.EmployeeID() == usersInfo["user"].email {
					err = enterpriseService.AcceptHiringInvitation(ctx, castedEvent.EventID)
					if err != nil {
						t.Error(err)
					}
				}
			}

		})
	t.Run(`You can't send a hiring invitation to a non existent user`,
		func(t *testing.T) {
			err := enterpriseService.InviteToProject(
				ctx,
				enterpriseID,
				usersInfo["owner"].email,
				"nonexistent@gmail.com",
				"MANAGER",
			)
			if err == nil {
				t.Error("expected an error")
			}
		})
	t.Run(`Only the owner can send a hiring invitation`, func(t *testing.T) {
		err := enterpriseService.InviteToProject(
			ctx,
			enterpriseID,
			usersInfo["user"].email,
			usersInfo["user"].email,
			"MANAGER",
		)
		if err == nil {
			t.Error("expected an error")
		}
	})
	t.Run(`A hiring invitation to a previously hired employee should fail`,
		func(t *testing.T) {
			err = enterpriseService.InviteToProject(
				ctx,
				enterpriseID,
				usersInfo["owner"].email,
				usersInfo["employee"].email,
				"MANAGER",
			)
			if err == nil {
				t.Error("expected an error")
			}
			if err != nil {
				t.Logf("Error at hiring previously hired: %v", err)
			}
		})
	t.Cleanup(func() {
		err := enterpriseRepository.Remove(ctx, "DocuSign")
		if err != nil {
			t.Error(err)
		}
		for _, userInfo := range usersInfo {
			err := userRepository.Remove(ctx, userInfo.email)
			if err != nil {
				t.Error(err)
			}
		}
		err = employeesRepository.Remove(ctx, usersInfo["employee"].email)
		if err != nil {
			t.Error(err)
		}
		err = employeesRepository.Remove(ctx, usersInfo["user"].email)
		if err != nil {
			t.Error(err)
		}

	})

}
