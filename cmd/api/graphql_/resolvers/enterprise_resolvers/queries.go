package enterprise_resolvers

import (
	"go-complaint/application/application_services"
	"go-complaint/application/queries"

	"github.com/graphql-go/graphql"
)

func EnterpriseResolver(params graphql.ResolveParams) (interface{}, error) {
	enterpriseQuery := queries.EnterpriseQuery{
		EnterpriseName: params.Args["id"].(string),
	}
	enterprise, err := enterpriseQuery.Enterprise(params.Context)
	if err != nil {
		return nil, err
	}
	return enterprise, nil
}

func IsEnterpriseNameAvailableResolver(params graphql.ResolveParams) (interface{}, error) {
	enterpriseQuery := queries.EnterpriseQuery{
		EnterpriseName: params.Args["id"].(string),
	}
	return enterpriseQuery.IsEnterpriseNameAvailable(params.Context)
}

func HiringProccesesResolver(params graphql.ResolveParams) (interface{}, error) {
	_, err := application_services.AuthorizationApplicationServiceInstance().ResourceAccess(
		params.Context,
		"Enterprise",
		params.Args["id"].(string),
		application_services.READ,
		"MANAGER", "OWNER",
	)
	if err != nil {
		return nil, err
	}
	enterpriseQuery := queries.EnterpriseQuery{
		EnterpriseName: params.Args["id"].(string),
		Term:           params.Args["query"].(string),
		Offset:         params.Args["offset"].(int),
		Limit:          params.Args["limit"].(int),
	}
	pendingHires, err := enterpriseQuery.HiringProcceses(params.Context)
	if err != nil {
		return nil, err
	}
	return pendingHires, nil
}

func OnlineUsersResolver(params graphql.ResolveParams) (interface{}, error) {

	credentials, err := application_services.AuthorizationApplicationServiceInstance().ResourceAccess(
		params.Context,
		"Enterprise",
		params.Args["id"].(string),
		application_services.READ,
		"ASSISTANT", "MANAGER", "OWNER",
	)
	if err != nil {
		return nil, err
	}
	enterpriseQuery := queries.EnterpriseQuery{
		EnterpriseName: params.Args["id"].(string),
		UserID:         credentials.Email,
	}
	onlineUsers, err := enterpriseQuery.OnlineUsers(params.Context)
	if err != nil {
		return nil, err
	}

	return onlineUsers, nil
}

func EnterpriseChatResolver(params graphql.ResolveParams) (interface{}, error) {
	_, err := application_services.AuthorizationApplicationServiceInstance().ResourceAccess(
		params.Context,
		"Enterprise",
		params.Args["enterpriseID"].(string),
		application_services.READ,
		"ASSISTANT", "MANAGER", "OWNER",
	)
	if err != nil {
		return nil, err
	}
	enterpriseQuery := queries.EnterpriseQuery{
		EnterpriseName: params.Args["enterpriseID"].(string),
		ChatID:         params.Args["chatID"].(string),
	}
	chat, err := enterpriseQuery.EnterpriseChat(params.Context)
	if err != nil {
		return nil, err
	}
	return chat, nil
}
