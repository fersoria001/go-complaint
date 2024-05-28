package enterprise

import (
	"encoding/json"
	"go-complaint/domain/model/common"
	"strings"
	"time"

	"github.com/google/uuid"
)

// Package enterprise
// <<Domain event>>
type EmployeeWaitingForApproval struct {
	enterpriseName string
	employeeID     string
	managerID      string
	invitationID   uuid.UUID
	occurredOn     time.Time
}

func NewEmployeeWaitingForApproval(employeeID, managerID string, invitationID uuid.UUID) *EmployeeWaitingForApproval {
	enterpriseName := strings.Split(employeeID, "-")[0]
	return &EmployeeWaitingForApproval{
		enterpriseName: enterpriseName,
		employeeID:     employeeID,
		managerID:      managerID,
		occurredOn:     time.Now(),
	}
}

func (e *EmployeeWaitingForApproval) OccurredOn() time.Time {
	return e.occurredOn
}

func (e *EmployeeWaitingForApproval) MarshalJSON() ([]byte, error) {
	j, err := json.Marshal(struct {
		EnterpriseName string `json:"enterprise_name"`
		EmployeeID     string `json:"employee_id"`
		ManagerID      string `json:"manager_id"`
		OccurredOn     string `json:"occurred_on"`
		InvitationID   string `json:"invitation_id"`
	}{
		EnterpriseName: e.enterpriseName,
		EmployeeID:     e.employeeID,
		ManagerID:      e.managerID,
		OccurredOn:     common.StringDate(e.occurredOn),
		InvitationID:   e.invitationID.String(),
	})
	if err != nil {
		return nil, err
	}
	return j, nil
}

func (e *EmployeeWaitingForApproval) UnmarshalJSON(data []byte) error {
	aux := struct {
		EnterpriseName string `json:"enterprise_name"`
		EmployeeID     string `json:"employee_id"`
		ManagerID      string `json:"manager_id"`
		OccurredOn     string `json:"occurred_on"`
		InvitationID   string `json:"invitation_id"`
	}{}
	err := json.Unmarshal(data, &aux)
	if err != nil {
		return err
	}
	e.enterpriseName = aux.EnterpriseName
	e.employeeID = aux.EmployeeID
	e.managerID = aux.ManagerID
	e.occurredOn, err = common.ParseDate(aux.OccurredOn)
	if err != nil {
		return err
	}
	e.invitationID, err = uuid.Parse(aux.InvitationID)
	if err != nil {
		return err
	}
	return nil
}

func (e *EmployeeWaitingForApproval) EnterpriseName() string {
	return e.enterpriseName
}

func (e *EmployeeWaitingForApproval) EmployeeID() string {
	return e.employeeID
}

func (e *EmployeeWaitingForApproval) ManagerID() string {
	return e.managerID
}

func (e *EmployeeWaitingForApproval) InvitationID() uuid.UUID {
	return e.invitationID
}
