package identity_test

import (
	"context"
	"go-complaint/domain"
	"go-complaint/domain/model/identity"
	"go-complaint/test/mock_data"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUser_NewUser(t *testing.T) {
	for _, v := range mock_data.NewUsers {
		person, err := identity.NewPerson(
			v.Person.Id,
			v.Person.Email,
			v.Person.ProfileImg,
			v.Person.Genre,
			v.Person.Pronoun,
			v.Person.FirstName,
			v.Person.LastName,
			v.Person.Phone,
			v.Person.BirthDate,
			v.Person.Address,
		)
		assert.Nil(t, err)
		assert.Equal(t, v.Person.Id, person.Id())
		assert.Equal(t, v.Person.Email, person.Email())
		assert.Equal(t, v.Person.ProfileImg, person.ProfileIMG())
		assert.Equal(t, v.Person.Genre, person.Genre())
		assert.Equal(t, v.Person.Pronoun, person.Pronoun())
		assert.Equal(t, v.Person.FirstName, person.FirstName())
		assert.Equal(t, v.Person.LastName, person.LastName())
		assert.Equal(t, v.Person.Phone, person.Phone())
		assert.Equal(t, v.Person.BirthDate.StringRepresentation(), person.BirthDate().StringRepresentation())
		assert.Equal(t, v.Person.Address, person.Address())
		user, err := identity.NewUser(
			v.Id,
			v.UserName,
			v.Password,
			v.RegisterDate,
			person,
			v.IsConfirmed,
			v.UserRoles,
		)
		assert.Nil(t, err)
		assert.Equal(t, v.Id, user.Id())
		assert.Equal(t, v.UserName, user.UserName())
		assert.Equal(t, v.Password, user.Password())
		assert.Equal(t, v.RegisterDate.StringRepresentation(), user.RegisterDate().StringRepresentation())
		assert.Equal(t, v.IsConfirmed, user.IsConfirmed())
		assert.Equal(t, v.UserRoles.Cardinality(), user.UserRoles().Cardinality())
		assert.NotNil(t, user.Person)
	}
}

func TestUser_CreateUser(t *testing.T) {
	ctx := context.Background()
	for _, v := range mock_data.NewUsers {
		person, err := identity.NewPerson(
			v.Person.Id,
			v.Person.Email,
			v.Person.ProfileImg,
			v.Person.Genre,
			v.Person.Pronoun,
			v.Person.FirstName,
			v.Person.LastName,
			v.Person.Phone,
			v.Person.BirthDate,
			v.Person.Address,
		)
		assert.Nil(t, err)
		c := 0
		domain.DomainEventPublisherInstance().Subscribe(
			domain.DomainEventSubscriber{
				HandleEvent: func(event domain.DomainEvent) error {
					if _, ok := event.(*identity.UserCreated); !ok {
						t.Fatalf("Incorrect EventType in Subscriber")
					}
					c++
					return nil
				},
				SubscribedToEventType: func() reflect.Type {
					return reflect.TypeOf(&identity.UserCreated{})
				},
			},
		)
		user, err := identity.CreateUser(
			ctx,
			v.Id,
			v.UserName,
			v.Password,
			mock_data.EmailVerificationToken,
			v.RegisterDate,
			person,
		)
		assert.Nil(t, err)
		assert.Equal(t, v.Id, user.Id())
		assert.Equal(t, v.UserName, user.UserName())
		assert.Equal(t, v.Password, user.Password())
		assert.Equal(t, v.RegisterDate.StringRepresentation(), user.RegisterDate().StringRepresentation())
		assert.Equal(t, v.IsConfirmed, user.IsConfirmed())
		assert.Equal(t, 0, user.UserRoles().Cardinality())
		assert.NotNil(t, user.Person)
		assert.Equal(t, 1, c)
	}
}

func TestUser_AddRole(t *testing.T) {
	ctx := context.Background()
	for _, v := range mock_data.NewUsers {
		person, err := identity.NewPerson(
			v.Person.Id,
			v.Person.Email,
			v.Person.ProfileImg,
			v.Person.Genre,
			v.Person.Pronoun,
			v.Person.FirstName,
			v.Person.LastName,
			v.Person.Phone,
			v.Person.BirthDate,
			v.Person.Address,
		)
		assert.Nil(t, err)
		user, err := identity.NewUser(
			v.Id,
			v.UserName,
			v.Password,
			v.RegisterDate,
			person,
			v.IsConfirmed,
			v.UserRoles,
		)
		assert.Nil(t, err)
		user.AddRole(ctx, v.RoleToAdd.GetRole(), v.RoleToAdd.EnterpriseId(), v.RoleToAdd.EnterpriseName())
		assert.Equal(t, 1, user.UserRoles().Cardinality())
	}
}
