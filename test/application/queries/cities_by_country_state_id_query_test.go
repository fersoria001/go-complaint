package queries_test

import (
	"context"
	"go-complaint/application/queries"
	"go-complaint/test/mock_data"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCitiesByCountryStateIdQuery(t *testing.T) {
	ctx := context.Background()
	q := queries.NewCitiesByCountryStateIdQuery(mock_data.CountryState.ID())
	cities, err := q.Execute(ctx)
	assert.Nil(t, err)
	assert.Greater(t, len(cities), 0)
}
