package graphql_arguments

import "github.com/graphql-go/graphql"

var AnswerFeedbackArgument = graphql.FieldConfigArgument{
	"feedbackOD": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"answerBody": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
}
