package complaint

import (
	"encoding/json"
	"go-complaint/domain/model/common"
	"time"

	"github.com/google/uuid"
)

type EnterpriseComplaintClosed struct {
	complaintID  uuid.UUID
	enterpriseID string
	occurredOn   time.Time
}

func NewEnterpriseComplaintClosed(complaintID uuid.UUID, enterpriseID string) *EnterpriseComplaintClosed {
	return &EnterpriseComplaintClosed{
		complaintID:  complaintID,
		enterpriseID: enterpriseID,
		occurredOn:   time.Now(),
	}
}

func (ecc *EnterpriseComplaintClosed) OccurredOn() time.Time {

	return ecc.occurredOn
}

func (ecc *EnterpriseComplaintClosed) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		ComplaintID  string `json:"complaint_id"`
		EnterpriseID string `json:"enterprise_id"`
		OccurredOn   string `json:"occurred_on"`
	}{
		ComplaintID:  ecc.complaintID.String(),
		EnterpriseID: ecc.enterpriseID,
		OccurredOn:   common.StringDate(ecc.occurredOn),
	})
}

func (ecc *EnterpriseComplaintClosed) UnmarshalJSON(data []byte) error {
	aux := &struct {
		ComplaintID  string `json:"complaint_id"`
		EnterpriseID string `json:"enterprise_id"`
		OccurredOn   string `json:"occurred_on"`
	}{}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	occurredOn, err := common.ParseDate(aux.OccurredOn)
	if err != nil {
		return err
	}
	complaintID, err := uuid.Parse(aux.ComplaintID)
	if err != nil {
		return err
	}
	ecc.complaintID = complaintID
	ecc.enterpriseID = aux.EnterpriseID
	ecc.occurredOn = occurredOn
	return nil
}
