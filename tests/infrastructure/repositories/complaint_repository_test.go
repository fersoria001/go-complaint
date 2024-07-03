package repositories_test

import (
	"context"
	"go-complaint/infrastructure/persistence/datasource"
	"go-complaint/infrastructure/persistence/repositories"
	"go-complaint/tests"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComplaintSave(t *testing.T) {
	setup()
	// Arrange
	complaintRepository := repositories.NewComplaintRepository(
		datasource.PublicSchema(),
	)
	complaint := tests.Complaint1
	// Act
	err := complaintRepository.Save(
		context.Background(),
		complaint,
	)
	// Assert
	assert.Nil(t, err)
}

func TestComplaintGet(t *testing.T) {
	setup()
	// Arrange
	complaintRepository := repositories.NewComplaintRepository(
		datasource.PublicSchema(),
	)
	// Act
	dbComplaint, err := complaintRepository.Get(
		context.Background(),
		tests.ComplaintID,
	)
	// Assert
	assert.Nil(t, err)
	assert.NotNil(t, dbComplaint)
	assert.Equal(t, tests.Complaint1.ID(), dbComplaint.ID())
	assert.Equal(t, tests.Complaint1.AuthorID(), dbComplaint.AuthorID())
	assert.Equal(t, tests.Complaint1.ReceiverID(), dbComplaint.ReceiverID())
	assert.Equal(t, tests.Complaint1.Status(), dbComplaint.Status())
	assert.Equal(t, tests.Complaint1.Message(), dbComplaint.Message())
	assert.Equal(t, tests.Complaint1.Rating(), dbComplaint.Rating())
	assert.Equal(t, tests.Complaint1.CreatedAt().StringRepresentation(), dbComplaint.CreatedAt().StringRepresentation())
	assert.Equal(t, tests.Complaint1.Replies(), dbComplaint.Replies())
}

func TestComplaintUpdate(t *testing.T) {
	setup()
	// Arrange
	ctx := context.Background()
	complaintRepository := repositories.NewComplaintRepository(
		datasource.PublicSchema(),
	)
	// Act
	dbComplaint, err := complaintRepository.Get(
		ctx,
		tests.ComplaintID,
	)
	dbComplaint.AddReply(tests.ReceiverReply1)
	err1 := complaintRepository.Update(
		ctx,
		dbComplaint,
	)
	dbComplaint.AddReply(tests.AuthorReply1)
	err2 := complaintRepository.Update(
		ctx,
		dbComplaint,
	)
	// Assert
	assert.Nil(t, err)
	assert.Nil(t, err1)
	assert.Nil(t, err2)
}

func TestComplaintAddReply(t *testing.T) {
	setup()
	// Arrange
	ctx := context.Background()
	complaintRepository := repositories.NewComplaintRepository(
		datasource.PublicSchema(),
	)
	// Act
	dbComplaint, err := complaintRepository.Get(
		ctx,
		tests.ComplaintID,
	)
	dbComplaint.AddReply(tests.ReceiverReply2)
	err1 := complaintRepository.Update(
		ctx,
		dbComplaint,
	)
	dbComplaint.AddReply(tests.AuthorReply2)
	err2 := complaintRepository.Update(
		ctx,
		dbComplaint,
	)
	// Assert
	assert.Nil(t, err)
	assert.Nil(t, err1)
	assert.Nil(t, err2)
}
