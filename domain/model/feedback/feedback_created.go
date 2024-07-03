package feedback

import (
	"encoding/json"
	"go-complaint/domain/model/common"
	"time"

	"github.com/google/uuid"
)

type FeedbackCreated struct {
	feedbackID   uuid.UUID
	complaintID  uuid.UUID
	enterpriseID string
	occurredOn   time.Time
}

func NewFeedbackCreated(
	feedbackID uuid.UUID,
	complaintID uuid.UUID,
	enterpriseID string,
) *FeedbackCreated {
	return &FeedbackCreated{
		feedbackID:   feedbackID,
		complaintID:  complaintID,
		enterpriseID: enterpriseID,
		occurredOn:   time.Now(),
	}
}

func (f FeedbackCreated) OccurredOn() time.Time {
	return f.occurredOn
}

func (f FeedbackCreated) FeedbackID() uuid.UUID {
	return f.feedbackID
}

func (f FeedbackCreated) ComplaintID() uuid.UUID {
	return f.complaintID
}

func (f FeedbackCreated) EnterpriseID() string {
	return f.enterpriseID
}

func (f *FeedbackCreated) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		FeedbackID   string `json:"feedback_id"`
		ComplaintID  string `json:"complaint_id"`
		EnterpriseID string `json:"enterprise_id"`
		OccurredOn   string `json:"occurred_on"`
	}{
		FeedbackID:   f.feedbackID.String(),
		ComplaintID:  f.complaintID.String(),
		EnterpriseID: f.enterpriseID,
		OccurredOn:   common.StringDate(f.occurredOn),
	})
}

func (f *FeedbackCreated) UnmarshalJSON(data []byte) error {
	var feedbackCreated struct {
		FeedbackID   string `json:"feedback_id"`
		ComplaintID  string `json:"complaint_id"`
		EnterpriseID string `json:"enterprise_id"`
		OccurredOn   string `json:"occurred_on"`
	}
	err := json.Unmarshal(data, &feedbackCreated)
	if err != nil {
		return err
	}
	f.occurredOn, err = common.ParseDate(feedbackCreated.OccurredOn)
	if err != nil {
		return err
	}
	f.feedbackID, err = uuid.Parse(feedbackCreated.FeedbackID)
	if err != nil {
		return err
	}
	f.complaintID, err = uuid.Parse(feedbackCreated.ComplaintID)
	if err != nil {
		return err
	}
	f.enterpriseID = feedbackCreated.EnterpriseID
	return nil
}
