package complaint

import (
	"encoding/json"
	"go-complaint/domain/model/common"
	"time"

	"github.com/google/uuid"
)

// Package complaint
// << Domain Event >>
type ComplaintRated struct {
	complaintID     uuid.UUID
	ratedBy         string
	assistantUserID string
	occurredOn      time.Time
}

func NewComplaintRated(
	complaintID uuid.UUID,
	ratedBy string,
	assistantUserID string,
	occurredOn time.Time,
) *ComplaintRated {
	return &ComplaintRated{
		complaintID:     complaintID,
		ratedBy:         ratedBy,
		assistantUserID: assistantUserID,
		occurredOn:      occurredOn,
	}
}

func (cr *ComplaintRated) OccurredOn() time.Time {
	return cr.occurredOn
}

func (cr *ComplaintRated) MarshalJSON() ([]byte, error) {
	j, err := json.Marshal(struct {
		ComplaintID     string
		RatedBy         string
		AssistantUserID string
		OccurredOn      string
		TypeName        string
	}{
		ComplaintID:     cr.complaintID.String(),
		RatedBy:         cr.ratedBy,
		AssistantUserID: cr.assistantUserID,
		OccurredOn:      common.StringDate(cr.occurredOn),
	})
	if err != nil {
		return nil, err
	}
	return j, nil
}

func (cr *ComplaintRated) UnmarshalJSON(data []byte) error {
	var err error
	aux := struct {
		ComplaintID     string
		RatedBy         string
		AssistantUserID string
		OccurredOn      string
	}{}
	if err = json.Unmarshal(data, &aux); err != nil {
		return err
	}
	cr.complaintID, err = uuid.Parse(aux.ComplaintID)
	if err != nil {
		return err
	}
	cr.assistantUserID = aux.AssistantUserID
	cr.ratedBy = aux.RatedBy
	cr.occurredOn, err = common.ParseDate(aux.OccurredOn)
	if err != nil {
		return err
	}
	return nil
}

func (cr *ComplaintRated) ComplaintID() uuid.UUID {
	return cr.complaintID
}

func (cr *ComplaintRated) RatedBy() string {
	return cr.ratedBy
}

func (cr *ComplaintRated) AssistantUserID() string {
	return cr.assistantUserID
}
