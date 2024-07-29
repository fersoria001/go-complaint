package identity

import (
	"encoding/json"
	"go-complaint/domain/model/common"
	"time"

	"github.com/google/uuid"
)

type RoleAdded struct {
	userId       uuid.UUID
	enterpriseId uuid.UUID
	roleID       string
	occurredOn   time.Time
}

func NewRoleAdded(
	userId,
	enterpriseId uuid.UUID,
	roleID string,
) *RoleAdded {
	return &RoleAdded{
		userId:       userId,
		enterpriseId: enterpriseId,
		roleID:       roleID,
		occurredOn:   time.Now(),
	}
}

func (ra *RoleAdded) OccurredOn() time.Time {
	return ra.occurredOn
}

func (ra *RoleAdded) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		UserId       uuid.UUID `json:"user_id"`
		EnterpriseId uuid.UUID `json:"enterprise_id"`
		RoleID       string    `json:"role_id"`
		OccurredOn   string    `json:"occurred_on"`
	}{
		UserId:       ra.userId,
		EnterpriseId: ra.enterpriseId,
		RoleID:       ra.roleID,
		OccurredOn:   common.StringDate(ra.occurredOn),
	})
}

func (ra *RoleAdded) UnmarshalJSON(data []byte) error {
	var aux struct {
		UserId       uuid.UUID `json:"user_id"`
		EnterpriseId uuid.UUID `json:"enterprise_id"`
		RoleID       string    `json:"role_id"`
		OccurredOn   string    `json:"occurred_on"`
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	ra.userId = aux.UserId
	ra.enterpriseId = aux.EnterpriseId
	ra.roleID = aux.RoleID
	ocurrOn, err := common.ParseDate(aux.OccurredOn)
	if err != nil {
		return err
	}
	ra.occurredOn = ocurrOn
	return nil
}
