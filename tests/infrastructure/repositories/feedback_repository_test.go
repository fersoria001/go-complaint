package repositories_test

import (
	"context"
	"go-complaint/infrastructure/persistence/datasource"
	"go-complaint/infrastructure/persistence/repositories"
	"go-complaint/tests"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFeedbackSave(t *testing.T) {
	// Arrange
	ctx := context.Background()
	feedbackRepository := repositories.NewFeedbackRepository(
		datasource.PublicSchema(),
	)
	// Act
	err := feedbackRepository.Save(ctx, tests.Feedback1)
	// Assert
	assert.Nil(t, err)
}

func TestFeedbackGet(t *testing.T) {
	// Arrange
	ctx := context.Background()
	feedbackRepository := repositories.NewFeedbackRepository(
		datasource.PublicSchema(),
	)
	// Act
	dbFeedback, err := feedbackRepository.Get(
		ctx,
		tests.Feedback1ID,
	)
	// Assert
	assert.Nil(t, err)
	assert.NotNil(t, dbFeedback)
}
