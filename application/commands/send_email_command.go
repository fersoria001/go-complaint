package commands

import (
	"bytes"
	"context"
	"fmt"
	"go-complaint/domain/model/email"
	"go-complaint/infrastructure"
	projectpath "go-complaint/project_path"
	"html/template"
	"log"
	"path/filepath"
	"strconv"
)

type SendEmailCommand struct {
	ToEmail           string
	ToName            string
	Text              string
	ConfirmationToken string
	ConfirmationCode  int
	RandomPassword    string
}

func (c SendEmailCommand) VerifySignIn(ctx context.Context) error {
	if c.ToEmail == "" || c.ToName == "" || c.ConfirmationCode == 0 {
		log.Println("error at verify sign in email, bad request")
		return nil
	}
	confirmationCode := strconv.Itoa(c.ConfirmationCode)
	tmpl, err := template.ParseFiles(filepath.Join(projectpath.EmailsPath, "verify_sign_in.html"))
	if err != nil {
		fmt.Println("error verify sign in template", err)
	}
	var buff bytes.Buffer
	tmpl.Execute(&buff, map[string]string{"ConfirmationCode": confirmationCode})
	html := buff.String()
	infrastructure.Send(ctx, email.Email{
		Recipient: c.ToEmail,
		Subject:   "Verify your sign-in action in Go-Complaint",
		HtmlBody:  html,
	})
	return nil
}

func (c SendEmailCommand) SignIn(ctx context.Context) error {
	if c.ToEmail == "" || c.ToName == "" {
		return nil
	}
	// service := infrastructure.EmailServiceInstance()
	html := "<div>"
	html += `We detected a new log in into your account if it wasn't you contact us to
	owner@go-complaint.com`
	html += `</div>`
	infrastructure.Send(ctx, email.Email{
		Recipient: c.ToEmail,
		Subject:   "New log-in to Go-Complaint",
		HtmlBody:  html,
	})
	return nil
}

func (c SendEmailCommand) PasswordChanged(ctx context.Context) error {
	if c.ToEmail == "" || c.ToName == "" {
		return nil
	}
	tmpl, err := template.ParseFiles(filepath.Join(projectpath.EmailsPath, "password_changed.html"))
	if err != nil {
		log.Println("error parsing welcome template", err)
	}
	var buff bytes.Buffer
	tmpl.Execute(&buff, map[string]string{})
	html := buff.String()
	infrastructure.Send(ctx, email.Email{
		Recipient: c.ToEmail,
		Subject:   "Your account password has changed.",
		HtmlBody:  html,
	})
	return nil
}

func (c SendEmailCommand) HiringInvitationSent(ctx context.Context) error {
	if c.ToEmail == "" || c.ToName == "" {
		return nil
	}
	hiringInvitationsLink := "https://www.go-complaint.com"
	receiverName := c.ToName
	tmpl, err := template.ParseFiles(filepath.Join(projectpath.EmailsPath, "hiring_invitation_sent.html"))
	if err != nil {
		log.Println("error parsing welcome template", err)
	}
	var buff bytes.Buffer
	tmpl.Execute(&buff, map[string]string{
		"ReceiverName":          receiverName,
		"HiringInvitationsLink": hiringInvitationsLink})
	html := buff.String()
	infrastructure.Send(ctx, email.Email{
		Recipient: c.ToEmail,
		Subject:   "You have been received a hiring invitation at Go-Complaint!",
		HtmlBody:  html,
	})
	return nil
}
