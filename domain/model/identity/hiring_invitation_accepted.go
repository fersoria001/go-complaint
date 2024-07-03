package identity

import (
	"encoding/json"
	"go-complaint/domain/model/common"
	"time"
)

type HiringInvitationAccepted struct {
	enterpriseID     string
	invitedUserID    string
	proposedPosition RolesEnum
	occurredOn       time.Time
}

func NewHiringInvitationAccepted(
	enterpriseID,
	invitedUserID string,
	proposedPosition RolesEnum,
) *HiringInvitationAccepted {
	return &HiringInvitationAccepted{
		enterpriseID:     enterpriseID,
		invitedUserID:    invitedUserID,
		proposedPosition: proposedPosition,
		occurredOn:       time.Now(),
	}
}

func (event HiringInvitationAccepted) OccurredOn() time.Time {
	return event.occurredOn
}

func (event HiringInvitationAccepted) EnterpriseID() string {
	return event.enterpriseID
}

func (event HiringInvitationAccepted) InvitedUserID() string {
	return event.invitedUserID
}

func (event HiringInvitationAccepted) ProposedPosition() RolesEnum {
	return event.proposedPosition
}

func (event *HiringInvitationAccepted) MarshalJSON() ([]byte, error) {
	commonDate := common.NewDate(event.occurredOn)
	stringDate := commonDate.StringRepresentation()
	return json.Marshal(struct {
		EnterpriseID     string `json:"enterprise_id"`
		InvitedUserID    string `json:"invited_user_id"`
		ProposedPosition string `json:"proposed_position"`
		OccurredOn       string `json:"occurred_on"`
	}{
		EnterpriseID:     event.enterpriseID,
		InvitedUserID:    event.invitedUserID,
		ProposedPosition: event.proposedPosition.String(),
		OccurredOn:       stringDate,
	})

}

func (event *HiringInvitationAccepted) UnmarshalJSON(data []byte) error {
	var raw struct {
		EnterpriseID     string `json:"enterprise_id"`
		InvitedUserID    string `json:"invited_user_id"`
		ProposedPosition string `json:"proposed_position"`
		OccurredOn       string `json:"occurred_on"`
	}
	err := json.Unmarshal(data, &raw)
	if err != nil {
		return err
	}
	event.enterpriseID = raw.EnterpriseID
	event.invitedUserID = raw.InvitedUserID
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
