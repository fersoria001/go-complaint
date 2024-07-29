package feedback

import (
	"encoding/json"
	"go-complaint/domain/model/common"
	"time"

	"github.com/google/uuid"
)

type FeedbackDone struct {
	feedbackId   uuid.UUID
	complaintId  uuid.UUID
	enterpriseId uuid.UUID
	reviewedId   uuid.UUID
	occurredOn   time.Time
}

func NewFeedbackDone(
	feedbackId,
	complaintId,
	enterpriseId,
	reviewedId uuid.UUID,
	occurredOn time.Time,
) *FeedbackDone {
	return &FeedbackDone{
		feedbackId:   feedbackId,
		complaintId:  complaintId,
		enterpriseId: enterpriseId,
		reviewedId:   reviewedId,
		occurredOn:   occurredOn,
	}
}

func (f *FeedbackDone) ReviewedId() uuid.UUID {
	return f.reviewedId
}

func (f *FeedbackDone) FeedbackId() uuid.UUID {
	return f.feedbackId
}

func (f *FeedbackDone) ComplaintId() uuid.UUID {
	return f.complaintId
}

func (f *FeedbackDone) EnterpriseId() uuid.UUID {
	return f.enterpriseId
}

func (f *FeedbackDone) OccurredOn() time.Time {
	return f.occurredOn
}

func (f *FeedbackDone) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		FeedbackId   uuid.UUID `json:"feedback_id"`
		ComplaintId  uuid.UUID `json:"complaint_id"`
		EnterpriseId uuid.UUID `json:"enterprise_id"`
		ReviewedId   uuid.UUID `json:"reviewed_id"`
		OccurredOn   string    `json:"occurred_on"`
	}{
		FeedbackId:   f.feedbackId,
		ComplaintId:  f.complaintId,
		EnterpriseId: f.enterpriseId,
		ReviewedId:   f.reviewedId,
		OccurredOn:   common.StringDate(f.occurredOn),
	})
}

func (f *FeedbackDone) UnmarshalJSON(data []byte) error {
	aux := struct {
		FeedbackId   uuid.UUID `json:"feedback_id"`
		ComplaintId  uuid.UUID `json:"complaint_id"`
		EnterpriseId uuid.UUID `json:"enterprise_id"`
		ReviewedId   uuid.UUID `json:"reviewed_id"`
		OccurredOn   string    `json:"occurred_on"`
	}{}
	err := json.Unmarshal(data, &aux)
	if err != nil {
		return err
	}
	f.reviewedId = aux.ReviewedId
	f.feedbackId = aux.FeedbackId
	f.complaintId = aux.ComplaintId
	f.enterpriseId = aux.EnterpriseId
	f.occurredOn, err = common.ParseDate(aux.OccurredOn)
	if err != nil {
		return err
	}
	return nil
}
