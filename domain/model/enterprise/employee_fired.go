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
	emitedBy         string
	id               uuid.UUID
	enterpriseID     string
	userID           string
	hiringDate       time.Time
	approvedHiring   bool
	approvedHiringAt time.Time
	position         Position
	occurredOn       time.Time
}

func NewEmployeeFired(emitedBy string, employee Employee) *EmployeeFired {
	return &EmployeeFired{
		emitedBy:         emitedBy,
		id:               employee.ID(),
		enterpriseID:     employee.EnterpriseID(),
		userID:           employee.Email(),
		hiringDate:       employee.HiringDate().Date(),
		approvedHiring:   employee.ApprovedHiring(),
		approvedHiringAt: employee.ApprovedHiringAt().Date(),
		position:         employee.Position(),
		occurredOn:       time.Now(),
	}
}

func (ef *EmployeeFired) EmitedBy() string {
	return ef.emitedBy
}

func (ef *EmployeeFired) ID() uuid.UUID {
	return ef.id
}

func (ef *EmployeeFired) EnterpriseID() string {
	return ef.enterpriseID
}

func (ef *EmployeeFired) UserID() string {
	return ef.userID
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
		ID               string `json:"id"`
		EmitedBy         string `json:"emited_by"`
		EnterpriseID     string `json:"enterprise_id"`
		UserID           string `json:"user_id"`
		HiringDate       string `json:"hiring_date"`
		ApprovedHiring   bool   `json:"approved_hiring"`
		ApprovedHiringAt string `json:"approved_hiring_at"`
		Position         string `json:"position"`
		OccurredOn       string `json:"occurred_on"`
	}{
		ID:               ef.id.String(),
		EmitedBy:         ef.emitedBy,
		EnterpriseID:     ef.enterpriseID,
		UserID:           ef.userID,
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
		ID               string `json:"id"`
		EmitedBy         string `json:"emited_by"`
		EnterpriseID     string `json:"enterprise_id"`
		UserID           string `json:"user_id"`
		HiringDate       string `json:"hiring_date"`
		ApprovedHiring   bool   `json:"approved_hiring"`
		ApprovedHiringAt string `json:"approved_hiring_at"`
		Position         string `json:"position"`
		OccurredOn       string `json:"occurred_on"`
	}{}
	err := json.Unmarshal(data, &aux)
	if err != nil {
		return err
	}
	ef.id, err = uuid.Parse(aux.ID)
	if err != nil {
		return err
	}
	ef.emitedBy = aux.EmitedBy
	ef.enterpriseID = aux.EnterpriseID
	ef.userID = aux.UserID

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
	if position == NOT_EXISTS {
		return ErrPositionNotExists
	}
	ef.position = position
	ef.occurredOn, err = common.ParseDate(aux.OccurredOn)
	if err != nil {
		return err
	}
	return nil
}
