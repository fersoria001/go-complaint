package enterprise_resolvers

import (
	"fmt"
	"go-complaint/application/application_services"
	"go-complaint/application/queries"

	"github.com/graphql-go/graphql"
)

func EnterpriseResolver(params graphql.ResolveParams) (interface{}, error) {
	enterpriseQuery := queries.EnterpriseQuery{
		EnterpriseName: params.Args["ID"].(string),
	}
	enterprise, err := enterpriseQuery.Enterprise(params.Context)
	if err != nil {
		return nil, err
	}
	return enterprise, nil
}

func IsEnterpriseNameAvailableResolver(params graphql.ResolveParams) (interface{}, error) {
	enterpriseQuery := queries.EnterpriseQuery{
		EnterpriseName: params.Args["name"].(string),
	}
	return enterpriseQuery.IsEnterpriseNameAvailable(params.Context)
}

func HiringInvitationsAcceptedResolver(params graphql.ResolveParams) (interface{}, error) {
	currentUser, err := application_services.AuthorizationApplicationServiceInstance().Credentials(params.Context)
	if err != nil {
		return nil, err
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
	enterpriseQuery := queries.EnterpriseQuery{
		EnterpriseName: params.Args["id"].(string),
	}
	pendingHires, err := enterpriseQuery.HiringInvitationsAccepted(params.Context)
	if err != nil {
		return nil, err
	}
	return pendingHires, nil
}
