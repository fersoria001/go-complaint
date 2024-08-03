package dto

import (
	"go-complaint/domain/model/common"
	"go-complaint/domain/model/enterprise"
)

type HiringProccess struct {
	Id         string     `json:"id"`
	Enterprise *Recipient `json:"enterprise"`
	User       *Recipient `json:"user"`
	Role       string     `json:"role"`
	Status     string     `json:"status"`
	Reason     string     `json:"reason"`
	EmitedBy   *Recipient `json:"emitedBy"`
	OccurredOn string     `json:"occurredOn"`
	LastUpdate string     `json:"lastUpdate"`
	UpdatedBy  *Recipient `json:"UpdatedBy"`
}

func NewHiringProccess(
	obj enterprise.HiringProccess,
) *HiringProccess {
	return &HiringProccess{
		Id:         obj.Id().String(),
		Enterprise: NewRecipient(obj.Enterprise()),
		User:       NewRecipient(obj.User()),
		Role:       obj.Role().String(),
		Status:     obj.Status().String(),
		Reason:     obj.Reason(),
		EmitedBy:   NewRecipient(obj.EmitedBy()),
		OccurredOn: common.StringDate(obj.OccurredOn()),
		LastUpdate: common.StringDate(obj.LastUpdate()),
		UpdatedBy:  NewRecipient(obj.UpdatedBy()),
	}
}
