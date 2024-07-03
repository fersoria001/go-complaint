package employee_resolvers

import (
	"go-complaint/application/application_services"
	"go-complaint/application/commands"

	"github.com/graphql-go/graphql"
)

func LeaveEnterpriseResolver(p graphql.ResolveParams) (interface{}, error) {
	_, err := application_services.AuthorizationApplicationServiceInstance().ResourceAccess(
		p.Context,
		"rid",
		p.Args["enterpriseName"].(string),
		application_services.WRITE,
		"MANAGER", "ASSISTANT",
	)
	if err != nil {
		return false, err
	}
	c := commands.EmployeeCommand{
		EmployeeID:   p.Args["employeeID"].(string),
		EnterpriseID: p.Args["enterpriseName"].(string),
	}
	err = c.LeaveEnterprise(p.Context)
	if err != nil {
		return false, err
	}
	return true, nil
}
