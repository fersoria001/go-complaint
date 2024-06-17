package graphql_types

import "github.com/graphql-go/graphql"

var ComplaintMessageType = graphql.NewObject(graphql.ObjectConfig{
	Name: "ComplaintMessage",
	Fields: graphql.Fields{
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"description": &graphql.Field{
			Type: graphql.String,
		},
		"body": &graphql.Field{
			Type: graphql.String,
		},
	},
})
