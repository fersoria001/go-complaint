package graphql_subscriptions

import (
	"context"
	"go-complaint/domain/model/enterprise"
	"strings"
	"sync"

	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

type Subscriber struct {
	ID            int
	Send          chan DataMessage
	Conn          *websocket.Conn
	Authenticated bool
	Mutex         sync.Mutex
	RequestString string
	OperationID   string
	CloseFunc     func()
}

func NewSubscriber(
	id int,
	conn *websocket.Conn,
	closeFunc func(),
) *Subscriber {
	return &Subscriber{
		ID:        id,
		Conn:      conn,
		Send:      make(chan DataMessage),
		CloseFunc: closeFunc,
	}
}

// sub.OperationID?sub.OperationID?
func (sub *Subscriber) SubscribedToOperationType(operationID string) bool {
	s := strings.Split(sub.OperationID, "?")
	if chatId, err := enterprise.NewChatID(operationID); err == nil {
		for i := range s {
			if s[i] == chatId.String() || s[i] == chatId.Reverse().String() ||
				s[i] == chatId.Partial() || s[i] == chatId.ReversePartial() {
				return true
			}
		}
	}
	for i := range s {
		if s[i] == operationID {
			return true
		}
	}

	return false
}

func (sub *Subscriber) SetSubscription(operationID string) {

	sub.OperationID = operationID
}

func (client *Subscriber) Listen(ctx context.Context) error {
	for {

		select {
		case msg := <-client.Send:
			err := wsjson.Write(ctx, client.Conn, msg)
			if err != nil {
				return err
			}
		case <-ctx.Done():
			return ErrContextDone
		}
	}
}
