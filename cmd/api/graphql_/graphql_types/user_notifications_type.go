package graphql_types

import "github.com/graphql-go/graphql"

var NotificationsType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Notifications",
	Fields: graphql.Fields{
		"ownerID": &graphql.Field{
			Type: graphql.String,
		},
		"thumbnail": &graphql.Field{
			Type: graphql.String,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"content": &graphql.Field{
			Type: graphql.String,
		},
		"link": &graphql.Field{
			Type: graphql.String,
		},
		"seen": &graphql.Field{
			Type: graphql.Boolean,
		},
	},
})
