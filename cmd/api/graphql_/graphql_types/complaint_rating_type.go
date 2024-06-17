package graphql_types

import "github.com/graphql-go/graphql"

var ComplaintRatingType = graphql.NewObject(graphql.ObjectConfig{
	Name: "ComplaintRating",
	Fields: graphql.Fields{
		"rate": &graphql.Field{
			Type: graphql.Int,
		},
		"comment": &graphql.Field{
			Type: graphql.String,
		},
	},
})
