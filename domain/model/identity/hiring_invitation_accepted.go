package identity

import (
	"encoding/json"
	"go-complaint/domain/model/common"
	"time"

	"github.com/google/uuid"
)

type HiringInvitationAccepted struct {
	enterpriseId     uuid.UUID
	invitedUserId    uuid.UUID
	proposedPosition RolesEnum
	occurredOn       time.Time
}

func NewHiringInvitationAccepted(
	enterpriseId,
	invitedUserId uuid.UUID,
	proposedPosition RolesEnum,
) *HiringInvitationAccepted {
	return &HiringInvitationAccepted{
		enterpriseId:     enterpriseId,
		invitedUserId:    invitedUserId,
		proposedPosition: proposedPosition,
		occurredOn:       time.Now(),
	}
}

func (event HiringInvitationAccepted) OccurredOn() time.Time {
	return event.occurredOn
}

func (event HiringInvitationAccepted) EnterpriseId() uuid.UUID {
	return event.enterpriseId
}

func (event HiringInvitationAccepted) InvitedUserId() uuid.UUID {
	return event.invitedUserId
}

func (event HiringInvitationAccepted) ProposedPosition() RolesEnum {
	return event.proposedPosition
}

func (event *HiringInvitationAccepted) MarshalJSON() ([]byte, error) {
	commonDate := common.NewDate(event.occurredOn)
	stringDate := commonDate.StringRepresentation()
	return json.Marshal(struct {
		EnterpriseId     uuid.UUID `json:"enterprise_id"`
		InvitedUserId    uuid.UUID `json:"invited_user_id"`
		ProposedPosition string    `json:"proposed_position"`
		OccurredOn       string    `json:"occurred_on"`
	}{
		EnterpriseId:     event.enterpriseId,
		InvitedUserId:    event.invitedUserId,
		ProposedPosition: event.proposedPosition.String(),
		OccurredOn:       stringDate,
	})

}

func (event *HiringInvitationAccepted) UnmarshalJSON(data []byte) error {
	var raw struct {
		EnterpriseId     uuid.UUID `json:"enterprise_id"`
		InvitedUserId    uuid.UUID `json:"invited_user_id"`
		ProposedPosition string    `json:"proposed_position"`
		OccurredOn       string    `json:"occurred_on"`
	}
	err := json.Unmarshal(data, &raw)
	if err != nil {
		return err
	}
	event.enterpriseId = raw.EnterpriseId
	event.invitedUserId = raw.InvitedUserId
	event.proposedPosition, err = ParseRole(raw.ProposedPosition)
	if err != nil {
		return err
	}
	date, err := common.NewDateFromString(raw.OccurredOn)
	if err != nil {
		return err
	}
	event.occurredOn = date.Date()
	return nil
}
