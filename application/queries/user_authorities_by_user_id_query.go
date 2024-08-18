package queries

import (
	"context"
	"go-complaint/dto"
	"go-complaint/infrastructure/persistence/repositories"

	"github.com/google/uuid"
)

type UserAuthoritiesByUserIdQuery struct {
	Id string `json:"id"`
}

func NewUserAuthoritiesByUserIdQuery(id string) *UserAuthoritiesByUserIdQuery {
	return &UserAuthoritiesByUserIdQuery{
		Id: id,
	}
}

func (q UserAuthoritiesByUserIdQuery) Execute(ctx context.Context) ([]dto.GrantedAuthority, error) {
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
	return dto.NewGrantedAuthorities(user.Authorities()), nil
}
