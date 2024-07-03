package graphql_types

import "github.com/graphql-go/graphql"

var EnterpriseChatType = graphql.NewObject(graphql.ObjectConfig{
	Name: "EnterpriseChat",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"replies": &graphql.Field{
			Type: graphql.NewList(ChatReplyType),
		},
	},
})

var ChatReplyType = graphql.NewObject(graphql.ObjectConfig{
	Name: "ChatReply",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"chatID": &graphql.Field{
			Type: graphql.String,
		},
		"user": &graphql.Field{
			Type: UserType,
		},
		"content": &graphql.Field{
			Type: graphql.String,
		},
		"seen": &graphql.Field{
			Type: graphql.Boolean,
		},
		"createdAt": &graphql.Field{
			Type: graphql.String,
		},
		"updatedAt": &graphql.Field{
			Type: graphql.String,
		},
	},
})
