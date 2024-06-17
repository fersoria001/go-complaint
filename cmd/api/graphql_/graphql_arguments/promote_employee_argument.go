package graphql_arguments

import "github.com/graphql-go/graphql"

var PromoteEmployeeArgument = graphql.FieldConfigArgument{
	"enterpriseName": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"employeeID": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"position": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
}
