package enterprise

import (
	"encoding/json"
	"go-complaint/domain/model/common"
	"time"

	"github.com/google/uuid"
)

// Package enterprise
// <<Domain event>> implements domain.DomainEvent
type EmployeeLeaved struct {
	id               uuid.UUID
	enterpriseID     string
	userID           string
	hiringDate       time.Time
	approvedHiring   bool
	approvedHiringAt time.Time
	position         Position
	occurredOn       time.Time
}

func NewEmployeeLeaved(employee Employee) *EmployeeLeaved {
	return &EmployeeLeaved{
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

func (ef *EmployeeLeaved) ID() uuid.UUID {
	return ef.id
}

func (ef *EmployeeLeaved) EnterpriseID() string {
	return ef.enterpriseID
}

func (ef *EmployeeLeaved) UserID() string {
	return ef.userID
}

func (ef *EmployeeLeaved) HiringDate() time.Time {
	return ef.hiringDate
}

func (ef *EmployeeLeaved) ApprovedHiring() bool {
	return ef.approvedHiring
}

func (ef *EmployeeLeaved) ApprovedHiringAt() time.Time {
	return ef.approvedHiringAt
}

func (ef *EmployeeLeaved) Position() Position {
	return ef.position
}

func (ef *EmployeeLeaved) OccurredOn() time.Time {
	return ef.occurredOn
}

func (ef *EmployeeLeaved) MarshalJSON() ([]byte, error) {
	j, err := json.Marshal(struct {
		ID               string `json:"id"`
		EnterpriseID     string `json:"enterprise_id"`
		UserID           string `json:"user_id"`
		HiringDate       string `json:"hiring_date"`
		ApprovedHiring   bool   `json:"approved_hiring"`
		ApprovedHiringAt string `json:"approved_hiring_at"`
		Position         string `json:"position"`
		OccurredOn       string `json:"occurred_on"`
	}{
		ID:               ef.id.String(),
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

func (ef *EmployeeLeaved) UnmarshalJSON(data []byte) error {
	aux := struct {
		ID               string `json:"id"`
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
	ef.userID = aux.UserID
	ef.id, err = uuid.Parse(aux.ID)
	if err != nil {
		return err
	}
	ef.enterpriseID = aux.EnterpriseID

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
