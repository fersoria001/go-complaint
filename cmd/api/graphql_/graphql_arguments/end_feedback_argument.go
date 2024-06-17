package graphql_arguments

import "github.com/graphql-go/graphql"

var EndFeedbackArgument = graphql.FieldConfigArgument{
	"complaintID": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"reviewerID": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
}
