package graphql_arguments

import "github.com/graphql-go/graphql"

var AnswerFeedbackArgument = graphql.FieldConfigArgument{
	"feedbackID": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"answerBody": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
}
