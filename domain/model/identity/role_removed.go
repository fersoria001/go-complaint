package identity

import (
	"encoding/json"
	"go-complaint/domain/model/common"
	"time"
)

type RoleRemoved struct {
	userID       string
	enterpriseID string
	roleID       string
	occurredOn   time.Time
}

func NewRoleRemoved(
	userID,
	enterpriseID,
	roleID string,
) *RoleRemoved {
	return &RoleRemoved{
		userID:       userID,
		enterpriseID: enterpriseID,
		roleID:       roleID,
		occurredOn:   time.Now(),
	}
}

func (rr *RoleRemoved) OccurredOn() time.Time {
	return rr.occurredOn
}

func (rr *RoleRemoved) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		UserID       string `json:"user_id"`
		EnterpriseID string `json:"enterprise_id"`
		RoleID       string `json:"role_id"`
		OccurredOn   string `json:"occurred_on"`
	}{
		UserID:       rr.userID,
		EnterpriseID: rr.enterpriseID,
		RoleID:       rr.roleID,
		OccurredOn:   common.StringDate(rr.occurredOn),
	})
}

func (rr *RoleRemoved) UnmarshalJSON(data []byte) error {
	var aux struct {
		UserID       string `json:"user_id"`
		EnterpriseID string `json:"enterprise_id"`
		RoleID       string `json:"role_id"`
		OccurredOn   string `json:"occurred_on"`
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	rr.userID = aux.UserID
	rr.enterpriseID = aux.EnterpriseID
	rr.roleID = aux.RoleID
	ocurrOn, err := common.ParseDate(aux.OccurredOn)
	if err != nil {
		return err
	}
	rr.occurredOn = ocurrOn
	return nil
}
