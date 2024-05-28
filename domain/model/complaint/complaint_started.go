package complaint

import (
	"encoding/json"
	"go-complaint/domain/model/common"
	"time"

	"github.com/google/uuid"
)

type ComplaintStarted struct {
	complaintID uuid.UUID
	replyID     uuid.UUID
	occurredOn  time.Time
}

func NewComplaintStarted(complaintID, replyID uuid.UUID, occuredOn time.Time) *ComplaintStarted {
	return &ComplaintStarted{
		complaintID: complaintID,
		replyID:     replyID,
		occurredOn:  occuredOn,
	}
}

func (cs *ComplaintStarted) OccurredOn() time.Time {
	return cs.occurredOn
}

func (cs *ComplaintStarted) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ComplaintID string `json:"complaint_id"`
		ReplyID     string `json:"reply_id"`
		OccurredOn  string `json:"occurred_on"`
	}{
		ComplaintID: cs.complaintID.String(),
		ReplyID:     cs.replyID.String(),
		OccurredOn:  common.StringDate(cs.occurredOn),
	})
}

func (cs *ComplaintStarted) UnmarshalJSON(data []byte) error {
	var err error
	var aux struct {
		ComplaintID string `json:"complaint_id"`
		ReplyID     string `json:"reply_id"`
		OccurredOn  string `json:"occurred_on"`
	}
	if err = json.Unmarshal(data, &aux); err != nil {
		return err
	}
	cs.complaintID, err = uuid.Parse(aux.ComplaintID)
	if err != nil {
		return err
	}
	cs.replyID, err = uuid.Parse(aux.ReplyID)
	if err != nil {
		return err
	}
	occurredOn, err := common.ParseDate(aux.OccurredOn)
	if err != nil {
		return err
	}
	cs.occurredOn = occurredOn
	return nil
}
