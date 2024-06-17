package graphql_types

import "github.com/graphql-go/graphql"

var ComplaintReceiverType = graphql.NewObject(graphql.ObjectConfig{
	Name: "ComplaintReceiver",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"fullName": &graphql.Field{
			Type: graphql.String,
		},
		"thumbnail": &graphql.Field{
			Type: graphql.String,
		},
	},
})
