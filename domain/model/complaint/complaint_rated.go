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
	complaintId     uuid.UUID
	ratedBy         uuid.UUID
	assistantUserId uuid.UUID
	occurredOn      time.Time
}

func NewComplaintRated(
	complaintId uuid.UUID,
	ratedBy uuid.UUID,
	assistantUserId uuid.UUID,
	occurredOn time.Time,
) *ComplaintRated {
	return &ComplaintRated{
		complaintId:     complaintId,
		ratedBy:         ratedBy,
		assistantUserId: assistantUserId,
		occurredOn:      occurredOn,
	}
}

func (cr *ComplaintRated) OccurredOn() time.Time {
	return cr.occurredOn
}

func (cr *ComplaintRated) MarshalJSON() ([]byte, error) {
	j, err := json.Marshal(struct {
		ComplaintId     uuid.UUID
		RatedBy         uuid.UUID
		AssistantUserId uuid.UUID
		OccurredOn      string
	}{
		ComplaintId:     cr.complaintId,
		RatedBy:         cr.ratedBy,
		AssistantUserId: cr.assistantUserId,
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
		ComplaintId     uuid.UUID
		RatedBy         uuid.UUID
		AssistantUserId uuid.UUID
		OccurredOn      string
	}{}
	if err = json.Unmarshal(data, &aux); err != nil {
		return err
	}
	cr.complaintId = aux.ComplaintId
	cr.assistantUserId = aux.AssistantUserId
	cr.ratedBy = aux.RatedBy
	cr.occurredOn, err = common.ParseDate(aux.OccurredOn)
	if err != nil {
		return err
	}
	return nil
}

func (cr *ComplaintRated) ComplaintId() uuid.UUID {
	return cr.complaintId
}

func (cr *ComplaintRated) RatedBy() uuid.UUID {
	return cr.ratedBy
}

func (cr *ComplaintRated) AssistantUserId() uuid.UUID {
	return cr.assistantUserId
}
