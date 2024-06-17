package user_resolvers

import (
	"go-complaint/application/application_services"
	"go-complaint/application/queries"

	"github.com/graphql-go/graphql"
)

func SignInResolver(params graphql.ResolveParams) (interface{}, error) {
	userQuery := queries.UserQuery{
		Email:      params.Args["email"].(string),
		Password:   params.Args["password"].(string),
		RememberMe: params.Args["rememberMe"].(bool),
	}
	token, err := userQuery.SignIn(params.Context)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func LoginResolver(params graphql.ResolveParams) (interface{}, error) {
	token, err := application_services.AuthorizationApplicationServiceInstance().JWTToken(params.Context)
	if err != nil {
		return nil, err
	}
	userQuery := queries.UserQuery{
		Token:            token,
		ConfirmationCode: params.Args["confirmationCode"].(int),
	}
	newToken, err := userQuery.Login(params.Context)
	if err != nil {
		return nil, err
	}
	return newToken.Token, nil
}

func UserDescriptorResolver(params graphql.ResolveParams) (interface{}, error) {
	currentUser, err := application_services.AuthorizationApplicationServiceInstance().Credentials(params.Context)
	if err != nil {
		return nil, err
	}
	return currentUser, nil
}

func UserResolver(params graphql.ResolveParams) (interface{}, error) {
	userQuery := queries.UserQuery{
		Email: params.Args["ID"].(string),
	}
	user, err := userQuery.User(params.Context)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func HiringInvitationsResolver(params graphql.ResolveParams) (interface{}, error) {
	//authorize
	currentUser, err := application_services.AuthorizationApplicationServiceInstance().Credentials(params.Context)
	if err != nil {
		return nil, err
	}
	//query
	uq := queries.UserQuery{
		Email: currentUser.Email,
	}
	//return
	return uq.HiringInvitations(params.Context)
}

func UsersForHiringResolver(params graphql.ResolveParams) (interface{}, error) {
	//authorize
	_, err := application_services.AuthorizationApplicationServiceInstance().ResourceAccess(
		params.Context,
		"rid",
		params.Args["id"].(string),
		application_services.READ,
		"MANAGER", "OWNER",
	)
	if err != nil {
		return false, err
	}
	//query
	uq := queries.EnterpriseQuery{
		EnterpriseName: params.Args["id"].(string),
	}
	//return
	return uq.UsersForHiring(params.Context)
}
