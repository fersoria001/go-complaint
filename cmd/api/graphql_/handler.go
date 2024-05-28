package graphql_

import (
	"context"
	"encoding/json"
	"go-complaint/application"
	"go-complaint/domain"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
)

type postData struct {
	Query     string                 `json:"query"`
	Operation string                 `json:"operation"`
	Variables map[string]interface{} `json:"variables"`
}

func ExecuteBody(ctx context.Context, p postData, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Context:        ctx,
		Schema:         schema,
		RequestString:  p.Query,
		VariableValues: p.Variables,
		OperationName:  p.Operation,
	})
	return result
}

func ExecuteAndEncodeBody(w http.ResponseWriter, r *http.Request) {
	publisher := domain.DomainEventPublisherInstance()
	publisher.Reset()
	eventProccesor := application.NewEventProcessor()
	publisher.Subscribe(eventProccesor.Subscriber())
	var p postData
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		w.WriteHeader(400)
		return
	}
	result := ExecuteBody(r.Context(), p, Schema)
	if err := json.NewEncoder(w).Encode(result); err != nil {
		log.Println("could not write result to response: ", err)
	}
}
