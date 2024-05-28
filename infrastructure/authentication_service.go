package infrastructure

import (
	"context"
	"go-complaint/domain/model/identity"
	"go-complaint/dto"
	"go-complaint/infrastructure/persistence/repositories"
)

type AuthenticationService struct {
	repository *repositories.UserRepository
}

func NewAuthenticationService(repository *repositories.UserRepository) *AuthenticationService {
	return &AuthenticationService{
		repository: repository,
	}
}

// this belongs to infrastructure layer
func (is *AuthenticationService) AuthenticateUser(ctx context.Context, email, password string,
	rememberMe bool) (*dto.UserDescriptor, error) {
	var (
		err  error
		user *identity.User
	)
	user, err = is.repository.Get(ctx, email)
	if err != nil {
		return nil, err
	}
	encryptionService := NewEncryptionService()
	err = encryptionService.Compare(user.Password(), password)
	if err != nil {
		return nil, err
	}
	userDescriptor, err := dto.NewUserDescriptor(
		ctx,
		user.Email(),
		user.Person().FullName(),
		user.ProfileIMG(),
		rememberMe)
	if err != nil {
		return nil, err
	}
	return userDescriptor, nil
}
