package repositories_test

import (
	"context"
	"go-complaint/domain/model/feedback"
	"go-complaint/infrastructure/persistence/repositories"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func setup() {
}

func TestCreateComment(t *testing.T) {
	setup()
	ctx := context.Background()
	repository := repositories.MapperRegistryInstance().Get("Feedback").(repositories.FeedbackRepository)
	complaintID := uuid.New()
	//reviewer := tests.Manager.User
	enterpriseID := "enterpriseID"
	//reply := *tests.ReceiverReply1
	fid := uuid.New()
	//rrid := uuid.New()
	//color := "#375701"
	//rr := feedback.NewReplyReviewEntity(rrid, fid, *reviewer, color)
	f := feedback.NewFeedbackEntity(fid, complaintID, enterpriseID)
	t.Run("Create feedback", func(t *testing.T) {
		err := repository.Save(ctx, f)
		assert.Nil(t, err)
	})
	t.Run(("Get empty feedback"), func(t *testing.T) {
		_, err := repository.Get(ctx, f.ID())
		assert.Nil(t, err)
	})

}
