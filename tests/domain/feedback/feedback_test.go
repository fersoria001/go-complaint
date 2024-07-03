package feedback_test

import (
	"context"
	"go-complaint/domain/model/feedback"
	"go-complaint/tests"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestFeedback(t *testing.T) {
	ctx := context.Background()
	complaintID := uuid.New()
	reviewer := tests.Manager.User
	enterpriseID := "enterpriseID"
	reply := *tests.ReceiverReply1
	reply1 := *tests.ReceiverReply2
	reply2 := *tests.AuthorReply1
	reply3 := *tests.AuthorReply2
	fid := uuid.New()
	rrid := uuid.New()
	color := "#375701"
	rr := feedback.NewReplyReviewEntity(rrid, fid, *reviewer, color)
	f := feedback.NewFeedbackEntity(fid, complaintID, enterpriseID)

	t.Run("Add an empty reply review", func(t *testing.T) {
		err := f.AddReplyReview(ctx, rr)
		assert.Nil(t, err)
		assert.Equal(t, 1, f.ReplyReviews().Cardinality())
	})
	t.Run("Remove a replyReview", func(t *testing.T) {
		id, err := f.RemoveReplyReview(rr)
		assert.Nil(t, err)
		assert.Equal(t, 0, f.ReplyReviews().Cardinality())
		assert.Equal(t, rrid, id)
	})
	t.Run("Add a reply", func(t *testing.T) {
		err := f.AddReply(ctx, color, reply)
		assert.NotNil(t, err)
		assert.Equal(t, 0, f.ReplyReviews().Cardinality())
		err = f.AddReplyReview(ctx, rr)
		assert.Nil(t, err)
		assert.Equal(t, 1, f.ReplyReviews().Cardinality())
		err = f.AddReply(ctx, color, reply)
		assert.Nil(t, err)
		assert.Equal(t, 1, f.ReplyReviews().Cardinality())
		rr, err := f.ReplyReview(color)
		assert.Nil(t, err)
		assert.Equal(t, 1, rr.Replies().Cardinality())
	})
	t.Run("Remove a reply", func(t *testing.T) {
		err := f.RemoveReply(color, reply)
		assert.Nil(t, err)
		assert.Equal(t, 1, f.ReplyReviews().Cardinality())
		replyReview, err := f.ReplyReview(color)
		assert.Nil(t, err)
		assert.Equal(t, 0, replyReview.Replies().Cardinality())
	})
	t.Run("Add replies", func(t *testing.T) {
		err := f.AddReply(ctx, color, reply)
		assert.Nil(t, err)
		err = f.AddReply(ctx, color, reply1)
		assert.Nil(t, err)
		err = f.AddReply(ctx, color, reply2)
		assert.Nil(t, err)
		err = f.AddReply(ctx, color, reply3)
		assert.Nil(t, err)
	})
	t.Run("Remove all replies", func(t *testing.T) {
		replyReview, err := f.ReplyReview(color)
		assert.Nil(t, err)
		err = f.RemoveReply(color, reply)
		assert.Nil(t, err)
		assert.Equal(t, 3, replyReview.Replies().Cardinality())
		err = f.RemoveReply(color, reply1)
		assert.Nil(t, err)
		assert.Equal(t, 2, replyReview.Replies().Cardinality())
		err = f.RemoveReply(color, reply2)
		assert.Nil(t, err)
		assert.Equal(t, 1, replyReview.Replies().Cardinality())
		err = f.RemoveReply(color, reply3)
		assert.Nil(t, err)
		assert.Equal(t, 0, replyReview.Replies().Cardinality())
	})
}
