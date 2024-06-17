package server

import (
	"context"
	"encoding/json"
	"errors"
	"log"

	"go-complaint/cmd/api/graphql_"
	"go-complaint/cmd/server/graphql_subscriptions"
	"go-complaint/infrastructure"
	"net/http"
	"sync"

	"github.com/gorilla/csrf"
	"github.com/graphql-go/graphql"
	"nhooyr.io/websocket"
)

type postData struct {
	Query     string                 `json:"query"`
	Operation string                 `json:"operation"`
	Variables map[string]interface{} `json:"variables"`
}

func PublisherHandler(w http.ResponseWriter, r *http.Request) {
	var p map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		r.Body.Close()
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	operationID, ok := p["operation_id"].(string)
	if !ok {
		http.Error(w, "missing operation_id", http.StatusBadRequest)
		return
	}
	err = graphql_subscriptions.SubscriptionsPublisherInstance().Publish(r.Context(), operationID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func GraphQLHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-CSRF-Token", csrf.Token(r))
	var p postData = postData{}
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		r.Body.Close()
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	result := graphql.Do(graphql.Params{
		Context:        r.Context(),
		Schema:         graphql_.Schema,
		RequestString:  p.Query,
		VariableValues: p.Variables,
		OperationName:  p.Operation,
	})
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
	queue := infrastructure.PushNotificationInMemoryQueueInstance()
	queue.SendAll(r.Context())
}

func SubscriptionsHandler(w http.ResponseWriter, r *http.Request) {
	err := subscribe(r.Context(), w, r)
	if errors.Is(err, context.Canceled) {
		return
	}
	if websocket.CloseStatus(err) == websocket.StatusNormalClosure ||
		websocket.CloseStatus(err) == websocket.StatusGoingAway {
		return
	}
	if err != nil {
		return
	}
}

func subscribe(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var conn *websocket.Conn
	conn1, err := websocket.Accept(w, r, &websocket.AcceptOptions{
		OriginPatterns: []string{"localhost:5173"},
	})
	if err != nil {
		return err
	}
	var closed bool
	var mu sync.Mutex
	subscriber := graphql_subscriptions.NewSubscriber(
		infrastructure.InMemoryCacheInstance().NextID(),
		conn1,
		func() {
			mu.Lock()
			defer mu.Unlock()
			closed = true
			if conn != nil {
				conn.Close(
					websocket.StatusPolicyViolation,
					"connection too slow to keep up with messages",
				)
			}
		},
	)
	err = graphql_subscriptions.SubscriptionsPublisherInstance().Subscribe(ctx, subscriber)
	if err != nil {
		log.Printf("error subscribing: %v", err)
		return err
	}
	defer func() {
		graphql_subscriptions.SubscriptionsPublisherInstance().Unsubscribe(subscriber)
	}()
	mu.Lock()
	if closed {
		mu.Unlock()
		http.Error(
			w,
			http.StatusText(http.StatusServiceUnavailable),
			http.StatusServiceUnavailable,
		)
	}
	conn = conn1
	mu.Unlock()
	defer conn.CloseNow()
	ctx = conn.CloseRead(ctx)
	err = subscriber.Listen(ctx)
	if err != nil {
		return err
	}
	return nil
}
