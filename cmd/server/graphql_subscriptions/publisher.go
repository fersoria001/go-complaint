package graphql_subscriptions

import (
	"context"
	"fmt"
	"go-complaint/application/application_services"
	"go-complaint/dto"
	"go-complaint/infrastructure/cache"
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

func (p *SubscriptionsPublisher) Background(ch chan cache.Request) {
	for {
		req := <-ch

		dataMessage := DataMessage{
			Type:        DATA.String(),
			OperationID: req.Key,
			Payload:     req.Payload,
		}
		ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*1))
		p.Publish(ctx, dataMessage)
		cancel()
	}
}

func (p *SubscriptionsPublisher) Publish(
	ctx context.Context, data DataMessage) error {
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
		if client.SubscribedToOperationType(data.OperationID) {
			log.Println("is subscribed send to", data.OperationID)
			client.Send <- data
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
			if pair := cache.RecognizeId(subscriber.OperationID); pair != nil {
				cache.RequestChannel2 <- cache.Request{
					Type:    cache.WRITE,
					Payload: pair,
					Key:     "ENTERPRISE_USER",
				}
				cache.RequestChannel <- cache.Request{
					Type: cache.WRITE,
					Payload: &cache.Pair{
						One: pair.Two,
						Two: dto.ONLINE.String(),
					},
					Key: fmt.Sprintf("chat:%s", pair.One),
				}
			}
			return nil
		}
	}
}
func (p *SubscriptionsPublisher) Unsubscribe(subscriber *Subscriber) {
	if pair := cache.RecognizeId(subscriber.OperationID); pair != nil {
		cache.RequestChannel2 <- cache.Request{
			Type:    cache.DELETE,
			Key:     "ENTERPRISE_USER",
			Payload: pair,
		}
		cache.RequestChannel <- cache.Request{
			Type: cache.WRITE,
			Payload: &cache.Pair{
				One: pair.Two,
				Two: dto.OFFLINE.String(),
			},
			Key: fmt.Sprintf("chat:%s", pair.One),
		}
	}
	p.subscribers.Delete(subscriber.ID)
}
