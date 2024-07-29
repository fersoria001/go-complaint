package complaint

import (
	"encoding/json"
	"go-complaint/domain/model/common"
	"time"

	"github.com/google/uuid"
)

type ComplaintStarted struct {
	complaintId uuid.UUID
	replyId     uuid.UUID
	occurredOn  time.Time
}

func NewComplaintStarted(complaintId, replyId uuid.UUID, occuredOn time.Time) *ComplaintStarted {
	return &ComplaintStarted{
		complaintId: complaintId,
		replyId:     replyId,
		occurredOn:  occuredOn,
	}
}

func (cs *ComplaintStarted) OccurredOn() time.Time {
	return cs.occurredOn
}

func (cs *ComplaintStarted) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ComplaintId uuid.UUID `json:"complaint_id"`
		ReplyId     uuid.UUID `json:"reply_id"`
		OccurredOn  string    `json:"occurred_on"`
	}{
		ComplaintId: cs.complaintId,
		ReplyId:     cs.replyId,
		OccurredOn:  common.StringDate(cs.occurredOn),
	})
}

func (cs *ComplaintStarted) UnmarshalJSON(data []byte) error {
	var err error
	var aux struct {
		ComplaintId uuid.UUID `json:"complaint_id"`
		ReplyId     uuid.UUID `json:"reply_id"`
		OccurredOn  string    `json:"occurred_on"`
	}
	if err = json.Unmarshal(data, &aux); err != nil {
		return err
	}
	cs.complaintId = aux.ComplaintId
	cs.replyId = aux.ReplyId
	occurredOn, err := common.ParseDate(aux.OccurredOn)
	if err != nil {
		return err
	}
	cs.occurredOn = occurredOn
	return nil
}
