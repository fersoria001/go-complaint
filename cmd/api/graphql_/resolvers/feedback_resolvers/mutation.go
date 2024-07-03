package feedback_resolvers

import (
	"go-complaint/application/application_services"
	"go-complaint/application/commands"

	"github.com/graphql-go/graphql"
)

func CreateFeedbackResolver(p graphql.ResolveParams) (interface{}, error) {
	_, err := application_services.AuthorizationApplicationServiceInstance().ResourceAccess(
		p.Context,
		"Enterprise",
		p.Args["enterpriseID"].(string),
		application_services.WRITE,
		"MANAGER", "ASSISTANT", "OWNER",
	)
	if err != nil {
		return false, err
	}
	command := commands.FeedbackCommand{
		EnterpriseID: p.Args["enterpriseID"].(string),
		ComplaintID:  p.Args["complaintID"].(string),
		Color:        p.Args["color"].(string),
	}
	err = command.CreateFeedback(p.Context)
	if err != nil {
		return false, err
	}
	return true, nil
}

func AddReplyResolver(p graphql.ResolveParams) (interface{}, error) {
	_, err := application_services.AuthorizationApplicationServiceInstance().ResourceAccess(
		p.Context,
		"Enterprise",
		p.Args["enterpriseID"].(string),
		application_services.WRITE,
		"MANAGER", "ASSISTANT", "OWNER",
	)
	if err != nil {
		return false, err
	}
	replies := p.Args["repliesID"].([]interface{})
	var repliesID []string
	for _, reply := range replies {
		repliesID = append(repliesID, reply.(string))
	}
	command := commands.FeedbackCommand{
		FeedbackID: p.Args["feedbackID"].(string),
		ReviewerID: p.Args["reviewerID"].(string),
		Color:      p.Args["color"].(string),
		Replies:    repliesID,
	}
	err = command.AddReply(p.Context)
	if err != nil {
		return false, err
	}
	return true, nil
}

func RemoveReplyResolver(p graphql.ResolveParams) (interface{}, error) {
	_, err := application_services.AuthorizationApplicationServiceInstance().ResourceAccess(
		p.Context,
		"Enterprise",
		p.Args["enterpriseID"].(string),
		application_services.WRITE,
		"MANAGER", "ASSISTANT", "OWNER",
	)
	if err != nil {
		return false, err
	}
	replies := p.Args["repliesID"].([]interface{})
	var repliesID []string
	for _, reply := range replies {
		repliesID = append(repliesID, reply.(string))
	}
	command := commands.FeedbackCommand{
		FeedbackID: p.Args["feedbackID"].(string),
		Color:      p.Args["color"].(string),
		Replies:    repliesID,
	}
	err = command.RemoveReply(p.Context)
	if err != nil {
		return false, err
	}
	return true, nil
}

func AddCommentResolver(p graphql.ResolveParams) (interface{}, error) {
	_, err := application_services.AuthorizationApplicationServiceInstance().ResourceAccess(
		p.Context,
		"Enterprise",
		p.Args["enterpriseID"].(string),
		application_services.WRITE,
		"MANAGER", "ASSISTANT", "OWNER",
	)
	if err != nil {
		return false, err
	}
	command := commands.FeedbackCommand{
		FeedbackID: p.Args["feedbackID"].(string),
		Color:      p.Args["color"].(string),
		Comment:    p.Args["comment"].(string),
	}
	err = command.AddComment(p.Context)
	if err != nil {
		return false, err
	}
	return true, nil
}

func DeleteCommentResolver(p graphql.ResolveParams) (interface{}, error) {
	_, err := application_services.AuthorizationApplicationServiceInstance().ResourceAccess(
		p.Context,
		"Enterprise",
		p.Args["enterpriseID"].(string),
		application_services.WRITE,
		"MANAGER", "ASSISTANT", "OWNER",
	)
	if err != nil {
		return false, err
	}
	command := commands.FeedbackCommand{
		FeedbackID: p.Args["feedbackID"].(string),
		Color:      p.Args["color"].(string),
	}
	err = command.DeleteComment(p.Context)
	if err != nil {
		return false, err
	}
	return true, nil
}

func EndFeedbackResolver(p graphql.ResolveParams) (interface{}, error) {
	credentials, err := application_services.AuthorizationApplicationServiceInstance().ResourceAccess(
		p.Context,
		"Enterprise",
		p.Args["enterpriseID"].(string),
		application_services.WRITE,
		"MANAGER", "ASSISTANT",
	)
	if err != nil {
		return false, err
	}
	c := commands.FeedbackCommand{
		FeedbackID: p.Args["feedbackID"].(string),
		ReviewerID: credentials.Email,
	}
	err = c.EndFeedback(p.Context)
	if err != nil {
		return false, err
	}
	return true, nil
}

func AnswerFeedbackResolver(p graphql.ResolveParams) (interface{}, error) {
	credentials, err := application_services.AuthorizationApplicationServiceInstance().Credentials(
		p.Context)
	if err != nil {
		return false, err
	}
	command := commands.FeedbackCommand{
		FeedbackID: p.Args["feedbackID"].(string),
		ReviewerID: credentials.Email,
		AnswerBody: p.Args["answerBody"].(string),
	}
	err = command.AnswerFeedback(p.Context)
	if err != nil {
		return false, err
	}
	return true, nil
}
