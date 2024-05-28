package complaint

import (
	"encoding/json"
	"go-complaint/domain/model/common"
	"time"

	"github.com/google/uuid"
)

type WaitingForReview struct {
	complaintID  uuid.UUID
	enterpriseID string
	assistantID  string
	occurredOn   time.Time
}

func NewWaitingForReview(complaintID uuid.UUID, enterpriseID, assistantID string) *WaitingForReview {
	return &WaitingForReview{
		complaintID:  complaintID,
		enterpriseID: enterpriseID,
		assistantID:  assistantID,
		occurredOn:   time.Now(),
	}
}

func (wfr *WaitingForReview) OccurredOn() time.Time {
	return wfr.occurredOn
}

func (wfr *WaitingForReview) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		ComplaintID  string `json:"complaint_id"`
		EnterpriseID string `json:"enterprise_id"`
		AssistantID  string `json:"assistant_id"`
		OccurredOn   string `json:"occurred_on"`
	}{
		ComplaintID:  wfr.complaintID.String(),
		EnterpriseID: wfr.enterpriseID,
		AssistantID:  wfr.assistantID,
		OccurredOn:   common.StringDate(wfr.occurredOn),
	})
}

func (wfr *WaitingForReview) UnmarshalJSON(data []byte) error {
	aux := &struct {
		ComplaintID  string `json:"complaint_id"`
		EnterpriseID string `json:"enterprise_id"`
		AssistantID  string `json:"assistant_id"`
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
	wfr.complaintID = complaintID
	wfr.enterpriseID = aux.EnterpriseID
	wfr.assistantID = aux.AssistantID
	wfr.occurredOn = occurredOn
	return nil
}
