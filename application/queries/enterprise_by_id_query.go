package queries

import (
	"context"
	"go-complaint/dto"
	"go-complaint/infrastructure/persistence/repositories"

	"github.com/google/uuid"
)

type EnterpriseByIdQuery struct {
	Id string `json:"id"`
}

func NewEnterpriseByIdQuery(id string) *EnterpriseByIdQuery {
	return &EnterpriseByIdQuery{
		Id: id,
	}
}

func (q EnterpriseByIdQuery) Execute(ctx context.Context) (*dto.Enterprise, error) {
	repository, ok := repositories.MapperRegistryInstance().Get("Enterprise").(repositories.EnterpriseRepository)
	if !ok {
		return nil, ErrWrongTypeAssertion
	}
	id, err := uuid.Parse(q.Id)
	if err != nil {
		return nil, err
	}
	e, err := repository.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return dto.NewEnterprise(e), nil
}
