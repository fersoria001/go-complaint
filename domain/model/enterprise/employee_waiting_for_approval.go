package enterprise

import (
	"encoding/json"
	"go-complaint/domain/model/common"
	"time"

	"github.com/google/uuid"
)

// Package enterprise
// <<Domain event>>
type EmployeeWaitingForApproval struct {
	enterpriseId     uuid.UUID
	invitedUserId    uuid.UUID
	proposedPosition Position
	invitationId     uuid.UUID
	occurredOn       time.Time
}

func NewEmployeeWaitingForApproval(
	enterpriseId,
	invitedUserId uuid.UUID,
	proposedPosition Position,
	invitationId uuid.UUID,
) *EmployeeWaitingForApproval {
	return &EmployeeWaitingForApproval{
		enterpriseId:     enterpriseId,
		invitedUserId:    invitedUserId,
		proposedPosition: proposedPosition,
		invitationId:     invitationId,
		occurredOn:       time.Now(),
	}
}

func (e *EmployeeWaitingForApproval) OccurredOn() time.Time {
	return e.occurredOn
}

func (e *EmployeeWaitingForApproval) MarshalJSON() ([]byte, error) {
	j, err := json.Marshal(struct {
		EnterpriseId     uuid.UUID `json:"enterprise_id"`
		InvitedUserId    uuid.UUID `json:"invited_user_id"`
		ProposedPosition string    `json:"proposed_position"`
		OccurredOn       string    `json:"occurred_on"`
		InvitationId     uuid.UUID `json:"invitation_id"`
	}{
		EnterpriseId:     e.enterpriseId,
		InvitedUserId:    e.invitedUserId,
		ProposedPosition: e.proposedPosition.String(),
		OccurredOn:       common.StringDate(e.occurredOn),
		InvitationId:     e.invitationId,
	})
	if err != nil {
		return nil, err
	}
	return j, nil
}

func (e *EmployeeWaitingForApproval) UnmarshalJSON(data []byte) error {
	aux := struct {
		EnterpriseId     uuid.UUID `json:"enterprise_id"`
		InvitedUserId    uuid.UUID `json:"invited_user_id"`
		ProposedPosition string    `json:"proposed_position"`
		OccurredOn       string    `json:"occurred_on"`
		InvitationID     string    `json:"invitation_id"`
	}{}
	err := json.Unmarshal(data, &aux)
	if err != nil {
		return err
	}
	e.invitedUserId = aux.InvitedUserId
	e.enterpriseId = aux.EnterpriseId
	e.occurredOn, err = common.ParseDate(aux.OccurredOn)
	if err != nil {
		return err
	}
	e.invitationId, err = uuid.Parse(aux.InvitationID)
	if err != nil {
		return err
	}
	position := ParsePosition(aux.ProposedPosition)
	if position < 0 {
		return ErrPositionNotExists
	}
	e.proposedPosition = position
	return nil
}

func (e *EmployeeWaitingForApproval) EnterpriseId() uuid.UUID {
	return e.enterpriseId
}

func (e *EmployeeWaitingForApproval) InvitationId() uuid.UUID {
	return e.invitationId
}

func (e *EmployeeWaitingForApproval) InvitedUserId() uuid.UUID {
	return e.invitedUserId
}

func (e *EmployeeWaitingForApproval) ProposedPosition() Position {
	return e.proposedPosition
}
