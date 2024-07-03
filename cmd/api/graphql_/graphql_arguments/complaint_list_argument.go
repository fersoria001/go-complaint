package graphql_arguments

import "github.com/graphql-go/graphql"

var ComplaintListArgument = graphql.FieldConfigArgument{
	"id": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"status": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	"term": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	"afterDate": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	"beforeDate": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	"limit": &graphql.ArgumentConfig{
		Type: graphql.Int,
	},
	"offset": &graphql.ArgumentConfig{
		Type: graphql.Int,
	},
}
