package graphql_arguments

import "github.com/graphql-go/graphql"

var EnterpriseEventArgument = graphql.FieldConfigArgument{
	"enterpriseName": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"eventID": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"reason": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
}
