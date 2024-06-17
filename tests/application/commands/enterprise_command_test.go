package commands_test

import (
	"context"
	"go-complaint/application"
	"go-complaint/application/application_services"
	"go-complaint/application/commands"
	"go-complaint/application/queries"
	"go-complaint/domain"
	"go-complaint/domain/model/identity"
	"go-complaint/erros"
	"go-complaint/infrastructure"
	"go-complaint/infrastructure/persistence/repositories"
	"go-complaint/tests"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func loginUser() (context.Context, error) {
	ctx := context.Background()
	userQuery := queries.UserQuery{
		Email:      tests.UserRegisterAndVerifyCommands["2"].Email,
		Password:   tests.UserRegisterAndVerifyCommands["2"].Password,
		RememberMe: true,
	}
	domain.DomainEventPublisherInstance().Subscribe(domain.DomainEventSubscriber{
		HandleEvent: func(event domain.DomainEvent) error {
			if userSignedIn, ok := event.(*identity.UserSignedIn); ok {
				userQuery.ConfirmationCode = userSignedIn.ConfirmationCode()
				return nil
			}
			return &erros.ValueNotFoundError{}
		},
		SubscribedToEventType: func() reflect.Type {
			return reflect.TypeOf(&identity.UserSignedIn{})
		},
	})
	unconfirmedToken, err := userQuery.SignIn(ctx)
	if err != nil {
		return nil, err
	}
	userQuery.Token = unconfirmedToken.Token
	jwtToken, err := userQuery.Login(ctx)
	if err != nil {
		return nil, err
	}
	userQuery.Token = jwtToken.Token
	ctx, err = application_services.AuthorizationApplicationServiceInstance().Authorize(
		ctx,
		jwtToken.Token)
	if err != nil {
		return nil, err
	}
	return ctx, nil
}
func TestEnterpriseCommand(t *testing.T) {
	// Arrange
	ctx, err := loginUser()
	assert.Nil(t, err)
	mapper := repositories.MapperRegistryInstance().Get("Enterprise")
	enterpriseRepository, ok := mapper.(repositories.EnterpriseRepository)
	assert.True(t, ok)
	assert.NotNil(t, enterpriseRepository)
	mapper = repositories.MapperRegistryInstance().Get("User")
	userRepository, ok := mapper.(repositories.UserRepository)
	assert.True(t, ok)
	assert.NotNil(t, userRepository)
	newEnterpriseCommand := tests.CreateEnterpriseCommands["Spoon company"]
	// Act
	err = newEnterpriseCommand.Register(ctx)
	assert.Nil(t, err)
	dbEnterprise, err := enterpriseRepository.Get(ctx, newEnterpriseCommand.Name)
	assert.Nil(t, err)
	user, err := userRepository.Get(ctx, dbEnterprise.Owner())
	assert.Nil(t, err)
	authorities := user.Authorities()
	assert.Equal(t, 1, len(authorities[dbEnterprise.Name()]))
}

func TestEnterpriseCommandCommand_UpdateProfile(t *testing.T) {
	// Arrange
	ctx := context.Background()
	mapper := repositories.MapperRegistryInstance().Get("Enterprise")
	enterpriseRepository, ok := mapper.(repositories.EnterpriseRepository)
	if !ok {
		t.Error("Error")
	}
	enterpriseCommand := commands.EnterpriseCommand{
		OwnerID:        tests.CreateEnterpriseCommands["Spoon company"].OwnerID,
		Name:           "Spoon company",
		LogoIMG:        "/random-img.jpg",
		BannerIMG:      "/random-banner.jpg",
		Website:        "www.spooncompany.sr",
		Email:          "spoon-company-sr@enterprise.com",
		Phone:          "10987654321",
		CountryID:      210,
		CountryStateID: 2840,
		CityID:         104833,
		IndustryID:     99,
	}

	// Act
	t.Run("logoIMG", func(t *testing.T) {
		enterpriseCommand.UpdateType = "logoIMG"
		err1 := enterpriseCommand.UpdateEnterprise(ctx)
		assert.Nil(t, err1)
	})
	t.Run("bannerIMG", func(t *testing.T) {
		enterpriseCommand.UpdateType = "bannerIMG"
		err1 := enterpriseCommand.UpdateEnterprise(ctx)
		assert.Nil(t, err1)
	})
	t.Run("website", func(t *testing.T) {
		enterpriseCommand.UpdateType = "website"
		err1 := enterpriseCommand.UpdateEnterprise(ctx)
		assert.Nil(t, err1)
	})
	t.Run("email", func(t *testing.T) {
		enterpriseCommand.UpdateType = "email"
		err1 := enterpriseCommand.UpdateEnterprise(ctx)
		assert.Nil(t, err1)
	})
	t.Run("phone", func(t *testing.T) {
		enterpriseCommand.UpdateType = "phone"
		err1 := enterpriseCommand.UpdateEnterprise(ctx)
		assert.Nil(t, err1)
	})
	t.Run("country", func(t *testing.T) {
		enterpriseCommand.UpdateType = "country"
		err1 := enterpriseCommand.UpdateEnterprise(ctx)
		assert.Nil(t, err1)
	})
	t.Run("countryState", func(t *testing.T) {
		enterpriseCommand.UpdateType = "countryState"
		err1 := enterpriseCommand.UpdateEnterprise(ctx)
		assert.Nil(t, err1)
	})
	t.Run("city", func(t *testing.T) {
		enterpriseCommand.UpdateType = "city"
		err1 := enterpriseCommand.UpdateEnterprise(ctx)
		assert.Nil(t, err1)
	})
	dbEnterprise, _ := enterpriseRepository.Get(ctx, enterpriseCommand.Name)
	// Assert
	assert.Equal(t, enterpriseCommand.LogoIMG, dbEnterprise.LogoIMG())
	assert.Equal(t, enterpriseCommand.BannerIMG, dbEnterprise.BannerIMG())
	assert.Equal(t, enterpriseCommand.Website, dbEnterprise.Website())
	assert.Equal(t, enterpriseCommand.Email, dbEnterprise.Email())
	assert.Equal(t, enterpriseCommand.Phone, dbEnterprise.Phone())
	assert.Equal(t, enterpriseCommand.CountryID, dbEnterprise.Address().Country().ID())
	assert.Equal(t, enterpriseCommand.CountryStateID, dbEnterprise.Address().CountryState().ID())
	assert.Equal(t, enterpriseCommand.CityID, dbEnterprise.Address().City().ID())
}

func TestEnterpriseCommandCommand_InviteToProject(t *testing.T) {
	// Arrange
	ctx, err := loginUser()
	application.EventProcessorInstance().ResetDomainEventPublisher()
	assert.Nil(t, err)
	enterpriseCommand := commands.EnterpriseCommand{
		Name:      tests.CreateEnterpriseCommands["Spoon company"].Name,
		Position:  "MANAGER",
		ProposeTo: tests.UserRegisterAndVerifyCommands["1"].Email,
	}
	queue := infrastructure.EmailServiceInstance()
	// Act
	err = enterpriseCommand.InviteToProject(ctx)
	queue.SendAll(ctx)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, 1, len(queue.SentLog()))
	assert.Equal(t, 0, queue.Queued())
}
