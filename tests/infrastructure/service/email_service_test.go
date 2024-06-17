package infrastructure_test

import (
	"context"
	"go-complaint/domain/model/email"
	"go-complaint/infrastructure"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmailServiceSend(t *testing.T) {
	// Arrange
	ctx := context.Background()
	service := infrastructure.EmailServiceInstance()
	email := &email.Email{
		TemplateID: "vywj2lpd9kpl7oqz",
		From: struct {
			Email string `json:"email"`
			Name  string `json:"name"`
		}{
			Email: "go-complaint.com",
			Name:  "Go-Complaint",
		},
		To: struct {
			Email string `json:"email"`
			Name  string `json:"name"`
		}{
			Email: "bercho001@gmail.com",
			Name:  "Fernando Agustin Soria",
		},
		Subject: "Welcome " + "Fernando Agustin Soria" + " to Go-Complaint",
		Personalization: []map[string]interface{}{
			{
				"email": "go-complaint.com",
				"data": map[string]interface{}{
					"name":              "Go-Complaint",
					"confirmationToken": "token",
					"account": map[string]interface{}{
						"name": "Go-Complaint",
					},
				},
			},
		},
	}
	// Act
	service.QueueEmail(email)
	service.SendAll(ctx)
	sentLog := service.SentLog()
	// Assert
	assert.Equal(t, 1, len(sentLog))
}
