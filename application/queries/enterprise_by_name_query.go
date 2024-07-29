package queries

import (
	"context"
	"go-complaint/dto"
	"go-complaint/infrastructure/persistence/finders/find_enterprise"
	"go-complaint/infrastructure/persistence/repositories"
)

type EnterpriseByNameQuery struct {
	Name string `json:"name"`
}

func NewEnterpriseByNameQuery(name string) *EnterpriseByNameQuery {
	return &EnterpriseByNameQuery{
		Name: name,
	}
}

func (q EnterpriseByNameQuery) Execute(ctx context.Context) (*dto.Enterprise, error) {
	repository, ok := repositories.MapperRegistryInstance().Get("Enterprise").(repositories.EnterpriseRepository)
	if !ok {
		return nil, ErrWrongTypeAssertion
	}
	e, err := repository.Find(ctx, find_enterprise.ByName(q.Name))
	if err != nil {
		return nil, err
	}
	return dto.NewEnterprise(e), nil
}
