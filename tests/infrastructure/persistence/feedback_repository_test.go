package persistence_test

import (
	"context"
	"go-complaint/domain/model/common"
	"go-complaint/domain/model/feedback"
	"go-complaint/infrastructure/persistence/datasource"
	"go-complaint/infrastructure/persistence/repositories"
	"go-complaint/tests"
	"testing"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/google/uuid"
)

func TestFeedbackRepository(t *testing.T) {
	//setup
	ctx := context.Background()
	feedbackRepository := repositories.NewFeedbackRepository(datasource.FeedbackSchema())
	commonDate := common.NewDate(time.Now())
	aComplaintUUID := uuid.New()
	anEmployeeID := tests.NewEmployeeID("enterpriseName", "manager@gmail.com", "user@gmail.com")
	anotherEmployeeID := tests.NewEmployeeID("enterpriseName", "manager@gmail.com", "anotheruser@gmail.com")
	domainReply, err := feedback.NewReply(
		"SenderID",
		"SenderIMG",
		"SenderName Fullname",
		"Body of the reply",
		commonDate,
		false,
		commonDate,
		commonDate,
	)
	if err != nil {
		t.Errorf("Error creating reply: %v", err)
	}
	domainReview, err := feedback.NewReview(
		anEmployeeID,
		"ManagerIMG",
		"Manager Fullname",
		commonDate,
		"Comment of the review",
	)
	if err != nil {
		t.Errorf("Error creating review: %v", err)
	}
	randomsUUID := make([]uuid.UUID, 10)
	aMapsetWithAllDomainReplyReviews := mapset.NewSet[*feedback.ReplyReview]()
	for i := 0; i < 10; i++ {
		randomsUUID[i] = uuid.New()
		domainReplyReview, err := feedback.NewReplyReview(
			randomsUUID[i],
			aComplaintUUID,
			domainReply,
			domainReview,
		)
		if err != nil {
			t.Errorf("Error creating reply review: %v", err)
		}
		aMapsetWithAllDomainReplyReviews.Add(domainReplyReview)
	}

	//given
	domainFeedbackWithFilledReplyReviews, err := feedback.NewFeedback(
		aComplaintUUID,
		anEmployeeID,
		anotherEmployeeID,
		aMapsetWithAllDomainReplyReviews,
	)
	if err != nil {
		t.Errorf("Error creating feedback with filled reply reviews: %v", err)
	}

	//when
	t.Run("Save a feedback with filled reply reviews", func(t *testing.T) {
		err := feedbackRepository.Save(ctx, domainFeedbackWithFilledReplyReviews)
		if err != nil {
			t.Errorf("Error saving feedback with filled reply reviews: %v", err)
		}
	})

	t.Run("Get a feedback with filled reply reviews", func(t *testing.T) {
		f, err := feedbackRepository.Get(ctx, aComplaintUUID.String())
		if err != nil {
			t.Errorf("Error getting feedback with filled reply reviews: %v", err)
		}
		dbFeedback, ok := f.(*feedback.Feedback)
		if !ok {
			t.Errorf("Error getting feedback with filled reply reviews: %v", err)
		}
		if dbFeedback.ComplaintID() != aComplaintUUID {
			t.Errorf("Error getting feedback with filled reply reviews: %v", err)
		}
		if dbFeedback.ReviewerID() != anEmployeeID {
			t.Errorf("Error getting feedback with filled reply reviews: %v", err)
		}
		if dbFeedback.ReplyReview().Cardinality() != 10 {
			t.Errorf("Error getting feedback with filled reply reviews: %v", err)
		}
	})

	t.Cleanup(func() {

		err := feedbackRepository.Remove(ctx, aComplaintUUID.String())
		if err != nil {
			t.Errorf("Error removing feedback: %v", err)
		}

	})
}
