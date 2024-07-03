package graphql_types

import "github.com/graphql-go/graphql"

var ReplyReviewType = graphql.NewObject(graphql.ObjectConfig{
	Name: "ReplyReview",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"feedbackID": &graphql.Field{
			Type: graphql.String,
		},
		"reviewer": &graphql.Field{
			Type: UserType,
		},
		"replies": &graphql.Field{
			Type: graphql.NewList(ComplaintReplyType),
		},
		"review": &graphql.Field{
			Type: ReviewType,
		},
		"color": &graphql.Field{
			Type: graphql.String,
		},
		"createdAt": &graphql.Field{
			Type: graphql.String,
		},
	},
})
