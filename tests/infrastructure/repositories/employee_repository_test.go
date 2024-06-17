package repositories_test

import (
	"context"
	"go-complaint/infrastructure/persistence/datasource"
	employeefindall "go-complaint/infrastructure/persistence/finders/employee_findall"
	"go-complaint/infrastructure/persistence/repositories"
	"go-complaint/tests"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindAllByEnterpriseID(t *testing.T) {
	// Arrange
	ctx := context.Background()
	employeeRepository := repositories.NewEmployeeRepository(
		datasource.PublicSchema(),
	)
	// Act
	employees, err := employeeRepository.FindAll(
		ctx,
		employeefindall.NewByEnterpriseID(tests.Enterprise1.Name()),
	)
	// Assert
	assert.Nil(t, err)
	assert.NotNil(t, employees)
	assert.Greater(t, employees.Cardinality(), 0)

}
