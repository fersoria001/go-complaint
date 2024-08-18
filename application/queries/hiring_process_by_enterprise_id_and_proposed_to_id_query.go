package queries

import (
	"context"
	"go-complaint/dto"
	"go-complaint/infrastructure/persistence/finders/find_enterprise"
	"go-complaint/infrastructure/persistence/finders/find_hiring_proccess"
	"go-complaint/infrastructure/persistence/repositories"

	"github.com/google/uuid"
)

type HiringProcessByEnterpriseNameAndProposedToQuery struct {
	ProposedId     string `json:"proposedId"`
	EnterpriseName string `json:"enterpriseName"`
}

func NewHiringProcessByEnterpriseNameAndProposedToQuery(proposedId, enterpriseName string) *HiringProcessByEnterpriseNameAndProposedToQuery {
	return &HiringProcessByEnterpriseNameAndProposedToQuery{
		ProposedId:     proposedId,
		EnterpriseName: enterpriseName,
	}
}

func (q HiringProcessByEnterpriseNameAndProposedToQuery) Execute(ctx context.Context) (*dto.HiringProccess, error) {
	proposedId, err := uuid.Parse(q.ProposedId)
	if err != nil {
		return nil, err
	}
	reg := repositories.MapperRegistryInstance()
	enterpriseRepository, ok := reg.Get("Enterprise").(repositories.EnterpriseRepository)
	if !ok {
		return nil, ErrWrongTypeAssertion
	}
	dbE, err := enterpriseRepository.Find(ctx, find_enterprise.ByName(q.EnterpriseName))
	if err != nil {
		return nil, err
	}
	enterpriseId := dbE.Id()
	r, ok := reg.Get("HiringProccess").(repositories.HiringProccessRepository)
	if !ok {
		return nil, ErrWrongTypeAssertion
	}
	hiringProcess, err := r.Find(ctx, find_hiring_proccess.ByUserIdAndEnterpriseId(proposedId, enterpriseId))
	if err != nil {
		return nil, err
	}
	return dto.NewHiringProccess(*hiringProcess), nil
}
