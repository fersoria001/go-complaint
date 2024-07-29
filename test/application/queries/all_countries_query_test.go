package queries_test

import (
	"context"
	"go-complaint/application/queries"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAllCountriesQuery_Execute(t *testing.T) {
	ctx := context.Background()
	q := queries.NewAllCountriesQuery()
	countries, err := q.Execute(ctx)
	assert.Nil(t, err)
	assert.NotNil(t, countries)
	assert.Greater(t, len(countries), 0)
}
