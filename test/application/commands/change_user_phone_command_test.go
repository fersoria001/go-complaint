package commands_test

import (
	"context"
	"go-complaint/application/commands"
	"go-complaint/domain"
	"go-complaint/infrastructure/persistence/finders/find_user"
	"go-complaint/infrastructure/persistence/repositories"
	"go-complaint/test/mock_data"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChangeUserPhoneCommand_Setup(t *testing.T) {
	ctx := context.Background()
	repository, ok := repositories.MapperRegistryInstance().Get("User").(repositories.UserRepository)
	assert.True(t, ok)
	recipientRepository, ok := repositories.MapperRegistryInstance().Get("Recipient").(repositories.RecipientRepository)
	assert.True(t, ok)
	for _, v := range mock_data.NewUsers {
		domain.DomainEventPublisherInstance().Reset()
		c := commands.NewRegisterUserCommand(
			v.UserName,
			v.Password,
			v.Person.FirstName,
			v.Person.LastName,
			v.Person.Phone,
			v.Person.Phone,
			v.Person.BirthDate.StringRepresentation(),
			v.Person.Phone,
			v.Person.ProfileImg,
			v.Person.Address.Country().ID(),
			v.Person.Address.CountryState().ID(),
			v.Person.Address.City().ID(),
		)
		err := c.Execute(ctx)
		assert.Nil(t, err)
		user, err := repository.Find(ctx, find_user.ByUsername(v.UserName))
		assert.Nil(t, err)
		assert.NotNil(t, user)
		userRecipient, err := recipientRepository.Get(ctx, user.Id())
		assert.Nil(t, err)
		assert.NotNil(t, userRecipient)

	}
}

func TestChangeUserPhoneCommand_Execute(t *testing.T) {
	TestChangeUserPhoneCommand_Setup(t)
	ctx := context.Background()
	repository, ok := repositories.MapperRegistryInstance().Get("User").(repositories.UserRepository)
	assert.True(t, ok)
	for _, v := range mock_data.NewUsers {
		user, err := repository.Find(ctx, find_user.ByUsername(v.UserName))
		assert.Nil(t, err)
		p := user.Phone()
		newP := "098765432119"
		assert.NotEqual(t, p, newP)
		c := commands.NewChangeUserPhoneCommand(user.Id().String(), newP)
		err = c.Execute(ctx)
		assert.Nil(t, err)
		user, err = repository.Find(ctx, find_user.ByUsername(v.UserName))
		assert.Nil(t, err)
		assert.NotEqual(t, p, user.Phone())
		assert.Equal(t, newP, user.Phone())
	}

	t.Cleanup(func() {
		ctx := context.Background()
		repository, ok := repositories.MapperRegistryInstance().Get("User").(repositories.UserRepository)
		assert.True(t, ok)
		recipientRepository, ok := repositories.MapperRegistryInstance().Get("Recipient").(repositories.RecipientRepository)
		assert.True(t, ok)
		for _, v := range mock_data.NewUsers {
			user, err := repository.Find(ctx, find_user.ByUsername(v.UserName))
			assert.Nil(t, err)
			err = repository.Remove(ctx, user.Id())
			assert.Nil(t, err)
			err = recipientRepository.Remove(ctx, user.Id())
			assert.Nil(t, err)
		}
	})
}
