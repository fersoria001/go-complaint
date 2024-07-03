package feedback_resolvers

import (
	"go-complaint/application/application_services"
	"go-complaint/application/queries"

	"github.com/graphql-go/graphql"
)

func FeedbackByComplaintIDResolver(params graphql.ResolveParams) (interface{}, error) {
	_, err := application_services.AuthorizationApplicationServiceInstance().ResourceAccess(
		params.Context,
		"Complaint",
		params.Args["id"].(string),
		application_services.WRITE,
		"MANAGER", "ASSISTANT", "OWNER",
	)
	if err != nil {
		return false, err
	}
	query := queries.FeedbackQuery{
		ComplaintID: params.Args["id"].(string),
	}
	c, err := query.FindByComplaintID(params.Context)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func FeedbackByIDResolver(params graphql.ResolveParams) (interface{}, error) {
	_, err := application_services.AuthorizationApplicationServiceInstance().ResourceAccess(
		params.Context,
		"Feedback",
		params.Args["id"].(string),
		application_services.WRITE,
		"MANAGER", "ASSISTANT", "OWNER",
	)
	if err != nil {

		return false, err
	}
	query := queries.FeedbackQuery{
		FeedbackID: params.Args["id"].(string),
	}
	c, err := query.Feedback(params.Context)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func FeedbackByRevieweeIDResolver(params graphql.ResolveParams) (interface{}, error) {
	_, err := application_services.AuthorizationApplicationServiceInstance().ResourceAccess(
		params.Context,
		"rid",
		params.Args["id"].(string),
		application_services.WRITE,
		"MANAGER", "ASSISTANT", "OWNER",
	)
	if err != nil {
		return false, err
	}
	query := queries.FeedbackQuery{
		RevieweeID: params.Args["id"].(string),
	}
	c, err := query.FindByRevieweeID(params.Context)
	if err != nil {
		return nil, err
	}
	return c, nil
}
