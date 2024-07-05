package server

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"

	"go-complaint/application"
	"go-complaint/cmd/api/graphql_"
	"go-complaint/cmd/server/graphql_subscriptions"
	"go-complaint/domain"
	"go-complaint/infrastructure/cache"
	"net/http"
	"sync"

	"github.com/graphql-go/graphql"
	"nhooyr.io/websocket"
)

type postData struct {
	Query     string                 `json:"query"`
	Operation string                 `json:"operation"`
	Variables map[string]interface{} `json:"variables"`
}

func PublisherHandler(w http.ResponseWriter, r *http.Request) {
	operationID := r.Header.Get("subscription-id")
	if operationID == "" {
		http.Error(w, "operation-id header is required", http.StatusBadRequest)
		return
	}
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
	if result.HasErrors() {
		http.Error(w, "error processing request", http.StatusInternalServerError)
		return
	}
	payload, ok := cache.InMemoryCacheInstance().GetPublish(operationID)
	if !ok {
		http.Error(w, fmt.Errorf("no subscriber found for operation %s", operationID).Error(), http.StatusInternalServerError)
		return
	}
	dataMessage := graphql_subscriptions.DataMessage{
		Type:        graphql_subscriptions.DATA.String(),
		OperationID: operationID,
		Payload:     payload,
	}
	err = graphql_subscriptions.SubscriptionsPublisherInstance().Publish(r.Context(), dataMessage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func GraphQLHandler(w http.ResponseWriter, r *http.Request) {
	domain.DomainEventPublisherInstance().Reset()
	ep := application.EventProcessor{}
	domain.DomainEventPublisherInstance().Subscribe(
		domain.DomainEventSubscriber{
			HandleEvent:           ep.HandleEvent,
			SubscribedToEventType: ep.SubscribedToEventType,
		},
	)

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

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(result)
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
	secWebsocketProtocolHeader := r.Header.Get("Sec-Websocket-Protocol")
	protocols := strings.Split(secWebsocketProtocolHeader, ",")
	originsEnv := os.Getenv("ORIGIN")
	origins := strings.Split(originsEnv, ",")
	conn1, err := websocket.Accept(w, r, &websocket.AcceptOptions{
		OriginPatterns: origins,
		Subprotocols:   protocols,
	})
	if err != nil {
		return err
	}
	var closed bool
	var mu sync.Mutex
	subscriber := graphql_subscriptions.NewSubscriber(
		cache.InMemoryCacheInstance().NextID(),
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
