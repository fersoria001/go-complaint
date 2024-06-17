package queries_test

import (
	"context"
	"go-complaint/application/queries"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnterpriseQuery_IsEnterpriseNameAvailable(t *testing.T) {
	// Arrange
	enterpriseQuery := queries.EnterpriseQuery{
		EnterpriseName: "Spoon company",
	}
	// Act
	isAvailable, err := enterpriseQuery.IsEnterpriseNameAvailable(context.Background())
	// Assert
	assert.Nil(t, err)
	assert.False(t, isAvailable)
}
