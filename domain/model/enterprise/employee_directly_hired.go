package enterprise

import (
	"encoding/json"
	"go-complaint/domain/model/common"
	"time"
)

// Package enterprise
// <<Domain event>> Implements domain.DomainEvent
type EmployeeDirectlyHired struct {
	enterpriseName string
	employeeID     string
	userID         string
	occurredOn     time.Time
}

func NewEmployeeDirectlyHired(enterpriseName, employeeID, userID string) *EmployeeDirectlyHired {
	return &EmployeeDirectlyHired{
		enterpriseName: enterpriseName,
		employeeID:     employeeID,
		userID:         userID,
		occurredOn:     time.Now(),
	}
}

func (eh *EmployeeDirectlyHired) OccurredOn() time.Time {
	return eh.occurredOn
}

func (eh *EmployeeDirectlyHired) MarshalJSON() ([]byte, error) {
	j, err := json.Marshal(struct {
		EnterpriseName string `json:"enterprise_name"`
		EmployeeID     string `json:"employee_id"`
		UserID         string `json:"user_id"`
		OccurredOn     string `json:"occurred_on"`
	}{
		EnterpriseName: eh.enterpriseName,
		EmployeeID:     eh.employeeID,
		UserID:         eh.userID,
		OccurredOn:     common.StringDate(eh.occurredOn),
	})
	if err != nil {
		return nil, err
	}
	return j, nil
}

func (eh *EmployeeDirectlyHired) UnmarshalJSON(data []byte) error {
	aux := struct {
		EnterpriseName string `json:"enterprise_name"`
		EmployeeID     string `json:"employee_id"`
		UserID         string `json:"user_id"`
		OccurredOn     string `json:"occurred_on"`
	}{}
	err := json.Unmarshal(data, &aux)
	if err != nil {
		return err
	}
	eh.enterpriseName = aux.EnterpriseName
	eh.employeeID = aux.EmployeeID
	eh.userID = aux.UserID
	eh.occurredOn, err = common.ParseDate(aux.OccurredOn)
	if err != nil {
		return err
	}
	return nil
}
