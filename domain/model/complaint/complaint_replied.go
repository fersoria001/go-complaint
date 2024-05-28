package complaint

import (
	"encoding/json"
	"go-complaint/domain/model/common"
	"time"

	"github.com/google/uuid"
)

// Package complaint
// << Domain Event >>
type ComplaintReplied struct {
	complaintID uuid.UUID
	replyID     uuid.UUID
	occurredOn  time.Time
}

func NewComplaintReplied(
	complaintID,
	replyID uuid.UUID,
	occurredOn time.Time,
) *ComplaintReplied {
	return &ComplaintReplied{
		complaintID: complaintID,
		replyID:     replyID,
		occurredOn:  occurredOn,
	}
}

func (cr *ComplaintReplied) OccurredOn() time.Time {
	return cr.occurredOn
}

func (cr *ComplaintReplied) MarshalJSON() ([]byte, error) {
	j, err := json.Marshal(struct {
		ComplaintID string
		ReplyID     string
		Sender      string
		Receiver    string
		OccurredOn  string
		TypeName    string
	}{
		ComplaintID: cr.complaintID.String(),
		ReplyID:     cr.replyID.String(),
		OccurredOn:  common.StringDate(cr.occurredOn),
	})
	if err != nil {
		return nil, err
	}
	return j, nil
}

func (cr *ComplaintReplied) UnmarshalJSON(data []byte) error {
	var err error
	aux := struct {
		ComplaintID string
		ReplyID     string
		Sender      string
		Receiver    string
		OccurredOn  string
	}{}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	cr.complaintID, err = uuid.Parse(aux.ComplaintID)
	if err != nil {
		return err
	}
	cr.replyID, err = uuid.Parse(aux.ReplyID)
	if err != nil {
		return err
	}
	cr.occurredOn, err = common.ParseDate(aux.OccurredOn)
	if err != nil {
		return err
	}
	return nil
}
