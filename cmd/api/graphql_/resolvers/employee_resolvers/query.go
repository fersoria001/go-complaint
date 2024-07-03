package employee_resolvers

import (
	"go-complaint/application/application_services"
	"go-complaint/application/queries"
	"go-complaint/dto"

	"github.com/graphql-go/graphql"
)

func EmployeeResolver(p graphql.ResolveParams) (interface{}, error) {
	_, err := application_services.AuthorizationApplicationServiceInstance().ResourceAccess(
		p.Context,
		"Enterprise",
		p.Args["enterpriseName"].(string),
		application_services.READ,
		"MANAGER", "ASSISTANT",
	)
	if err != nil {
		return dto.Employee{}, err
	}
	q := queries.EmployeeQuery{
		EmployeeID: p.Args["employeeID"].(string),
	}
	emp, err := q.Employee(p.Context)
	if err != nil {
		return dto.Employee{}, err
	}
	return emp, nil
}

func EmployeesResolver(p graphql.ResolveParams) (interface{}, error) {
	_, err := application_services.AuthorizationApplicationServiceInstance().ResourceAccess(
		p.Context,
		"Enterprise",
		p.Args["enterpriseName"].(string),
		application_services.READ,
		"MANAGER", "ASSISTANT",
	)
	if err != nil {
		return []dto.Employee{}, err
	}
	q := queries.EmployeeQuery{
		EnterpriseID: p.Args["enterpriseName"].(string),
	}
	emp, err := q.Employees(p.Context)
	if err != nil {
		return []dto.Employee{}, err
	}
	return emp, nil
}

func SolvedComplaintsResolver(params graphql.ResolveParams) (interface{}, error) {
	_, err := application_services.AuthorizationApplicationServiceInstance().Credentials(params.Context)
	if err != nil {
		return false, err
	}
	query := queries.EmployeeQuery{
		EmployeeID: params.Args["id"].(string),
	}
	complaints, err := query.SolvedComplaints(params.Context)
	if err != nil {
		return nil, err
	}
	return complaints, nil
}
