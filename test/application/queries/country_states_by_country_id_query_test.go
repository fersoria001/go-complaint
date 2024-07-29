package queries_test

import (
	"context"
	"go-complaint/application/queries"
	"go-complaint/test/mock_data"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountryStatesByCountryIdQuery_Execute(t *testing.T) {
	ctx := context.Background()
	q := queries.NewCountryStatesByCountryIdQuery(mock_data.Country.ID())
	countryStates, err := q.Execute(ctx)
	assert.Nil(t, err)
	assert.Greater(t, len(countryStates), 0)
}
