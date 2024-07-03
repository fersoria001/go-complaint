package graphql_

import (
	"go-complaint/cmd/api/graphql_/graphql_arguments"
	"go-complaint/cmd/api/graphql_/graphql_types"
	"go-complaint/cmd/api/graphql_/resolvers/notification_resolvers"

	"github.com/graphql-go/graphql"
)

var subscription = graphql.NewObject(graphql.ObjectConfig{
	Name: "Subscription",
	Fields: graphql.Fields{
		"notifications": &graphql.Field{
			Type:    graphql.NewList(graphql_types.NotificationsType),
			Args:    graphql_arguments.StringID,
			Resolve: notification_resolvers.NotificationsResolver,
		},
	},
})
