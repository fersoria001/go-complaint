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

type ComplaintsRatedByAuthorIdQuery struct {
	AuthorId string `json:"authorId"`
}

func NewComplaintsRatedByAuthorIdQuery(authorId string) *ComplaintsRatedByAuthorIdQuery {
	return &ComplaintsRatedByAuthorIdQuery{
		AuthorId: authorId,
	}
}

func (q ComplaintsRatedByAuthorIdQuery) Execute(ctx context.Context) ([]*dto.Complaint, error) {
	authorId, err := uuid.Parse(q.AuthorId)
	if err != nil {
		return nil, err
	}
	reg := repositories.MapperRegistryInstance()
	r, ok := reg.Get("Complaint").(repositories.ComplaintRepository)
	if !ok {
		return nil, ErrWrongTypeAssertion
	}
	result := make([]*dto.Complaint, 0)
	c, err := r.FindAll(ctx, find_all_complaints.ByAuthorAndStatusIn(authorId, []string{complaint.CLOSED.String(), complaint.IN_HISTORY.String()}))
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
