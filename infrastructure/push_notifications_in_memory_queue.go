package infrastructure

import (
	"bytes"
	"context"
	"encoding/json"
	"go-complaint/infrastructure/queue"
	"net/http"
	"sync"
	"time"
)

var pushNotificationInMemoryQueueInstance *PushNotificationInMemoryQueue
var pushNotificationInMemoryQueueOnce sync.Once

func PushNotificationInMemoryQueueInstance() *PushNotificationInMemoryQueue {
	pushNotificationInMemoryQueueOnce.Do(func() {
		pushNotificationInMemoryQueueInstance = NewPushNotificationInmemoryQueue()
	})
	return pushNotificationInMemoryQueueInstance
}

type Operation struct {
	ID        string ` json:"operation_id" `
	Operation interface{}
}
type Result struct {
	StatusCode   int
	ResponseBody string
}

// email queue instance
type PushNotificationInMemoryQueue struct {
	queue   *queue.LinkedQueue[Operation]
	sentLog map[string]Result
	queued  int
}

func NewPushNotificationInmemoryQueue() *PushNotificationInMemoryQueue {
	return &PushNotificationInMemoryQueue{
		queue:   queue.NewLinkedQueue[Operation](),
		sentLog: make(map[string]Result),
		queued:  0,
	}
}
func (es *PushNotificationInMemoryQueue) Queued() int {
	return es.queued
}
func (es *PushNotificationInMemoryQueue) QueueNotification(id Operation) {
	es.queue.Enqueue(id)
	es.queued++
}

func (es *PushNotificationInMemoryQueue) SentLog() map[string]Result {
	return es.sentLog
}

func (es *PushNotificationInMemoryQueue) SendAll(ctx context.Context) {
	for i := 0; i <= es.queue.Length(); i++ {
		id, err := es.queue.Dequeue()
		if err != nil {
			break
		}
		sentOperationID, result := es.Send(ctx, id)
		es.sentLog[sentOperationID] = result
	}
}

func (es *PushNotificationInMemoryQueue) Send(ctx context.Context, id Operation) (string, Result) {
	j, err := json.Marshal(id)
	if err != nil {
		return id.ID, Result{
			StatusCode:   500,
			ResponseBody: err.Error(),
		}
	}
	b := bytes.NewBuffer(j)
	sendCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	InMemoryCacheInstance().Set(id.ID, id.Operation)
	request, err := http.NewRequestWithContext(
		sendCtx,
		http.MethodPost,
		"http://localhost:8080/publish",
		b,
	)
	if err != nil {
		return id.ID, Result{
			StatusCode:   500,
			ResponseBody: err.Error(),
		}
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer mlsn.0557f4217143328c73149ad91c7455121924f188c63af0fe093b42feb3fa1de1")
	body, err := http.DefaultClient.Do(request)
	if err != nil {
		return id.ID, Result{
			StatusCode:   500,
			ResponseBody: err.Error(),
		}
	}
	var responseBody []byte
	_, err = body.Body.Read(responseBody)
	if err != nil {
		return id.ID, Result{
			StatusCode:   body.StatusCode,
			ResponseBody: string(responseBody),
		}
	}
	return id.ID, Result{
		StatusCode:   body.StatusCode,
		ResponseBody: string(responseBody),
	}
}
