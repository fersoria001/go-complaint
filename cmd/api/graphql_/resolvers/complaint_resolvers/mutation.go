package complaint_resolvers

import (
	"fmt"
	"go-complaint/application/application_services"
	"go-complaint/application/commands"

	"github.com/graphql-go/graphql"
)

func SendComplaintResolver(params graphql.ResolveParams) (interface{}, error) {
	_, err := application_services.AuthorizationApplicationServiceInstance().Credentials(params.Context)
	if err != nil {
		return false, err
	}
	c := commands.ComplaintCommand{
		AuthorID:    params.Args["authorID"].(string),
		ReceiverID:  params.Args["receiverID"].(string),
		Title:       params.Args["title"].(string),
		Description: params.Args["description"].(string),
		Content:     params.Args["content"].(string),
	}
	err = c.SendNew(params.Context)
	if err != nil {
		return false, err
	}
	return true, nil
}

func ReplyComplaintResolver(params graphql.ResolveParams) (interface{}, error) {
	currentUser, err := application_services.AuthorizationApplicationServiceInstance().Credentials(params.Context)
	if err != nil {
		return false, err
	}
	enterpriseID, ok := params.Args["replyEnterpriseID"].(string)
	if !ok {
		enterpriseID = ""
	} else {
		hasPermissions := false
		if currentUser.Email != enterpriseID {
			for _, v := range currentUser.GrantedAuthorities {
				if v.EnterpriseID == enterpriseID {
					hasPermissions = true
					break
				}
			}
		}
		if !hasPermissions {
			return nil, fmt.Errorf("you don't have permissions to access this resource")
		}
	}
	c := commands.ComplaintCommand{
		ID:                params.Args["complaintID"].(string),
		ReplyAuthorID:     params.Args["replyAuthorID"].(string),
		ReplyBody:         params.Args["replyBody"].(string),
		ReplyEnterpriseID: enterpriseID,
	}
	err = c.Reply(params.Context)
	if err != nil {
		return false, err
	}
	return true, nil
}

func SendForReviewingResolver(params graphql.ResolveParams) (interface{}, error) {
	currentUser, err := application_services.AuthorizationApplicationServiceInstance().ResourceAccess(
		params.Context,
		"Complaint",
		params.Args["id"].(string),
		application_services.WRITE,
		"MANAGER", "ASSISTANT", "OWNER",
	)
	if err != nil {
		return false, err
	}
	c := commands.ComplaintCommand{
		ID:     params.Args["id"].(string),
		UserID: currentUser.Email,
	}
	err = c.SendForReviewing(params.Context)
	if err != nil {
		return false, err
	}
	return true, nil
}

func RateComplaintResolver(params graphql.ResolveParams) (interface{}, error) {
	currentUser, err := application_services.AuthorizationApplicationServiceInstance().ResourceAccess(
		params.Context,
		"Complaint",
		params.Args["id"].(string),
		application_services.READ,
		"MANAGER", "ASSISTANT", "OWNER",
	)
	if err != nil {
		return false, err
	}
	c := commands.ComplaintCommand{
		ID:      params.Args["id"].(string),
		UserID:  currentUser.Email,
		Rating:  params.Args["rate"].(int),
		Comment: params.Args["comment"].(string),
	}
	err = c.RateComplaint(params.Context)
	if err != nil {
		return false, err
	}
	return true, nil
}
