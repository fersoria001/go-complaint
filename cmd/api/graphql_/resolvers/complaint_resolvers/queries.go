package complaint_resolvers

import (
	"fmt"
	"go-complaint/application/application_services"
	"go-complaint/application/queries"

	"github.com/graphql-go/graphql"
)

func FindComplaintReceiversResolver(params graphql.ResolveParams) (interface{}, error) {
	_, err := application_services.AuthorizationApplicationServiceInstance().Credentials(params.Context)
	if err != nil {
		return nil, err
	}
	query := queries.ComplaintQuery{
		Term: params.Args["term"].(string),
	}
	receivers, err := query.FindReceivers(params.Context)
	if err != nil {
		return nil, err
	}
	return receivers, nil
}

func ComplaintInboxResolver(params graphql.ResolveParams) (interface{}, error) {
	currentUser, err := application_services.AuthorizationApplicationServiceInstance().Credentials(params.Context)
	if err != nil {
		return false, err
	}
	requestID := params.Args["id"].(string)
	hasPermissions := false
	if currentUser.Email != requestID {
		for _, v := range currentUser.GrantedAuthorities {
			if v.EnterpriseID == requestID {
				hasPermissions = true
				break
			}
		}
	}
	if !hasPermissions {
		return nil, fmt.Errorf("you don't have permissions to access this resource")
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
func ComplaintHistoryResolver(params graphql.ResolveParams) (interface{}, error) {
	currentUser, err := application_services.AuthorizationApplicationServiceInstance().Credentials(params.Context)
	if err != nil {
		return false, err
	}
	requestID := params.Args["id"].(string)
	hasPermissions := false
	if currentUser.Email != requestID {
		for _, v := range currentUser.GrantedAuthorities {
			if v.EnterpriseID == requestID {
				hasPermissions = true
				break
			}
		}
	}
	if !hasPermissions {
		return nil, fmt.Errorf("you don't have permissions to access this resource")
	}
	query := queries.ComplaintQuery{
		UserID: params.Args["id"].(string),
		Status: params.Args["status"].(string),
		Limit:  params.Args["limit"].(int),
		Offset: params.Args["offset"].(int),
	}
	complaints, err := query.History(params.Context)
	if err != nil {
		return nil, err
	}
	return complaints, nil
}

func ComplaintsSentResolver(params graphql.ResolveParams) (interface{}, error) {
	currentUser, err := application_services.AuthorizationApplicationServiceInstance().Credentials(params.Context)
	if err != nil {
		return false, err
	}
	requestID := params.Args["id"].(string)
	hasPermissions := false
	if currentUser.Email != requestID {
		for _, v := range currentUser.GrantedAuthorities {
			if v.EnterpriseID == requestID {
				hasPermissions = true
				break
			}
		}
	}
	if !hasPermissions {
		return nil, fmt.Errorf("you don't have permissions to access this resource")
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
	currentUser, err := application_services.AuthorizationApplicationServiceInstance().ResourceAccess(
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
		UserID: currentUser.Email,
	}
	complaints, err := query.PendingComplaintReviews(params.Context)
	if err != nil {
		return nil, err
	}
	return complaints, nil
}
