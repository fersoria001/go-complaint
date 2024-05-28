package persistence_test

import (
	"context"
	"go-complaint/domain/model/common"
	"go-complaint/domain/model/complaint"
	"go-complaint/infrastructure/persistence/datasource"
	"go-complaint/infrastructure/persistence/repositories"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestReplyRepository(t *testing.T) {
	ctx := context.Background()
	schema := datasource.ComplaintSchema()
	repository := repositories.NewReplyRepository(schema)
	complaintID := uuid.New()
	replyID := uuid.New()
	date := common.NewDate(time.Now())
	senderID := "senderID"
	reply, err := complaint.NewReply(
		replyID,
		complaintID,
		senderID,
		"image.jpg",
		"firstName lastName",
		"message",
		false,
		date,
		date,
		date,
	)
	if err != nil {
		t.Error(err)
	}
	t.Run("A reply is saved, recovered and deleted by it's own identity", func(t *testing.T) {
		err := repository.Save(ctx, reply)
		if err != nil {
			t.Error(err)
		}
		recoveredReply, err := repository.Get(ctx, replyID.String())
		if err != nil {
			t.Error(err)
		}
		err = repository.Remove(ctx, recoveredReply.ID().String())
		if err != nil {
			t.Error(err)
		}
	})
	t.Run("A reply is saved, recovered it's read status is updated and then deleted", func(t *testing.T) {
		err := repository.Save(ctx, reply)
		if err != nil {
			t.Error(err)
		}
		recoveredReply, err := repository.Get(ctx, replyID.String())
		if err != nil {
			t.Error(err)
		}
		recoveredReply.MarkAsRead()
		err = repository.Update(ctx, recoveredReply)
		if err != nil {
			t.Error(err)
		}
		recoveredReply, err = repository.Get(ctx, replyID.String())
		if err != nil {
			t.Error(err)
		}
		if !recoveredReply.Read() {
			t.Error("The reply should be read")
		}
		err = repository.Remove(ctx, recoveredReply.ID().String())
		if err != nil {
			t.Error(err)
		}
	})
	t.Run("A reply is saved and recovered by the complaint identity", func(t *testing.T) {
		err := repository.Save(ctx, reply)
		if err != nil {
			t.Error(err)
		}
		recoveredReplies, err := repository.FindByComplaintID(ctx, complaintID.String())
		if err != nil {
			t.Error(err)
		}
		if len(recoveredReplies) != 1 {
			t.Error("The number of replies should be 1")
		}
		err = repository.Remove(ctx, recoveredReplies[0].ID().String())
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("A reply is saved and ask for count, then remove", func(t *testing.T) {
		err := repository.Save(ctx, reply)
		if err != nil {
			t.Error(err)
		}
		recoveredReplies, err := repository.FindByComplaintID(ctx, complaintID.String())
		if err != nil {
			t.Error(err)
		}
		if len(recoveredReplies) != 1 {
			t.Error("The number of replies should be 1")
		}
		count, err := repository.Count(ctx, complaintID.String())
		if err != nil {
			t.Error(err)
		}
		if count != 1 {
			t.Error("The count should be 1")
		}
		err = repository.Remove(ctx, recoveredReplies[0].ID().String())
		if err != nil {
			t.Error(err)
		}
	})

}
