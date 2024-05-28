package models

import (
	"go-complaint/domain/model/complaint"

	"github.com/google/uuid"
)

type Complaint struct {
	ID          uuid.UUID
	AuthorID    string
	ReceiverID  string
	Status      string
	Title       string
	Description string
	Content     string
	Rate        int
	Comment     string
	CreatedAt   string
	UpdatedAt   string
}

func (c *Complaint) Columns() Columns {
	return Columns{
		"id",
		"author_id",
		"receiver_id",
		"complaint_status",
		"title",
		"descriptionn",
		"body",
		"rating_rate",
		"rating_comment",
		"created_at",
		"updated_at",
	}
}

func (c *Complaint) Values() Values {
	return Values{
		&c.ID,
		&c.AuthorID,
		&c.ReceiverID,
		&c.Status,
		&c.Title,
		&c.Description,
		&c.Content,
		&c.Rate,
		&c.Comment,
		&c.CreatedAt,
		&c.UpdatedAt,
	}
}

func (c *Complaint) Args() string {
	return "$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11"
}

func (c *Complaint) Table() string {
	return "complaints"
}

func NewComplaint(domain *complaint.Complaint) Complaint {
	return Complaint{
		ID:          domain.ID(),
		AuthorID:    domain.AuthorID(),
		ReceiverID:  domain.ReceiverID(),
		Status:      domain.Status().String(),
		Title:       domain.Message().Title(),
		Description: domain.Message().Description(),
		Content:     domain.Message().Body(),
		Rate:        domain.Rating().Rate(),
		Comment:     domain.Rating().Comment(),
		CreatedAt:   domain.CreatedAt().StringRepresentation(),
		UpdatedAt:   domain.UpdatedAt().StringRepresentation(),
	}
}
