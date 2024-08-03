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

func TestChangeUserAddressCommand_Setup(t *testing.T) {
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
			v.Person.FirstName,
			v.Person.FirstName,
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

func TestChangeUserAddressCommand_Execute(t *testing.T) {
	TestChangeUserAddressCommand_Setup(t)
	ctx := context.Background()
	repository, ok := repositories.MapperRegistryInstance().Get("User").(repositories.UserRepository)
	assert.True(t, ok)
	for _, v := range mock_data.NewUsers {
		user, err := repository.Find(ctx, find_user.ByUsername(v.UserName))
		assert.Nil(t, err)
		c := commands.NewChangeUserAddressCommand(user.Id().String(),
			mock_data.Country1.ID(), mock_data.CountryState1.ID(), mock_data.City1.ID())
		err = c.Execute(ctx)
		assert.Nil(t, err)
		user, err = repository.Find(ctx, find_user.ByUsername(v.UserName))
		assert.Nil(t, err)
		assert.Equal(t, mock_data.Country1.ID(), user.Address().Country().ID())
		assert.Equal(t, mock_data.CountryState1.ID(), user.Address().CountryState().ID())
		assert.Equal(t, mock_data.City1.ID(), user.Address().City().ID())
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
