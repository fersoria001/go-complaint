package queries_test

import (
	"context"
	"go-complaint/application/application_services"
	"go-complaint/application/commands"
	"go-complaint/application/queries"
	"go-complaint/domain"
	"go-complaint/infrastructure/cache"
	"go-complaint/infrastructure/persistence/repositories"
	"go-complaint/test/mock_data"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoginQuery_Setup(t *testing.T) {
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

func TestLoginQuery_Execute(t *testing.T) {
	TestLoginQuery_Setup(t)
	ctx := context.Background()
	for _, v := range mock_data.NewUsers {
		q := queries.NewSignInQuery(v.UserName, v.Password, false)
		token, err := q.Execute(ctx)
		assert.Nil(t, err)
		assert.NotNil(t, token.Token)
		_, ok := cache.InMemoryInstance().Get(token.Token)
		assert.True(t, ok)
		q1 := queries.NewLoginQuery(token.Token, 9999999)
		loginToken, err := q1.Execute(ctx)
		assert.Nil(t, err)
		assert.NotNil(t, loginToken.Token)
		_, ok = cache.InMemoryInstance().Get(token.Token)
		assert.False(t, ok)
		svc := application_services.AuthorizationApplicationServiceInstance()
		_, err = svc.Authorize(ctx, loginToken.Token)
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
