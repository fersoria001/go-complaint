package dto

import (
	"go-complaint/domain/model/common"
	"go-complaint/domain/model/enterprise"
)

type UserNotifications struct {
	Count            int                `json:"count"`
	HiringInvitation []HiringInvitation `json:"hiring_invitation"`
}

type HiringInvitation struct {
	EventID          string `json:"event_id"`
	EnterpriseID     string `json:"enterprise_id"`
	ProposedPosition string `json:"proposed_position"`
	FirstName        string `json:"first_name"`
	LastName         string `json:"last_name"`
	Email            string `json:"email"`
	Phone            string `json:"phone"`
	Age              int    `json:"age"`
	OccurredOn       string `json:"occurred_on"`
	Seen             bool   `json:"seen"`
}

func NewHiringInvitation(eventID string, seen bool, domainEvent enterprise.HiringInvitationSent) HiringInvitation {
	stringDate := common.NewDate(domainEvent.OccurredOn()).StringRepresentation()
	return HiringInvitation{
		EventID:          eventID,
		EnterpriseID:     domainEvent.EnterpriseID(),
		ProposedPosition: domainEvent.ProposalPosition().String(),
		FirstName:        domainEvent.FirstName(),
		LastName:         domainEvent.LastName(),
		Email:            domainEvent.Email(),
		Phone:            domainEvent.Phone(),
		Age:              domainEvent.Age(),
		OccurredOn:       stringDate,
		Seen:             seen,
	}
}
