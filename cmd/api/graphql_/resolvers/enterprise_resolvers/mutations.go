package enterprise_resolvers

import (
	"go-complaint/application/application_services"
	"go-complaint/application/commands"
	"strings"

	"github.com/graphql-go/graphql"
)

func CreateEnterpriseResolver(params graphql.ResolveParams) (interface{}, error) {
	currentUser, err := application_services.AuthorizationApplicationServiceInstance().Credentials(params.Context)
	if err != nil {
		return false, err
	}
	command := commands.EnterpriseCommand{
		OwnerID: currentUser.Email,
		Name:    params.Args["name"].(string),
		Website: params.Args["website"].(string),
		Email:   params.Args["email"].(string),
		//		PhoneCode:      params.Args["phoneCode"].(string),
		Phone:          params.Args["phone"].(string),
		CountryID:      params.Args["countryID"].(int),
		CountryStateID: params.Args["countryStateID"].(int),
		CityID:         params.Args["cityID"].(int),
		IndustryID:     params.Args["industryID"].(int),
		FoundationDate: params.Args["foundationDate"].(string),
	}
	err = command.Register(params.Context)
	if err != nil {
		return false, err
	}
	return true, nil
}

func UpdateEnterpriseResolver(params graphql.ResolveParams) (interface{}, error) {
	currentUser, err := application_services.AuthorizationApplicationServiceInstance().ResourceAccess(
		params.Context,
		"Enterprise",
		params.Args["enterpriseID"].(string),
		application_services.WRITE,
		"OWNER",
	)
	if err != nil {
		return nil, err
	}
	command := commands.EnterpriseCommand{
		OwnerID:    currentUser.Email,
		Name:       params.Args["enterpriseID"].(string),
		UpdateType: params.Args["updateType"].(string),
	}
	switch command.UpdateType {
	case "logoIMG":
		command.LogoIMG = params.Args["value"].(string)
	case "bannerIMG":
		command.BannerIMG = params.Args["value"].(string)
	case "website":
		command.Website = params.Args["value"].(string)
	case "email":
		command.Email = params.Args["value"].(string)
	case "phone":
		command.Phone = params.Args["value"].(string)
	case "country":
		command.CountryID = params.Args["numberValue"].(int)
	case "countryState":
		command.CountryStateID = params.Args["numberValue"].(int)
	case "city":
		command.CityID = params.Args["numberValue"].(int)
	default:
		return false, command.UpdateEnterprise(params.Context)
	}
	err = command.UpdateEnterprise(params.Context)
	if err != nil {
		return false, err
	}
	return true, nil
}

func InviteToEnterpriseResolver(params graphql.ResolveParams) (interface{}, error) {
	currentUser, err := application_services.AuthorizationApplicationServiceInstance().ResourceAccess(
		params.Context,
		"Enterprise",
		params.Args["enterpriseName"].(string),
		application_services.WRITE,
		"OWNER", "MANAGER",
	)
	if err != nil {
		return false, err
	}
	command := commands.EnterpriseCommand{
		OwnerID:   currentUser.Email,
		Name:      params.Args["enterpriseName"].(string),
		ProposeTo: params.Args["proposeTo"].(string),
		Position:  params.Args["proposedPosition"].(string),
	}

	err = command.InviteToProject(params.Context)
	if err != nil {
		return false, err
	}
	return true, nil

}

func HireEmployeeResolver(params graphql.ResolveParams) (interface{}, error) {
	currentUser, err := application_services.AuthorizationApplicationServiceInstance().ResourceAccess(
		params.Context,
		"Enterprise",
		params.Args["enterpriseName"].(string),
		application_services.WRITE,
		"OWNER", "MANAGER",
	)
	if err != nil {
		return false, err
	}
	command := commands.EnterpriseCommand{
		Name:    params.Args["enterpriseName"].(string),
		EventID: params.Args["eventID"].(string),
		OwnerID: currentUser.Email,
	}
	err = command.HireEmployee(params.Context)
	if err != nil {
		return false, err
	}
	return true, nil
}

func CancelHiringProccessResolver(params graphql.ResolveParams) (interface{}, error) {
	currentUser, err := application_services.AuthorizationApplicationServiceInstance().ResourceAccess(
		params.Context,
		"Enterprise",
		params.Args["enterpriseName"].(string),
		application_services.WRITE,
		"OWNER", "MANAGER",
	)
	if err != nil {
		return false, err
	}
	reason := ""
	if params.Args["reason"] != nil {
		reason = params.Args["reason"].(string)
	}

	command := commands.EnterpriseCommand{
		OwnerID:      currentUser.Email,
		Name:         params.Args["enterpriseName"].(string),
		EventID:      params.Args["eventID"].(string),
		CancelReason: reason,
	}
	err = command.CancelHiringProccess(params.Context)
	if err != nil {
		return false, err
	}
	return true, nil
}

func FireEmployeeResolver(params graphql.ResolveParams) (interface{}, error) {
	currentUser, err := application_services.AuthorizationApplicationServiceInstance().ResourceAccess(
		params.Context,
		"Enterprise",
		params.Args["enterpriseName"].(string),
		application_services.WRITE,
		"OWNER", "MANAGER",
	)
	if err != nil {
		return false, err
	}
	command := commands.EnterpriseCommand{
		OwnerID:    currentUser.Email,
		Name:       params.Args["enterpriseName"].(string),
		EmployeeID: params.Args["employeeID"].(string),
	}
	err = command.FireEmployee(params.Context)
	if err != nil {
		return false, err
	}
	return true, nil
}

func PromoteEmployeeResolver(params graphql.ResolveParams) (interface{}, error) {
	currentUser, err := application_services.AuthorizationApplicationServiceInstance().ResourceAccess(
		params.Context,
		"Enterprise",
		params.Args["enterpriseName"].(string),
		application_services.WRITE,
		"OWNER", "MANAGER",
	)

	if err != nil {
		return false, err
	}
	command := commands.EnterpriseCommand{
		OwnerID:    currentUser.Email,
		Name:       params.Args["enterpriseName"].(string),
		EmployeeID: params.Args["employeeID"].(string),
		Position:   params.Args["position"].(string),
	}
	err = command.PromoteEmployee(params.Context)
	if err != nil {
		return false, err
	}
	return true, nil
}

func ChatReplyResolver(params graphql.ResolveParams) (interface{}, error) {
	_, err := application_services.AuthorizationApplicationServiceInstance().ResourceAccess(
		params.Context,
		"Enterprise",
		params.Args["enterpriseName"].(string),
		application_services.WRITE,
		"ASSISTANT", "OWNER", "MANAGER",
	)
	if err != nil {
		return false, err
	}
	command := commands.EnterpriseCommand{
		ID:       params.Args["id"].(string),
		SenderID: params.Args["senderID"].(string),
		Content:  params.Args["content"].(string),
	}
	err = command.ReplyChat(params.Context)
	if err != nil {
		return false, err
	}
	return true, nil
}

func MarkChatReplyAsSeenResolver(params graphql.ResolveParams) (interface{}, error) {
	_, err := application_services.AuthorizationApplicationServiceInstance().ResourceAccess(
		params.Context,
		"Enterprise",
		params.Args["enterpriseName"].(string),
		application_services.WRITE,
		"ASSISTANT", "OWNER", "MANAGER",
	)
	if err != nil {
		return false, err
	}
	ids := params.Args["repliesID"].(string)
	s := strings.Split(ids, ",")

	command := commands.EnterpriseCommand{
		ID:        params.Args["chatID"].(string),
		RepliesID: s,
	}
	err = command.MarkAsSeen(params.Context)
	if err != nil {
		return false, err
	}
	return true, nil
}
