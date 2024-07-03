package graphql_types

import "github.com/graphql-go/graphql"

var HiringProccessListType = graphql.NewObject(graphql.ObjectConfig{
	Name: "HiringProccessList",
	Fields: graphql.Fields{
		"hiringProccesses": &graphql.Field{
			Type: graphql.NewList(HiringProccessType),
		},
		"count": &graphql.Field{
			Type: graphql.Int,
		},
		"currentOffset": &graphql.Field{
			Type: graphql.Int,
		},
		"currentLimit": &graphql.Field{
			Type: graphql.Int,
		},
	},
})
var HiringProccessType = graphql.NewObject(graphql.ObjectConfig{
	Name: "HiringProccess",
	Fields: graphql.Fields{
		"eventID": &graphql.Field{
			Type: graphql.String,
		},
		"user": &graphql.Field{
			Type: UserType,
		},
		"position": &graphql.Field{
			Type: graphql.String,
		},
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"emitedBy": &graphql.Field{
			Type: UserType,
		},
		"occurredOn": &graphql.Field{
			Type: graphql.String,
		},
		"lastUpdate": &graphql.Field{
			Type: graphql.String,
		},
		"reason": &graphql.Field{
			Type: graphql.String,
		},
	},
})
