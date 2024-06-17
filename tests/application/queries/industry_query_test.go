package queries_test

import (
	"context"
	"go-complaint/application/queries"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndustryQuery(t *testing.T) {
	// Arrange
	ctx := context.Background()
	industryQuery := queries.IndustryQuery{}
	// Act
	allIndustries, err := industryQuery.AllIndustries(ctx)
	// Assert
	assert.Nil(t, err)
	assert.NotNil(t, allIndustries)
	assert.Greater(t, len(allIndustries), 0)
}
