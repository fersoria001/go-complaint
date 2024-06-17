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
	reviewedID  string
	answerID    uuid.UUID
	occurredOn  time.Time
}

func NewFeedbackReplied(
	feedbackID uuid.UUID,
	complaintID uuid.UUID,
	reviewedID string,
	answerID uuid.UUID,
) *FeedbackReplied {
	return &FeedbackReplied{
		feedbackID:  feedbackID,
		complaintID: complaintID,
		reviewedID:  reviewedID,
		answerID:    answerID,
		occurredOn:  time.Now(),
	}
}

func (f *FeedbackReplied) OccurredOn() time.Time {
	return f.occurredOn
}

func (f *FeedbackReplied) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		FeedbackID  string `json:"feedback_id"`
		ComplaintID string `json:"complaint_id"`
		ReviewedID  string `json:"reviewed_id"`
		AnswerID    string `json:"answer_id"`
		OccurredOn  string `json:"occurred_on"`
	}{
		FeedbackID:  f.feedbackID.String(),
		ComplaintID: f.complaintID.String(),
		ReviewedID:  f.reviewedID,
		AnswerID:    f.answerID.String(),
		OccurredOn:  common.StringDate(f.occurredOn),
	})
}

func (f *FeedbackReplied) UnmarshalJSON(data []byte) error {
	aux := struct {
		FeedbackID  string `json:"feedback_id"`
		ComplaintID string `json:"complaint_id"`
		ReviewedID  string `json:"reviewed_id"`
		AnswerID    string `json:"answer_id"`
		OccurredOn  string `json:"occurred_on"`
	}{}
	err := json.Unmarshal(data, &aux)
	if err != nil {
		return err
	}
	f.feedbackID, err = uuid.Parse(aux.FeedbackID)
	if err != nil {
		return err
	}
	f.complaintID, err = uuid.Parse(aux.ComplaintID)
	if err != nil {
		return err
	}

	f.reviewedID = aux.ReviewedID
	f.answerID, err = uuid.Parse(aux.AnswerID)
	if err != nil {
		return err
	}
	f.occurredOn, err = common.ParseDate(aux.OccurredOn)
	if err != nil {
		return err
	}
	return nil
}
