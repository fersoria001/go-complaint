package graphql_types

import "github.com/graphql-go/graphql"

var ComplaintListType = graphql.NewObject(graphql.ObjectConfig{
	Name: "ComplaintList",
	Fields: graphql.Fields{
		"complaints": &graphql.Field{
			Type: graphql.NewList(ComplaintType),
		},
		"count": &graphql.Field{
			Type: graphql.Int,
		},
		"currentLimit": &graphql.Field{
			Type: graphql.Int,
		},
		"currentOffset": &graphql.Field{
			Type: graphql.Int,
		},
	},
})
