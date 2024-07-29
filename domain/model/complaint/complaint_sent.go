package complaint

import (
	"encoding/json"
	"go-complaint/domain/model/common"
	"time"

	"github.com/google/uuid"
)

// Package domain
// Event implement the DomainEvent interface
type ComplaintSent struct {
	occurredOn  time.Time
	complaintId uuid.UUID
	authorId    uuid.UUID
	receiverId  uuid.UUID
}

func NewComplaintSent(complaintId, authorId, receiverId uuid.UUID, occurredOn time.Time) *ComplaintSent {
	c := new(ComplaintSent)
	c.authorId = authorId
	c.receiverId = receiverId
	c.complaintId = complaintId
	c.occurredOn = occurredOn
	return c
}

func (c *ComplaintSent) OccurredOn() time.Time {
	return c.occurredOn
}

func (c *ComplaintSent) MarshalJSON() ([]byte, error) {
	j, err := json.Marshal(struct {
		ComplaintId uuid.UUID
		AuthorId    uuid.UUID
		ReceiverId  uuid.UUID
		OccurredOn  string
	}{
		OccurredOn:  common.StringDate(c.occurredOn),
		ComplaintId: c.complaintId,
		AuthorId:    c.authorId,
		ReceiverId:  c.receiverId,
	})
	if err != nil {
		return nil, err
	}
	return j, nil
}

func (c *ComplaintSent) UnmarshalJSON(data []byte) error {
	aux := struct {
		ComplaintId uuid.UUID
		AuthorId    uuid.UUID
		ReceiverId  uuid.UUID
		OccurredOn  string
	}{}
	err := json.Unmarshal(data, &aux)
	if err != nil {
		return err
	}
	c.occurredOn, err = common.ParseDate(aux.OccurredOn)
	if err != nil {
		return err
	}
	c.complaintId = aux.ComplaintId
	c.authorId = aux.AuthorId
	c.receiverId = aux.ReceiverId
	return nil
}
