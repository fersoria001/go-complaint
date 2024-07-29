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
	feedbackId  uuid.UUID
	complaintId uuid.UUID
	reviewerId  uuid.UUID
	reviewedId  uuid.UUID
	occurredOn  time.Time
}

func NewAddedFeedback(
	feedbackId,
	complaintId,
	reviewerId,
	reviewedId uuid.UUID,
) *AddedFeedback {
	return &AddedFeedback{
		feedbackId:  feedbackId,
		complaintId: complaintId,
		reviewerId:  reviewerId,
		reviewedId:  reviewedId,
		occurredOn:  time.Now(),
	}
}

func (f *AddedFeedback) OccurredOn() time.Time {
	return f.occurredOn
}

func (f *AddedFeedback) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		FeedbackId  uuid.UUID `json:"feedback_id"`
		ComplaintId uuid.UUID `json:"complaint_id"`
		ReviewerId  uuid.UUID `json:"reviewer_id"`
		ReviewedId  uuid.UUID `json:"reviewed_id"`
		OccurredOn  string    `json:"occurred_on"`
	}{
		FeedbackId:  f.feedbackId,
		ComplaintId: f.complaintId,
		ReviewerId:  f.reviewerId,
		ReviewedId:  f.reviewedId,
		OccurredOn:  common.StringDate(f.occurredOn),
	})
}

func (f *AddedFeedback) UnmarshalJSON(data []byte) error {
	aux := struct {
		FeedbackId  uuid.UUID `json:"feedback_id"`
		ComplaintId uuid.UUID `json:"complaint_id"`
		ReviewerId  uuid.UUID `json:"reviewer_id"`
		ReviewedId  uuid.UUID `json:"reviewed_id"`
		OccurredOn  string    `json:"occurred_on"`
	}{}
	err := json.Unmarshal(data, &aux)
	if err != nil {
		return err
	}
	f.feedbackId = aux.FeedbackId
	f.complaintId = aux.ComplaintId
	f.reviewerId = aux.ReviewerId
	f.reviewedId = aux.ReviewedId
	f.occurredOn, err = common.ParseDate(aux.OccurredOn)
	if err != nil {
		return err
	}
	return nil
}

func (f AddedFeedback) FeedbackId() uuid.UUID {
	return f.feedbackId
}

func (f AddedFeedback) ComplaintId() uuid.UUID {
	return f.complaintId
}

func (f AddedFeedback) ReviewerId() uuid.UUID {
	return f.reviewerId
}

func (f AddedFeedback) ReviewedId() uuid.UUID {
	return f.reviewedId
}
