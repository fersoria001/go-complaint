package repositories_test

import (
	"go-complaint/infrastructure/persistence/repositories"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapperRegistry(t *testing.T) {
	// Arrange
	instance := repositories.MapperRegistryInstance()
	// Act
	mapper := instance.Get("Complaint")
	m, ok := mapper.(repositories.ComplaintRepository)
	// Assert
	assert.True(t, ok)
	assert.NotNil(t, m)
}
