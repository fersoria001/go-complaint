package queries_test

import (
	"context"
	"go-complaint/application/queries"
	"go-complaint/domain/model/enterprise"
	"go-complaint/infrastructure/persistence/repositories"
	"go-complaint/test/mock_data"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnterpriseByNameQuery_Execute(t *testing.T) {
	ctx := context.Background()
	repository, ok := repositories.MapperRegistryInstance().Get("Enterprise").(repositories.EnterpriseRepository)
	assert.True(t, ok)
	for _, v := range mock_data.NewEnterprises {
		industry, err := enterprise.NewIndustry(v.Industry.Id, v.Industry.Name)
		assert.Nil(t, err)
		e, err := enterprise.NewEnterprise(
			v.Id,
			v.OwnerId,
			v.Name,
			v.LogoImg,
			v.BannerImg,
			v.Website,
			v.Email,
			v.Phone,
			v.Address,
			industry,
			v.RegisterAt,
			v.UpdatedAt,
			v.FoundationDate,
			v.Employees,
		)
		assert.Nil(t, err)
		err = repository.Save(ctx, e)
		assert.Nil(t, err)
		q := queries.NewEnterpriseByNameQuery(e.Name())
		dbE, err := q.Execute(ctx)
		assert.Nil(t, err)
		assert.NotNil(t, dbE)
		assert.Equal(t, v.Id.String(), dbE.Id)
		assert.Equal(t, v.OwnerId.String(), dbE.OwnerID)
		assert.Equal(t, v.Name, dbE.Name)
		assert.Equal(t, v.LogoImg, dbE.LogoIMG)
		assert.Equal(t, v.BannerImg, dbE.BannerIMG)
		assert.Equal(t, v.Website, dbE.Website)
		assert.Equal(t, v.Email, dbE.Email)
		assert.Equal(t, v.Phone, dbE.Phone)
		assert.Equal(t, v.Address.Country().Name(), dbE.Address.Country)
		assert.Equal(t, v.Address.CountryState().Name(), dbE.Address.County)
		assert.Equal(t, v.Address.City().Name(), dbE.Address.City)
		assert.Equal(t, v.Industry.Name, dbE.Industry)
		assert.Equal(t, v.Employees.Cardinality(), len(dbE.Employees))
		t.Cleanup(func() {
			err = repository.Remove(ctx, e.Id())
			assert.Nil(t, err)
		})
	}
}
