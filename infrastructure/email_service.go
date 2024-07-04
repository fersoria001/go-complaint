package infrastructure

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"go-complaint/domain/model/email"
	"log"
	"net/http"
	"sync"
	"time"
)

type To struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}
type From struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}
type Email struct {
	From            *From                    `json:"from"`
	To              []*To                    `json:"to"`
	Subject         string                   `json:"subject"`
	Html            string                   `json:"html"`
	Personalization []map[string]interface{} `json:"personalization"`
}

var emailServiceInstance *EmailService
var emailQueueOnce sync.Once

func EmailServiceInstance() *EmailService {
	emailQueueOnce.Do(func() {
		emailServiceInstance = NewEmailService()
	})
	return emailServiceInstance
}

type EmailService struct {
}

func NewEmailService() *EmailService {
	return &EmailService{}
}

func (es *EmailService) Send(ctx context.Context, obj email.Email) {
	sender := "owner@go-complaint.com"
	to := make([]*To, 0)
	to = append(to, &To{
		Email: obj.Recipient,
	})
	input := &Email{
		From: &From{
			Email: sender,
			Name:  "Go Complaint",
		},
		To:      to,
		Subject: obj.Subject,
		Html:    obj.HtmlBody,
	}
	j, err := json.Marshal(input)
	if err != nil {
		fmt.Sprintln("error marshaling", err)
		return
	}
	b := bytes.NewBuffer(j)
	sendCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	request, err := http.NewRequestWithContext(
		sendCtx,
		"POST",
		"https://api.mailersend.com/v1/email",
		b,
	)
	if err != nil {
		fmt.Sprintln("error at create request", err)
		return
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer mlsn.0557f4217143328c73149ad91c7455121924f188c63af0fe093b42feb3fa1de1")
	body, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Sprintln("error at send request", err)
		return
	}
	log.Println("Status", body.Status)
	msgID := body.Header.Get("X-Message-Id")
	paused := body.Header.Get("x-send-paused")
	var responseBody map[string]interface{}
	err = json.NewDecoder(body.Body).Decode(&responseBody)
	if err != nil {
		log.Println("Error decoding res", err)
	}
	log.Println("response body", responseBody)
	fmt.Println("Email Sent to address: " + obj.Recipient)
	fmt.Println("msgID", msgID)
	fmt.Println("paused?", paused)
}

// func (es *EmailService) Send(ctx context.Context, email email.Email) {
// 	cfg, err := config.LoadDefaultConfig(ctx)
// 	if err != nil {
// 		log.Fatal("error", err)
// 	}
// 	charSet := "UTF-8"
// 	sender := "owner@go-complaint.com"
// 	svc := sesv2.NewFromConfig(cfg)
// 	input := &sesv2.SendEmailInput{
// 		Content: &types.EmailContent{
// 			Simple: &types.Message{
// 				Body: &types.Body{
// 					Html: &types.Content{
// 						Data:    &email.HtmlBody,
// 						Charset: &charSet,
// 					},
// 				},
// 				Subject: &types.Content{
// 					Data:    &sender,
// 					Charset: &charSet,
// 				},
// 			},
// 		},
// 		Destination: &types.Destination{
// 			ToAddresses: []string{email.Recipient},
// 		},
// 		FeedbackForwardingEmailAddress: &sender,
// 		FromEmailAddress:               &sender,
// 		ReplyToAddresses:               []string{sender},
// 	}
// 	result, err := svc.SendEmail(ctx, input)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	} else {
// 		fmt.Println("Email Sent to address: " + email.Recipient)
// 		fmt.Println(result.ResultMetadata)
// 	}
// }
