package address_resolvers

import (
	"go-complaint/application/queries"

	"github.com/graphql-go/graphql"
)

func CountriesResolver(params graphql.ResolveParams) (interface{}, error) {
	addressQuery := queries.AddressQuery{}
	return addressQuery.AllCountries(params.Context)
}

func CountryStatesResolver(params graphql.ResolveParams) (interface{}, error) {
	addressQuery := queries.AddressQuery{
		CountryID: params.Args["ID"].(int),
	}
	return addressQuery.ProvideCountryStateByCountryID(params.Context)
}

func CitiesResolver(params graphql.ResolveParams) (interface{}, error) {
	addressQuery := queries.AddressQuery{
		CountryStateID: params.Args["ID"].(int),
	}
	return addressQuery.ProvideStateCitiesByStateID(params.Context)
}
