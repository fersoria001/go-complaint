package queries

import (
	"context"
	"go-complaint/domain/model/common"
	"go-complaint/dto"
	countrystatefindall "go-complaint/infrastructure/persistence/finders/country_state_findall"
	"go-complaint/infrastructure/persistence/repositories"
	"slices"
	"strings"
)

type CountryStatesByCountryIdQuery struct {
	Id int `json:"id"`
}

func NewCountryStatesByCountryIdQuery(id int) *CountryStatesByCountryIdQuery {
	return &CountryStatesByCountryIdQuery{
		Id: id,
	}
}

func (q CountryStatesByCountryIdQuery) Execute(ctx context.Context) ([]*dto.CountryState, error) {
	repository, ok := repositories.MapperRegistryInstance().Get("CountryState").(repositories.CountryStateRepository)
	if !ok {
		return nil, ErrWrongTypeAssertion
	}
	countryStates, err := repository.FindAll(ctx, countrystatefindall.NewByCountryID(q.Id))
	if err != nil {
		return nil, err
	}
	s := countryStates.ToSlice()
	slices.SortStableFunc(s, func(a, b common.CountryState) int {
		return strings.Compare(a.Name(), b.Name())
	})
	var countryStateDTOs []*dto.CountryState
	for _, countryState := range s {
		countryStateDTOs = append(countryStateDTOs, dto.NewCountryState(countryState))
	}
	return countryStateDTOs, nil
}
