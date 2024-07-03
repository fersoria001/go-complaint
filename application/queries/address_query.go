package queries

import (
	"context"
	"go-complaint/dto"
	countryfindall "go-complaint/infrastructure/persistence/finders/country_findall"
	countrystatefindall "go-complaint/infrastructure/persistence/finders/country_state_findall"
	statecitiesfindall "go-complaint/infrastructure/persistence/finders/state_cities_findall"
	"go-complaint/infrastructure/persistence/repositories"
)

type AddressQuery struct {
	CountryID      int
	CountryStateID int
}

func (addressQuery AddressQuery) AllCountries(
	ctx context.Context,
) ([]dto.Country, error) {
	mapper := repositories.MapperRegistryInstance().Get("Country")
	countryRepository, ok := mapper.(repositories.CountryRepository)
	if !ok {
		return nil, repositories.ErrWrongTypeAssertion
	}
	countries, err := countryRepository.FindAll(
		ctx,
		countryfindall.NewCountries(),
	)
	if err != nil {
		return nil, err
	}
	var countryDTOs []dto.Country
	for country := range countries.Iter() {
		countryDTOs = append(countryDTOs, dto.NewCountry(country))
	}
	return countryDTOs, nil
}

func (addressQuery AddressQuery) ProvideCountryStateByCountryID(
	ctx context.Context,
) ([]dto.CountryState, error) {
	mapper := repositories.MapperRegistryInstance().Get("CountryState")
	countryStateRepository, ok := mapper.(repositories.CountryStateRepository)
	if !ok {
		return nil, repositories.ErrWrongTypeAssertion
	}
	countryStates, err := countryStateRepository.FindAll(
		ctx,
		countrystatefindall.NewByCountryID(addressQuery.CountryID),
	)
	if err != nil {
		return nil, err
	}
	var countryStateDTOs []dto.CountryState
	for countryState := range countryStates.Iter() {
		countryStateDTOs = append(countryStateDTOs, dto.NewCountryState(countryState))
	}
	return countryStateDTOs, nil
}

func (addressQuery AddressQuery) ProvideStateCitiesByStateID(
	ctx context.Context,
) ([]dto.City, error) {
	mapper := repositories.MapperRegistryInstance().Get("City")
	stateCitiesRepository, ok := mapper.(repositories.StateCitiesRepository)
	if !ok {
		return nil, repositories.ErrWrongTypeAssertion
	}
	stateCities, err := stateCitiesRepository.FindAll(
		ctx,
		statecitiesfindall.NewByStateID(addressQuery.CountryStateID),
	)
	if err != nil {
		return nil, err
	}
	var cityDTOs []dto.City
	for city := range stateCities.Iter() {
		cityDTOs = append(cityDTOs, dto.NewCity(city))
	}
	return cityDTOs, nil
}
