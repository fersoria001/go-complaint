package complaint_resolvers

import (
	"go-complaint/application/application_services"
	"go-complaint/application/commands"
	"strings"

	"github.com/graphql-go/graphql"
)

func SendComplaintResolver(params graphql.ResolveParams) (interface{}, error) {
	_, err := application_services.AuthorizationApplicationServiceInstance().Credentials(params.Context)
	if err != nil {
		return false, err
	}
	c := commands.ComplaintCommand{
		AuthorID:           params.Args["authorID"].(string),
		ReceiverID:         params.Args["receiverID"].(string),
		ReceiverFullName:   params.Args["receiverFullName"].(string),
		ReceiverProfileIMG: params.Args["receiverProfileIMG"].(string),
		Title:              params.Args["title"].(string),
		Description:        params.Args["description"].(string),
		Content:            params.Args["content"].(string),
	}
	err = c.SendNew(params.Context)
	if err != nil {
		return false, err
	}
	return true, nil
}

func ReplyComplaintResolver(params graphql.ResolveParams) (interface{}, error) {
	_, err := application_services.AuthorizationApplicationServiceInstance().ResourceAccess(
		params.Context,
		"Complaint",
		params.Args["complaintID"].(string),
		application_services.WRITE,
		"MANAGER", "ASSISTANT", "OWNER",
	)
	if err != nil {
		return false, err
	}
	c := commands.ComplaintCommand{
		ID:                params.Args["complaintID"].(string),
		ReplyAuthorID:     params.Args["replyAuthorID"].(string),
		ReplyBody:         params.Args["replyBody"].(string),
		ReplyEnterpriseID: params.Args["replyEnterpriseID"].(string),
	}
	err = c.Reply(params.Context)
	if err != nil {
		return false, err
	}
	return true, nil
}

func MarkAsSeenResolver(params graphql.ResolveParams) (interface{}, error) {
	_, err := application_services.AuthorizationApplicationServiceInstance().ResourceAccess(
		params.Context,
		"Complaint",
		params.Args["complaintID"].(string),
		application_services.WRITE,
		"MANAGER", "ASSISTANT", "OWNER",
	)
	if err != nil {
		return false, err
	}
	replies := params.Args["ids"].(string)
	ids := strings.Split(replies, ",")
	c := commands.ComplaintCommand{
		ID:  params.Args["complaintID"].(string),
		IDS: ids,
	}
	err = c.MarkAsSeen(params.Context)
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
		params.Args["complaintId"].(string),
		application_services.WRITE,
		"MANAGER", "ASSISTANT", "OWNER",
	)
	if err != nil {
		return false, err
	}
	c := commands.ComplaintCommand{
		ID:      params.Args["complaintId"].(string),
		EventID: params.Args["eventId"].(string),
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
