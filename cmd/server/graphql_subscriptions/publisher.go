package graphql_subscriptions

import (
	"context"
	"go-complaint/application/application_services"
	"log"
	"sync"
	"time"

	"nhooyr.io/websocket/wsjson"
)

var publisherInstance *SubscriptionsPublisher
var publisherOnce sync.Once

func SubscriptionsPublisherInstance() *SubscriptionsPublisher {
	publisherOnce.Do(func() {
		publisherInstance = &SubscriptionsPublisher{
			subscribers: sync.Map{},
		}
	})
	return publisherInstance
}

type SubscriptionsPublisher struct {
	subscribers sync.Map
	publishing  bool
}

func (p *SubscriptionsPublisher) Publish(
	ctx context.Context, operationID string) error {
	if p.publishing {
		return nil
	}

	p.publishing = true
	ticker := time.NewTicker(time.Millisecond * 100)
	defer func() {
		ticker.Stop()
		p.publishing = false
	}()
	<-ticker.C

	p.subscribers.Range(func(key, value interface{}) bool {
		client := value.(*Subscriber)
		log.Printf("clientOperationType: %v", client.OperationID)
		if client.SubscribedToOperationType(operationID) {
			log.Printf("sent one: %v", operationID)
			client.Send <- operationID
		}
		return true
	})
	return nil
}

func (p *SubscriptionsPublisher) Subscribe(ctx context.Context, subscriber *Subscriber) error {
	for {
		var msg ConnectionACKMessage
		err := wsjson.Read(ctx, subscriber.Conn, &msg)
		if err != nil {
			return err
		}
		if msg.Type == CONNECTIONACK.String() {
			if msg.Payload.Token == "" {
				return ErrAuthenticationFailed
			}
			_, err := application_services.AuthorizationApplicationServiceInstance().Authorize(
				ctx,
				msg.Payload.Token,
			)
			if err != nil {
				return err
			}

			subscriber.Authenticated = true
			subscriber.SetSubscription(msg.Payload.SubscriptionID)
			subscriber.RequestString = msg.Payload.Query
			p.subscribers.Store(subscriber.ID, subscriber)
			return nil
		}
	}
}
func (p *SubscriptionsPublisher) Unsubscribe(subscriber *Subscriber) {
	p.subscribers.Delete(subscriber.ID)
}

// func (p *SubscriptionsPublisher) ProccessFile(client *Subscriber,
// 	msg SubmitFileMessage) (Message, error) {
// 	var response = FileSubmittedMessage{
// 		Email:        client.User.Email,
// 		EnterpriseID: client.Enterprise.Name,
// 		File:         msg.File,
// 		Success:      false,
// 		TriesLeft:    msg.Try,
// 	}
// 	fileName := client.User.Email + "_" + msg.FileName
// 	filePayload, err := infrastructure.NewFilePayload(
// 		fileName,
// 		msg.FileType,
// 		msg.File,
// 	)
// 	if err != nil {
// 		return Message{}, err
// 	}
// 	err = filePayload.Save()
// 	if err != nil {
// 		if errors.Is(err, infrastructure.ErrFileAlreadyExists) {
// 			response.Success = false
// 		} else {
// 			return Message{}, err
// 		}
// 	} else {
// 		response.FileName = fileName
// 		response.Success = true
// 	}
// 	if msg.Try != UNLIMITED_TRIES {
// 		response.TriesLeft--
// 	}
// 	j, err := json.Marshal(response)
// 	if err != nil {
// 		return Message{}, err
// 	}
// 	return Message{
// 		Type:    FileSubmited.String(),
// 		Payload: j,
// 	}, nil
// }

// case NewMessage.String():
// 	p.subscribers.Range(func(key, value interface{}) bool {
// 		client := value.(*Subscriber)
// 		client.Send <- message
// 		return true
// 	})
// case SubmitFile.String():
// 	var msg SubmitFileMessage
// 	err := json.Unmarshal(message.Payload, &msg)
// 	if err != nil {
// 		return err
// 	}
// 	subscriber, err := p.GetSubscriber(message.SubscriptionID)
// 	if err != nil {
// 		return err
// 	}
// 	result, err := p.ProccessFile(subscriber, msg)
// 	if err != nil {
// 		return err
// 	}
// 	p.subscribers.Range(func(key, value interface{}) bool {
// 		client := value.(*Subscriber)
// 		client.Send <- result
// 		return true
// 	})
// case CloseRoom.String():
// 	var msg CloseRoomMessage
// 	err := json.Unmarshal(message.Payload, &msg)
// 	if err != nil {
// 		return err
// 	}
// 	subscriber, err := p.GetSubscriber(message.SubscriptionID)
// 	if err != nil {
// 		return err
// 	}
// 	if subscriber.Authenticated && subscriber.Enterprise.Name == msg.EnterpriseID && subscriber.User.Email == msg.Email {
// 		response := RoomClosedMessage{
// 			Email:        subscriber.User.Email,
// 			EnterpriseID: subscriber.Enterprise.Name,
// 			RoomID:       msg.RoomID,
// 		}
// 		j, err := json.Marshal(response)
// 		if err != nil {
// 			return err
// 		}
// 		result := Message{
// 			Type:    RoomClosed.String(),
// 			Payload: j,
// 		}
// 		p.subscribers.Range(func(key, value interface{}) bool {
// 			client := value.(*Subscriber)
// 			client.Send <- result
// 			return true
// 		})
// 	}

// }
