package feedback

import (
	"go-complaint/domain/model/common"
	"go-complaint/erros"
	"time"

	"github.com/google/uuid"
)

// Package feedback
// <<Entity>> Answer
// Answer is a struct that represents the answer of a feedback.
// a feedback can have multiple answers from different users that
// are related to the complaint trough the feedback, thus it's part of the feedback aggregate.
type Answer struct {
	id         uuid.UUID
	feedbackID uuid.UUID
	senderID   string
	senderIMG  string
	senderName string
	body       string
	createdAt  common.Date
	read       bool
	readAt     common.Date
	updatedAt  common.Date
}

func (a *Answer) MarkAsRead() error {
	if a.read {
		return &erros.InvalidStateError{}
	}
	a.read = true
	a.readAt = common.NewDate(time.Now())
	return nil
}

func NewAnswer(
	id uuid.UUID,
	feedbackID uuid.UUID,
	senderID string,
	senderIMG string,
	senderName string,
	body string,
	createdAt common.Date,
	read bool, readAt common.Date,
	updatedAt common.Date,
) (*Answer, error) {
	var answer *Answer = new(Answer)
	err := answer.setID(id)
	if err != nil {
		return nil, err
	}
	err = answer.setFeedbackID(feedbackID)
	if err != nil {
		return nil, err
	}
	err = answer.setSenderID(senderID)
	if err != nil {
		return nil, err
	}
	err = answer.setSenderIMG(senderIMG)
	if err != nil {
		return nil, err
	}
	err = answer.setSenderName(senderName)
	if err != nil {
		return nil, err
	}
	err = answer.setBody(body)
	if err != nil {
		return nil, err
	}
	err = answer.setCreatedAt(createdAt)
	if err != nil {
		return nil, err
	}
	answer.setRead(read)
	err = answer.setReadAt(readAt)
	if err != nil {
		return nil, err
	}
	err = answer.setUpdatedAt(updatedAt)
	if err != nil {
		return nil, err
	}
	return answer, nil
}

func (a *Answer) setID(id uuid.UUID) error {
	if id == uuid.Nil {
		return &erros.NullValueError{}
	}
	a.id = id
	return nil
}

func (a *Answer) setFeedbackID(feedbackID uuid.UUID) error {
	if feedbackID == uuid.Nil {
		return &erros.NullValueError{}
	}
	a.feedbackID = feedbackID
	return nil
}

func (a *Answer) setSenderID(senderID string) error {
	if senderID == "" {
		return &erros.NullValueError{}
	}
	a.senderID = senderID
	return nil
}

func (a *Answer) setSenderIMG(senderIMG string) error {
	if senderIMG == "" {
		return &erros.NullValueError{}
	}
	a.senderIMG = senderIMG
	return nil
}

func (a *Answer) setSenderName(senderName string) error {
	if senderName == "" {
		return &erros.NullValueError{}
	}
	a.senderName = senderName
	return nil
}

func (a *Answer) setBody(body string) error {
	if body == "" {
		return &erros.NullValueError{}
	}
	a.body = body
	return nil
}

func (a *Answer) setCreatedAt(createdAt common.Date) error {
	if createdAt == (common.Date{}) {
		return &erros.NullValueError{}
	}
	a.createdAt = createdAt
	return nil

}

func (a *Answer) setReadAt(readAt common.Date) error {
	if readAt == (common.Date{}) {
		return &erros.NullValueError{}
	}
	a.readAt = readAt
	return nil
}

func (a *Answer) setUpdatedAt(updatedAt common.Date) error {
	if updatedAt == (common.Date{}) {
		return &erros.NullValueError{}
	}
	a.updatedAt = updatedAt
	return nil
}

func (a *Answer) setRead(read bool) {
	a.read = read
}

func (a *Answer) ID() uuid.UUID {
	return a.id
}

func (a *Answer) FeedbackID() uuid.UUID {
	return a.feedbackID
}

func (a *Answer) SenderID() string {
	return a.senderID
}

func (a *Answer) SenderIMG() string {
	return a.senderIMG
}

func (a *Answer) SenderName() string {
	return a.senderName
}

func (a *Answer) Body() string {
	return a.body
}

func (a *Answer) CreatedAt() common.Date {
	return a.createdAt
}

func (a *Answer) Read() bool {
	return a.read
}

func (a *Answer) ReadAt() common.Date {
	return a.readAt
}

func (a *Answer) UpdatedAt() common.Date {
	return a.updatedAt
}
