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

type HiringProccessByUserIdQuery struct {
	UserId string `json:"userId"`
}

func NewHiringProccessByUserIdQuery(userId string) *HiringProccessByUserIdQuery {
	return &HiringProccessByUserIdQuery{
		UserId: userId,
	}
}

func (q HiringProccessByUserIdQuery) Execute(ctx context.Context) ([]*dto.HiringProccess, error) {
	userId, err := uuid.Parse(q.UserId)
	if err != nil {
		return nil, err
	}
	reg := repositories.MapperRegistryInstance()
	r, ok := reg.Get("HiringProccess").(repositories.HiringProccessRepository)
	if !ok {
		return nil, ErrWrongTypeAssertion
	}
	result := make([]*dto.HiringProccess, 0)
	hiringProccesses, err := r.FindAll(ctx, find_all_hiring_proccesses.ByUserId(userId))
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
