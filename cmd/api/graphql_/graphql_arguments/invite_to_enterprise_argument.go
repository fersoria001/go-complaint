package graphql_arguments

import "github.com/graphql-go/graphql"

var InviteToEnterpriseArgument = graphql.FieldConfigArgument{
	"enterpriseName": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"proposedPosition": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"proposeTo": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
}
