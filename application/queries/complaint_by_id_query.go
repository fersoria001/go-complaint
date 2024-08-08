package queries

import (
	"context"
	"go-complaint/dto"
	"go-complaint/infrastructure/persistence/repositories"

	"github.com/google/uuid"
)

type ComplaintByIdQuery struct {
	ComplaintId string `json:"complaintId"`
}

func NewComplaintByIdQuery(complaintId string) *ComplaintByIdQuery {
	return &ComplaintByIdQuery{
		ComplaintId: complaintId,
	}
}

func (c ComplaintByIdQuery) Execute(ctx context.Context) (*dto.Complaint, error) {
	complaintId, err := uuid.Parse(c.ComplaintId)
	if err != nil {
		return nil, err
	}
	reg := repositories.MapperRegistryInstance()
	r, ok := reg.Get("Complaint").(repositories.ComplaintRepository)
	if !ok {
		return nil, ErrWrongTypeAssertion
	}
	dbc, err := r.Get(ctx, complaintId)
	if err != nil {
		return nil, err
	}
	return dto.NewComplaint(*dbc), nil
}
