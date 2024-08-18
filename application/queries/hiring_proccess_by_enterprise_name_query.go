package queries

import (
	"context"
	"errors"
	"go-complaint/dto"
	"go-complaint/infrastructure/persistence/finders/find_all_hiring_proccesses"
	"go-complaint/infrastructure/persistence/finders/find_enterprise"
	"go-complaint/infrastructure/persistence/repositories"

	"github.com/jackc/pgx/v5"
)

type HiringProccessByEnterpriseNameQuery struct {
	EnterpriseName string `json:"enterpriseName"`
}

func NewHiringProccessByEnterpriseNameQuery(enterpriseName string) *HiringProccessByEnterpriseNameQuery {
	return &HiringProccessByEnterpriseNameQuery{
		EnterpriseName: enterpriseName,
	}
}

func (q HiringProccessByEnterpriseNameQuery) Execute(ctx context.Context) ([]*dto.HiringProccess, error) {
	reg := repositories.MapperRegistryInstance()
	enterpriseRepository, ok := reg.Get("Enterprise").(repositories.EnterpriseRepository)
	if !ok {
		return nil, ErrWrongTypeAssertion
	}
	r, ok := reg.Get("HiringProccess").(repositories.HiringProccessRepository)
	if !ok {
		return nil, ErrWrongTypeAssertion
	}
	result := make([]*dto.HiringProccess, 0)
	dbE, err := enterpriseRepository.Find(ctx, find_enterprise.ByName(q.EnterpriseName))
	if err != nil {
		return nil, err
	}
	hiringProccesses, err := r.FindAll(ctx, find_all_hiring_proccesses.ByEnterpriseId(dbE.Id()))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return result, nil
		}
		return nil, err
	}
	for _, v := range hiringProccesses {
		result = append(result, dto.NewHiringProccess(*v))
	}
	return result, nil
}
