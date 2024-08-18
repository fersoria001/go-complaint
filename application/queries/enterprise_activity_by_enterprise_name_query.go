package queries

import (
	"context"
	"errors"
	"go-complaint/dto"
	"go-complaint/infrastructure/persistence/finders/find_all_enterprise_activity"
	"go-complaint/infrastructure/persistence/repositories"

	"github.com/jackc/pgx/v5"
)

type EnterpriseActivityByEnterpriseNameQuery struct {
	EnterpriseName string `json:"enterpriseName"`
}

func NewEnterpriseActivityByEnterpriseNameQuery(enterpriseName string) *EnterpriseActivityByEnterpriseNameQuery {
	return &EnterpriseActivityByEnterpriseNameQuery{
		EnterpriseName: enterpriseName,
	}
}

func (c EnterpriseActivityByEnterpriseNameQuery) Execute(ctx context.Context) ([]*dto.EnterpriseActivity, error) {
	reg := repositories.MapperRegistryInstance()
	r, ok := reg.Get("EnterpriseActivity").(repositories.EnterpriseActivityRepository)
	if !ok {
		return nil, ErrWrongTypeAssertion
	}
	results := make([]*dto.EnterpriseActivity, 0)
	dbEa, err := r.FindAll(ctx, find_all_enterprise_activity.ByEnterpriseName(c.EnterpriseName))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return results, nil
		}
		return nil, err
	}
	for i := range dbEa {
		results = append(results, dto.NewEnterpriseActivity(*dbEa[i]))
	}
	return results, nil
}
