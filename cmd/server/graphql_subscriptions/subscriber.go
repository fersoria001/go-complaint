package graphql_subscriptions

import (
	"context"
	"go-complaint/cmd/api/graphql_"
	"log"
	"sync"

	"github.com/graphql-go/graphql"
	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

type Subscriber struct {
	ID            int
	Send          chan string
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
		Send:      make(chan string),
		CloseFunc: closeFunc,
	}
}

func (sub *Subscriber) SubscribedToOperationType(operationID string) bool {
	log.Printf("sub.OperationID: %v, operationID: %v == %v", sub.OperationID, operationID, sub.OperationID == operationID)
	return sub.OperationID == operationID
}

func (sub *Subscriber) SetSubscription(operationID string) {
	sub.OperationID = operationID
}

func (client *Subscriber) Listen(ctx context.Context) error {
	for {
		select {
		case <-client.Send:
			log.Printf("client.Send: ")
			gqlResult := graphql.Do(graphql.Params{
				Context:       ctx,
				Schema:        graphql_.Schema,
				RequestString: client.RequestString,
			})
			result := GraphQLResultMessage{
				Type:    DATA.String(),
				Payload: gqlResult,
			}
			err := wsjson.Write(ctx, client.Conn, result)
			if err != nil {
				return err
			}
		case <-ctx.Done():
			return ErrContextDone
		}
	}
}
