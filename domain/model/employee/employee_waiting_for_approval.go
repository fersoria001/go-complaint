package employee

import (
	"encoding/json"
	"go-complaint/domain/model/common"
	"go-complaint/domain/model/enterprise"
	"time"

	"github.com/google/uuid"
)

// Package enterprise
// <<Domain event>>
type EmployeeWaitingForApproval struct {
	enterpriseID     string
	invitedUserID    string
	proposedPosition enterprise.Position
	invitationID     uuid.UUID
	occurredOn       time.Time
}

func NewEmployeeWaitingForApproval(
	enterpriseID string,
	invitedUserID string,
	proposedPosition enterprise.Position,
	invitationID uuid.UUID,
) *EmployeeWaitingForApproval {
	return &EmployeeWaitingForApproval{
		enterpriseID:     enterpriseID,
		invitedUserID:    invitedUserID,
		proposedPosition: proposedPosition,
		invitationID:     invitationID,
		occurredOn:       time.Now(),
	}
}

func (e *EmployeeWaitingForApproval) OccurredOn() time.Time {
	return e.occurredOn
}

func (e *EmployeeWaitingForApproval) MarshalJSON() ([]byte, error) {
	j, err := json.Marshal(struct {
		EnterpriseID     string `json:"enterprise_id"`
		InvitedUserID    string `json:"invited_user_id"`
		ProposedPosition string `json:"proposed_position"`
		OccurredOn       string `json:"occurred_on"`
		InvitationID     string `json:"invitation_id"`
	}{
		EnterpriseID:     e.enterpriseID,
		InvitedUserID:    e.invitedUserID,
		ProposedPosition: e.proposedPosition.String(),
		OccurredOn:       common.StringDate(e.occurredOn),
		InvitationID:     e.invitationID.String(),
	})
	if err != nil {
		return nil, err
	}
	return j, nil
}

func (e *EmployeeWaitingForApproval) UnmarshalJSON(data []byte) error {
	aux := struct {
		EnterpriseID     string `json:"enterprise_id"`
		InvitedUserID    string `json:"invited_user_id"`
		ProposedPosition string `json:"proposed_position"`
		OccurredOn       string `json:"occurred_on"`
		InvitationID     string `json:"invitation_id"`
	}{}
	err := json.Unmarshal(data, &aux)
	if err != nil {
		return err
	}
	e.invitedUserID = aux.InvitedUserID
	e.enterpriseID = aux.EnterpriseID
	e.occurredOn, err = common.ParseDate(aux.OccurredOn)
	if err != nil {
		return err
	}
	e.invitationID, err = uuid.Parse(aux.InvitationID)
	if err != nil {
		return err
	}
	position := enterprise.ParsePosition(aux.ProposedPosition)
	if position == enterprise.NOT_EXISTS {
		return enterprise.ErrPositionNotExists
	}
	e.proposedPosition = position
	return nil
}

func (e *EmployeeWaitingForApproval) EnterpriseID() string {
	return e.enterpriseID
}

func (e *EmployeeWaitingForApproval) InvitationID() uuid.UUID {
	return e.invitationID
}

func (e *EmployeeWaitingForApproval) InvitedUserID() string {
	return e.invitedUserID
}

func (e *EmployeeWaitingForApproval) ProposedPosition() enterprise.Position {
	return e.proposedPosition
}
