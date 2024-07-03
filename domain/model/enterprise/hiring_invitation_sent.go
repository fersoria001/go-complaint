package enterprise

import (
	"encoding/json"
	"go-complaint/domain/model/common"
	"time"
)

// userID is emitBy
type HiringInvitationSent struct {
	enterpriseID     string
	userID           string
	proposedTo       string
	proposalPosition Position
	occurredOn       time.Time
}

func NewHiringInvitationSent(
	enterpriseID string,
	userID string,
	proposedTo string,
	proposalPosition Position) *HiringInvitationSent {
	return &HiringInvitationSent{
		enterpriseID:     enterpriseID,
		userID:           userID,
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
		EnterpriseID     string `json:"enterprise_id"`
		UserID           string `json:"user_id"`
		ProposedTo       string `json:"proposed_to"`
		ProposalPosition string `json:"position_proposal"`
		OccurredOn       string `json:"occurred_on"`
	}{
		EnterpriseID:     h.enterpriseID,
		UserID:           h.userID,
		ProposedTo:       h.proposedTo,
		ProposalPosition: h.proposalPosition.String(),
		OccurredOn:       common.StringDate(h.occurredOn),
	})
}

func (h *HiringInvitationSent) UnmarshalJSON(data []byte) error {
	aux := &struct {
		EnterpriseID     string `json:"enterprise_id"`
		UserID           string `json:"user_id"`
		ProposedTo       string `json:"proposed_to"`
		ProposalPosition string `json:"position_proposal"`
		OccurredOn       string `json:"occurred_on"`
	}{}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	h.enterpriseID = aux.EnterpriseID
	h.userID = aux.UserID
	h.proposedTo = aux.ProposedTo
	h.proposalPosition = ParsePosition(aux.ProposalPosition)
	if h.proposalPosition == NOT_EXISTS {
		return ErrPositionNotExists
	}
	occurredOn, err := common.ParseDate(aux.OccurredOn)
	if err != nil {
		return err
	}
	h.occurredOn = occurredOn
	return nil
}

func (h *HiringInvitationSent) EnterpriseID() string {
	return h.enterpriseID
}

func (h *HiringInvitationSent) UserID() string {
	return h.userID
}

func (h *HiringInvitationSent) ProposedTo() string {
	return h.proposedTo
}

func (h *HiringInvitationSent) ProposalPosition() Position {
	return h.proposalPosition
}
