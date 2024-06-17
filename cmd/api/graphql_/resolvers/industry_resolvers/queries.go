package industry_resolvers

import (
	"go-complaint/application/queries"

	"github.com/graphql-go/graphql"
)

func IndustriesResolver(params graphql.ResolveParams) (interface{}, error) {
	industryQuery := queries.IndustryQuery{}
	return industryQuery.AllIndustries(params.Context)
}
