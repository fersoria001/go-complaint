package queries

import (
	"context"
	"go-complaint/domain/model/complaint"
	"go-complaint/dto"
	"go-complaint/infrastructure/persistence/finders/find_all_complaint_data"
	"go-complaint/infrastructure/persistence/repositories"

	"github.com/google/uuid"
)

type ComplaintDataByOwnerIdQuery struct {
	Id string `json:"id"`
}

func NewComplaintDataByOwnerIdQuery(id string) *ComplaintDataByOwnerIdQuery {
	return &ComplaintDataByOwnerIdQuery{
		Id: id,
	}
}

func (q ComplaintDataByOwnerIdQuery) Execute(ctx context.Context) (*dto.ComplaintsInfo, error) {
	id, err := uuid.Parse(q.Id)
	if err != nil {
		return nil, err
	}
	registry := repositories.MapperRegistryInstance()
	repository, ok := registry.Get("ComplaintData").(repositories.ComplaintDataRepository)
	if !ok {
		return nil, ErrWrongTypeAssertion
	}
	c, err := repository.FindAll(ctx, find_all_complaint_data.ByOwnerId(id))
	if err != nil {
		return nil, err
	}
	received := make([]*dto.ComplaintData, 0)
	resolved := make([]*dto.ComplaintData, 0)
	reviewed := make([]*dto.ComplaintData, 0)
	sent := make([]*dto.ComplaintData, 0)
	for _, v := range c {
		switch v.DataType() {
		case complaint.RECEIVED:
			received = append(received, dto.NewComplaintData(*v))
		case complaint.SENT:
			sent = append(sent, dto.NewComplaintData(*v))
		case complaint.RESOLVED:
			resolved = append(resolved, dto.NewComplaintData(*v))
		case complaint.REVIEWED:
			reviewed = append(reviewed, dto.NewComplaintData(*v))
		}
	}
	return &dto.ComplaintsInfo{
		Received: received,
		Resolved: resolved,
		Reviewed: reviewed,
		Sent:     sent,
	}, nil
}
