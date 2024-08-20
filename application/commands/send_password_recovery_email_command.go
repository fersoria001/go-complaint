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

type SendPasswordRecoveryEmailCommand struct {
	To                string `json:"to"`
	NewRandomPassword string `json:"newRandomPassword"`
}

func NewSendPasswordRecoveryEmail(to string, newRandomPassword string) *SendPasswordRecoveryEmailCommand {
	return &SendPasswordRecoveryEmailCommand{
		To:                to,
		NewRandomPassword: newRandomPassword,
	}
}

func (c SendPasswordRecoveryEmailCommand) Execute(ctx context.Context) error {
	tmpl, err := template.ParseFiles(filepath.Join(projectpath.EmailsPath, "password_recovery.html"))
	if err != nil {
		return fmt.Errorf("error parsing welcome template: %w", err)
	}
	var buff bytes.Buffer
	tmpl.Execute(&buff, map[string]string{"RandomPassword": c.NewRandomPassword})
	html := buff.String()
	infrastructure.Send(ctx, email.Email{
		Recipient: c.To,
		Subject:   "Go-Complaint password recovery",
		HtmlBody:  html,
	})
	return nil
}
