package graphql_types

import "github.com/graphql-go/graphql"

var HiringInvitationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "HiringInvitation",
	Fields: graphql.Fields{
		"eventID": &graphql.Field{
			Type: graphql.String,
		},
		"enterpriseID": &graphql.Field{
			Type: graphql.String,
		},
		"proposedPosition": &graphql.Field{
			Type: graphql.String,
		},
		"ownerID": &graphql.Field{
			Type: graphql.String,
		},
		"fullName": &graphql.Field{
			Type: graphql.String,
		},
		"enterpriseEmail": &graphql.Field{
			Type: graphql.String,
		},
		"enterprisePhone": &graphql.Field{
			Type: graphql.String,
		},
		"enterpriseLogoIMG": &graphql.Field{
			Type: graphql.String,
		},
		"occurredOn": &graphql.Field{
			Type: graphql.String,
		},
		"seen": &graphql.Field{
			Type: graphql.String,
		},
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"reason": &graphql.Field{
			Type: graphql.String,
		},
	},
})
