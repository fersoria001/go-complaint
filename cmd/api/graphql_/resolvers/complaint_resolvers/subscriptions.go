package complaint_resolvers

import (
	"go-complaint/application/queries"

	"github.com/graphql-go/graphql"
)

func ComplaintLastReplyResolver(params graphql.ResolveParams) (interface{}, error) {
	q := queries.ComplaintQuery{
		ID: params.Args["id"].(string),
	}
	reply, err := q.ComplaintLastReply(params.Context)
	if err != nil {
		return nil, err
	}
	return reply, nil
}
