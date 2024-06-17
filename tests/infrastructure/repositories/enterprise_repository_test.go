package repositories_test

import (
	"context"
	"go-complaint/infrastructure/persistence/datasource"
	"go-complaint/infrastructure/persistence/repositories"
	"go-complaint/tests"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSave(t *testing.T) {
	// Arrange
	ctx := context.Background()
	enterpriseRepository := repositories.NewEnterpriseRepository(datasource.PublicSchema())
	newEnterprise := tests.Enterprise1
	// Act
	err := enterpriseRepository.Save(ctx, newEnterprise)
	// Assert
	assert.Nil(t, err)
}

func TestGet(t *testing.T) {
	// Arrange
	ctx := context.Background()
	enterpriseRepository := repositories.NewEnterpriseRepository(datasource.PublicSchema())
	// Act
	dbEnterprise, err := enterpriseRepository.Get(ctx, tests.Enterprise1.Name())
	// Assert
	assert.Nil(t, err)
	assert.NotNil(t, dbEnterprise)
	assert.NotNil(t, dbEnterprise.Employees())
	assert.NotNil(t, dbEnterprise.Owner())
}

func TestUpdate(t *testing.T) {
	// Arrange
	ctx := context.Background()
	enterpriseRepository := repositories.NewEnterpriseRepository(datasource.PublicSchema())
	// Act
	dbEnterprise, err := enterpriseRepository.Get(ctx, tests.Enterprise1.Name())
	dbEnterprise.AddEmployee(tests.Asisstant)
	err1 := enterpriseRepository.Update(ctx, dbEnterprise)
	dbEnterprise.AddEmployee(tests.Manager)
	err2 := enterpriseRepository.Update(ctx, dbEnterprise)
	//
	assert.Nil(t, err)
	assert.Equal(t, 2, dbEnterprise.Employees().Cardinality())
	assert.Nil(t, err1)
	assert.Nil(t, err2)

}
