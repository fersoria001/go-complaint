package queries

import (
	"context"
	"errors"
	"go-complaint/dto"
	"go-complaint/infrastructure/persistence/finders/find_all_hiring_proccesses"
	"go-complaint/infrastructure/persistence/repositories"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type HiringProccessByEnterpriseIdQuery struct {
	EnterpriseId string `json:"enterpriseId"`
}

func NewHiringProccessByEnterpriseIdQuery(enterpriseId string) *HiringProccessByEnterpriseIdQuery {
	return &HiringProccessByEnterpriseIdQuery{
		EnterpriseId: enterpriseId,
	}
}

func (q HiringProccessByEnterpriseIdQuery) Execute(ctx context.Context) ([]*dto.HiringProccess, error) {
	enterpriseId, err := uuid.Parse(q.EnterpriseId)
	if err != nil {
		return nil, err
	}
	reg := repositories.MapperRegistryInstance()
	r, ok := reg.Get("HiringProccess").(repositories.HiringProccessRepository)
	if !ok {
		return nil, ErrWrongTypeAssertion
	}
	result := make([]*dto.HiringProccess, 0)
	hiringProccesses, err := r.FindAll(ctx, find_all_hiring_proccesses.ByEnterpriseId(enterpriseId))
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
