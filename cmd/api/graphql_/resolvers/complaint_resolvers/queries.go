package complaint_resolvers

import (
	"go-complaint/application/application_services"
	"go-complaint/application/queries"

	"github.com/graphql-go/graphql"
)

func ComplaintsReceivedInfoResolver(params graphql.ResolveParams) (interface{}, error) {
	_, err := application_services.AuthorizationApplicationServiceInstance().ResourceAccess(
		params.Context,
		"rid",
		params.Args["id"].(string),
		application_services.READ,
		"MANAGER", "ASSISTANT", "OWNER",
	)
	if err != nil {
		return false, err
	}
	query := queries.ComplaintQuery{
		ID: params.Args["id"].(string),
	}
	info, err := query.ComplaintsReceivedInfo(params.Context)
	if err != nil {
		return nil, err
	}
	return info, nil
}

func IsValidComplaintReceiverResolver(params graphql.ResolveParams) (interface{}, error) {
	_, err := application_services.AuthorizationApplicationServiceInstance().Credentials(params.Context)
	if err != nil {
		return false, err
	}
	query := queries.ComplaintQuery{
		ReceiverID: params.Args["id"].(string),
	}
	valid := query.IsValidComplaintReceiver(params.Context)
	return valid, nil

}
func FindComplaintReceiversResolver(params graphql.ResolveParams) (interface{}, error) {
	_, err := application_services.AuthorizationApplicationServiceInstance().Credentials(params.Context)
	if err != nil {
		return nil, err
	}
	term := ""
	if params.Args["term"] != nil {
		term = params.Args["term"].(string)
	}
	query := queries.ComplaintQuery{
		ID:   params.Args["id"].(string),
		Term: term,
	}
	receivers, err := query.FindReceivers(params.Context)
	if err != nil {
		return nil, err
	}
	return receivers, nil
}

func FindAuthorByIDResolver(params graphql.ResolveParams) (interface{}, error) {
	_, err := application_services.AuthorizationApplicationServiceInstance().Credentials(params.Context)
	if err != nil {
		return nil, err
	}
	query := queries.ComplaintQuery{
		AuthorID: params.Args["id"].(string),
	}
	author, err := query.FindAuthor(params.Context)
	if err != nil {
		return nil, err
	}
	return author, nil
}

func ComplaintInboxResolver(params graphql.ResolveParams) (interface{}, error) {
	_, err := application_services.AuthorizationApplicationServiceInstance().ResourceAccess(
		params.Context,
		"rid",
		params.Args["id"].(string),
		application_services.READ,
		"MANAGER", "ASSISTANT", "OWNER",
	)
	if err != nil {
		return false, err
	}
	query := queries.ComplaintQuery{
		ReceiverID: params.Args["id"].(string),
		Limit:      params.Args["limit"].(int),
		Offset:     params.Args["offset"].(int),
	}
	complaints, err := query.Inbox(params.Context)
	if err != nil {
		return nil, err
	}
	return complaints, nil
}
func ComplaintInboxSearchResolver(params graphql.ResolveParams) (interface{}, error) {
	_, err := application_services.AuthorizationApplicationServiceInstance().ResourceAccess(
		params.Context,
		"rid",
		params.Args["id"].(string),
		application_services.READ,
		"MANAGER", "ASSISTANT", "OWNER",
	)
	if err != nil {
		return false, err
	}
	query := queries.ComplaintQuery{
		ReceiverID: params.Args["id"].(string),
		Term:       params.Args["term"].(string),
		AfterDate:  params.Args["afterDate"].(string),
		BeforeDate: params.Args["beforeDate"].(string),
		Limit:      params.Args["limit"].(int),
		Offset:     params.Args["offset"].(int),
	}
	complaints, err := query.InboxSearch(params.Context)
	if err != nil {
		return nil, err
	}
	return complaints, nil
}
func ComplaintSentSearchResolver(params graphql.ResolveParams) (interface{}, error) {
	_, err := application_services.AuthorizationApplicationServiceInstance().ResourceAccess(
		params.Context,
		"rid",
		params.Args["id"].(string),
		application_services.READ,
		"MANAGER", "ASSISTANT", "OWNER",
	)
	if err != nil {
		return false, err
	}
	query := queries.ComplaintQuery{
		AuthorID:   params.Args["id"].(string),
		Term:       params.Args["term"].(string),
		AfterDate:  params.Args["afterDate"].(string),
		BeforeDate: params.Args["beforeDate"].(string),
		Limit:      params.Args["limit"].(int),
		Offset:     params.Args["offset"].(int),
	}
	complaints, err := query.SentSearch(params.Context)
	if err != nil {
		return nil, err
	}
	return complaints, nil
}

func ComplaintHistoryResolver(params graphql.ResolveParams) (interface{}, error) {
	_, err := application_services.AuthorizationApplicationServiceInstance().ResourceAccess(
		params.Context,
		"rid",
		params.Args["id"].(string),
		application_services.READ,
		"MANAGER", "ASSISTANT", "OWNER",
	)
	if err != nil {
		return false, err
	}
	query := queries.ComplaintQuery{
		ReceiverID: params.Args["id"].(string),
		Term:       params.Args["term"].(string),
		AfterDate:  params.Args["afterDate"].(string),
		BeforeDate: params.Args["beforeDate"].(string),
		Limit:      params.Args["limit"].(int),
		Offset:     params.Args["offset"].(int),
	}
	complaints, err := query.History(params.Context)
	if err != nil {
		return nil, err
	}
	return complaints, nil
}

func ComplaintsSentResolver(params graphql.ResolveParams) (interface{}, error) {
	_, err := application_services.AuthorizationApplicationServiceInstance().ResourceAccess(
		params.Context,
		"rid",
		params.Args["id"].(string),
		application_services.READ,
		"MANAGER", "ASSISTANT", "OWNER",
	)
	if err != nil {
		return false, err
	}

	query := queries.ComplaintQuery{
		AuthorID: params.Args["id"].(string),
		Limit:    params.Args["limit"].(int),
		Offset:   params.Args["offset"].(int),
	}
	complaints, err := query.Sent(params.Context)
	if err != nil {
		return nil, err
	}
	return complaints, nil
}

func ComplaintResolver(params graphql.ResolveParams) (interface{}, error) {
	_, err := application_services.AuthorizationApplicationServiceInstance().Credentials(params.Context)
	if err != nil {
		return false, err
	}
	query := queries.ComplaintQuery{
		ID: params.Args["id"].(string),
	}
	c, err := query.Complaint(params.Context)
	if err != nil {
		return nil, err
	}
	return c, nil
}
func PendingComplaintReviewsResolver(params graphql.ResolveParams) (interface{}, error) {
	_, err := application_services.AuthorizationApplicationServiceInstance().ResourceAccess(
		params.Context,
		"rid",
		params.Args["id"].(string),
		application_services.READ,
		"MANAGER", "ASSISTANT", "OWNER",
	)
	if err != nil {
		return false, err
	}
	query := queries.ComplaintQuery{
		UserID: params.Args["id"].(string),
	}
	complaints, err := query.PendingComplaintReviews(params.Context)
	if err != nil {
		return nil, err
	}
	return complaints, nil
}
