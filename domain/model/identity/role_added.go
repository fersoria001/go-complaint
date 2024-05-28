package identity

import (
	"encoding/json"
	"go-complaint/domain/model/common"
	"time"
)

type RoleAdded struct {
	userID       string
	enterpriseID string
	roleID       string
	occurredOn   time.Time
}

func NewRoleAdded(
	userID,
	enterpriseID,
	roleID string,
) *RoleAdded {
	return &RoleAdded{
		userID:       userID,
		enterpriseID: enterpriseID,
		roleID:       roleID,
		occurredOn:   time.Now(),
	}
}

func (ra *RoleAdded) OccurredOn() time.Time {
	return ra.occurredOn
}

func (ra *RoleAdded) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		UserID       string `json:"user_id"`
		EnterpriseID string `json:"enterprise_id"`
		RoleID       string `json:"role_id"`
		OccurredOn   string `json:"occurred_on"`
	}{
		UserID:       ra.userID,
		EnterpriseID: ra.enterpriseID,
		RoleID:       ra.roleID,
		OccurredOn:   common.StringDate(ra.occurredOn),
	})
}

func (ra *RoleAdded) UnmarshalJSON(data []byte) error {
	var aux struct {
		UserID       string `json:"user_id"`
		EnterpriseID string `json:"enterprise_id"`
		RoleID       string `json:"role_id"`
		OccurredOn   string `json:"occurred_on"`
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	ra.userID = aux.UserID
	ra.enterpriseID = aux.EnterpriseID
	ra.roleID = aux.RoleID
	ocurrOn, err := common.ParseDate(aux.OccurredOn)
	if err != nil {
		return err
	}
	ra.occurredOn = ocurrOn
	return nil
}
