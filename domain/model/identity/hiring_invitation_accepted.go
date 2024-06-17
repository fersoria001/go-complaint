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
	stringDate := common.StringDate(event.occurredOn)
	return json.Marshal(map[string]interface{}{
		"enterprise_id":     event.enterpriseID,
		"invited_user_id":   event.invitedUserID,
		"proposed_position": event.proposedPosition,
		"occurred_on":       stringDate,
	})
}

func (event *HiringInvitationAccepted) UnmarshalJSON(data []byte) error {
	var raw map[string]interface{}
	err := json.Unmarshal(data, &raw)
	if err != nil {
		return err
	}
	event.enterpriseID = raw["enterprise_id"].(string)
	event.invitedUserID = raw["invited_user_id"].(string)
	event.proposedPosition = raw["proposed_position"].(RolesEnum)
	stringDate := raw["occurred_on"].(string)
	event.occurredOn, err = common.ParseDate(stringDate)
	if err != nil {
		return err
	}
	return nil
}
