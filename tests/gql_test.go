package tests

import (
	"context"
	"go-complaint/cmd/api/graphql_"
	"log"
	"testing"

	"github.com/graphql-go/graphql"
)

func TestGQL(t *testing.T) {
	query := "subscription {complaintLastReply(id: \"c80015e6-bf46-47c3-b48f-2ee932de11f8\") { id, complaintID, senderID, senderIMG, senderName, body, createdAt, read, readAt, updatedAt, isEnterprise, enterpriseID}}"
	result := graphql.Do(graphql.Params{
		Context:       context.Background(),
		Schema:        graphql_.Schema,
		RequestString: query,
	})
	log.Printf("gqlResult: %v", result)
}
