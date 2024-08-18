package queries

import (
	"context"
	"go-complaint/dto"
	"go-complaint/infrastructure/persistence/repositories"

	"github.com/google/uuid"
)

type HiringProcessByIdQuery struct {
	Id string `json:"id"`
}

func NewHiringProcessByIdQuery(id string) *HiringProcessByIdQuery {
	return &HiringProcessByIdQuery{
		Id: id,
	}
}

func (q HiringProcessByIdQuery) Execute(ctx context.Context) (*dto.HiringProccess, error) {
	reg := repositories.MapperRegistryInstance()
	id, err := uuid.Parse(q.Id)
	if err != nil {
		return nil, err
	}
	r, ok := reg.Get("HiringProccess").(repositories.HiringProccessRepository)
	if !ok {
		return nil, ErrWrongTypeAssertion
	}
	dbH, err := r.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return dto.NewHiringProccess(*dbH), nil
}
