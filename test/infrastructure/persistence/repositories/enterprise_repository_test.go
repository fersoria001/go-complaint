package repositories_test

import (
	"context"
	"go-complaint/domain/model/enterprise"
	"go-complaint/infrastructure/persistence/finders/find_enterprise"
	"go-complaint/infrastructure/persistence/repositories"
	"go-complaint/test/mock_data"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnterpriseRepository_Save(t *testing.T) {
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
		t.Cleanup(func() {
			err = repository.Remove(ctx, e.Id())
			assert.Nil(t, err)
		})
	}
}

func TestEnterpriseRepository_Get(t *testing.T) {
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
		dbE, err := repository.Get(ctx, e.Id())
		assert.Nil(t, err)
		assert.NotNil(t, dbE)
		assert.Equal(t, v.Id, dbE.Id())
		assert.Equal(t, v.OwnerId, dbE.OwnerId())
		assert.Equal(t, v.Name, dbE.Name())
		assert.Equal(t, v.LogoImg, dbE.LogoIMG())
		assert.Equal(t, v.BannerImg, dbE.BannerIMG())
		assert.Equal(t, v.Website, dbE.Website())
		assert.Equal(t, v.Email, dbE.Email())
		assert.Equal(t, v.Phone, dbE.Phone())
		assert.Equal(t, v.Address, dbE.Address())
		assert.Equal(t, industry, dbE.Industry())
		assert.Equal(t, v.UpdatedAt.StringRepresentation(), dbE.UpdatedAt().StringRepresentation())
		assert.Equal(t, v.FoundationDate.StringRepresentation(), dbE.FoundationDate().StringRepresentation())
		assert.Equal(t, v.Employees.Cardinality(), dbE.Employees().Cardinality())
		t.Cleanup(func() {
			err = repository.Remove(ctx, e.Id())
			assert.Nil(t, err)
		})
	}
}

func TestEnterpriseRepository_Find_ByName(t *testing.T) {
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
		dbE, err := repository.Find(ctx, find_enterprise.ByName(e.Name()))
		assert.Nil(t, err)
		assert.NotNil(t, dbE)
		assert.Equal(t, v.Id, dbE.Id())
		assert.Equal(t, v.OwnerId, dbE.OwnerId())
		assert.Equal(t, v.Name, dbE.Name())
		assert.Equal(t, v.LogoImg, dbE.LogoIMG())
		assert.Equal(t, v.BannerImg, dbE.BannerIMG())
		assert.Equal(t, v.Website, dbE.Website())
		assert.Equal(t, v.Email, dbE.Email())
		assert.Equal(t, v.Phone, dbE.Phone())
		assert.Equal(t, v.Address, dbE.Address())
		assert.Equal(t, industry, dbE.Industry())
		assert.Equal(t, v.UpdatedAt.StringRepresentation(), dbE.UpdatedAt().StringRepresentation())
		assert.Equal(t, v.FoundationDate.StringRepresentation(), dbE.FoundationDate().StringRepresentation())
		assert.Equal(t, v.Employees.Cardinality(), dbE.Employees().Cardinality())
		t.Cleanup(func() {
			err = repository.Remove(ctx, e.Id())
			assert.Nil(t, err)
		})
	}
}
