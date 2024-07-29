package feedback

import (
	"encoding/json"
	"go-complaint/domain/model/common"
	"time"

	"github.com/google/uuid"
)

/*
Package feedback
implements domain.DomainEvent interface
*/
type FeedbackReplied struct {
	feedbackID  uuid.UUID
	complaintID uuid.UUID
	senderID    uuid.UUID
	answerID    uuid.UUID
	occurredOn  time.Time
}

func NewFeedbackReplied(
	feedbackID uuid.UUID,
	complaintID uuid.UUID,
	senderID uuid.UUID,
	answerID uuid.UUID,
) *FeedbackReplied {
	return &FeedbackReplied{
		feedbackID:  feedbackID,
		complaintID: complaintID,
		senderID:    senderID,
		answerID:    answerID,
		occurredOn:  time.Now(),
	}
}

func (f *FeedbackReplied) OccurredOn() time.Time {
	return f.occurredOn
}

func (f *FeedbackReplied) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		FeedbackID  uuid.UUID `json:"feedback_id"`
		ComplaintID uuid.UUID `json:"complaint_id"`
		SenderID    uuid.UUID `json:"reviewed_id"`
		AnswerID    uuid.UUID `json:"answer_id"`
		OccurredOn  string    `json:"occurred_on"`
	}{
		FeedbackID:  f.feedbackID,
		ComplaintID: f.complaintID,
		SenderID:    f.senderID,
		AnswerID:    f.answerID,
		OccurredOn:  common.StringDate(f.occurredOn),
	})
}

func (f *FeedbackReplied) UnmarshalJSON(data []byte) error {
	aux := struct {
		FeedbackID  uuid.UUID `json:"feedback_id"`
		ComplaintID uuid.UUID `json:"complaint_id"`
		SenderID    uuid.UUID `json:"reviewed_id"`
		AnswerID    uuid.UUID `json:"answer_id"`
		OccurredOn  string    `json:"occurred_on"`
	}{}
	err := json.Unmarshal(data, &aux)
	if err != nil {
		return err
	}
	f.feedbackID = aux.FeedbackID
	f.complaintID = aux.ComplaintID
	f.senderID = aux.SenderID
	f.answerID = aux.AnswerID
	f.occurredOn, err = common.ParseDate(aux.OccurredOn)
	if err != nil {
		return err
	}
	return nil
}
