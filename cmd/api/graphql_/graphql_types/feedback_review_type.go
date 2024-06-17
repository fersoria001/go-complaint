package graphql_types

import "github.com/graphql-go/graphql"

var ReviewType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Review",
	Fields: graphql.Fields{
		"replyReviewID": &graphql.Field{
			Type: graphql.String,
		},
		"reviewerID": &graphql.Field{
			Type: graphql.String,
		},
		"reviewedAt": &graphql.Field{
			Type: graphql.String,
		},
		"comment": &graphql.Field{
			Type: graphql.String,
		},
	},
})
