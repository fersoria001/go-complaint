package dto

import (
	"go-complaint/domain/model/common"
	"go-complaint/domain/model/enterprise"
	"go-complaint/domain/model/identity"
)

type HiringInvitation struct {
	EventID           string `json:"event_id"`
	EnterpriseID      string `json:"enterprise_id"`
	ProposedPosition  string `json:"proposed_position"`
	FullName          string `json:"full_name"`
	EnterpriseEmail   string `json:"enterprise_email"`
	EnterprisePhone   string `json:"enterprise_phone"`
	EnterpriseLogoIMG string `json:"enterprise_logo_img"`
	OccurredOn        string `json:"occurred_on"`
	Seen              bool   `json:"seen"`
	Status            string `json:"status"`
}

func NewHiringInvitation(
	eventID string,
	seen bool,
	user identity.User,
	domainEnterprise enterprise.Enterprise,
	domainEvent enterprise.HiringInvitationSent,
) HiringInvitation {
	stringDate := common.NewDate(domainEvent.OccurredOn()).StringRepresentation()
	return HiringInvitation{
		EventID:           eventID,
		EnterpriseID:      domainEvent.EnterpriseID(),
		FullName:          user.FullName(),
		EnterpriseEmail:   domainEnterprise.Email(),
		EnterprisePhone:   domainEnterprise.Phone(),
		EnterpriseLogoIMG: domainEnterprise.LogoIMG(),
		ProposedPosition:  domainEvent.ProposalPosition().String(),
		OccurredOn:        stringDate,
		Seen:              seen,
	}
}
