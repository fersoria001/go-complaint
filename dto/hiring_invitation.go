package dto

import (
	"go-complaint/domain/model/common"
	"go-complaint/domain/model/enterprise"
	"go-complaint/domain/model/identity"
	"time"
)

type HiringInvitation struct {
	EventID           string `json:"event_id"`
	EnterpriseID      string `json:"enterprise_id"`
	ProposedPosition  string `json:"proposed_position"`
	OwnerID           string `json:"owner_id"`
	FullName          string `json:"full_name"`
	EnterpriseEmail   string `json:"enterprise_email"`
	EnterprisePhone   string `json:"enterprise_phone"`
	EnterpriseLogoIMG string `json:"enterprise_logo_img"`
	OccurredOn        string `json:"occurred_on"`
	Seen              bool   `json:"seen"`
	Status            string `json:"status"`
	Reason            string `json:"reason"`
}

func NewHiringInvitation(
	eventID string,
	domainEvent enterprise.HiringInvitationSent,
) *HiringInvitation {
	stringDate := common.NewDate(domainEvent.OccurredOn()).StringRepresentation()
	return &HiringInvitation{
		EventID:          eventID,
		OwnerID:          domainEvent.ProposedTo(),
		EnterpriseID:     domainEvent.EnterpriseID(),
		ProposedPosition: domainEvent.ProposalPosition().String(),
		OccurredOn:       stringDate,
		Seen:             false,
	}
}

func (h *HiringInvitation) SetOcurredOn(occurredOn time.Time) {
	h.OccurredOn = common.StringDate(occurredOn)
}
func (h *HiringInvitation) GetOcurredOn() time.Time {
	d, _ := common.ParseDate(h.OccurredOn)
	return d
}

func (hiringInvitation *HiringInvitation) SetReason(reason string) {
	hiringInvitation.Reason = reason
}

func (hiringInvitation *HiringInvitation) SetStatus(status string) {
	hiringInvitation.Status = status

}

func (hiringInvitation *HiringInvitation) SetUser(user identity.User) {
	hiringInvitation.FullName = user.FullName()

}

func (hiringInvitation *HiringInvitation) SetEnterprise(enterprise enterprise.Enterprise) {
	hiringInvitation.EnterpriseEmail = enterprise.Email()
	hiringInvitation.EnterprisePhone = enterprise.Phone()
	hiringInvitation.EnterpriseLogoIMG = enterprise.LogoIMG()
}
