package infrastructure

import (
	"bytes"
	"context"
	"encoding/json"
	"go-complaint/domain/model/email"
	"log"
	"net/http"
	"sync"
	"time"
)

var emailServiceInstance *EmailService
var emailQueueOnce sync.Once

func EmailServiceInstance() *EmailService {
	emailQueueOnce.Do(func() {
		emailServiceInstance = NewEmailService()
	})
	return emailServiceInstance
}

// email queue instance
type EmailService struct {
	emailQueueInstance chan *email.Email
	sentLog            map[string]interface{}
	queued             int
}

func NewEmailService() *EmailService {
	return &EmailService{
		emailQueueInstance: make(chan *email.Email),
		sentLog:            make(map[string]interface{}),
		queued:             0,
	}
}
func (es *EmailService) Queued() int {
	return es.queued
}
func (es *EmailService) QueueEmail(email *email.Email) {
	es.emailQueueInstance <- email
	es.queued++
}

func (es *EmailService) SentLog() map[string]interface{} {
	return es.sentLog
}

func (es *EmailService) SendAll(ctx context.Context) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	for {
		<-ticker.C
		select {
		case <-ctx.Done():
			return
		case email := <-es.emailQueueInstance:
			msgID, err := es.Send(ctx, *email)
			log.Println("Email sent with message ID: ", msgID, "Error: ", err)
			es.queued--
			es.sentLog[msgID] = struct {
				Error      error
				OccurredOn time.Time
			}{
				Error:      err,
				OccurredOn: time.Now(),
			}
		}
	}

}

func (es *EmailService) Send(ctx context.Context, email email.Email) (string, error) {
	j, err := json.Marshal(email)
	if err != nil {
		return "", err
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
		return "", err
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer mlsn.0557f4217143328c73149ad91c7455121924f188c63af0fe093b42feb3fa1de1")
	log.Println("Headers set")
	body, err := http.DefaultClient.Do(request)
	if err != nil {
		return "", err
	}
	log.Printf("Email sent with status code: %d", body.StatusCode)
	msgID := body.Header.Get("X-Message-Id")
	var responseBody []byte
	_, err = body.Body.Read(responseBody)
	if err != nil {
		return "", err
	}
	if string(responseBody) != "" {
		return msgID, &EmailError{
			StatusCode:   body.StatusCode,
			ResponseBody: string(responseBody),
		}
	}
	return msgID, nil
}
