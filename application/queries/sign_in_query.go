package queries

import (
	"context"
	"go-complaint/application"
	"go-complaint/application/application_services"
	"go-complaint/dto"
	"go-complaint/infrastructure"
	"go-complaint/infrastructure/cache"
	"os"
)

type SignInQuery struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	RememberMe bool   `json:"rememberMe"`
}

func NewSignInQuery(username, password string, rememberMe bool) *SignInQuery {
	return &SignInQuery{
		Username:   username,
		Password:   password,
		RememberMe: rememberMe,
	}
}

func (siq *SignInQuery) Execute(ctx context.Context) (*dto.JWTToken, error) {
	env := os.Getenv("ENVIRONMENT")
	authSvc := infrastructure.AuthenticationServiceInstance()
	jwtSvc := application_services.JWTApplicationServiceInstance()
	err := authSvc.AuthenticateUser(
		ctx,
		siq.Username,
		siq.Password,
		siq.RememberMe,
	)
	if err != nil {
		return nil, err
	}
	confirmationCode := application.CreateConfirmationCode()
	t, err := jwtSvc.GenerateJWTToken(confirmationCode)
	if err != nil {
		return nil, err
	}
	confirmation := application.NewLoginConfirmation(siq.Username, t, false)
	switch env {
	case "PROD":
		// 	commands.SendEmailCommand{
		// 		ToEmail:          user.Email(),
		// 		ToName:           user.FullName(),
		// 		ConfirmationCode: userSignedIn.ConfirmationCode(),
		// 	}.VerifySignIn(ctx)
	case "DEV":
	}
	cache.InMemoryInstance().Set(t.Token(), confirmation)
	return &dto.JWTToken{
		Token: t.Token(),
	}, nil
}
