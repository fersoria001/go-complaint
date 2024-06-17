package server_test

import (
	"context"
	"encoding/json"
	"fmt"
	"go-complaint/application/commands"
	"go-complaint/cmd/server"
	"go-complaint/cmd/server/graphql_subscriptions"
	"go-complaint/infrastructure"
	"go-complaint/tests"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

func TestPubSub(t *testing.T) {
	t.Parallel()
	url, closeFn := setup(t)
	defer closeFn()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()
	cl, err := newClient(ctx, url)
	assert.Nil(t, err)
	defer cl.Close()
	command := commands.ComplaintCommand{
		ReplyAuthorID: tests.UserRegisterAndVerifyCommands["1"].Email,
		ReplyBody:     tests.RepeatString("ja", 10),
		ID:            "c80015e6-bf46-47c3-b48f-2ee932de11f8",
	}
	err = command.Reply(ctx)
	assert.Nil(t, err)

	auth := graphql_subscriptions.ConnectionACKMessage{
		Type:        "connection_ack",
		OperationID: "mock-client@user.com",
		Payload: struct {
			Query          string `json:"query"`
			SubscriptionID string `json:"subscription_id"`
			Token          string `json:"token"`
			EnterpriseID   string `json:"enterprise_id"`
		}{
			Query:          "subscription {complaintLastReply(ID: \"c80015e6-bf46-47c3-b48f-2ee932de11f8\") { id, complaintID, senderID, senderIMG, senderName, body, createdAt, read, readAt, updatedAt, isEnterprise, enterpriseID}}",
			SubscriptionID: "complaintLastReply-c80015e6-bf46-47c3-b48f-2ee932de11f8",
			Token:          "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImJlcmNobzAwMUBnbWFpbC5jb20iLCJmdWxsX25hbWUiOiJGZXJuYW5kbyBBZ3VzdGluIFNvcmlhIiwicHJvZmlsZV9pbWciOiIvZGVmYXVsdC5qcGciLCJnZW5kZXIiOiJNQUxFIiwicHJvbm91biI6IkhlIiwiY2xpZW50X2RhdGEiOnsiaXAiOiIiLCJkZXZpY2UiOiIiLCJnZW9sb2NhbGl6YXRpb24iOnsibGF0aXR1ZGUiOjAsImxvbmdpdHVkZSI6MH0sImxvZ2luX2RhdGUiOiIxNzE4MzIzNDI1MTEwIn0sInJlbWVtYmVyX21lIjp0cnVlLCJhdXRob3JpdGllcyI6bnVsbCwiZXhwIjoxNzE4NDA5ODI1LCJpYXQiOjE3MTgzMjM0MjV9.1fWYrGVWP5YrYm_D0Ok6ifkWrixU3Ot83tgUOOAFWj0",
			EnterpriseID:   "",
		},
	}
	err = wsjson.Write(ctx, cl.c, auth)
	assert.Nil(t, err)
	infrastructure.PushNotificationInMemoryQueueInstance().SendAll(ctx)
	l := infrastructure.PushNotificationInMemoryQueueInstance().SentLog()
	log.Printf("log: %v", l)
	assert.Nil(t, err)
	msg, err := cl.nextMessage()
	assert.Nil(t, err)
	t.Logf("msg: %v", msg.Payload)
}

func (cl *client) publish(ctx context.Context, msg string) (err error) {
	defer func() {
		if err != nil {
			cl.c.Close(websocket.StatusInternalError, "publish failed")
		}
	}()

	req, _ := http.NewRequestWithContext(ctx, http.MethodPost, cl.url+"/graphql", strings.NewReader(msg))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusAccepted {
		return fmt.Errorf("publish request failed: %v", resp.StatusCode)
	}
	return nil
}
func setup(t *testing.T) (url string, closeFunc func()) {
	t.Log("Before all tests")
	mux := http.NewServeMux()
	mux.HandleFunc("/graphql", server.GraphQLHandler)
	mux.HandleFunc("/subscriptions", server.SubscriptionsHandler)
	mux.HandleFunc("/publish", server.PublisherHandler)
	ts := &ts{mux: mux}
	s := httptest.NewServer(ts)
	return s.URL, func() {
		s.Close()
	}
}

type ts struct {
	mux *http.ServeMux
}

func (t *ts) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.mux.ServeHTTP(w, r)
}

type client struct {
	url string
	c   *websocket.Conn
}

func newClient(ctx context.Context, url string) (*client, error) {
	c, _, err := websocket.Dial(ctx, url+"/subscriptions", nil)
	if err != nil {
		return nil, err
	}

	cl := &client{
		url: url,
		c:   c,
	}

	return cl, nil
}

func (cl *client) Close() error {
	return cl.c.Close(websocket.StatusNormalClosure, "")
}

func (cl *client) nextMessage() (graphql_subscriptions.GraphQLResultMessage, error) {
	typ, b, err := cl.c.Read(context.Background())
	if err != nil {
		return graphql_subscriptions.GraphQLResultMessage{}, err
	}

	if typ != websocket.MessageText {
		cl.c.Close(websocket.StatusUnsupportedData, "expected text message")
		return graphql_subscriptions.GraphQLResultMessage{}, fmt.Errorf("expected text message but got %v", typ)
	}

	var msg graphql_subscriptions.GraphQLResultMessage
	err = json.Unmarshal(b, &msg)
	if err != nil {
		return graphql_subscriptions.GraphQLResultMessage{}, err
	}

	return msg, nil
}
