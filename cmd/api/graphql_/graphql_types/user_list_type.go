package graphql_types

import "github.com/graphql-go/graphql"

var UserListType = graphql.NewObject(graphql.ObjectConfig{
	Name: "UserList",
	Fields: graphql.Fields{
		"users": &graphql.Field{
			Type: graphql.NewList(UserType),
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
