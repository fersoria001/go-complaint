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

func (c SendEmailCommand) Welcome(ctx context.Context) error {
	if c.ToEmail == "" || c.ToName == "" || c.ConfirmationToken == "" {
		log.Println("sending email error bad request", c.ToEmail, c.ToName)
		return nil
	}
	confirmationLink := fmt.Sprintf("https://api.go-complaint.com/confirmation-link?token=%s", c.ConfirmationToken)
	tmpl, err := template.ParseFiles(filepath.Join(projectpath.EmailsPath, "welcome.html"))
	if err != nil {
		fmt.Println("error parsing welcome template", err)
	}
	var buff bytes.Buffer
	tmpl.Execute(&buff, map[string]string{"ConfirmationLink": confirmationLink})
	html := buff.String()
	infrastructure.Send(ctx, email.Email{
		Recipient: c.ToEmail,
		Subject:   "Welcome " + c.ToName + " to Go-Complaint",
		HtmlBody:  html,
	})
	return nil
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
func (c SendEmailCommand) EmailVerified(ctx context.Context) error {
	if c.ToEmail == "" || c.ToName == "" {
		return nil
	}
	signInLink := "https://www.go-complaint.com/"
	tmpl, err := template.ParseFiles(filepath.Join(projectpath.EmailsPath, "email_verified.html"))
	if err != nil {
		fmt.Println("error parsing welcome template", err)
	}
	var buff bytes.Buffer
	tmpl.Execute(&buff, map[string]string{"SignInLink": signInLink})
	html := buff.String()
	infrastructure.Send(ctx, email.Email{
		Recipient: c.ToEmail,
		Subject:   c.ToName + "your email has been verified successfully at Go-Complaint",
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

func (c SendEmailCommand) PasswordRecovery(ctx context.Context) error {
	if c.ToEmail == "" || c.ToName == "" || c.RandomPassword == "" {
		log.Println("error at passwordRecoveryEmail bad request")
		return nil
	}
	randomPassword := c.RandomPassword
	tmpl, err := template.ParseFiles(filepath.Join(projectpath.EmailsPath, "password_recovery.html"))
	if err != nil {
		log.Println("error parsing welcome template", err)
	}
	var buff bytes.Buffer
	tmpl.Execute(&buff, map[string]string{"RandomPassword": randomPassword})
	html := buff.String()
	infrastructure.Send(ctx, email.Email{
		Recipient: c.ToEmail,
		Subject:   "Go-Complaint password recovery",
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

func (c SendEmailCommand) Contact(ctx context.Context) error {
	if c.ToEmail == "" {
		return nil
	}
	sender := c.ToEmail
	message := c.Text
	tmpl, err := template.ParseFiles(filepath.Join(projectpath.EmailsPath, "contact_email.html"))
	if err != nil {
		log.Println("error parsing welcome template", err)
	}
	var buff bytes.Buffer
	tmpl.Execute(&buff, map[string]string{
		"Sender":  sender,
		"Message": message})
	html := buff.String()
	infrastructure.Send(ctx, email.Email{
		Recipient: "bercho001@gmail.com",
		Subject:   "Contact email to Go-Complaint",
		HtmlBody:  html,
	})
	return nil
}
