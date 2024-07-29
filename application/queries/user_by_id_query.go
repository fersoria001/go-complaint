package queries

import (
	"context"
	"go-complaint/dto"
	"go-complaint/infrastructure/persistence/repositories"

	"github.com/google/uuid"
)

type UserByIdQuery struct {
	Id string `json:"id"`
}

func NewUserByIdQuery(id string) *UserByIdQuery {
	return &UserByIdQuery{
		Id: id,
	}
}

func (q UserByIdQuery) Execute(ctx context.Context) (*dto.User, error) {
	id, err := uuid.Parse(q.Id)
	if err != nil {
		return nil, err
	}
	userRepository, ok := repositories.MapperRegistryInstance().Get("User").(repositories.UserRepository)
	if !ok {
		return nil, ErrWrongTypeAssertion
	}
	user, err := userRepository.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return dto.NewUser(user), nil
}
