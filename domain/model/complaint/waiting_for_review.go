package complaint

import (
	"encoding/json"
	"go-complaint/domain/model/common"
	"time"

	"github.com/google/uuid"
)

type ComplaintSentForReview struct {
	complaintID uuid.UUID
	receiverID  string
	triggeredBy string
	authorID    string
	occurredOn  time.Time
}

func NewComplaintSentForReview(complaintID uuid.UUID, receiverID, authorID, triggeredBy string) *ComplaintSentForReview {
	return &ComplaintSentForReview{
		complaintID: complaintID,
		receiverID:  receiverID,
		triggeredBy: triggeredBy,
		authorID:    authorID,
		occurredOn:  time.Now(),
	}
}

func (wfr *ComplaintSentForReview) OccurredOn() time.Time {
	return wfr.occurredOn
}

func (wfr *ComplaintSentForReview) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		ComplaintID string `json:"complaint_id"`
		ReceiverID  string `json:"receiver_id"`
		TriggeredBy string `json:"triggered_by"`
		AuthorID    string `json:"author_id"`
		OccurredOn  string `json:"occurred_on"`
	}{
		ComplaintID: wfr.complaintID.String(),
		ReceiverID:  wfr.receiverID,
		TriggeredBy: wfr.triggeredBy,
		AuthorID:    wfr.authorID,
		OccurredOn:  common.StringDate(wfr.occurredOn),
	})
}

func (wfr *ComplaintSentForReview) UnmarshalJSON(data []byte) error {
	aux := &struct {
		ComplaintID string `json:"complaint_id"`
		ReceiverID  string `json:"receiver_id"`
		TriggeredBy string `json:"triggered_by"`
		AuthorID    string `json:"author_id"`
		OccurredOn  string `json:"occurred_on"`
	}{}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	occurredOn, err := common.ParseDate(aux.OccurredOn)
	if err != nil {
		return err
	}
	complaintID, err := uuid.Parse(aux.ComplaintID)
	if err != nil {
		return err
	}
	wfr.complaintID = complaintID
	wfr.receiverID = aux.ReceiverID
	wfr.triggeredBy = aux.TriggeredBy
	wfr.authorID = aux.AuthorID
	wfr.occurredOn = occurredOn
	return nil
}

func (wfr *ComplaintSentForReview) ComplaintID() uuid.UUID {
	return wfr.complaintID
}

func (wfr *ComplaintSentForReview) ReceiverID() string {
	return wfr.receiverID
}

func (wfr *ComplaintSentForReview) TriggeredBy() string {
	return wfr.triggeredBy
}

func (wfr *ComplaintSentForReview) AuthorID() string {
	return wfr.authorID
}
