package commands

import (
	"bytes"
	"context"
	"fmt"
	"go-complaint/domain/model/email"
	"go-complaint/infrastructure"
	projectpath "go-complaint/project_path"
	"html/template"
	"net/mail"
	"path/filepath"
)

type SendContactEmailCommand struct {
	From    string `json:"from"`
	Message string `json:"message"`
}

func (c SendContactEmailCommand) Execute(ctx context.Context) error {
	if c.From == "" {
		return fmt.Errorf("the from email is not present at send contact email")
	}
	if _, err := mail.ParseAddress(c.From); err != nil {
		return err
	}
	tmpl, err := template.ParseFiles(filepath.Join(projectpath.EmailsPath, "contact_email.html"))
	if err != nil {
		return err
	}
	var buff bytes.Buffer
	tmpl.Execute(&buff, map[string]string{
		"Sender":  c.From,
		"Message": c.Message})
	html := buff.String()
	infrastructure.Send(ctx, email.Email{
		Recipient: "owner@go-complaint.com",
		Subject:   "Contact email to Go-Complaint",
		HtmlBody:  html,
	})
	return nil
}
