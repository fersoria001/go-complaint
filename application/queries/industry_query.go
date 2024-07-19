package queries

import (
	"context"
	"go-complaint/dto"
	industryfindall "go-complaint/infrastructure/persistence/finders/industry_findall"
	"go-complaint/infrastructure/persistence/repositories"
	"slices"
	"strings"
)

type IndustryQuery struct {
}

func (industryQuery IndustryQuery) AllIndustries(
	ctx context.Context,
) ([]dto.Industry, error) {
	mapper := repositories.MapperRegistryInstance().Get("Industry")
	industryRepository, ok := mapper.(repositories.IndustryRepository)
	if !ok {
		return nil, repositories.ErrWrongTypeAssertion
	}
	industries, err := industryRepository.FindAll(
		ctx,
		industryfindall.NewIndustries(),
	)
	if err != nil {
		return nil, err
	}
	var industryDTOs []dto.Industry
	for industry := range industries.Iter() {
		industryDTOs = append(industryDTOs, dto.NewIndustry(industry))
	}
	slices.SortStableFunc(industryDTOs, func(a, b dto.Industry) int {
		return strings.Compare(a.Name, b.Name)
	})
	return industryDTOs, nil
}
