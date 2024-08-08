package queries

import (
	"context"
	"go-complaint/dto"
	"go-complaint/infrastructure/persistence/finders/find_complaint"
	"go-complaint/infrastructure/persistence/repositories"

	"github.com/google/uuid"
)

type ComplaintInWritingStatusQuery struct {
	AuthorId   string `json:"authorId"`
	ReceiverId string `json:"receiverId"`
}

func NewComplaintInWritingStatusQuery(authorId, receiverId string) *ComplaintInWritingStatusQuery {
	return &ComplaintInWritingStatusQuery{
		AuthorId:   authorId,
		ReceiverId: receiverId,
	}
}

func (c ComplaintInWritingStatusQuery) Execute(ctx context.Context) (*dto.Complaint, error) {
	authorId, err := uuid.Parse(c.AuthorId)
	if err != nil {
		return nil, err
	}
	receiverId, err := uuid.Parse(c.ReceiverId)
	if err != nil {
		return nil, err
	}
	reg := repositories.MapperRegistryInstance()
	complaintRepository, ok := reg.Get("Complaint").(repositories.ComplaintRepository)
	if !ok {
		return nil, ErrWrongTypeAssertion
	}
	dbc, err := complaintRepository.Find(ctx, find_complaint.ByAuthorAndReceiverAndWritingTrue(
		authorId,
		receiverId,
	))
	if err != nil {
		return nil, err
	}
	return dto.NewComplaint(*dbc), nil
}
