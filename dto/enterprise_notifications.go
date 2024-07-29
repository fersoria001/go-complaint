package dto

import (
	"go-complaint/domain/model/common"
	"go-complaint/domain/model/enterprise"
)

type EnterpriseNotifications struct {
	Count                      int                          `json:"count"`
	EmployeeWaitingForApproval []EmployeeWaitingForApproval `json:"employee_waiting_for_approval"`
	WaitingForReview           []WaitingForReview           `json:"waiting_for_review"`
}

type EmployeeWaitingForApproval struct {
	ID               string `json:"id"`
	EnterpriseID     string `json:"enterprise_id"`
	InvitedUserID    string `json:"invited_user_id"`
	ProposedPosition string `json:"proposed_position"`
	OccurredOn       string `json:"occurred_on"`
	InvitationID     string `json:"invitation_id"`
	Seen             bool   `json:"seen"`
}

func NewEmployeeWaitingForApproval(id string,
	seen bool, domainEvent enterprise.EmployeeWaitingForApproval) EmployeeWaitingForApproval {
	return EmployeeWaitingForApproval{
		ID:               id,
		EnterpriseID:     domainEvent.EnterpriseId().String(),
		InvitedUserID:    domainEvent.InvitedUserId().String(),
		ProposedPosition: domainEvent.ProposedPosition().String(),
		InvitationID:     domainEvent.InvitationId().String(),
		OccurredOn:       common.NewDate(domainEvent.OccurredOn()).StringRepresentation(),
		Seen:             seen,
	}
}
