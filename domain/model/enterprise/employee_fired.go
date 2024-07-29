package enterprise

import (
	"encoding/json"
	"go-complaint/domain/model/common"
	"time"

	"github.com/google/uuid"
)

// Package enterprise
// <<Domain event>> implements domain.DomainEvent
type EmployeeFired struct {
	emitedBy         uuid.UUID
	id               uuid.UUID
	enterpriseId     uuid.UUID
	userId           uuid.UUID
	hiringDate       time.Time
	approvedHiring   bool
	approvedHiringAt time.Time
	position         Position
	occurredOn       time.Time
}

func NewEmployeeFired(emitedBy uuid.UUID, employee *Employee) *EmployeeFired {
	return &EmployeeFired{
		emitedBy:         emitedBy,
		id:               employee.Id(),
		enterpriseId:     employee.EnterpriseId(),
		userId:           employee.Id(),
		hiringDate:       employee.HiringDate().Date(),
		approvedHiring:   employee.ApprovedHiring(),
		approvedHiringAt: employee.ApprovedHiringAt().Date(),
		position:         employee.Position(),
		occurredOn:       time.Now(),
	}
}

func (ef *EmployeeFired) EmitedBy() uuid.UUID {
	return ef.emitedBy
}

func (ef *EmployeeFired) Id() uuid.UUID {
	return ef.id
}

func (ef *EmployeeFired) EnterpriseId() uuid.UUID {
	return ef.enterpriseId
}

func (ef *EmployeeFired) UserId() uuid.UUID {
	return ef.userId
}

func (ef *EmployeeFired) HiringDate() time.Time {
	return ef.hiringDate
}

func (ef *EmployeeFired) ApprovedHiring() bool {
	return ef.approvedHiring
}

func (ef *EmployeeFired) ApprovedHiringAt() time.Time {
	return ef.approvedHiringAt
}

func (ef *EmployeeFired) Position() Position {
	return ef.position
}

func (ef *EmployeeFired) OccurredOn() time.Time {
	return ef.occurredOn
}

func (ef *EmployeeFired) MarshalJSON() ([]byte, error) {
	j, err := json.Marshal(struct {
		Id               uuid.UUID `json:"id"`
		EmitedBy         uuid.UUID `json:"emited_by"`
		EnterpriseId     uuid.UUID `json:"enterprise_id"`
		UserId           uuid.UUID `json:"user_id"`
		HiringDate       string    `json:"hiring_date"`
		ApprovedHiring   bool      `json:"approved_hiring"`
		ApprovedHiringAt string    `json:"approved_hiring_at"`
		Position         string    `json:"position"`
		OccurredOn       string    `json:"occurred_on"`
	}{
		Id:               ef.id,
		EmitedBy:         ef.emitedBy,
		EnterpriseId:     ef.enterpriseId,
		UserId:           ef.userId,
		HiringDate:       common.StringDate(ef.hiringDate),
		ApprovedHiring:   ef.approvedHiring,
		ApprovedHiringAt: common.StringDate(ef.approvedHiringAt),
		Position:         ef.position.String(),
		OccurredOn:       common.StringDate(ef.occurredOn),
	})
	if err != nil {
		return nil, err
	}
	return j, nil
}

func (ef *EmployeeFired) UnmarshalJSON(data []byte) error {
	aux := struct {
		Id               uuid.UUID `json:"id"`
		EmitedBy         uuid.UUID `json:"emited_by"`
		EnterpriseId     uuid.UUID `json:"enterprise_id"`
		UserId           uuid.UUID `json:"user_id"`
		HiringDate       string    `json:"hiring_date"`
		ApprovedHiring   bool      `json:"approved_hiring"`
		ApprovedHiringAt string    `json:"approved_hiring_at"`
		Position         string    `json:"position"`
		OccurredOn       string    `json:"occurred_on"`
	}{}
	err := json.Unmarshal(data, &aux)
	if err != nil {
		return err
	}
	ef.id = aux.Id
	ef.emitedBy = aux.EmitedBy
	ef.enterpriseId = aux.EnterpriseId
	ef.userId = aux.UserId

	ef.hiringDate, err = common.ParseDate(aux.HiringDate)
	if err != nil {
		return err
	}
	ef.approvedHiring = aux.ApprovedHiring
	ef.approvedHiringAt, err = common.ParseDate(aux.ApprovedHiringAt)
	if err != nil {
		return err
	}
	position := ParsePosition(aux.Position)
	if position < 0 {
		return ErrPositionNotExists
	}
	ef.position = position
	ef.occurredOn, err = common.ParseDate(aux.OccurredOn)
	if err != nil {
		return err
	}
	return nil
}
