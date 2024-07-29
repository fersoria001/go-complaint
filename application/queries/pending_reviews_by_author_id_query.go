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

type PendingReviewsByAuthorIdQuery struct {
	AuthorId string `json:"authorId"`
}

func NewPendingReviewsByAuthorIdQuery(authorId string) *PendingReviewsByAuthorIdQuery {
	return &PendingReviewsByAuthorIdQuery{
		AuthorId: authorId,
	}
}

func (q PendingReviewsByAuthorIdQuery) Execute(ctx context.Context) ([]*dto.Complaint, error) {
	reg := repositories.MapperRegistryInstance()
	r, ok := reg.Get("Complaint").(repositories.ComplaintRepository)
	if !ok {
		return nil, ErrWrongTypeAssertion
	}
	authorId, err := uuid.Parse(q.AuthorId)
	if err != nil {
		return nil, err
	}
	result := make([]*dto.Complaint, 0)
	p, err := r.FindAll(ctx, find_all_complaints.ByAuthorAndStatusIn(authorId, []string{
		complaint.IN_REVIEW.String(),
	}))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return result, nil
		}
		return nil, err
	}
	for _, v := range p {
		result = append(result, dto.NewComplaint(*v))
	}
	return result, nil
}
