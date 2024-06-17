package graphql_types

import "github.com/graphql-go/graphql"

var ComplaintType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Complaint",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.ID,
		},
		"authorID": &graphql.Field{
			Type: graphql.String,
		},
		"receiverID": &graphql.Field{
			Type: graphql.String,
		},
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"message": &graphql.Field{
			Type: ComplaintMessageType,
		},
		"rating": &graphql.Field{
			Type: ComplaintRatingType,
		},
		"createdAt": &graphql.Field{
			Type: graphql.String,
		},
		"updatedAt": &graphql.Field{
			Type: graphql.String,
		},
		"replies": &graphql.Field{
			Type: graphql.NewList(ComplaintReplyType),
		},
	},
})
