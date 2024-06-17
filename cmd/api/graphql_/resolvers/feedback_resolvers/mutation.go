package feedback_resolvers

import (
	"go-complaint/application/application_services"
	"go-complaint/application/commands"

	"github.com/graphql-go/graphql"
)

func AddReplyReviewResolver(p graphql.ResolveParams) (interface{}, error) {
	_, err := application_services.AuthorizationApplicationServiceInstance().ResourceAccess(
		p.Context,
		"Complaint",
		p.Args["id"].(string),
		application_services.WRITE,
		"MANAGER", "ASSISTANT", "OWNER",
	)
	if err != nil {
		return false, err
	}
	command := commands.FeedbackCommand{
		ComplaintID:      p.Args["id"].(string),
		ReviewedID:       p.Args["reviewedID"].(string),
		ReviewerID:       p.Args["reviewerID"].(string),
		ReviewComment:    p.Args["reviewComment"].(string),
		ReplyReviewColor: p.Args["replyReviewColor"].(string),
		Replies:          p.Args["repliesID"].([]string),
	}
	err = command.AddReplyReview(p.Context)
	if err != nil {
		return false, err
	}
	return true, nil
}

func EndFeedbackResolver(p graphql.ResolveParams) (interface{}, error) {
	credentials, err := application_services.AuthorizationApplicationServiceInstance().ResourceAccess(
		p.Context,
		"rid",
		p.Args["enterpriseID"].(string),
		application_services.WRITE,
		"MANAGER", "ASSISTANT",
	)
	if err != nil {
		return false, err
	}
	c := commands.FeedbackCommand{
		ComplaintID: p.Args["complaintID"].(string),
		ReviewerID:  credentials.Email,
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
