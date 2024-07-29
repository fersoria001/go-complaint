package queries_test

import (
	"context"
	"go-complaint/application"
	"go-complaint/application/commands"
	"go-complaint/application/queries"
	"go-complaint/domain"
	"go-complaint/infrastructure/cache"
	"go-complaint/infrastructure/persistence/repositories"
	"go-complaint/test/mock_data"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSignInQuery_Setup(t *testing.T) {
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

func TestSignInQuery_Execute(t *testing.T) {
	TestSignInQuery_Setup(t)
	ctx := context.Background()
	for _, v := range mock_data.NewUsers {
		q := queries.NewSignInQuery(v.UserName, v.Password, false)
		token, err := q.Execute(ctx)
		assert.Nil(t, err)
		assert.NotNil(t, token.Token)
		cached, ok := cache.InMemoryInstance().Get(token.Token)
		assert.True(t, ok)
		assert.NotNil(t, cached)
		_, ok = cached.(*application.LoginConfirmation)
		assert.True(t, ok)
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
