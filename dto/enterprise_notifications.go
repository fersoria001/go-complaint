package dto

import (
	"go-complaint/domain/model/common"
	"go-complaint/domain/model/enterprise"
)

type EnterpriseNotifications struct {
	Count                      int                          `json:"count"`
	EmployeeWaitingForApproval []EmployeeWaitingForApproval `json:"employee_waiting_for_approval"`
}

type EmployeeWaitingForApproval struct {
	ID             string `json:"id"`
	EnterpriseName string `json:"enterprise_name"`
	EmployeeID     string `json:"employee_id"`
	ManagerID      string `json:"manager_id"`
	OccurredOn     string `json:"occurred_on"`
	Seen           bool   `json:"seen"`
}

func NewEmployeeWaitingForApproval(id string,
	seen bool, domainEvent enterprise.EmployeeWaitingForApproval) EmployeeWaitingForApproval {
	return EmployeeWaitingForApproval{
		ID:             id,
		EnterpriseName: domainEvent.EnterpriseName(),
		EmployeeID:     domainEvent.EmployeeID(),
		ManagerID:      domainEvent.ManagerID(),
		OccurredOn:     common.NewDate(domainEvent.OccurredOn()).StringRepresentation(),
		Seen:           seen,
	}
}
