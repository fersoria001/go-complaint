package queries

import (
	"context"
	"go-complaint/dto"
	"go-complaint/infrastructure/persistence/finders/find_user"
	"go-complaint/infrastructure/persistence/repositories"
)

type UserByUsernameQuery struct {
	Username string `json:"username"`
}

func NewUserByUsernameQuery(username string) *UserByUsernameQuery {
	return &UserByUsernameQuery{
		Username: username,
	}
}

func (q UserByUsernameQuery) Execute(ctx context.Context) (*dto.User, error) {
	userRepository, ok := repositories.MapperRegistryInstance().Get("User").(repositories.UserRepository)
	if !ok {
		return nil, ErrWrongTypeAssertion
	}
	user, err := userRepository.Find(ctx, find_user.ByUsername(q.Username))
	if err != nil {
		return nil, err
	}
	return dto.NewUser(user), nil
}
