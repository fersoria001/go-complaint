package commands

import (
	"context"
	"go-complaint/application/commands"
	"go-complaint/domain"
	"go-complaint/domain/model/identity"
	"go-complaint/infrastructure/persistence/finders/find_enterprise"
	"go-complaint/infrastructure/persistence/repositories"
	"go-complaint/test/mock_data"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegisterEnterpriseCommand_Execute(t *testing.T) {
	ctx := context.Background()
	userRepository, ok := repositories.MapperRegistryInstance().Get("User").(repositories.UserRepository)
	assert.True(t, ok)
	enterpriseRepository, ok := repositories.MapperRegistryInstance().Get("Enterprise").(repositories.EnterpriseRepository)
	assert.True(t, ok)
	recipientRepository, ok := repositories.MapperRegistryInstance().Get("Recipient").(repositories.RecipientRepository)
	assert.True(t, ok)
	for _, v := range mock_data.NewRegisterEnterprises {
		domain.DomainEventPublisherInstance().Reset()
		person, err := identity.NewPerson(
			v.Owner.Person.Id,
			v.Owner.Person.Email,
			v.Owner.Person.ProfileImg,
			v.Owner.Person.Genre,
			v.Owner.Person.Pronoun,
			v.Owner.Person.FirstName,
			v.Owner.Person.LastName,
			v.Owner.Person.Phone,
			v.Owner.Person.BirthDate,
			v.Owner.Person.Address,
		)
		assert.Nil(t, err)
		user, err := identity.NewUser(
			v.Owner.Id,
			v.Owner.UserName,
			v.Owner.Password,
			v.Owner.RegisterDate,
			person,
			v.Owner.IsConfirmed,
			v.Owner.UserRoles,
		)
		assert.Nil(t, err)
		err = userRepository.Save(ctx, user)
		assert.Nil(t, err)
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
		err = c.Execute(ctx)
		assert.Nil(t, err)
		dbE, err := enterpriseRepository.Find(ctx, find_enterprise.ByName(v.Name))
		assert.Nil(t, err)
		assert.NotNil(t, dbE)
		assert.Equal(t, v.Owner.Id, dbE.OwnerId())
		assert.Equal(t, v.Name, dbE.Name())
		assert.Equal(t, v.LogoImg, dbE.LogoIMG())
		assert.Equal(t, v.BannerImg, dbE.BannerIMG())
		assert.Equal(t, v.Website, dbE.Website())
		assert.Equal(t, v.Email, dbE.Email())
		assert.Equal(t, v.Phone, dbE.Phone())
		assert.Equal(t, v.Address.Country().ID(), dbE.Address().Country().ID())
		assert.Equal(t, v.Address.CountryState().ID(), dbE.Address().CountryState().ID())
		assert.Equal(t, v.Address.City().ID(), dbE.Address().City().ID())
		assert.Equal(t, v.Industry.Id, dbE.Industry().ID())
		assert.Equal(t, v.Industry.Name, dbE.Industry().Name())
		assert.Equal(t, v.FoundationDate.StringRepresentation(), dbE.FoundationDate().StringRepresentation())
		assert.Equal(t, v.Employees.Cardinality(), dbE.Employees().Cardinality())
		dbOwner, err := userRepository.Get(ctx, v.Owner.Id)
		assert.Nil(t, err)
		assert.Equal(t, v.Owner.Id, dbOwner.Id())
		assert.Equal(t, v.Owner.UserName, dbOwner.UserName())
		assert.Equal(t, v.Owner.Password, dbOwner.Password())
		assert.Equal(t, v.Owner.RegisterDate.StringRepresentation(), dbOwner.RegisterDate().StringRepresentation())
		assert.Equal(t, v.Owner.IsConfirmed, dbOwner.IsConfirmed())
		assert.Equal(t, 1, dbOwner.UserRoles().Cardinality())
		assert.NotNil(t, dbOwner.Person)
		recipient, err := recipientRepository.Get(ctx, dbE.Id())
		assert.Nil(t, err)
		assert.NotNil(t, recipient)
		t.Cleanup(func() {
			err = enterpriseRepository.Remove(ctx, v.Id)
			assert.Nil(t, err)
			err = userRepository.Remove(ctx, v.Owner.Id)
			assert.Nil(t, err)
			err = recipientRepository.Remove(ctx, recipient.Id())
			assert.Nil(t, err)
		})
	}
}
