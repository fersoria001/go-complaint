package complaint

import (
	"encoding/json"
	"go-complaint/domain/model/common"
	"time"

	"github.com/google/uuid"
)

type ComplaintClosed struct {
	complaintID uuid.UUID
	authorID    string
	triggeredBy string
	occurredOn  time.Time
}

func NewComplaintClosed(complaintID uuid.UUID, authorID, triggeredBy string) *ComplaintClosed {
	return &ComplaintClosed{
		complaintID: complaintID,
		authorID:    authorID,
		triggeredBy: triggeredBy,
		occurredOn:  time.Now(),
	}
}

func (ecc ComplaintClosed) OccurredOn() time.Time {
	return ecc.occurredOn
}

func (ecc ComplaintClosed) ComplaintID() uuid.UUID {
	return ecc.complaintID
}

func (ecc ComplaintClosed) AuthorID() string {
	return ecc.authorID
}

func (ecc ComplaintClosed) TriggeredBy() string {
	return ecc.triggeredBy
}

func (ecc *ComplaintClosed) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		ComplaintID string `json:"complaint_id"`
		AuthorID    string `json:"author_id"`
		TriggeredBy string `json:"triggered_by"`
		OccurredOn  string `json:"occurred_on"`
	}{
		ComplaintID: ecc.complaintID.String(),
		AuthorID:    ecc.authorID,
		TriggeredBy: ecc.triggeredBy,
		OccurredOn:  common.StringDate(ecc.occurredOn),
	})
}

func (ecc *ComplaintClosed) UnmarshalJSON(data []byte) error {
	aux := &struct {
		ComplaintID string `json:"complaint_id"`
		AuthorID    string `json:"author_id"`
		TriggeredBy string `json:"triggered_by"`
		OccurredOn  string `json:"occurred_on"`
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
	ecc.authorID = aux.AuthorID
	ecc.triggeredBy = aux.TriggeredBy
	ecc.occurredOn = occurredOn
	return nil
}
