package complaint

import (
	"encoding/json"
	"go-complaint/domain/model/common"
	"time"

	"github.com/google/uuid"
)

// Package complaint
// << Domain Event >>
type ComplaintSentToHistory struct {
	complaintId uuid.UUID
	triggeredBy uuid.UUID
	occurredOn  time.Time
}

func NewComplaintSentToHistory(
	complaintId uuid.UUID,
	triggeredBy uuid.UUID,
) *ComplaintSentToHistory {
	return &ComplaintSentToHistory{
		complaintId: complaintId,
		triggeredBy: triggeredBy,
		occurredOn:  time.Now(),
	}
}

func (cr ComplaintSentToHistory) ComplaintId() uuid.UUID {
	return cr.complaintId
}

func (cr ComplaintSentToHistory) OccurredOn() time.Time {
	return cr.occurredOn
}

func (cr ComplaintSentToHistory) TriggeredBy() uuid.UUID {
	return cr.triggeredBy
}

func (cr *ComplaintSentToHistory) MarshalJSON() ([]byte, error) {
	j, err := json.Marshal(struct {
		ComplaintId uuid.UUID
		TriggeredBy uuid.UUID
		OccurredOn  string
	}{
		ComplaintId: cr.complaintId,
		TriggeredBy: cr.triggeredBy,
		OccurredOn:  common.StringDate(cr.occurredOn),
	})
	if err != nil {
		return nil, err
	}
	return j, nil
}

func (cr *ComplaintSentToHistory) UnmarshalJSON(data []byte) error {
	var err error
	aux := struct {
		ComplaintId uuid.UUID
		TriggeredBy uuid.UUID
		OccurredOn  string
	}{}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	cr.triggeredBy = aux.TriggeredBy
	cr.complaintId = aux.ComplaintId
	cr.occurredOn, err = common.ParseDate(aux.OccurredOn)
	if err != nil {
		return err
	}
	return nil
}
