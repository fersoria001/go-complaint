package commands_test

import (
	"context"
	"go-complaint/application/commands"
	"go-complaint/infrastructure/cache"
	"go-complaint/infrastructure/persistence/finders/find_all_feedback"
	"go-complaint/infrastructure/persistence/repositories"
	"go-complaint/tests"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFeedback(t *testing.T) {
	ctx := context.Background()
	repository := repositories.MapperRegistryInstance().Get("Feedback").(repositories.FeedbackRepository)
	complaintID := tests.ComplaintID
	reviewer := tests.Manager.User
	enterpriseID := "enterpriseID"
	reply := *tests.ReceiverReply1
	color := "#375701"
	go cache.Cache(cache.RequestChannel)
	t.Run("CreateFeedback", func(t *testing.T) {
		command := commands.FeedbackCommand{
			Color:        color,
			ComplaintID:  complaintID.String(),
			EnterpriseID: enterpriseID,
		}
		err := command.CreateFeedback(ctx)
		assert.Nil(t, err)
	})
	t.Run("AddReply", func(t *testing.T) {
		dbfs, err := repository.FindAll(ctx, find_all_feedback.ByComplaintID(complaintID))
		assert.Nil(t, err)
		assert.Equal(t, 1, dbfs.Cardinality())
		dbf, ok := dbfs.Pop()
		assert.True(t, ok)
		command := commands.FeedbackCommand{
			FeedbackID: dbf.ID().String(),
			ReviewerID: reviewer.Email(),
			Color:      color,
			Replies:    []string{reply.ID().String()},
		}
		err = command.AddReply(ctx)
		assert.Nil(t, err)
		dbf, err = repository.Get(ctx, dbf.ID())
		assert.Nil(t, err)
		assert.Equal(t, 1, dbf.ReplyReviews().Cardinality())
		dbfReplyReview, err := dbf.ReplyReview(color)
		assert.Nil(t, err)
		assert.Equal(t, 1, dbfReplyReview.Replies().Cardinality())
	})
	t.Run("AddComment", func(t *testing.T) {
		dbfs, err := repository.FindAll(ctx, find_all_feedback.ByComplaintID(complaintID))
		assert.Nil(t, err)
		assert.Equal(t, 1, dbfs.Cardinality())
		dbf, ok := dbfs.Pop()
		assert.True(t, ok)
		command := commands.FeedbackCommand{
			FeedbackID: dbf.ID().String(),
			Color:      color,
			Comment:    "comment",
		}
		err = command.AddComment(ctx)
		assert.Nil(t, err)
		dbf, err = repository.Get(ctx, dbf.ID())
		assert.Nil(t, err)
		dbfReplyReview, err := dbf.ReplyReview(color)
		assert.Nil(t, err)
		comment := dbfReplyReview.Review().Comment()
		assert.Equal(t, "comment", comment)
	})
	t.Run("DeleteComment: deletes a reply review", func(t *testing.T) {
		dbfs, err := repository.FindAll(ctx, find_all_feedback.ByComplaintID(complaintID))
		assert.Nil(t, err)
		assert.Equal(t, 1, dbfs.Cardinality())
		dbf, ok := dbfs.Pop()
		assert.True(t, ok)
		command := commands.FeedbackCommand{
			FeedbackID: dbf.ID().String(),
			Color:      color,
		}
		err = command.DeleteComment(ctx)
		assert.Nil(t, err)
		dbf, err = repository.Get(ctx, dbf.ID())
		assert.Nil(t, err)
		dbfReplyReview, err := dbf.ReplyReview(color)
		assert.NotNil(t, err)
		assert.Nil(t, dbfReplyReview)
	})

}
