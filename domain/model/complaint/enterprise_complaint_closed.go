package complaint

import (
	"encoding/json"
	"go-complaint/domain/model/common"
	"time"

	"github.com/google/uuid"
)

type ComplaintClosed struct {
	complaintId uuid.UUID
	authorId    uuid.UUID
	triggeredBy uuid.UUID
	occurredOn  time.Time
}

func NewComplaintClosed(complaintId, authorId uuid.UUID, triggeredBy uuid.UUID) *ComplaintClosed {
	return &ComplaintClosed{
		complaintId: complaintId,
		authorId:    authorId,
		triggeredBy: triggeredBy,
		occurredOn:  time.Now(),
	}
}

func (ecc ComplaintClosed) OccurredOn() time.Time {
	return ecc.occurredOn
}

func (ecc ComplaintClosed) ComplaintId() uuid.UUID {
	return ecc.complaintId
}

func (ecc ComplaintClosed) AuthorId() uuid.UUID {
	return ecc.authorId
}

func (ecc ComplaintClosed) TriggeredBy() uuid.UUID {
	return ecc.triggeredBy
}

func (ecc *ComplaintClosed) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		ComplaintId uuid.UUID `json:"complaint_id"`
		AuthorId    uuid.UUID `json:"author_id"`
		TriggeredBy uuid.UUID `json:"triggered_by"`
		OccurredOn  string    `json:"occurred_on"`
	}{
		ComplaintId: ecc.complaintId,
		AuthorId:    ecc.authorId,
		TriggeredBy: ecc.triggeredBy,
		OccurredOn:  common.StringDate(ecc.occurredOn),
	})
}

func (ecc *ComplaintClosed) UnmarshalJSON(data []byte) error {
	aux := &struct {
		ComplaintId uuid.UUID `json:"complaint_id"`
		AuthorId    uuid.UUID `json:"author_id"`
		TriggeredBy uuid.UUID `json:"triggered_by"`
		OccurredOn  string    `json:"occurred_on"`
	}{}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	occurredOn, err := common.ParseDate(aux.OccurredOn)
	if err != nil {
		return err
	}
	ecc.complaintId = aux.ComplaintId
	ecc.authorId = aux.AuthorId
	ecc.triggeredBy = aux.TriggeredBy
	ecc.occurredOn = occurredOn
	return nil
}
