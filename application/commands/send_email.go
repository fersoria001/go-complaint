package commands

import (
	"context"
	"go-complaint/domain/model/email"
	"go-complaint/infrastructure"
)

type SendEmailCommand struct {
	ToEmail           string
	ToName            string
	ConfirmationToken string
	ConfirmationCode  int
	RandomPassword    string
}

func (c SendEmailCommand) Welcome(ctx context.Context) error {
	if c.ToEmail == "" || c.ToName == "" || c.ConfirmationToken == "" {
		return nil
	}
	service := infrastructure.EmailServiceInstance()
	service.QueueEmail(&email.Email{
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
			Email: c.ToEmail,
			Name:  c.ToName,
		},
		Subject: "Welcome " + c.ToName + " to Go-Complaint",
		Personalization: []map[string]interface{}{
			{
				"email": "go-complaint.com",
				"data": map[string]interface{}{
					"name":              "Go-Complaint",
					"confirmationToken": c.ConfirmationToken,
					"account": map[string]interface{}{
						"name": "Go-Complaint",
					},
				},
			},
		},
	})
	return nil
}
func (c SendEmailCommand) VerifySignIn(ctx context.Context) error {
	if c.ToEmail == "" || c.ToName == "" || c.ConfirmationCode == 0 {
		return nil
	}
	service := infrastructure.EmailServiceInstance()
	service.QueueEmail(&email.Email{
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
			Email: c.ToEmail,
			Name:  c.ToName,
		},
		Subject: "Verify your sign-in action to Go-Complaint",
		Personalization: []map[string]interface{}{
			{
				"email": "go-complaint.com",
				"data": map[string]interface{}{
					"name":             "Go-Complaint",
					"confirmationCode": c.ConfirmationCode,
					"account": map[string]interface{}{
						"name": "Go-Complaint",
					},
				},
			},
		},
	})
	return nil
}
func (c SendEmailCommand) EmailVerified(ctx context.Context) error {
	if c.ToEmail == "" || c.ToName == "" {
		return nil
	}
	service := infrastructure.EmailServiceInstance()
	service.QueueEmail(&email.Email{
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
			Email: c.ToEmail,
			Name:  c.ToName,
		},
		Subject: "Welcome " + c.ToName + " to Go-Complaint",
		Personalization: []map[string]interface{}{
			{
				"email": "go-complaint.com",
				"data": map[string]interface{}{
					"name": "Go-Complaint",
					"account": map[string]interface{}{
						"name": "Go-Complaint",
					},
				},
			},
		},
	})
	return nil
}

func (c SendEmailCommand) SignIn(ctx context.Context) error {
	if c.ToEmail == "" || c.ToName == "" {
		return nil
	}
	service := infrastructure.EmailServiceInstance()
	service.QueueEmail(&email.Email{
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
			Email: c.ToEmail,
			Name:  c.ToName,
		},
		Subject: "Welcome " + c.ToName + " to Go-Complaint",
		Personalization: []map[string]interface{}{
			{
				"email": "go-complaint.com",
				"data": map[string]interface{}{
					"name": "Go-Complaint",
					"account": map[string]interface{}{
						"name": "Go-Complaint",
					},
				},
			},
		},
	})
	return nil
}

func (c SendEmailCommand) PasswordRecovery(ctx context.Context) error {
	if c.ToEmail == "" || c.ToName == "" || c.RandomPassword == "" {
		return nil
	}
	service := infrastructure.EmailServiceInstance()
	service.QueueEmail(&email.Email{
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
			Email: c.ToEmail,
			Name:  c.ToName,
		},
		Subject: "Welcome " + c.ToName + " to Go-Complaint",
		Personalization: []map[string]interface{}{
			{
				"email": "go-complaint.com",
				"data": map[string]interface{}{
					"name":           "Go-Complaint",
					"randomPassword": c.RandomPassword,
					"account": map[string]interface{}{
						"name": "Go-Complaint",
					},
				},
			},
		},
	})
	return nil
}

func (c SendEmailCommand) PasswordChanged(ctx context.Context) error {
	if c.ToEmail == "" || c.ToName == "" {
		return nil
	}
	service := infrastructure.EmailServiceInstance()
	service.QueueEmail(&email.Email{
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
			Email: c.ToEmail,
			Name:  c.ToName,
		},
		Subject: "Welcome " + c.ToName + " to Go-Complaint",
		Personalization: []map[string]interface{}{
			{
				"email": "go-complaint.com",
				"data": map[string]interface{}{
					"name": "Go-Complaint",
					"account": map[string]interface{}{
						"name": "Go-Complaint",
					},
				},
			},
		},
	})
	return nil
}

func (c SendEmailCommand) HiringInvitationSent(ctx context.Context) error {
	if c.ToEmail == "" || c.ToName == "" {
		return nil
	}
	service := infrastructure.EmailServiceInstance()
	service.QueueEmail(&email.Email{
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
			Email: c.ToEmail,
			Name:  c.ToName,
		},
		Subject: "Welcome " + c.ToName + " to Go-Complaint",
		Personalization: []map[string]interface{}{
			{
				"email": "go-complaint.com",
				"data": map[string]interface{}{
					"name": "Go-Complaint",
					"account": map[string]interface{}{
						"name": "Go-Complaint",
					},
				},
			},
		},
	})
	return nil
}
