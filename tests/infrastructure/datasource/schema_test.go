package datasource_test

import (
	"context"
	"go-complaint/infrastructure/persistence/datasource"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSchema(t *testing.T) {
	// Arrange
	ctx := context.Background()
	instance := datasource.PublicSchema()
	// Act
	conn, err1 := instance.Acquire(ctx)
	// Assert
	assert.Nil(t, err1)
	assert.NotNil(t, conn)
	conn.Release()
}
