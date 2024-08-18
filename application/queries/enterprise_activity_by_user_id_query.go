package queries

import (
	"context"
	"errors"
	"go-complaint/dto"
	"go-complaint/infrastructure/persistence/finders/find_all_enterprise_activity"
	"go-complaint/infrastructure/persistence/repositories"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type EnterpriseActivityByUserIdQuery struct {
	UserId string `json:"userId"`
}

func NewEnterpriseActivityByUserIdQuery(userId string) *EnterpriseActivityByUserIdQuery {
	return &EnterpriseActivityByUserIdQuery{
		UserId: userId,
	}
}

func (c EnterpriseActivityByUserIdQuery) Execute(ctx context.Context) ([]*dto.EnterpriseActivity, error) {
	userId, err := uuid.Parse(c.UserId)
	if err != nil {
		return nil, err
	}
	reg := repositories.MapperRegistryInstance()
	r, ok := reg.Get("EnterpriseActivity").(repositories.EnterpriseActivityRepository)
	if !ok {
		return nil, ErrWrongTypeAssertion
	}
	results := make([]*dto.EnterpriseActivity, 0)
	dbEa, err := r.FindAll(ctx, find_all_enterprise_activity.ByUserId(userId))
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
