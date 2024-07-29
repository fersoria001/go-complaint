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
	enterpriseId  uuid.UUID
	emitedBy      uuid.UUID
	employeeId    uuid.UUID
	employeeEmail string
	position      Position
	occurredOn    time.Time
}

func NewEmployeeHired(
	enterpriseId,
	emitedBy,
	employeeId uuid.UUID,
	employeeEmail string,
	position Position,
) *EmployeeHired {
	return &EmployeeHired{
		enterpriseId:  enterpriseId,
		emitedBy:      emitedBy,
		employeeId:    employeeId,
		employeeEmail: employeeEmail,
		position:      position,
		occurredOn:    time.Now(),
	}
}

func (eh EmployeeHired) EmitedBy() uuid.UUID {
	return eh.emitedBy
}

func (eh *EmployeeHired) EmployeeEmail() string {
	return eh.employeeEmail
}

func (eh *EmployeeHired) Position() Position {
	return eh.position
}

func (eh *EmployeeHired) EmployeeId() uuid.UUID {
	return eh.employeeId
}

func (eh *EmployeeHired) EnterpriseId() uuid.UUID {
	return eh.enterpriseId
}

func (eh *EmployeeHired) OccurredOn() time.Time {
	return eh.occurredOn
}

func (eh *EmployeeHired) MarshalJSON() ([]byte, error) {
	j, err := json.Marshal(struct {
		EnterpriseId  uuid.UUID `json:"enterprise_name"`
		EmitedBy      uuid.UUID `json:"emited_by"`
		EmployeeId    uuid.UUID `json:"employee_id"`
		EmployeeEmail string    `json:"employee_email"`
		Position      string    `json:"position"`
		OccurredOn    string    `json:"occurred_on"`
	}{
		EnterpriseId:  eh.enterpriseId,
		EmitedBy:      eh.emitedBy,
		EmployeeId:    eh.employeeId,
		EmployeeEmail: eh.employeeEmail,
		Position:      eh.position.String(),
		OccurredOn:    common.StringDate(eh.occurredOn),
	})
	if err != nil {
		return nil, err
	}
	return j, nil
}

func (eh *EmployeeHired) UnmarshalJSON(data []byte) error {
	aux := struct {
		EnterpriseId  uuid.UUID `json:"enterprise_name"`
		EmitedBy      uuid.UUID `json:"emited_by"`
		EmployeeId    uuid.UUID `json:"employee_id"`
		EmployeeEmail string    `json:"employee_email"`
		Position      string    `json:"position"`
		OccurredOn    string    `json:"occurred_on"`
	}{}
	err := json.Unmarshal(data, &aux)
	if err != nil {
		return err
	}
	eh.emitedBy = aux.EmitedBy
	eh.employeeEmail = aux.EmployeeEmail
	eh.position = ParsePosition(aux.Position)
	eh.enterpriseId = aux.EnterpriseId
	eh.employeeId = aux.EmployeeId
	eh.occurredOn, err = common.ParseDate(aux.OccurredOn)
	if err != nil {
		return err
	}
	return nil
}
