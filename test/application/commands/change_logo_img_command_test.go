package commands_test

import (
	"context"
	"fmt"
	"go-complaint/application/commands"
	"go-complaint/domain"
	"go-complaint/domain/model/identity"
	"go-complaint/infrastructure/persistence/finders/find_enterprise"
	"go-complaint/infrastructure/persistence/repositories"
	"go-complaint/test/mock_data"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestChangeLogoImgCommand_Setup(t *testing.T) {
	ctx := context.Background()
	userRepository, ok := repositories.MapperRegistryInstance().Get("User").(repositories.UserRepository)
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
	}
}

func TestChangeLogoImgCommand_Execute(t *testing.T) {
	TestChangeLogoImgCommand_Setup(t)
	ctx := context.Background()
	userRepository, ok := repositories.MapperRegistryInstance().Get("User").(repositories.UserRepository)
	assert.True(t, ok)
	enterpriseRepository, ok := repositories.MapperRegistryInstance().Get("Enterprise").(repositories.EnterpriseRepository)
	assert.True(t, ok)
	recipientRepository, ok := repositories.MapperRegistryInstance().Get("Recipient").(repositories.RecipientRepository)
	assert.True(t, ok)
	for _, v := range mock_data.NewRegisterEnterprises {
		dbEnterprise, err := enterpriseRepository.Find(ctx, find_enterprise.ByName(v.Name))
		assert.Nil(t, err)
		aTxtFile, err := os.CreateTemp(t.TempDir(), "a.txt")
		require.NoError(t, err)
		defer aTxtFile.Close()
		aTxtFile.WriteString(`test`)
		c := commands.NewChangeLogoImageCommand(
			dbEnterprise.Id().String(),
			"a.txt",
			aTxtFile,
		)
		err = c.Execute(ctx)
		assert.Nil(t, err)
		dbEnterprise, err = enterpriseRepository.Find(ctx, find_enterprise.ByName(v.Name))
		assert.Nil(t, err)
		dns := os.Getenv("DNS")
		resource := fmt.Sprintf("%s/%s", "logo_img", c.FileName)
		url := fmt.Sprintf("%s/%s", dns, resource)
		assert.Equal(t, url, dbEnterprise.LogoIMG())
	}
	t.Cleanup(func() {
		for _, v := range mock_data.NewRegisterEnterprises {
			dbEnterprise, err := enterpriseRepository.Find(ctx, find_enterprise.ByName(v.Name))
			assert.Nil(t, err)
			err = enterpriseRepository.Remove(ctx, dbEnterprise.Id())
			assert.Nil(t, err)
			err = userRepository.Remove(ctx, dbEnterprise.OwnerId())
			assert.Nil(t, err)
			err = recipientRepository.Remove(ctx, dbEnterprise.Id())
			assert.Nil(t, err)
		}
	})

}
