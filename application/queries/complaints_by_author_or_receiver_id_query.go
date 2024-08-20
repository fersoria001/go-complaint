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

type ComplaintsByAuthorOrReceiverIdQuery struct {
	Id string `json:"id"`
}

func NewComplaintsByAuthorOrReceiverIdQuery(id string) *ComplaintsByAuthorOrReceiverIdQuery {
	return &ComplaintsByAuthorOrReceiverIdQuery{
		Id: id,
	}
}

func (cbaoriq *ComplaintsByAuthorOrReceiverIdQuery) Execute(ctx context.Context) ([]*dto.Complaint, error) {
	id, err := uuid.Parse(cbaoriq.Id)
	if err != nil {
		return nil, err
	}
	reg := repositories.MapperRegistryInstance()
	repository, ok := reg.Get("Complaint").(repositories.ComplaintRepository)
	if !ok {
		return nil, ErrWrongTypeAssertion
	}
	validStatus := []string{}
	validStatus = append(validStatus, complaint.OPEN.String(), complaint.STARTED.String(), complaint.IN_DISCUSSION.String())
	// count, err := repository.Count(ctx, count_complaints.WhereAuthorOrReceiverId(id, validStatus))
	// if err != nil {
	// 	return nil, err
	// }
	// limit := 10
	// if cbaoriq.Cursor >= count {
	// 	return nil, ErrCursorOutOfRange
	// }
	result := make([]*dto.Complaint, 0)
	dbCs, err := repository.FindAll(ctx, find_all_complaints.ByAuthorOrReceiver(id, validStatus))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return result, nil
		}
		return nil, err
	}
	// nextCursor := cbaoriq.Cursor + len(dbCs)
	// if nextCursor >= count {
	// 	nextCursor = -1
	// }
	// prevCursor := cbaoriq.Cursor - len(dbCs)
	// if prevCursor < 0 {
	// 	prevCursor = -1
	// }
	for _, v := range dbCs {
		complaintDto := dto.NewComplaint(*v)
		result = append(result, complaintDto)
	}
	return result, nil
}
