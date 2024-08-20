package commands

import (
	"bytes"
	"context"
	"fmt"
	"go-complaint/domain/model/email"
	"go-complaint/infrastructure"
	projectpath "go-complaint/project_path"
	"html/template"
	"os"
	"path/filepath"
)

type SendEmailVerifiedEmailCommand struct {
	To   string `json:"to"`
	Name string `json:"name"`
}

func NewSendEmailVerifiedEmailCommand(to, name string) *SendEmailVerifiedEmailCommand {
	return &SendEmailVerifiedEmailCommand{
		To:   to,
		Name: name,
	}
}

func (c SendEmailVerifiedEmailCommand) Execute(ctx context.Context) error {
	rootUrl := os.Getenv("FRONT_END_URL")
	signInLink := fmt.Sprintf("%s/%s", rootUrl, "sign-in")
	tmpl, err := template.ParseFiles(filepath.Join(projectpath.EmailsPath, "email_verified.html"))
	if err != nil {
		fmt.Println("error parsing welcome template", err)
	}
	var buff bytes.Buffer
	tmpl.Execute(&buff, map[string]string{"SignInLink": signInLink})
	html := buff.String()
	infrastructure.Send(ctx, email.Email{
		Recipient: c.To,
		Subject:   c.Name + "your email has been verified successfully at Go-Complaint",
		HtmlBody:  html,
	})
	return nil
}
