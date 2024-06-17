package complaint

import (
	"go-complaint/domain/model/common"

	"github.com/google/uuid"
)

type SolvedBy struct {
	id          uuid.UUID
	complaintID uuid.UUID
	solvedByID  string
	solvedAt    common.Date
}

func NewSolvedBy(id, complaintID uuid.UUID, solvedByID string, solvedAt common.Date) *SolvedBy {
	return &SolvedBy{
		id:          id,
		complaintID: complaintID,
		solvedByID:  solvedByID,
		solvedAt:    solvedAt,
	}
}

func (s SolvedBy) ID() uuid.UUID {
	return s.id
}

func (s SolvedBy) ComplaintID() uuid.UUID {
	return s.complaintID
}

func (s SolvedBy) SolvedByID() string {
	return s.solvedByID
}

func (s SolvedBy) SolvedAt() common.Date {
	return s.solvedAt
}
