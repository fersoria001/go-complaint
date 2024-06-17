package graphql_arguments

import "github.com/graphql-go/graphql"

var SearchTermArgument = graphql.FieldConfigArgument{
	"term": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
}
