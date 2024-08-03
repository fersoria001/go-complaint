package enterprise

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type EmployeePromoted struct {
	enterpriseId     uuid.UUID
	managerId        uuid.UUID
	employeeUsername string
	prevPosition     Position
	newPosition      Position
	occurredOn       time.Time
}

func NewEmployeePromoted(enterpriseId, managerId uuid.UUID, employeeUsername string, prev, new Position) *EmployeePromoted {
	return &EmployeePromoted{
		enterpriseId:     enterpriseId,
		managerId:        managerId,
		employeeUsername: employeeUsername,
		prevPosition:     prev,
		newPosition:      new,
		occurredOn:       time.Now(),
	}
}

func (ep *EmployeePromoted) EnterpriseId() uuid.UUID {
	return ep.enterpriseId
}

func (ep *EmployeePromoted) ManagerId() uuid.UUID {
	return ep.managerId
}

func (ep *EmployeePromoted) EmployeeUsername() string {
	return ep.employeeUsername
}

func (ep *EmployeePromoted) PrevPosition() Position {
	return ep.prevPosition
}
func (ep *EmployeePromoted) NewPosition() Position {
	return ep.newPosition
}

func (ep *EmployeePromoted) OccurredOn() time.Time {
	return ep.occurredOn
}

func (ep *EmployeePromoted) MarshalJSON() ([]byte, error) {
	j, err := json.Marshal(struct {
		EnterpriseId     uuid.UUID
		ManagerId        uuid.UUID
		EmployeeUsername string
		PrevPosition     Position
		NewPosition      Position
		OccurredOn       time.Time
	}{
		EnterpriseId:     ep.enterpriseId,
		ManagerId:        ep.managerId,
		EmployeeUsername: ep.employeeUsername,
		PrevPosition:     ep.prevPosition,
		NewPosition:      ep.newPosition,
		OccurredOn:       ep.occurredOn,
	})
	if err != nil {
		return nil, err
	}
	return j, nil
}

func (ep *EmployeePromoted) UnmarshalJSON(data []byte) error {
	j := struct {
		EnterpriseId     uuid.UUID
		ManagerId        uuid.UUID
		EmployeeUsername string
		PrevPosition     Position
		NewPosition      Position
		OccurredOn       time.Time
	}{}
	err := json.Unmarshal(data, &j)
	if err != nil {
		return err
	}
	ep.enterpriseId = j.EnterpriseId
	ep.managerId = j.ManagerId
	ep.employeeUsername = j.EmployeeUsername
	ep.prevPosition = j.PrevPosition
	ep.newPosition = j.NewPosition
	ep.occurredOn = j.OccurredOn
	return nil
}
