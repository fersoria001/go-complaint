package identity

import (
	"encoding/json"
	"go-complaint/domain/model/common"
	"time"

	"github.com/google/uuid"
)

type RoleRemoved struct {
	userId       uuid.UUID
	enterpriseId uuid.UUID
	roleID       string
	occurredOn   time.Time
}

func NewRoleRemoved(
	userId,
	enterpriseId uuid.UUID,
	roleID string,
) *RoleRemoved {
	return &RoleRemoved{
		userId:       userId,
		enterpriseId: enterpriseId,
		roleID:       roleID,
		occurredOn:   time.Now(),
	}
}

func (rr *RoleRemoved) UserId() uuid.UUID {
	return rr.userId
}
func (rr *RoleRemoved) EnterpriseId() uuid.UUID {
	return rr.enterpriseId
}
func (rr *RoleRemoved) RoleId() string {
	return rr.roleID
}
func (rr *RoleRemoved) OccurredOn() time.Time {
	return rr.occurredOn
}

func (rr *RoleRemoved) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		UserId       uuid.UUID `json:"user_id"`
		EnterpriseId uuid.UUID `json:"enterprise_id"`
		RoleID       string    `json:"role_id"`
		OccurredOn   string    `json:"occurred_on"`
	}{
		UserId:       rr.userId,
		EnterpriseId: rr.enterpriseId,
		RoleID:       rr.roleID,
		OccurredOn:   common.StringDate(rr.occurredOn),
	})
}

func (rr *RoleRemoved) UnmarshalJSON(data []byte) error {
	var aux struct {
		UserId       uuid.UUID `json:"user_id"`
		EnterpriseId uuid.UUID `json:"enterprise_id"`
		RoleID       string    `json:"role_id"`
		OccurredOn   string    `json:"occurred_on"`
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	rr.userId = aux.UserId
	rr.enterpriseId = aux.EnterpriseId
	rr.roleID = aux.RoleID
	ocurrOn, err := common.ParseDate(aux.OccurredOn)
	if err != nil {
		return err
	}
	rr.occurredOn = ocurrOn
	return nil
}
