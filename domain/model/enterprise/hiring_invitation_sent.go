package enterprise

import (
	"encoding/json"
	"go-complaint/domain/model/common"
	"time"

	"github.com/google/uuid"
)

// userID is emitBy
type HiringInvitationSent struct {
	enterpriseId     uuid.UUID
	userId           uuid.UUID
	proposedTo       uuid.UUID
	proposalPosition Position
	occurredOn       time.Time
}

func NewHiringInvitationSent(
	enterpriseId,
	userId,
	proposedTo uuid.UUID,
	proposalPosition Position) *HiringInvitationSent {
	return &HiringInvitationSent{
		enterpriseId:     enterpriseId,
		userId:           userId,
		proposalPosition: proposalPosition,
		proposedTo:       proposedTo,
		occurredOn:       time.Now(),
	}
}

func (h *HiringInvitationSent) OccurredOn() time.Time {
	return h.occurredOn
}

func (h *HiringInvitationSent) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		EnterpriseId     uuid.UUID `json:"enterprise_id"`
		UserId           uuid.UUID `json:"user_id"`
		ProposedTo       uuid.UUID `json:"proposed_to"`
		ProposalPosition string    `json:"position_proposal"`
		OccurredOn       string    `json:"occurred_on"`
	}{
		EnterpriseId:     h.enterpriseId,
		UserId:           h.userId,
		ProposedTo:       h.proposedTo,
		ProposalPosition: h.proposalPosition.String(),
		OccurredOn:       common.StringDate(h.occurredOn),
	})
}

func (h *HiringInvitationSent) UnmarshalJSON(data []byte) error {
	aux := &struct {
		EnterpriseId     uuid.UUID `json:"enterprise_id"`
		UserId           uuid.UUID `json:"user_id"`
		ProposedTo       uuid.UUID `json:"proposed_to"`
		ProposalPosition string    `json:"position_proposal"`
		OccurredOn       string    `json:"occurred_on"`
	}{}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	h.enterpriseId = aux.EnterpriseId
	h.userId = aux.UserId
	h.proposedTo = aux.ProposedTo
	h.proposalPosition = ParsePosition(aux.ProposalPosition)
	if h.proposalPosition < 0 {
		return ErrPositionNotExists
	}
	occurredOn, err := common.ParseDate(aux.OccurredOn)
	if err != nil {
		return err
	}
	h.occurredOn = occurredOn
	return nil
}

func (h *HiringInvitationSent) EnterpriseId() uuid.UUID {
	return h.enterpriseId
}

func (h *HiringInvitationSent) UserID() uuid.UUID {
	return h.userId
}

func (h *HiringInvitationSent) ProposedTo() uuid.UUID {
	return h.proposedTo
}

func (h *HiringInvitationSent) ProposalPosition() Position {
	return h.proposalPosition
}
