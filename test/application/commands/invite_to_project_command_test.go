package commands_test

import (
	"context"
	"go-complaint/application/commands"
	"go-complaint/domain"
	"go-complaint/infrastructure/persistence/finders/find_enterprise"
	"go-complaint/infrastructure/persistence/finders/find_user"
	"go-complaint/infrastructure/persistence/repositories"
	"go-complaint/test/mock_data"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInviteToProjectCommand_Setup(t *testing.T) {
	ctx := context.Background()
	userRepository, ok := repositories.MapperRegistryInstance().Get("User").(repositories.UserRepository)
	assert.True(t, ok)
	enterpriseRepository, ok := repositories.MapperRegistryInstance().Get("Enterprise").(repositories.EnterpriseRepository)
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
		user, err := userRepository.Find(ctx, find_user.ByUsername(v.UserName))
		assert.Nil(t, err)
		assert.NotNil(t, user)
		userRecipient, err := recipientRepository.Get(ctx, user.Id())
		assert.Nil(t, err)
		assert.NotNil(t, userRecipient)
	}
	for _, v := range mock_data.NewRegisterEnterprises {
		domain.DomainEventPublisherInstance().Reset()
		c := commands.NewRegisterEnterpriseCommand(
			v.Owner.Id.String(),
			v.Name,
			v.LogoImg,
			v.BannerImg,
			v.Website,
			v.Email,
			v.Phone,
			v.FoundationDate.StringRepresentation(),
			v.Industry.Id,
			v.Address.Country().ID(),
			v.Address.CountryState().ID(),
			v.Address.City().ID(),
		)
		err := c.Execute(ctx)
		assert.Nil(t, err)
		dbE, err := enterpriseRepository.Find(ctx, find_enterprise.ByName(v.Name))
		assert.Nil(t, err)
		recipient, err := recipientRepository.Get(ctx, dbE.Id())
		assert.Nil(t, err)
		assert.NotNil(t, recipient)
	}
}

func TestInviteToProjectCommand_Execute(t *testing.T) {}
