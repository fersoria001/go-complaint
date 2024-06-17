package graphql_

import (
	"go-complaint/application/queries"
	"go-complaint/cmd/api/graphql_/graphql_arguments"
	"go-complaint/cmd/api/graphql_/graphql_types"
	"go-complaint/cmd/api/graphql_/resolvers/complaint_resolvers"
	"go-complaint/cmd/api/graphql_/resolvers/feedback_resolvers"

	"github.com/graphql-go/graphql"
)

var subscription = graphql.NewObject(graphql.ObjectConfig{
	Name: "Subscription",
	Fields: graphql.Fields{
		"notifications": &graphql.Field{
			Type: graphql_types.NotificationsType,
			Args: graphql_arguments.StringID,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				nq := queries.NotificationQuery{
					OwnerID: p.Args["ID"].(string),
				}
				notifications, err := nq.Notifications(p.Context)
				if err != nil {
					return nil, err
				}
				return notifications, nil
			},
		},
		"complaintLastReply": &graphql.Field{
			Type:    graphql_types.ComplaintReplyType,
			Args:    graphql_arguments.StringID,
			Resolve: complaint_resolvers.ComplaintLastReplyResolver,
		},
		"feedbackLastReply": &graphql.Field{
			Type:    graphql_types.FeedbackAnswerType,
			Args:    graphql_arguments.StringID,
			Resolve: feedback_resolvers.ComplaintLastReplyResolver,
		},
	},
})
