package graphql_arguments

import "github.com/graphql-go/graphql"

var ContactArgument = graphql.FieldConfigArgument{
	"email": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"text": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
}
