package graphql_types

import "github.com/graphql-go/graphql"

var FeedbackType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Feedback",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"complaintID": &graphql.Field{
			Type: graphql.String,
		},
		"reviewedID": &graphql.Field{
			Type: graphql.String,
		},
		"replyReview": &graphql.Field{
			Type: graphql.NewList(ReplyReviewType),
		},
		"feedbackAnswer": &graphql.Field{
			Type: graphql.NewList(FeedbackAnswerType),
		},
	},
})
