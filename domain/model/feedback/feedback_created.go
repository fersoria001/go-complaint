package feedback

import (
	"encoding/json"
	"go-complaint/domain/model/common"
	"time"

	"github.com/google/uuid"
)

type FeedbackCreated struct {
	feedbackId   uuid.UUID
	complaintId  uuid.UUID
	enterpriseId uuid.UUID
	occurredOn   time.Time
}

func NewFeedbackCreated(
	feedbackId,
	complaintId,
	enterpriseId uuid.UUID,
) *FeedbackCreated {
	return &FeedbackCreated{
		feedbackId:   feedbackId,
		complaintId:  complaintId,
		enterpriseId: enterpriseId,
		occurredOn:   time.Now(),
	}
}

func (f FeedbackCreated) OccurredOn() time.Time {
	return f.occurredOn
}

func (f FeedbackCreated) FeedbackId() uuid.UUID {
	return f.feedbackId
}

func (f FeedbackCreated) ComplaintId() uuid.UUID {
	return f.complaintId
}

func (f FeedbackCreated) EnterpriseId() uuid.UUID {
	return f.enterpriseId
}

func (f *FeedbackCreated) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		FeedbackId   uuid.UUID `json:"feedback_id"`
		ComplaintId  uuid.UUID `json:"complaint_id"`
		EnterpriseId uuid.UUID `json:"enterprise_id"`
		OccurredOn   string    `json:"occurred_on"`
	}{
		FeedbackId:   f.feedbackId,
		ComplaintId:  f.complaintId,
		EnterpriseId: f.enterpriseId,
		OccurredOn:   common.StringDate(f.occurredOn),
	})
}

func (f *FeedbackCreated) UnmarshalJSON(data []byte) error {
	var feedbackCreated struct {
		FeedbackId   uuid.UUID `json:"feedback_id"`
		ComplaintId  uuid.UUID `json:"complaint_id"`
		EnterpriseId uuid.UUID `json:"enterprise_id"`
		OccurredOn   string    `json:"occurred_on"`
	}
	err := json.Unmarshal(data, &feedbackCreated)
	if err != nil {
		return err
	}
	f.occurredOn, err = common.ParseDate(feedbackCreated.OccurredOn)
	if err != nil {
		return err
	}
	f.feedbackId = feedbackCreated.FeedbackId
	f.complaintId = feedbackCreated.ComplaintId
	f.enterpriseId = feedbackCreated.EnterpriseId
	return nil
}
