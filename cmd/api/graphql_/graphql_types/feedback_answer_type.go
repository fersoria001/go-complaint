package graphql_types

import "github.com/graphql-go/graphql"

var FeedbackAnswerType = graphql.NewObject(graphql.ObjectConfig{
	Name: "FeedbackAnswer",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"feedbackID": &graphql.Field{
			Type: graphql.String,
		},
		"senderID": &graphql.Field{
			Type: graphql.String,
		},
		"senderIMG": &graphql.Field{
			Type: graphql.String,
		},
		"senderName": &graphql.Field{
			Type: graphql.String,
		},
		"body": &graphql.Field{
			Type: graphql.String,
		},
		"createdAt": &graphql.Field{
			Type: graphql.String,
		},
		"read": &graphql.Field{
			Type: graphql.Boolean,
		},
		"readAt": &graphql.Field{
			Type: graphql.String,
		},
		"updatedAt": &graphql.Field{
			Type: graphql.String,
		},
		"isEnterprise": &graphql.Field{
			Type: graphql.Boolean,
		},
		"enterpriseID": &graphql.Field{
			Type: graphql.String,
		},
	},
})
