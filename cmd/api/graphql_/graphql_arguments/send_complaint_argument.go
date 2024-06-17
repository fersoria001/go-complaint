package graphql_arguments

import "github.com/graphql-go/graphql"

var SendComplaintArgument = graphql.FieldConfigArgument{
	"authorID": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"receiverID": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"title": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"description": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"content": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
}
