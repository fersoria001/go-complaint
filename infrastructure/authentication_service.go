package infrastructure

import (
	"context"
	"go-complaint/domain/model/identity"
	"go-complaint/infrastructure/persistence/finders/find_user"
	"go-complaint/infrastructure/persistence/repositories"
	"sync"
)

var authenticationServiceInstance *AuthenticationService
var authenticationServiceOnce sync.Once

func AuthenticationServiceInstance() *AuthenticationService {
	authenticationServiceOnce.Do(func() {
		mapper := repositories.MapperRegistryInstance().Get("User")
		repository, ok := mapper.(repositories.UserRepository)
		if !ok {
			panic("Error")
		}
		authenticationServiceInstance = NewAuthenticationService(repository)
	})
	return authenticationServiceInstance
}

type AuthenticationService struct {
	repository repositories.UserRepository
}

func NewAuthenticationService(repository repositories.UserRepository) *AuthenticationService {
	return &AuthenticationService{
		repository: repository,
	}
}

func (is AuthenticationService) AuthenticateUser(
	ctx context.Context,
	userName,
	password string,
	rememberMe bool,
) error {
	var (
		err  error
		user *identity.User
	)
	user, err = is.repository.Find(ctx, find_user.ByUsername(userName))
	if err != nil {
		return err
	}

	err = EncryptionServiceInstance().Compare(user.Password(), password)
	if err != nil {
		return err
	}
	return nil
}
