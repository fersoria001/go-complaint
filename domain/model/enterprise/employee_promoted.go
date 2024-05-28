package enterprise

import (
	"encoding/json"
	"go-complaint/domain/model/common"
	"time"
)

type EmployeePromoted struct {
	enterpriseID string
	managerID    string
	employeeID   string
	position     Position
	occurredOn   time.Time
}

func NewEmployeePromoted(enterpriseID, managerID, employeeID string, position Position) *EmployeePromoted {
	return &EmployeePromoted{
		enterpriseID: enterpriseID,
		managerID:    managerID,
		employeeID:   employeeID,
		position:     position,
		occurredOn:   time.Now(),
	}
}

func (ep *EmployeePromoted) OccurredOn() time.Time {
	return ep.occurredOn
}

func (ep *EmployeePromoted) MarshalJSON() ([]byte, error) {
	j, err := json.Marshal(struct {
		EnterpriseID string
		ManagerID    string
		EmployeeID   string
		Position     string
		OccurredOn   string
	}{
		EnterpriseID: ep.enterpriseID,
		ManagerID:    ep.managerID,
		EmployeeID:   ep.employeeID,
		Position:     ep.position.String(),
		OccurredOn:   common.StringDate(ep.occurredOn),
	})
	if err != nil {
		return nil, err
	}
	return j, nil
}

func (ep *EmployeePromoted) UnmarshalJSON(data []byte) error {
	j := struct {
		EnterpriseID string
		ManagerID    string
		EmployeeID   string
		Position     string
		OccurredOn   string
	}{}
	err := json.Unmarshal(data, &j)
	if err != nil {
		return err
	}
	ep.enterpriseID = j.EnterpriseID
	ep.managerID = j.ManagerID
	ep.employeeID = j.EmployeeID
	ep.position, err = ParsePosition(j.Position)
	if err != nil {
		return err
	}
	occurredOn, err := common.ParseDate(j.OccurredOn)
	if err != nil {
		return err
	}
	ep.occurredOn = occurredOn
	return nil
}
