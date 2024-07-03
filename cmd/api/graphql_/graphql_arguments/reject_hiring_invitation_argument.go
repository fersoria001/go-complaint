package graphql_arguments

import "github.com/graphql-go/graphql"

var RejectHiringInvitationArgument = graphql.FieldConfigArgument{
	"id": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"reason": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
}
