package cmd_test

import (
	"context"
	"encoding/json"
	"go-complaint/cmd/api/middleware"
	"go-complaint/cmd/server"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

type ts struct {
	mux *http.ServeMux
}

func (t *ts) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.mux.ServeHTTP(w, r)
}
func setup(t *testing.T) (url string, closeFunc func()) {
	t.Log("Before all tests")
	mux := http.NewServeMux()
	mux.HandleFunc("/graphql", middleware.CORS()(middleware.Chain(server.GraphQLHandler,
		middleware.AuthenticationMiddleware(),
	)))
	mux.HandleFunc("/subscriptions", server.SubscriptionsHandler)
	mux.HandleFunc("/publish", server.PublisherHandler)
	ts := &ts{mux: mux}
	s := httptest.NewServer(ts)

	return s.URL, func() {
		s.Close()
	}
}

type r struct {
	Query string `json:"query"`
}

func TestHandler(t *testing.T) {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(1*time.Second))
	defer cancel()
	url, closer := setup(t)
	defer closer()
	rq := r{Query: `mutation {ReplyComplaint(complaintID: "c80015e6-bf46-47c3-b48f-2ee932de11f8", replyAuthorID: "bercho001@gmail.com", replyBody: "JAAAAAAAAAAA") }`}
	j, err := json.Marshal(rq)
	if err != nil {
		t.Fatal(err)
	}
	rquesbody := strings.NewReader(string(j))
	req, _ := http.NewRequestWithContext(ctx, http.MethodPost, url+"/graphql", rquesbody)
	req.Header.Set("Authorization", string(`Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImJlcmNobzAwMUBnbWFpbC5jb20iLCJmdWxsX25hbWUiOiJGZXJuYW5kbyBBZ3VzdGluIFNvcmlhIiwicHJvZmlsZV9pbWciOiIvZGVmYXVsdC5qcGciLCJnZW5kZXIiOiJNQUxFIiwicHJvbm91biI6IkhlIiwiY2xpZW50X2RhdGEiOnsiaXAiOiIiLCJkZXZpY2UiOiIiLCJnZW9sb2NhbGl6YXRpb24iOnsibGF0aXR1ZGUiOjAsImxvbmdpdHVkZSI6MH0sImxvZ2luX2RhdGUiOiIxNzE4MzI4ODQ3NjA2In0sInJlbWVtYmVyX21lIjp0cnVlLCJhdXRob3JpdGllcyI6bnVsbCwiZXhwIjoxNzE4NDE1MjQ3LCJpYXQiOjE3MTgzMjg4NDd9.B4BQnmX_Qq2a3DBljYZmr59TLHtLtmlGxq4THU67ies`))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		t.Fatalf("request failed: %v", resp.Status)
	}
	var v map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&v)
	if err != nil {
		t.Fatal(err)
	}

}
