package graphql_arguments

import "github.com/graphql-go/graphql"

var RateComplaintArgument = graphql.FieldConfigArgument{
	"complaintId": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"eventId": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"rate": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.Int),
	},
	"comment": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
}
