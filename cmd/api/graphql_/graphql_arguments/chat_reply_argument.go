package graphql_arguments

import "github.com/graphql-go/graphql"

var ChatReplyArgument = graphql.FieldConfigArgument{
	"id": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	"enterpriseName": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	"senderID": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	"content": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
}
