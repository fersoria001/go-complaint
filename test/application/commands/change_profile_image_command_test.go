package commands

import (
	"context"
	"fmt"
	"go-complaint/application/commands"
	"go-complaint/domain"
	"go-complaint/infrastructure/persistence/finders/find_user"
	"go-complaint/infrastructure/persistence/repositories"
	"go-complaint/test/mock_data"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestChangeProfileImageCommand_Setup(t *testing.T) {
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
		user, err := repository.Find(ctx, find_user.ByUsername(v.UserName))
		assert.Nil(t, err)
		assert.NotNil(t, user)
		userRecipient, err := recipientRepository.Get(ctx, user.Id())
		assert.Nil(t, err)
		assert.NotNil(t, userRecipient)

	}
}

func TestChangeProfileImageCommand_Execute(t *testing.T) {
	TestChangeProfileImageCommand_Setup(t)
	ctx := context.Background()
	repository, ok := repositories.MapperRegistryInstance().Get("User").(repositories.UserRepository)
	assert.True(t, ok)
	for _, v := range mock_data.NewUsers {
		user, err := repository.Find(ctx, find_user.ByUsername(v.UserName))
		assert.Nil(t, err)
		aTxtFile, err := os.CreateTemp(t.TempDir(), "a.txt")
		require.NoError(t, err)
		defer aTxtFile.Close()
		aTxtFile.WriteString(`test`)
		c := commands.NewChangeProfileImageCommand(user.Id().String(), "a.txt", aTxtFile)
		err = c.Execute(ctx)
		assert.Nil(t, err)
		user, err = repository.Find(ctx, find_user.ByUsername(v.UserName))
		assert.Nil(t, err)
		dns := os.Getenv("DNS")
		resource := fmt.Sprintf("%s/%s", "profile_img", c.FileName)
		url := fmt.Sprintf("%s/%s", dns, resource)
		assert.Equal(t, url, user.ProfileIMG())
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
