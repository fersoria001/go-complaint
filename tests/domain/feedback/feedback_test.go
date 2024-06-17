package feedback_test

import (
	"go-complaint/domain/model/complaint"
	"go-complaint/domain/model/feedback"
	"go-complaint/tests"
	"testing"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestFeedback(t *testing.T) {
	// Arrange
	replyReviewID := uuid.New()
	replies := mapset.NewSet[complaint.Reply]()
	replies.Add(*tests.ReceiverReply1)
	review, err := feedback.NewReview(
		replyReviewID,
		tests.Manager.Email(),
		tests.CommonDate,
		"nice reply",
	)
	newReplyReview, err1 := feedback.NewReplyReview(
		replyReviewID,
		tests.ComplaintID,
		replies,
		review,
		"#375774",
	)
	replyReviews := mapset.NewSet[*feedback.ReplyReview]()
	replyReviews.Add(newReplyReview)
	feedbackAnswers := mapset.NewSet[*feedback.Answer]()
	// Act
	newFeedback, err2 := feedback.NewFeedback(
		uuid.New(),
		tests.ComplaintID,
		tests.Asisstant.Email(),
		replyReviews,
		feedbackAnswers,
	)
	// Assert
	assert.Nil(t, err)
	assert.Nil(t, err1)
	assert.Nil(t, err2)
	assert.NotNil(t, newFeedback)
}
