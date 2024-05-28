package models

import (
	"go-complaint/domain/model/feedback"

	"github.com/google/uuid"
)

type Answer struct {
	ID         uuid.UUID `json:"id"`
	FeedbackID uuid.UUID `json:"feedback_id"`
	SenderID   string    `json:"sender_id"`
	SenderIMG  string    `json:"sender_img"`
	SenderName string    `json:"sender_name"`
	Body       string    `json:"body"`
	CreatedAt  string    `json:"created_at"`
	Read       bool      `json:"read"`
	ReadAt     string    `json:"read_at"`
	UpdatedAt  string    `json:"updated_at"`
}

func NewAnswer(
	domainObject *feedback.Answer,
) *Answer {
	return &Answer{
		ID:         domainObject.ID(),
		FeedbackID: domainObject.FeedbackID(),
		SenderID:   domainObject.SenderID(),
		SenderIMG:  domainObject.SenderIMG(),
		SenderName: domainObject.SenderName(),
		Body:       domainObject.Body(),
		CreatedAt:  domainObject.CreatedAt().StringRepresentation(),
		Read:       domainObject.Read(),
		ReadAt:     domainObject.ReadAt().StringRepresentation(),
		UpdatedAt:  domainObject.UpdatedAt().StringRepresentation(),
	}
}

func (a *Answer) Columns() Columns {
	return Columns{
		"id",
		"feedback_id",
		"sender_id",
		"sender_img",
		"sender_name",
		"answer_body",
		"created_at",
		"read",
		"read_at",
		"updated_at",
	}
}

func (a *Answer) Values() Values {
	return Values{
		&a.ID,
		&a.FeedbackID,
		&a.SenderID,
		&a.SenderIMG,
		&a.SenderName,
		&a.Body,
		&a.CreatedAt,
		&a.Read,
		&a.ReadAt,
		&a.UpdatedAt,
	}
}

func (a *Answer) Table() string {
	return "answers"
}

func (a *Answer) Args() string {
	return "$1, $2, $3, $4, $5, $6, $7, $8, $9, $10"
}
