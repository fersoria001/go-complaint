package dto

import (
	"go-complaint/domain/model/common"
	"go-complaint/domain/model/enterprise"
)

type HiringProccess struct {
	Id         string     `json:"id"`
	Enterprise *Recipient `json:"enterprise"`
	User       *User      `json:"user"`
	Role       string     `json:"role"`
	Status     string     `json:"status"`
	Reason     string     `json:"reason"`
	EmitedBy   *Recipient `json:"emitedBy"`
	OccurredOn string     `json:"occurredOn"`
	LastUpdate string     `json:"lastUpdate"`
	UpdatedBy  *Recipient `json:"UpdatedBy"`
	Industry   Industry   `json:"industry"`
}

func NewHiringProccess(
	obj enterprise.HiringProccess,
) *HiringProccess {
	u := obj.User()
	return &HiringProccess{
		Id:         obj.Id().String(),
		Enterprise: NewRecipient(obj.Enterprise()),
		User:       NewUser(&u),
		Role:       obj.Role().String(),
		Status:     obj.Status().String(),
		Reason:     obj.Reason(),
		EmitedBy:   NewRecipient(obj.EmitedBy()),
		OccurredOn: common.StringDate(obj.OccurredOn()),
		LastUpdate: common.StringDate(obj.LastUpdate()),
		UpdatedBy:  NewRecipient(obj.UpdatedBy()),
		Industry:   NewIndustry(obj.Industry()),
	}
}
