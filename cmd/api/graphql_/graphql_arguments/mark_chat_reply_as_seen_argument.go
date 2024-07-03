package graphql_arguments

import "github.com/graphql-go/graphql"

var MarkChatReplyAsSeenArgument = graphql.FieldConfigArgument{
	"chatID": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	"enterpriseName": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	"repliesID": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
}
