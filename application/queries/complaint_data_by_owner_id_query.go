package queries

import (
	"context"
	"errors"
	"go-complaint/dto"
	"go-complaint/infrastructure/persistence/finders/find_all_complaint_data"
	"go-complaint/infrastructure/persistence/repositories"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type ComplaintDataByOwnerIdQuery struct {
	Id string `json:"id"`
}

func NewComplaintDataByOwnerIdQuery(id string) *ComplaintDataByOwnerIdQuery {
	return &ComplaintDataByOwnerIdQuery{
		Id: id,
	}
}

func (q ComplaintDataByOwnerIdQuery) Execute(ctx context.Context) ([]*dto.ComplaintData, error) {
	id, err := uuid.Parse(q.Id)
	if err != nil {
		return nil, err
	}
	registry := repositories.MapperRegistryInstance()
	repository, ok := registry.Get("ComplaintData").(repositories.ComplaintDataRepository)
	if !ok {
		return nil, ErrWrongTypeAssertion
	}
	results := make([]*dto.ComplaintData, 0)
	c, err := repository.FindAll(ctx, find_all_complaint_data.ByOwnerIdAndDataOwnership(id))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return results, nil
		}
		return nil, err
	}
	for i := range c {
		results = append(results, dto.NewComplaintData(*c[i]))
	}
	return results, nil
}
