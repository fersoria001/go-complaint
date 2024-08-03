package queries

import (
	"context"

	"go-complaint/application/application_services"
	"go-complaint/dto"
	"go-complaint/erros"
	"go-complaint/infrastructure/persistence/finders/find_user"
	"go-complaint/infrastructure/persistence/repositories"
)

type UserQuery struct {
	Email            string `json:"email"`
	Password         string `json:"password"`
	RememberMe       bool   `json:"rememberMe"`
	Token            string `json:"token"`
	ConfirmationCode int    `json:"confirmation_code"`
	EventID          string `json:"event_id"`
}

// avoid using this
func (userQuery UserQuery) UserDescriptor(
	ctx context.Context,
) (*dto.UserDescriptor, error) {
	if userQuery.Email == "" {
		return nil, &erros.ValueNotFoundError{}
	}
	clientData := application_services.AuthorizationApplicationServiceInstance().ClientData(ctx)
	user, err := repositories.MapperRegistryInstance().Get("User").(repositories.UserRepository).Find(ctx, find_user.ByUsername(userQuery.Email))
	if err != nil {
		return nil, err
	}
	return dto.NewUserDescriptor(clientData, *user), nil
}

// func (userQuery UserQuery) User(
// 	ctx context.Context,
// ) (dto.User, error) {
// 	if userQuery.Email == "" {
// 		return dto.User{}, ErrNilValue
// 	}
// 	mapper := repositories.MapperRegistryInstance().Get("User")
// 	userRepository, ok := mapper.(repositories.UserRepository)
// 	if !ok {
// 		return dto.User{}, repositories.ErrWrongTypeAssertion
// 	}
// 	user, err := userRepository.Find(ctx, find_user.ByUsername(userQuery.Email))
// 	if err != nil {
// 		return dto.User{}, err
// 	}
// 	if !user.IsConfirmed() {
// 		return dto.User{}, ErrUserNotConfirmed
// 	}
// 	return dto.NewUser(*user), nil
// }
