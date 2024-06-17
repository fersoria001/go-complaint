package graphql_types

import "github.com/graphql-go/graphql"

var ComplaintReplyType = graphql.NewObject(graphql.ObjectConfig{
	Name: "ComplaintReply",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.ID,
		},
		"complaintID": &graphql.Field{
			Type: graphql.ID,
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
		"complaintStatus": &graphql.Field{
			Type: graphql.String,
		},
	},
})
