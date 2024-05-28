package enterprise

import (
	"encoding/json"
	"go-complaint/domain/model/common"
	"time"
)

//Package enterprise
//<<Domain event>> Implements domain.DomainEvent
type EmployeeHired struct {
	enterpriseName string
	employeeID  string
	managerID  string
	occurredOn time.Time
}

func NewEmployeeHired(enterpriseName, employeeID, managerID string) *EmployeeHired {
	return &EmployeeHired{
		enterpriseName: enterpriseName,
		employeeID: employeeID,
		managerID: managerID,
		occurredOn: time.Now(),
	}
}

func (eh *EmployeeHired) OccurredOn() time.Time {
	return eh.occurredOn
}

func (eh *EmployeeHired) MarshalJSON() ([]byte, error) {
	j, err := json.Marshal(struct {
		EnterpriseName string `json:"enterprise_name"`
		EmployeeID string `json:"employee_id"`
		ManagerID string `json:"manager_id"`
		OccurredOn string `json:"occurred_on"`
	}{
		EnterpriseName: eh.enterpriseName,
		EmployeeID: eh.employeeID,
		ManagerID: eh.managerID,
		OccurredOn: common.StringDate(eh.occurredOn),
	})
	if err != nil {
		return nil, err
	}
	return j, nil
}

func (eh *EmployeeHired) UnmarshalJSON(data []byte) error {
	aux := struct {
		EnterpriseName string `json:"enterprise_name"`
		EmployeeID string `json:"employee_id"`
		ManagerID string `json:"manager_id"`
		OccurredOn string `json:"occurred_on"`
	}{}
	err := json.Unmarshal(data, &aux)
	if err != nil {
		return err
	}
	eh.enterpriseName = aux.EnterpriseName
	eh.employeeID = aux.EmployeeID
	eh.managerID = aux.ManagerID
	eh.occurredOn, err = common.ParseDate(aux.OccurredOn)
	if err != nil {
		return err
	}
	return nil
}