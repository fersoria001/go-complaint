package graphql_types

import "github.com/graphql-go/graphql"

var ComplaintInfoType = graphql.NewObject(graphql.ObjectConfig{
	Name: "ComplaintInfo",
	Fields: graphql.Fields{
		"complaintsReceived": &graphql.Field{
			Type: graphql.Int,
		},
		"complaintsResolved": &graphql.Field{
			Type: graphql.Int,
		},
		"complaintsReviewed": &graphql.Field{
			Type: graphql.Int,
		},
		"complaintsPending": &graphql.Field{
			Type: graphql.Int,
		},
		"averageRating": &graphql.Field{
			Type: graphql.Float,
		},
	},
})
