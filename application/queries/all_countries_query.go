package queries

import (
	"context"
	"go-complaint/domain/model/common"
	"go-complaint/dto"
	countryfindall "go-complaint/infrastructure/persistence/finders/country_findall"
	"go-complaint/infrastructure/persistence/repositories"
	"slices"
	"strings"
)

type AllCountriesQuery struct {
}

func NewAllCountriesQuery() *AllCountriesQuery {
	return &AllCountriesQuery{}
}

func (q AllCountriesQuery) Execute(ctx context.Context) ([]*dto.Country, error) {
	repository, ok := repositories.MapperRegistryInstance().Get("Country").(repositories.CountryRepository)
	if !ok {
		return nil, ErrWrongTypeAssertion
	}
	countries, err := repository.FindAll(ctx, countryfindall.NewCountries())
	if err != nil {
		return nil, err
	}
	s := countries.ToSlice()
	slices.SortStableFunc(s, func(a, b common.Country) int {
		return strings.Compare(a.Name(), b.Name())
	})
	var countryDTOs []*dto.Country
	for _, country := range s {
		countryDTOs = append(countryDTOs, dto.NewCountry(country))
	}
	return countryDTOs, nil
}
