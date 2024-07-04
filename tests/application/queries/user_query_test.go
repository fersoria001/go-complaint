package queries_test

import (
	"context"
	"go-complaint/application/application_services"
	"go-complaint/application/queries"
	"go-complaint/domain"
	"go-complaint/domain/model/identity"
	"go-complaint/erros"
	"go-complaint/tests"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserQuery_SignIn(t *testing.T) {

	// Arrange
	ctx := context.Background()
	userQuery := queries.UserQuery{
		Email:      "bercho001@gmail.com",
		Password:   "Password1",
		RememberMe: true,
	}
	// Act
	jwtToken, err := userQuery.SignIn(ctx)
	// Assert
	assert.Nil(t, err)
	assert.NotNil(t, jwtToken)
}

func TestUserQuery_TwoStepLogin(t *testing.T) {
	// Arrange
	ctx := context.Background()

	userQuery := queries.UserQuery{
		Email:      "bercho001@gmail.com",
		Password:   "Password1",
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
	userQuery.Token = unconfirmedToken.Token
	// Act
	jwtToken, err1 := userQuery.Login(ctx)
	// Assert
	assert.Nil(t, err)
	assert.Nil(t, err1)
	assert.NotNil(t, jwtToken)
	t.Logf(jwtToken.Token)
}

func TestUserQuery_User(t *testing.T) {
	// Arrange
	ctx := context.Background()
	userQuery := queries.UserQuery{
		Email: tests.UserRegisterAndVerifyCommands["2"].Email,
	}
	// Act
	user, err := userQuery.User(ctx)
	// Assert
	assert.Nil(t, err)
	assert.NotNil(t, user)
}
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
func TestUserHiringInvitations(t *testing.T) {
	// Arrange
	ctx, err := loginUser()
	assert.Nil(t, err)
	userQuery := queries.UserQuery{Email: tests.UserRegisterAndVerifyCommands["2"].Email}
	invitations, err := userQuery.HiringInvitations(ctx)
	assert.Nil(t, err)
	assert.NotNil(t, invitations)

}
