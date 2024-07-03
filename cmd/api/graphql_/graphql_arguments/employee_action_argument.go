package graphql_arguments

import "github.com/graphql-go/graphql"

var EmployeeActionArgument = graphql.FieldConfigArgument{
	"enterpriseName": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	"employeeID": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
}
