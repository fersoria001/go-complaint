package infrastructure_test

import (
	"context"
	"go-complaint/application/commands"
	"go-complaint/domain"
	"go-complaint/infrastructure"
	"go-complaint/infrastructure/persistence/repositories"
	"go-complaint/test/mock_data"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuthenticationService_Setup(t *testing.T) {
	ctx := context.Background()
	for _, v := range mock_data.NewUsers {
		domain.DomainEventPublisherInstance().Reset()
		c := commands.NewRegisterUserCommand(
			v.UserName,
			v.Password,
			v.Person.FirstName,
			v.Person.LastName,
			v.Person.Genre,
			v.Person.Pronoun,
			v.Person.BirthDate.StringRepresentation(),
			v.Person.Phone,
			v.Person.ProfileImg,
			v.Person.Address.Country().ID(),
			v.Person.Address.CountryState().ID(),
			v.Person.Address.City().ID(),
		)
		err := c.Execute(ctx)
		assert.Nil(t, err)
	}
}

func TestAuthenticationService_AuthenticateUser(t *testing.T) {
	TestAuthenticationService_Setup(t)
	ctx := context.Background()
	svc := infrastructure.AuthenticationServiceInstance()
	for _, v := range mock_data.NewUsers {
		err := svc.AuthenticateUser(ctx, v.UserName, v.Password, false)
		assert.Nil(t, err)
	}
	t.Cleanup(func() {
		reg := repositories.MapperRegistryInstance()
		userRepository, ok := reg.Get("User").(repositories.UserRepository)
		assert.True(t, ok)
		recipientRepository, ok := reg.Get("Recipient").(repositories.RecipientRepository)
		assert.True(t, ok)
		for _, v := range mock_data.NewUsers {
			err := userRepository.Remove(ctx, v.Id)
			assert.Nil(t, err)
			err = recipientRepository.Remove(ctx, v.Id)
			assert.Nil(t, err)
		}
	})
}
