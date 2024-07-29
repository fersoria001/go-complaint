package complaint

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type ComplaintCreated struct {
	id                   uuid.UUID
	authorId             uuid.UUID
	authorName           string
	authorThumbnail      string
	authorIsEnterprise   bool
	receiverId           uuid.UUID
	receiverName         string
	receiverThumbnail    string
	receiverIsEnterprise bool
	status               Status
	occurredOn           time.Time
}

func NewComplaintCreated(
	id,
	authorId uuid.UUID,
	authorName,
	authorThumbnail string,
	authorIsEnterprise bool,
	receiverId uuid.UUID,
	receiverName,
	receiverThumbnail string,
	receiverIsEnterprise bool,
	status Status,
	occurredOn time.Time,
) *ComplaintCreated {
	return &ComplaintCreated{
		id:                   id,
		authorId:             authorId,
		authorName:           authorName,
		authorThumbnail:      authorThumbnail,
		authorIsEnterprise:   authorIsEnterprise,
		receiverId:           receiverId,
		receiverName:         receiverName,
		receiverThumbnail:    receiverThumbnail,
		receiverIsEnterprise: receiverIsEnterprise,
		status:               status,
		occurredOn:           occurredOn,
	}
}

func (cc ComplaintCreated) Id() uuid.UUID {
	return cc.id
}

func (cc ComplaintCreated) OccurredOn() time.Time {
	return cc.occurredOn
}

func (cc *ComplaintCreated) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Id                   uuid.UUID `json:"id"`
		AuthorId             uuid.UUID `json:"authorId"`
		AuthorName           string    `json:"authorName"`
		AuthorThumbnail      string    `json:"authorThumbnail"`
		AuthorIsEnterprise   bool      `json:"authorIsEnterprise"`
		ReceiverId           uuid.UUID `json:"receiverId"`
		ReceiverName         string    `json:"receiverName"`
		ReceiverIsEnterprise bool      `json:"receiverIsEnterprise"`
		Status               Status    `json:"status"`
		OccurredOn           time.Time `json:"occurredOn"`
	}{
		Id:                   cc.id,
		AuthorId:             cc.authorId,
		AuthorName:           cc.authorName,
		AuthorThumbnail:      cc.authorThumbnail,
		AuthorIsEnterprise:   cc.authorIsEnterprise,
		ReceiverId:           cc.receiverId,
		ReceiverName:         cc.receiverName,
		ReceiverIsEnterprise: cc.receiverIsEnterprise,
		Status:               cc.status,
		OccurredOn:           cc.occurredOn,
	})
}

func (cc *ComplaintCreated) UnmarshalJSON(data []byte) error {
	aux := struct {
		Id                   uuid.UUID `json:"id"`
		AuthorId             uuid.UUID `json:"authorId"`
		AuthorName           string    `json:"authorName"`
		AuthorThumbnail      string    `json:"authorThumbnail"`
		AuthorIsEnterprise   bool      `json:"authorIsEnterprise"`
		ReceiverId           uuid.UUID `json:"receiverId"`
		ReceiverName         string    `json:"receiverName"`
		ReceiverIsEnterprise bool      `json:"receiverIsEnterprise"`
		Status               Status    `json:"status"`
		OccurredOn           time.Time `json:"occurredOn"`
	}{}
	err := json.Unmarshal(data, &aux)
	if err != nil {
		return err
	}
	cc.id = aux.Id
	cc.authorId = aux.AuthorId
	cc.authorName = aux.AuthorName
	cc.authorThumbnail = aux.AuthorThumbnail
	cc.authorIsEnterprise = aux.AuthorIsEnterprise
	cc.receiverId = aux.ReceiverId
	cc.receiverName = aux.ReceiverName
	cc.receiverIsEnterprise = aux.ReceiverIsEnterprise
	cc.status = aux.Status
	cc.occurredOn = aux.OccurredOn
	return nil
}
