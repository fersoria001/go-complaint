package identity

import (
	"encoding/json"
	"go-complaint/domain/model/common"
	"time"

	"github.com/google/uuid"
)

type HiringInvitationRejected struct {
	enterpriseId     uuid.UUID
	invitedUserId    uuid.UUID
	rejectionReason  string
	proposedPosition RolesEnum
	occurredOn       time.Time
}

func NewHiringInvitationRejected(
	enterpriseId,
	invitedUserId uuid.UUID,
	rejectionReason string,
	proposedPosition RolesEnum,
) *HiringInvitationRejected {
	return &HiringInvitationRejected{
		enterpriseId:     enterpriseId,
		invitedUserId:    invitedUserId,
		rejectionReason:  rejectionReason,
		proposedPosition: proposedPosition,
		occurredOn:       time.Now(),
	}
}

func (event HiringInvitationRejected) OccurredOn() time.Time {
	return event.occurredOn
}

func (event HiringInvitationRejected) EnterpriseId() uuid.UUID {
	return event.enterpriseId
}

func (event HiringInvitationRejected) InvitedUserId() uuid.UUID {
	return event.invitedUserId
}

func (event HiringInvitationRejected) ProposedPosition() RolesEnum {
	return event.proposedPosition
}

func (event HiringInvitationRejected) RejectionReason() string {
	return event.rejectionReason
}

func (event *HiringInvitationRejected) MarshalJSON() ([]byte, error) {
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

func (event *HiringInvitationRejected) UnmarshalJSON(data []byte) error {
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
