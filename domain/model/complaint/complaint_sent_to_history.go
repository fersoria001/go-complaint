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
	complaintID uuid.UUID
	triggeredBy string
	occurredOn  time.Time
}

func NewComplaintSentToHistory(
	complaintID uuid.UUID,
	triggeredBy string,
) *ComplaintSentToHistory {
	return &ComplaintSentToHistory{
		complaintID: complaintID,
		triggeredBy: triggeredBy,
		occurredOn:  time.Now(),
	}
}

func (cr ComplaintSentToHistory) ComplaintID() uuid.UUID {
	return cr.complaintID
}

func (cr ComplaintSentToHistory) OccurredOn() time.Time {
	return cr.occurredOn
}

func (cr ComplaintSentToHistory) TriggeredBy() string {
	return cr.triggeredBy
}

func (cr *ComplaintSentToHistory) MarshalJSON() ([]byte, error) {
	j, err := json.Marshal(struct {
		ComplaintID string
		TriggeredBy string
		OccurredOn  string
	}{
		ComplaintID: cr.complaintID.String(),
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
		ComplaintID string
		TriggeredBy string
		OccurredOn  string
	}{}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	cr.triggeredBy = aux.TriggeredBy

	cr.complaintID, err = uuid.Parse(aux.ComplaintID)
	if err != nil {
		return err
	}
	cr.occurredOn, err = common.ParseDate(aux.OccurredOn)
	if err != nil {
		return err
	}
	return nil
}
