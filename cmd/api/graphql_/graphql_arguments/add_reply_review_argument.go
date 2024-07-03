package graphql_arguments

import "github.com/graphql-go/graphql"

var FeedbackArgument = graphql.FieldConfigArgument{
	"complaintID": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	"feedbackID": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	"enterpriseID": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"reviewerID": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	"comment": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	"color": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	"repliesID": &graphql.ArgumentConfig{
		Type: graphql.NewList(graphql.String),
	},
}
