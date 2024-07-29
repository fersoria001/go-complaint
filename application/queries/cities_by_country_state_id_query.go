package queries

import (
	"context"
	"go-complaint/domain/model/common"
	"go-complaint/dto"
	statecitiesfindall "go-complaint/infrastructure/persistence/finders/state_cities_findall"
	"go-complaint/infrastructure/persistence/repositories"
	"slices"
	"strings"
)

type CitiesByCountryStateIdQuery struct {
	Id int `json:"id"`
}

func NewCitiesByCountryStateIdQuery(id int) *CitiesByCountryStateIdQuery {
	return &CitiesByCountryStateIdQuery{
		Id: id,
	}
}

func (q CitiesByCountryStateIdQuery) Execute(ctx context.Context) ([]*dto.City, error) {
	repository, ok := repositories.MapperRegistryInstance().Get("City").(repositories.StateCitiesRepository)
	if !ok {
		return nil, ErrWrongTypeAssertion
	}
	cities, err := repository.FindAll(ctx, statecitiesfindall.NewByStateID(q.Id))
	if err != nil {
		return nil, err
	}
	s := cities.ToSlice()
	slices.SortStableFunc(s, func(a, b common.City) int {
		return strings.Compare(a.Name(), b.Name())
	})
	var cityDTOs []*dto.City
	for _, city := range s {
		cityDTOs = append(cityDTOs, dto.NewCity(city))
	}
	return cityDTOs, nil
}
