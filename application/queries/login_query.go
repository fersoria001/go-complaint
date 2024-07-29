package queries

import (
	"context"
	"go-complaint/application"
	"go-complaint/application/application_services"
	"go-complaint/dto"
	"go-complaint/infrastructure/cache"
	"go-complaint/infrastructure/persistence/finders/find_user"
	"go-complaint/infrastructure/persistence/repositories"
)

type LoginQuery struct {
	Token            string
	ConfirmationCode int
}

func NewLoginQuery(token string, confirmationCode int) *LoginQuery {
	return &LoginQuery{
		Token:            token,
		ConfirmationCode: confirmationCode,
	}
}

func (q LoginQuery) Execute(ctx context.Context) (*dto.JWTToken, error) {
	confirmation, ok := cache.InMemoryInstance().Get(q.Token)
	if !ok {
		return nil, ErrConfirmationNotFound
	}
	loginConfirmation, ok := confirmation.(*application.LoginConfirmation)
	if !ok {
		return nil, ErrWrongTypeAssertion
	}
	if loginConfirmation.IsConfirmed() {
		cache.InMemoryInstance().Delete(q.Token)
		return nil, ErrConfirmationAlreadyDone
	}
	code, err := application_services.JWTApplicationServiceInstance().ParseConfirmationCode(q.Token)
	if err != nil {
		return nil, err
	}
	if code.Code != q.ConfirmationCode {
		err := loginConfirmation.RetryConfirmation()
		if err != nil {
			return nil, err
		}
		return nil, ErrConfirmationCodeNotMatch
	}
	loginConfirmation.Confirm()
	cache.InMemoryInstance().Delete(q.Token)
	reg := repositories.MapperRegistryInstance()
	userRepository, ok := reg.Get("User").(repositories.UserRepository)
	if !ok {
		return nil, repositories.ErrWrongTypeAssertion
	}
	user, err := userRepository.Find(ctx, find_user.ByUsername(loginConfirmation.Email()))
	if err != nil {
		return nil, err
	}
	clientData := application_services.AuthorizationApplicationServiceInstance().ClientData(ctx)
	userDescriptor := dto.NewUserDescriptor(clientData, *user)
	token, err := application_services.JWTApplicationServiceInstance().GenerateJWTToken(userDescriptor)
	if err != nil {
		return nil, err
	}
	return &dto.JWTToken{Token: token.Token()}, nil
}
