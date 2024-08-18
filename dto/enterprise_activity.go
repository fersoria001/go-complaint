package dto

import (
	"go-complaint/domain/model/common"
	"go-complaint/domain/model/enterprise"
)

type EnterpriseActivity struct {
	Id             string    `json:"id"`
	User           Recipient `json:"user"`
	ActivityId     string    `json:"activityId"`
	EnterpriseId   string    `json:"enterpriseId"`
	EnterpriseName string    `json:"enterpriseName"`
	OccurredOn     string    `json:"occurredOn"`
	ActivityType   string    `json:"activityType"`
}

func NewEnterpriseActivity(obj enterprise.EnterpriseActivity) *EnterpriseActivity {
	return &EnterpriseActivity{
		Id:             obj.Id().String(),
		User:           *NewRecipient(obj.User()),
		ActivityId:     obj.ActivityId().String(),
		EnterpriseId:   obj.EnterpriseId().String(),
		EnterpriseName: obj.EnterpriseName(),
		OccurredOn:     common.StringDate(obj.OccurredOn()),
		ActivityType:   obj.ActivityType().String(),
	}
}
