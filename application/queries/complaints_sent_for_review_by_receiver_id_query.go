package queries

import (
	"context"
	"errors"
	"go-complaint/domain/model/complaint"
	"go-complaint/dto"
	"go-complaint/infrastructure/persistence/finders/find_all_complaints"
	"go-complaint/infrastructure/persistence/repositories"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type ComplaintsSentForReviewByReceiverIdQuery struct {
	ReceiverId string `json:"receiverId"`
}

func NewComplaintsSentForReviewByReceiverIdQuery(receiverId string) *ComplaintsSentForReviewByReceiverIdQuery {
	return &ComplaintsSentForReviewByReceiverIdQuery{
		ReceiverId: receiverId,
	}
}

func (q ComplaintsSentForReviewByReceiverIdQuery) Execute(ctx context.Context) ([]*dto.Complaint, error) {
	receiverId, err := uuid.Parse(q.ReceiverId)
	if err != nil {
		return nil, err
	}
	reg := repositories.MapperRegistryInstance()
	r, ok := reg.Get("Complaint").(repositories.ComplaintRepository)
	if !ok {
		return nil, ErrWrongTypeAssertion
	}
	result := make([]*dto.Complaint, 0)
	c, err := r.FindAll(ctx, find_all_complaints.ByReceiverAndStatusIn(receiverId, []string{complaint.IN_REVIEW.String()}))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return result, nil
		}
		return nil, err
	}
	for _, v := range c {
		result = append(result, dto.NewComplaint(*v))
	}
	return result, nil
}
