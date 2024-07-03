package graphql_arguments

import "github.com/graphql-go/graphql"

var EnterpriseChatArgument = graphql.FieldConfigArgument{
	"enterpriseID": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"chatID": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
}
