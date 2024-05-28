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
	complaintID uuid.UUID
	authorID    string
	receiverID  string
}

func NewComplaintSent(complaintID uuid.UUID, authorID, receiverID string, occurredOn time.Time) *ComplaintSent {
	var c *ComplaintSent = &ComplaintSent{}
	c.authorID = authorID
	c.receiverID = receiverID
	c.complaintID = complaintID
	c.occurredOn = occurredOn
	return c
}


func (c *ComplaintSent) OccurredOn() time.Time {
	return c.occurredOn
}

func (c *ComplaintSent) MarshalJSON() ([]byte, error) {
	j, err := json.Marshal(struct {
		OccurredOn  string
		ComplaintID string
		AuthorID    string
		ReceiverID  string
	}{
		OccurredOn:  common.StringDate(c.occurredOn),
		ComplaintID: c.complaintID.String(),
		AuthorID:    c.authorID,
		ReceiverID:  c.receiverID,
	})
	if err != nil {
		return nil, err
	}
	return j, nil
}

func (c *ComplaintSent) UnmarshalJSON(data []byte) error {
	aux := struct {
		OccurredOn  string
		ComplaintID string
		AuthorID    string
		ReceiverID  string

	}{}
	err := json.Unmarshal(data, &aux)
	if err != nil {
		return err
	}
	c.occurredOn, err = common.ParseDate(aux.OccurredOn)
	if err != nil {
		return err
	}
	c.complaintID, err = uuid.Parse(aux.ComplaintID)
	if err != nil {
		return err
	}
	c.authorID = aux.AuthorID
	c.receiverID = aux.ReceiverID

	return nil
}
