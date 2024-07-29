package complaint

import (
	"encoding/json"
	"go-complaint/domain/model/common"
	"time"

	"github.com/google/uuid"
)

type ComplaintSentForReview struct {
	complaintId uuid.UUID
	receiverId  uuid.UUID
	triggeredBy uuid.UUID
	authorId    uuid.UUID
	occurredOn  time.Time
}

func NewComplaintSentForReview(complaintId, receiverId, authorId, triggeredBy uuid.UUID) *ComplaintSentForReview {
	return &ComplaintSentForReview{
		complaintId: complaintId,
		receiverId:  receiverId,
		triggeredBy: triggeredBy,
		authorId:    authorId,
		occurredOn:  time.Now(),
	}
}

func (wfr *ComplaintSentForReview) OccurredOn() time.Time {
	return wfr.occurredOn
}

func (wfr *ComplaintSentForReview) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		ComplaintId uuid.UUID `json:"complaint_id"`
		ReceiverId  uuid.UUID `json:"receiver_id"`
		TriggeredBy uuid.UUID `json:"triggered_by"`
		AuthorId    uuid.UUID `json:"author_id"`
		OccurredOn  string    `json:"occurred_on"`
	}{
		ComplaintId: wfr.complaintId,
		ReceiverId:  wfr.receiverId,
		TriggeredBy: wfr.triggeredBy,
		AuthorId:    wfr.authorId,
		OccurredOn:  common.StringDate(wfr.occurredOn),
	})
}

func (wfr *ComplaintSentForReview) UnmarshalJSON(data []byte) error {
	aux := &struct {
		ComplaintId uuid.UUID `json:"complaint_id"`
		ReceiverId  uuid.UUID `json:"receiver_id"`
		TriggeredBy uuid.UUID `json:"triggered_by"`
		AuthorId    uuid.UUID `json:"author_id"`
		OccurredOn  string    `json:"occurred_on"`
	}{}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	occurredOn, err := common.ParseDate(aux.OccurredOn)
	if err != nil {
		return err
	}
	wfr.complaintId = aux.ComplaintId
	wfr.receiverId = aux.ReceiverId
	wfr.triggeredBy = aux.TriggeredBy
	wfr.authorId = aux.AuthorId
	wfr.occurredOn = occurredOn
	return nil
}

func (wfr *ComplaintSentForReview) ComplaintId() uuid.UUID {
	return wfr.complaintId
}

func (wfr *ComplaintSentForReview) ReceiverId() uuid.UUID {
	return wfr.receiverId
}

func (wfr *ComplaintSentForReview) TriggeredBy() uuid.UUID {
	return wfr.triggeredBy
}

func (wfr *ComplaintSentForReview) AuthorId() uuid.UUID {
	return wfr.authorId
}
