package graphql_arguments

import "github.com/graphql-go/graphql"

var ReplyComplaintArgument = graphql.FieldConfigArgument{
	"complaintID": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"replyAuthorID": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"replyBody": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"replyEnterpriseID": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
}
