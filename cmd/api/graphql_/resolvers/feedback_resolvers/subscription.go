package feedback_resolvers

import (
	"go-complaint/application/queries"

	"github.com/graphql-go/graphql"
)

func ComplaintLastReplyResolver(params graphql.ResolveParams) (interface{}, error) {
	q := queries.FeedbackQuery{
		ReplyID: params.Args["id"].(string),
	}
	reply, err := q.FeedbackLastReply(params.Context)
	if err != nil {
		return nil, err
	}
	return reply, nil
}
