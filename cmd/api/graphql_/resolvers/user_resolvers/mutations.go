package user_resolvers

import (
	"go-complaint/application/application_services"
	"go-complaint/application/commands"

	"github.com/graphql-go/graphql"
)

func CreateUserResolver(p graphql.ResolveParams) (interface{}, error) {
	command := commands.UserCommand{
		Email:          p.Args["email"].(string),
		Password:       p.Args["password"].(string),
		FirstName:      p.Args["firstName"].(string),
		LastName:       p.Args["lastName"].(string),
		Gender:         p.Args["gender"].(string),
		Pronoun:        p.Args["pronoun"].(string),
		BirthDate:      p.Args["birthDate"].(string),
		PhoneCode:      p.Args["phoneCode"].(string),
		Phone:          p.Args["phone"].(string),
		CountryID:      p.Args["country"].(int),
		CountryStateID: p.Args["county"].(int),
		CityID:         p.Args["city"].(int),
	}
	err := command.Register(p.Context)
	if err != nil {
		return false, err
	}
	return true, nil
}

func VerifyEmailResolver(params graphql.ResolveParams) (interface{}, error) {
	userCommand := commands.UserCommand{
		EmailVerificationToken: params.Args["id"].(string),
	}
	err := userCommand.VerifyEmail(params.Context)
	if err != nil {
		return false, err
	}
	return true, nil
}

func RecoverPasswordResolver(params graphql.ResolveParams) (interface{}, error) {
	userCommand := commands.UserCommand{
		Email: params.Args["id"].(string),
	}
	err := userCommand.RecoverPassword(params.Context)
	if err != nil {
		return false, err
	}
	return true, nil
}

func ChangePasswordResolver(params graphql.ResolveParams) (interface{}, error) {
	currentUser, err := application_services.AuthorizationApplicationServiceInstance().Credentials(params.Context)
	if err != nil {
		return false, err
	}
	userCommand := commands.UserCommand{
		Email:       currentUser.Email,
		OldPassword: params.Args["oldPassword"].(string),
		Password:    params.Args["newPassword"].(string),
	}
	err = userCommand.ChangePassword(params.Context)
	if err != nil {
		return false, err
	}
	return true, nil
}

func UpdateUserResolver(params graphql.ResolveParams) (interface{}, error) {
	currentUser, err := application_services.AuthorizationApplicationServiceInstance().Credentials(params.Context)
	if err != nil {
		return false, err
	}
	updateType := params.Args["updateType"].(string)
	command := commands.UserCommand{
		UpdateType: updateType,
		Email:      currentUser.Email,
	}
	switch updateType {
	case "pronoun":
		command.Pronoun = params.Args["value"].(string)
	case "gender":
		command.Gender = params.Args["value"].(string)
	case "profileIMG":
		command.ProfileIMG = params.Args["value"].(string)
	case "firstName":
		command.FirstName = params.Args["value"].(string)
	case "lastName":
		command.LastName = params.Args["value"].(string)
	case "phone":
		command.Phone = params.Args["value"].(string)
	case "country":
		command.CountryID = params.Args["numberValue"].(int)
	case "countryState":
		command.CountryStateID = params.Args["numberValue"].(int)
	case "city":
		command.CityID = params.Args["numberValue"].(int)
	default:
		return false, command.UpdatePersonalData(params.Context)
	}
	err = command.UpdatePersonalData(params.Context)
	if err != nil {
		return false, err
	}
	return true, nil
}

func AcceptEnterpriseInvitationResolver(params graphql.ResolveParams) (interface{}, error) {
	userID, err := application_services.AuthorizationApplicationServiceInstance().Credentials(params.Context)
	if err != nil {
		return false, err
	}
	uc := commands.UserCommand{
		Email:   userID.Email,
		EventID: params.Args["id"].(string),
	}
	err = uc.AcceptHiringInvitation(
		params.Context,
	)
	if err != nil {
		return false, err
	}

	return true, nil
}

func RejectEnterpriseInvitationResolver(params graphql.ResolveParams) (interface{}, error) {
	userID, err := application_services.AuthorizationApplicationServiceInstance().Credentials(params.Context)
	if err != nil {
		return false, err
	}
	reason, ok := params.Args["reason"].(string)
	if !ok {
		reason = ""
	}
	uc := commands.UserCommand{
		Email:          userID.Email,
		RejectedReason: reason,
		EventID:        params.Args["id"].(string),
	}
	err = uc.RejectHiringInvitation(
		params.Context,
	)
	if err != nil {
		return false, err
	}

	return true, nil
}

func MarkNotificationAsReadResolver(params graphql.ResolveParams) (interface{}, error) {
	currentUser, err := application_services.AuthorizationApplicationServiceInstance().Credentials(params.Context)
	if err != nil {
		return false, err
	}
	notificationCommand := commands.NotificationCommand{
		ID:      params.Args["id"].(string),
		OwnerID: currentUser.Email,
	}
	err = notificationCommand.MarkAsRead(params.Context)
	if err != nil {
		return false, err
	}
	return true, nil
}
