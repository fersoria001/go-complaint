package commands

import (
	"context"
	"fmt"
	"go-complaint/domain/model/email"
	"go-complaint/infrastructure"
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
		return nil
	}
	service := infrastructure.EmailServiceInstance()
	confirmationLink := fmt.Sprintf("\"https://www.go-complaint.com/confirmation-link?token=%s\"", c.ConfirmationToken)
	html := `
	<div>
	<div> Welcome to Go Complaint! </div>
	<p>To confirm your email you have to click in the following </p>`
	html += "<a href=" + confirmationLink + ">link</a>"
	html += `<p> If you encounter any issues with the confirmation link contact to owner@go-complaint.com </p>`
	html += `</div>`
	service.Send(ctx, email.Email{
		Recipient: c.ToEmail,
		Subject:   "Welcome " + c.ToName + " to Go-Complaint",
		HtmlBody:  html,
	})
	return nil
}
func (c SendEmailCommand) VerifySignIn(ctx context.Context) error {
	if c.ToEmail == "" || c.ToName == "" || c.ConfirmationCode == 0 {
		return nil
	}
	service := infrastructure.EmailServiceInstance()
	html := `<div>`
	html += `You have attempt to sign-in into Go Complaint
	to complete your login process enter this confirmation code`
	html += fmt.Sprintf("<div>%s</div>", c.ConfirmationToken)
	html += `</div>`
	service.Send(ctx, email.Email{
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
	service := infrastructure.EmailServiceInstance()
	html := `<div>`
	html += `Welcome!`
	html += `<div> You have been verified your email in Go Complain, now you are able to`
	html += fmt.Sprintf("<a href=\"%s\">sign-in</a>", "https://www.go-complaint.com/sign-in")
	html += `</div>`
	html += `</div>`
	service.Send(ctx, email.Email{
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
	service := infrastructure.EmailServiceInstance()
	html := "<div>"
	html += `We detected a new log in into your account if it wasn't you contact us to
	owner@go-complaint.com`
	html += `</div>`
	service.Send(ctx, email.Email{
		Recipient: c.ToEmail,
		Subject:   "New log-in to Go-Complaint",
		HtmlBody:  html,
	})
	return nil
}

func (c SendEmailCommand) PasswordRecovery(ctx context.Context) error {
	if c.ToEmail == "" || c.ToName == "" || c.RandomPassword == "" {
		return nil
	}
	service := infrastructure.EmailServiceInstance()
	html := `<div>`
	html += `You have request to reset your password`
	html += fmt.Sprintf("your new randomly generated password is: %s", c.RandomPassword)
	html += `<div>Use it to log into your account</div>`
	html += `<div> If you did not request this action on your account 
	or you find trouble login in contact us to 
	owner@go-complaint.com</div>`
	html += `</div>`
	service.Send(ctx, email.Email{
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
	service := infrastructure.EmailServiceInstance()
	html := `<div>`
	html += `You has changed your password from your account configuration.`
	html += `<div> If you didn't trigger this action on your account, contact us 
	to owner@go-complaint.com</div>`
	html += `</div>`
	service.Send(ctx, email.Email{
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
	service := infrastructure.EmailServiceInstance()
	html := `<div>`
	html += fmt.Sprintf(`Hi %s you have been received a hiring invitation at Go Complaint,
	you can review your invitations <a href=\"%s\">here</a>`, c.ToName,
		"https://www.go-complaint.com/hiring-invitations")
	html += `</div>`
	service.Send(ctx, email.Email{
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
	service := infrastructure.EmailServiceInstance()
	html := `<div>`
	html += fmt.Sprintf(`You have been received a contact message from %s at Go Complaint `, c.ToEmail)
	html += fmt.Sprintf(`<div>%s</div>`, c.Text)
	html += `</div>`
	service.Send(ctx, email.Email{
		Recipient: "bercho001@gmail.com",
		Subject:   "Contact email to Go-Complaint",
		HtmlBody:  html,
	})
	return nil
}
