package enterprise_resolvers

import (
	"go-complaint/application/application_services"
	"go-complaint/application/commands"

	"github.com/graphql-go/graphql"
)

func CreateEnterpriseResolver(params graphql.ResolveParams) (interface{}, error) {
	currentUser, err := application_services.AuthorizationApplicationServiceInstance().Credentials(params.Context)
	if err != nil {
		return false, err
	}
	command := commands.EnterpriseCommand{
		OwnerID:        currentUser.Email,
		Name:           params.Args["name"].(string),
		Website:        params.Args["website"].(string),
		Email:          params.Args["email"].(string),
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
	currentUser, err := application_services.AuthorizationApplicationServiceInstance().Credentials(params.Context)
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
		command.LogoIMG = params.Args["logoIMG"].(string)
	case "bannerIMG":
		command.BannerIMG = params.Args["bannerIMG"].(string)
	case "website":
		command.Website = params.Args["website"].(string)
	case "email":
		command.Email = params.Args["email"].(string)
	case "phone":
		command.Phone = params.Args["phone"].(string)
	case "country":
		command.CountryID = params.Args["countryID"].(int)
	case "countryState":
		command.CountryStateID = params.Args["countryStateID"].(int)
	case "city":
		command.CityID = params.Args["cityID"].(int)
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
	currentUser, err := application_services.AuthorizationApplicationServiceInstance().Credentials(params.Context)
	if err != nil {
		return nil, err
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
	currentUser, err := application_services.AuthorizationApplicationServiceInstance().Credentials(params.Context)
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
	currentUser, err := application_services.AuthorizationApplicationServiceInstance().Credentials(params.Context)
	if err != nil {
		return false, err
	}
	command := commands.EnterpriseCommand{
		OwnerID: currentUser.Email,
		Name:    params.Args["enterpriseName"].(string),
		EventID: params.Args["eventID"].(string),
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
		"OWNER",
	)
	if err != nil {
		return false, err
	}
	command := commands.EnterpriseCommand{
		OwnerID: currentUser.Email,
		Name:    params.Args["enterpriseName"].(string),
		EventID: params.Args["employeeID"].(string),
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
		OwnerID:  currentUser.Email,
		Name:     params.Args["enterpriseName"].(string),
		EventID:  params.Args["employeeID"].(string),
		Position: params.Args["position"].(string),
	}
	err = command.PromoteEmployee(params.Context)
	if err != nil {
		return false, err
	}
	return true, nil
}
