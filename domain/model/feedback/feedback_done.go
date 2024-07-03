package feedback

import (
	"encoding/json"
	"go-complaint/domain/model/common"
	"time"

	"github.com/google/uuid"
)

type FeedbackDone struct {
	feedbackID   uuid.UUID
	complaintID  uuid.UUID
	enterpriseID string
	reviewedID   string
	occurredOn   time.Time
}

func NewFeedbackDone(
	feedbackID,
	complaintID uuid.UUID,
	enterpriseID string,
	reviewedID string,
	occurredOn time.Time,
) *FeedbackDone {
	return &FeedbackDone{
		feedbackID:   feedbackID,
		complaintID:  complaintID,
		enterpriseID: enterpriseID,
		reviewedID:   reviewedID,
		occurredOn:   occurredOn,
	}
}

func (f *FeedbackDone) ReviewedID() string {
	return f.reviewedID
}

func (f *FeedbackDone) FeedbackID() uuid.UUID {
	return f.feedbackID
}

func (f *FeedbackDone) ComplaintID() uuid.UUID {
	return f.complaintID
}

func (f *FeedbackDone) EnterpriseID() string {
	return f.enterpriseID
}

func (f *FeedbackDone) OccurredOn() time.Time {
	return f.occurredOn
}

func (f *FeedbackDone) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		FeedbackID   string `json:"feedback_id"`
		ComplaintID  string `json:"complaint_id"`
		EnterpriseID string `json:"enterprise_id"`
		ReviewedID   string `json:"reviewed_id"`
		OccurredOn   string `json:"occurred_on"`
	}{
		FeedbackID:   f.feedbackID.String(),
		ComplaintID:  f.complaintID.String(),
		EnterpriseID: f.enterpriseID,
		ReviewedID:   f.reviewedID,
		OccurredOn:   common.StringDate(f.occurredOn),
	})
}

func (f *FeedbackDone) UnmarshalJSON(data []byte) error {
	aux := struct {
		FeedbackID   string `json:"feedback_id"`
		ComplaintID  string `json:"complaint_id"`
		EnterpriseID string `json:"enterprise_id"`
		ReviewedID   string `json:"reviewed_id"`
		OccurredOn   string `json:"occurred_on"`
	}{}
	err := json.Unmarshal(data, &aux)
	if err != nil {
		return err
	}
	f.reviewedID = aux.ReviewedID
	f.feedbackID, err = uuid.Parse(aux.FeedbackID)
	if err != nil {
		return err
	}
	f.complaintID, err = uuid.Parse(aux.ComplaintID)
	if err != nil {
		return err
	}
	f.enterpriseID = aux.EnterpriseID
	f.occurredOn, err = common.ParseDate(aux.OccurredOn)
	if err != nil {
		return err
	}
	return nil
}
