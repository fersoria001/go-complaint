package infrastructure_test

import (
	"go-complaint/infrastructure"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInMemoryCache(t *testing.T) {
	// Arrange
	cache := infrastructure.InMemoryCacheInstance()
	// Act
	cache.Set("key", "value")
	value, ok := cache.Get("key")
	// Assert
	assert.True(t, ok)
	assert.Equal(t, "value", value)
}
