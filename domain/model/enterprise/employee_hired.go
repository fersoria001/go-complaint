package enterprise

import (
	"encoding/json"
	"go-complaint/domain/model/common"
	"time"

	"github.com/google/uuid"
)

// Package enterprise
// <<Domain event>> Implements domain.DomainEvent
type EmployeeHired struct {
	enterpriseName string
	employeeID     uuid.UUID
	employeeEmail  string
	position       Position
	occurredOn     time.Time
}

func NewEmployeeHired(
	enterpriseName string,
	employeeID uuid.UUID,
	employeeEmail string,
	position Position,
) *EmployeeHired {
	return &EmployeeHired{
		enterpriseName: enterpriseName,
		employeeID:     employeeID,
		occurredOn:     time.Now(),
	}
}

func (eh *EmployeeHired) EmployeeEmail() string {
	return eh.employeeEmail
}

func (eh *EmployeeHired) Position() Position {
	return eh.position
}

func (eh *EmployeeHired) EmployeeID() uuid.UUID {
	return eh.employeeID
}

func (eh *EmployeeHired) EnterpriseName() string {
	return eh.enterpriseName
}

func (eh *EmployeeHired) OccurredOn() time.Time {
	return eh.occurredOn
}

func (eh *EmployeeHired) MarshalJSON() ([]byte, error) {
	j, err := json.Marshal(struct {
		EnterpriseName string `json:"enterprise_name"`
		EmployeeID     string `json:"employee_id"`
		EmployeeEmail  string `json:"employee_email"`
		Position       string `json:"position"`
		OccurredOn     string `json:"occurred_on"`
	}{
		EnterpriseName: eh.enterpriseName,
		EmployeeID:     eh.employeeID.String(),
		EmployeeEmail:  eh.employeeEmail,
		Position:       eh.position.String(),
		OccurredOn:     common.StringDate(eh.occurredOn),
	})
	if err != nil {
		return nil, err
	}
	return j, nil
}

func (eh *EmployeeHired) UnmarshalJSON(data []byte) error {
	aux := struct {
		EnterpriseName string `json:"enterprise_name"`
		EmployeeID     string `json:"employee_id"`
		EmployeeEmail  string `json:"employee_email"`
		Position       string `json:"position"`
		OccurredOn     string `json:"occurred_on"`
	}{}
	err := json.Unmarshal(data, &aux)
	if err != nil {
		return err
	}
	eh.employeeEmail = aux.EmployeeEmail
	eh.position = ParsePosition(aux.Position)
	eh.enterpriseName = aux.EnterpriseName
	eh.employeeID, err = uuid.Parse(aux.EmployeeID)
	if err != nil {
		return err
	}
	eh.occurredOn, err = common.ParseDate(aux.OccurredOn)
	if err != nil {
		return err
	}
	return nil
}
