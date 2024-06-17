package commands_test

import (
	"context"
	"go-complaint/application/commands"
	"go-complaint/application/queries"
	"go-complaint/domain"
	"go-complaint/domain/model/identity"
	"go-complaint/infrastructure/persistence/repositories"
	"go-complaint/tests"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestUserCommand_RegisterUser(t *testing.T) {
	// Arrange
	ctx := context.Background()
	mapper := repositories.MapperRegistryInstance().Get("User")
	repository, ok := mapper.(repositories.UserRepository)
	if !ok {
		t.Error("Error")
	}
	// Act
	command := tests.UserRegisterCommands["1"]
	time.Sleep(1 * time.Second)
	err := command.Register(ctx)
	assert.Nil(t, err)

	dbUser, err := repository.Get(ctx, command.Email)
	assert.Nil(t, err)

	assert.NotNil(t, dbUser)
	assert.Equal(t, command.Email, dbUser.Email())
	assert.Equal(t, command.FirstName, dbUser.FirstName())
	assert.Equal(t, command.LastName, dbUser.LastName())
	assert.Equal(t, command.BirthDate, dbUser.BirthDate().StringRepresentation())
	assert.Equal(t, false, dbUser.IsConfirmed())
	assert.Equal(t, command.Phone, dbUser.Phone())
	assert.Equal(t, command.CountryID, dbUser.Address().Country().ID())
	assert.Equal(t, command.CountryStateID, dbUser.Address().CountryState().ID())
	assert.Equal(t, command.CityID, dbUser.Address().City().ID())

	// Assert
}

func TestUserCommand_VerifyEmail(t *testing.T) {
	// Arrange
	ctx := context.Background()
	mapper := repositories.MapperRegistryInstance().Get("User")
	userRepository, ok := mapper.(repositories.UserRepository)
	if !ok {
		t.Error("Error")
	}
	command := tests.UserRegisterAndVerifyCommands["2"]
	domain.DomainEventPublisherInstance().Subscribe(
		domain.DomainEventSubscriber{
			HandleEvent: func(event domain.DomainEvent) error {
				if castedEvent, ok := event.(*identity.UserCreated); ok {
					verificationToken := castedEvent.ConfirmationToken()
					tests.InMemoryCacheInstance().Set(
						command.Email,
						verificationToken,
					)
					return nil
				}
				return nil
			},
			SubscribedToEventType: func() reflect.Type {
				return reflect.TypeOf(&identity.UserCreated{})
			},
		},
	)
	// Act
	err := command.Register(ctx)
	time.Sleep(1 * time.Second)
	assert.Nil(t, err)
	token, ok := tests.InMemoryCacheInstance().Get(command.Email)
	assert.True(t, ok)
	strToken, ok := token.(string)
	assert.True(t, ok)
	command.EmailVerificationToken = strToken
	err = command.VerifyEmail(ctx)
	assert.Nil(t, err)
	verifiedUser, err := userRepository.Get(ctx, command.Email)
	assert.Nil(t, err)
	assert.NotNil(t, verifiedUser)
	assert.True(t, verifiedUser.IsConfirmed())
	tests.InMemoryCacheInstance().Delete(command.Email)

}

func TestUserCommand_RecoverPassword(t *testing.T) {
	// Arrange
	ctx := context.Background()
	mapper := repositories.MapperRegistryInstance().Get("User")
	userRepository, ok := mapper.(repositories.UserRepository)
	if !ok {
		t.Error("Error")
	}
	userCommand := tests.UserRegisterAndVerifyCommands["3"]
	domain.DomainEventPublisherInstance().Subscribe(
		domain.DomainEventSubscriber{
			HandleEvent: func(event domain.DomainEvent) error {
				if castedEvent, ok := event.(*identity.UserCreated); ok {
					verificationToken := castedEvent.ConfirmationToken()
					tests.InMemoryCacheInstance().Set(
						userCommand.Email,
						verificationToken,
					)
					return nil
				}
				return nil
			},
			SubscribedToEventType: func() reflect.Type {
				return reflect.TypeOf(&identity.UserCreated{})
			},
		},
	)
	err := userCommand.Register(ctx)
	time.Sleep(1 * time.Second)
	token, ok := tests.InMemoryCacheInstance().Get(userCommand.Email)
	strToken, ok1 := token.(string)
	userCommand.EmailVerificationToken = strToken
	err1 := userCommand.VerifyEmail(ctx)
	verifiedUser, err2 := userRepository.Get(ctx, userCommand.Email)
	userCommand.Email = verifiedUser.Email()
	userQuery := queries.UserQuery{
		Email:    userCommand.Email,
		Password: userCommand.Password,
	}
	// Act
	err3 := userCommand.RecoverPassword(ctx)
	_, err4 := userQuery.Login(ctx)
	// Assert
	assert.Nil(t, err)
	assert.True(t, ok)
	assert.NotNil(t, token)
	assert.True(t, ok1)
	assert.NotNil(t, strToken)
	assert.Nil(t, err1)
	assert.Nil(t, err2)
	assert.NotNil(t, verifiedUser)
	assert.True(t, verifiedUser.IsConfirmed())
	tests.InMemoryCacheInstance().Delete(userCommand.Email)
	assert.Nil(t, err3)
	assert.NotNil(t, err4)
}

func TestUserCommand_ChangePassword(t *testing.T) {
	// Arrange
	ctx := context.Background()
	mapper := repositories.MapperRegistryInstance().Get("User")
	userRepository, ok := mapper.(repositories.UserRepository)
	if !ok {
		t.Error("Error")
	}
	err := tests.UserRegisterAndVerifyCommands["4"].Register(ctx)
	assert.Nil(t, err)
	verifiedUser, err := userRepository.Get(ctx, tests.UserRegisterAndVerifyCommands["4"].Email)
	if err != nil {
		t.Fatal("Error", err)
	}
	userQuery := queries.UserQuery{
		Email:    verifiedUser.Email(),
		Password: "Password1",
	}
	userCommand := commands.UserCommand{
		Email:       verifiedUser.Email(),
		Password:    "Password2",
		OldPassword: "Password1",
	}
	// Act
	err1 := userCommand.ChangePassword(ctx)
	_, err2 := userQuery.SignIn(ctx)
	userQuery.Password = "Password2"
	_, err3 := userQuery.SignIn(ctx)
	// Assert
	assert.Nil(t, err)
	assert.Nil(t, err1)
	assert.NotNil(t, err2)
	assert.Nil(t, err3)
}

func TestUserCommand_UpdateProfile(t *testing.T) {
	// Arrange
	ctx := context.Background()
	mapper := repositories.MapperRegistryInstance().Get("User")
	userRepository, ok := mapper.(repositories.UserRepository)
	if !ok {
		t.Error("Error")
	}
	verifiedUser, err := userRepository.Get(ctx, tests.UserRegisterAndVerifyCommands["2"].Email)
	if err != nil {
		t.Fatal("Error", err)
	}
	userCommand := commands.UserCommand{
		Email:          verifiedUser.Email(),
		Pronoun:        "she",
		Gender:         "female",
		FirstName:      "Fernanda",
		LastName:       "Landa",
		ProfileIMG:     "/cool-direct-image-url.jpg",
		Phone:          "10987654321",
		CountryID:      15,
		CountryStateID: 2065,
		CityID:         1602,
	}
	// Act
	//Assert
	t.Run("pronoun", func(t *testing.T) {
		userCommand.UpdateType = "pronoun"
		err1 := userCommand.UpdatePersonalData(ctx)
		assert.Nil(t, err1)
	})
	t.Run("gender", func(t *testing.T) {
		userCommand.UpdateType = "gender"
		err1 := userCommand.UpdatePersonalData(ctx)
		assert.Nil(t, err1)
	})
	t.Run("firstName", func(t *testing.T) {
		userCommand.UpdateType = "firstName"
		err1 := userCommand.UpdatePersonalData(ctx)
		assert.Nil(t, err1)
	})
	t.Run("lastName", func(t *testing.T) {
		userCommand.UpdateType = "lastName"
		err1 := userCommand.UpdatePersonalData(ctx)
		assert.Nil(t, err1)
	})
	t.Run("profileIMG", func(t *testing.T) {
		userCommand.UpdateType = "profileIMG"
		err1 := userCommand.UpdatePersonalData(ctx)
		assert.Nil(t, err1)
	})
	t.Run("phone", func(t *testing.T) {
		userCommand.UpdateType = "phone"
		err1 := userCommand.UpdatePersonalData(ctx)
		assert.Nil(t, err1)
	})
	t.Run("country", func(t *testing.T) {
		userCommand.UpdateType = "country"
		err1 := userCommand.UpdatePersonalData(ctx)
		assert.Nil(t, err1)
	})
	t.Run("countryState", func(t *testing.T) {
		userCommand.UpdateType = "countryState"
		err1 := userCommand.UpdatePersonalData(ctx)
		assert.Nil(t, err1)
	})
	t.Run("city", func(t *testing.T) {
		userCommand.UpdateType = "city"
		err1 := userCommand.UpdatePersonalData(ctx)
		assert.Nil(t, err1)
	})
}
