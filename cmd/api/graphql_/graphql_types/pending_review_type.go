package graphql_types

import "github.com/graphql-go/graphql"

var PendingReviewType = graphql.NewObject(graphql.ObjectConfig{
	Name: "PendingReview",
	Fields: graphql.Fields{
		"eventID": &graphql.Field{
			Type: graphql.String,
		},
		"complaint": &graphql.Field{
			Type: ComplaintType,
		},
		"triggeredBy": &graphql.Field{
			Type: UserType,
		},
	},
})
