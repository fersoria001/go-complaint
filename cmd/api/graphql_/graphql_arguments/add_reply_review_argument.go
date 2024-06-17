package graphql_arguments

import "github.com/graphql-go/graphql"

var AddReplyReviewArgument = graphql.FieldConfigArgument{
	"id": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"reviewedID": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"reviewerID": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"reviewComment": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"replyReviewColor": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"repliesID": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.NewList(graphql.String)),
	},
}
