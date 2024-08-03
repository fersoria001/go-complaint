package commands_test

import (
	"context"
	"go-complaint/application"
	"go-complaint/application/commands"
	"go-complaint/domain"
	"go-complaint/infrastructure/persistence/finders/find_all_hiring_proccesses"
	"go-complaint/infrastructure/persistence/finders/find_enterprise"
	"go-complaint/infrastructure/persistence/finders/find_user"
	"go-complaint/infrastructure/persistence/repositories"
	"go-complaint/test/mock_data"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInviteToProjectCommand_Setup(t *testing.T) {
	ctx := context.Background()
	reg := repositories.MapperRegistryInstance()
	userRepository, ok := reg.Get("User").(repositories.UserRepository)
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
	}
	for _, v := range mock_data.NewRegisterEnterprises {
		domain.DomainEventPublisherInstance().Reset()
		owner, err := userRepository.Find(ctx, find_user.ByUsername(v.Owner.UserName))
		assert.Nil(t, err)
		c := commands.NewRegisterEnterpriseCommand(
			owner.Id().String(),
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

func TestInviteToProjectCommand_Execute(t *testing.T) {
	TestInviteToProjectCommand_Setup(t)
	ctx := context.Background()
	hiringProccessRepository, ok := repositories.MapperRegistryInstance().Get("HiringProccess").(repositories.HiringProccessRepository)
	assert.True(t, ok)
	userRepository, ok := repositories.MapperRegistryInstance().Get("User").(repositories.UserRepository)
	assert.True(t, ok)
	enterpriseRepository, ok := repositories.MapperRegistryInstance().Get("Enterprise").(repositories.EnterpriseRepository)
	assert.True(t, ok)
	recipientRepository, ok := repositories.MapperRegistryInstance().Get("Recipient").(repositories.RecipientRepository)
	assert.True(t, ok)
	pub := application.ApplicationMessagePublisherInstance()
	counter := 0
	for _, v := range mock_data.NewHiringProccesses {
		dbUser, err := userRepository.Find(ctx, find_user.ByUsername(v.User.SubjectEmail))
		assert.Nil(t, err)

		ch := make(chan application.ApplicationMessage)
		pub.Subscribe(&application.Subscriber{
			Id:   dbUser.Id().String(),
			Send: ch,
		})
		go func() {
			for {
				m := <-ch
				if m.Id() == dbUser.Id().String() {
					counter += 1
				}
			}
		}()
		dbEnterprise, err := enterpriseRepository.Find(ctx, find_enterprise.ByName(v.Enterprise.SubjectName))
		assert.Nil(t, err)

		dbEmitedBy, err := userRepository.Find(ctx, find_user.ByUsername(v.EmitedBy.SubjectEmail))
		assert.Nil(t, err)
		c := commands.NewInviteToProjectCommand(
			dbEnterprise.Id().String(),
			v.Role.String(),
			dbUser.Id().String(),
			dbEmitedBy.Id().String(),
		)
		err = c.Execute(ctx)
		assert.Nil(t, err)
	}
	for _, v := range mock_data.NewHiringProccesses {
		dbUser, err := userRepository.Find(ctx, find_user.ByUsername(v.User.SubjectEmail))
		assert.Nil(t, err)
		dbHp, err := hiringProccessRepository.FindAll(ctx, find_all_hiring_proccesses.ByUserId(dbUser.Id()))
		assert.Nil(t, err)
		assert.NotNil(t, dbHp)
		assert.GreaterOrEqual(t, len(dbHp), 1)
	}
	assert.GreaterOrEqual(t, counter, 1)
	t.Cleanup(func() {
		for _, v := range mock_data.NewHiringProccesses {
			dbUser, err := userRepository.Find(ctx, find_user.ByUsername(v.User.SubjectEmail))
			assert.Nil(t, err)
			dbHp, err := hiringProccessRepository.FindAll(ctx, find_all_hiring_proccesses.ByUserId(dbUser.Id()))
			assert.Nil(t, err)
			assert.NotNil(t, dbHp)
			assert.GreaterOrEqual(t, len(dbHp), 1)
			for _, dbH := range dbHp {
				err = hiringProccessRepository.Remove(ctx, dbH.Id())
				assert.Nil(t, err)
			}
		}

		for _, v := range mock_data.NewRegisterEnterprises {
			dbEnterprise, err := enterpriseRepository.Find(ctx, find_enterprise.ByName(v.Name))
			assert.Nil(t, err)
			err = enterpriseRepository.Remove(ctx, dbEnterprise.Id())
			assert.Nil(t, err)
			err = recipientRepository.Remove(ctx, dbEnterprise.Id())
			assert.Nil(t, err)
		}

		for _, v := range mock_data.NewUsers {
			dbUser, err := userRepository.Find(ctx, find_user.ByUsername(v.UserName))
			assert.Nil(t, err)
			err = userRepository.Remove(ctx, dbUser.Id())
			assert.Nil(t, err)
			err = recipientRepository.Remove(ctx, dbUser.Id())
			assert.Nil(t, err)
		}
	})
}
