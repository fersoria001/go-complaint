package commands

import (
	"bytes"
	"context"
	"fmt"
	"go-complaint/domain/model/email"
	"go-complaint/infrastructure"
	projectpath "go-complaint/project_path"
	"html/template"
	"path/filepath"
)

type SendEmailValidationCommand struct {
	To              string `json:"to"`
	Name            string `json:"name"`
	ValidationToken string `json:"validationToken"`
}

func NewSendEmailValidationCommand(to, name, validationToken string) *SendEmailValidationCommand {
	return &SendEmailValidationCommand{
		To:              to,
		Name:            name,
		ValidationToken: validationToken,
	}
}

func (c SendEmailValidationCommand) Execute(ctx context.Context) error {
	confirmationLink := fmt.Sprintf("https://api.go-complaint.com/validation-link?token=%s", c.ValidationToken)
	tmpl, err := template.ParseFiles(filepath.Join(projectpath.EmailsPath, "welcome.html"))
	if err != nil {
		return fmt.Errorf("error parsing welcome template %w", err)
	}
	var buff bytes.Buffer
	tmpl.Execute(&buff, map[string]string{"ConfirmationLink": confirmationLink})
	html := buff.String()
	infrastructure.Send(ctx, email.Email{
		Recipient: c.To,
		Subject:   c.Name + "Welcome " + " to Go-Complaint",
		HtmlBody:  html,
	})
	return nil
}
