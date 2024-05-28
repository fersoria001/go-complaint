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
type AddedFeedback struct {
	feedbackID  uuid.UUID
	complaintID uuid.UUID
	reviewerID  string
	reviewedID  string
	occurredOn  time.Time
}

func NewAddedFeedback(
	feedbackID uuid.UUID,
	complaintID uuid.UUID,
	reviewerID string,
	reviewedID string,
) *AddedFeedback {
	return &AddedFeedback{
		feedbackID:  feedbackID,
		complaintID: complaintID,
		reviewerID:  reviewerID,
		reviewedID:  reviewedID,
		occurredOn:  time.Now(),
	}
}

func (f *AddedFeedback) OccurredOn() time.Time {
	return f.occurredOn
}

func (f *AddedFeedback) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		FeedbackID  string `json:"feedback_id"`
		ComplaintID string `json:"complaint_id"`
		ReviewerID  string `json:"reviewer_id"`
		ReviewedID  string `json:"reviewed_id"`
		OccurredOn  string `json:"occurred_on"`
	}{
		FeedbackID:  f.feedbackID.String(),
		ComplaintID: f.complaintID.String(),
		ReviewerID:  f.reviewerID,
		ReviewedID:  f.reviewedID,
		OccurredOn:  common.StringDate(f.occurredOn),
	})
}

func (f *AddedFeedback) UnmarshalJSON(data []byte) error {
	aux := struct {
		FeedbackID  string `json:"feedback_id"`
		ComplaintID string `json:"complaint_id"`
		ReviewerID  string `json:"reviewer_id"`
		ReviewedID  string `json:"reviewed_id"`
		OccurredOn  string `json:"occurred_on"`
	}{}

	err := json.Unmarshal(data, &aux)
	if err != nil {
		return err
	}
	feedbackID, err := uuid.Parse(aux.FeedbackID)
	if err != nil {
		return err
	}
	f.feedbackID = feedbackID
	complaintID, err := uuid.Parse(aux.ComplaintID)
	if err != nil {
		return err
	}
	f.complaintID = complaintID
	f.reviewerID = aux.ReviewerID
	f.reviewedID = aux.ReviewedID
	f.occurredOn, err = common.ParseDate(aux.OccurredOn)
	if err != nil {
		return err
	}

	return nil
}
