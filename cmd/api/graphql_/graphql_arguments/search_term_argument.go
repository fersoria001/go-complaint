package graphql_arguments

import "github.com/graphql-go/graphql"

var SearchTermArgument = graphql.FieldConfigArgument{
	"id": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"term": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
}
