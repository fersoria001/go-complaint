package application_test

import (
	"context"
	"go-complaint/application"
	"go-complaint/domain/model/common"
	"go-complaint/domain/model/feedback"
	"go-complaint/infrastructure/persistence/datasource"
	"go-complaint/infrastructure/persistence/repositories"
	"go-complaint/tests"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestFeedbackService(t *testing.T) {
	//setup
	ctx := context.Background()
	feedbackRepository := repositories.NewFeedbackRepository(datasource.FeedbackSchema())
	answerRepository := repositories.NewAnswerRepository(datasource.FeedbackSchema())
	feedbackService := application.NewFeedbackService(feedbackRepository, answerRepository)
	stringDate := common.NewDate(time.Now()).StringRepresentation()
	complaintID := uuid.New()
	reviewerID := tests.NewEmployeeID("enterpriseName", "manager@gmail.com", "reviewer@gmail.com")
	reviewedID := tests.NewEmployeeID("enterpriseName", "manager@gmail.com", "reviewed@gmail.com")
	complaintSenderID := "user@gmail.com"
	//test
	t.Run(`Create a feedback when it not exists`, func(t *testing.T) {
		for i := 0; i < 10; i++ {
			err := feedbackService.CreateAFeedback(
				ctx,
				complaintID.String(),
				reviewerID,
				reviewedID,
				"reviewerIMG.jpg",
				"reviewerName",
				complaintSenderID,
				"complaintSenderIMG.jpg",
				"complaintSenderName",
				"complaint body",
				stringDate,
				false,
				stringDate,
				stringDate,
				"This answer review comment",
			)
			if err != nil {
				t.Errorf("Error creating feedback: %v", err)
			}
		}
		dbObj, err := feedbackRepository.Get(ctx, complaintID.String())
		if err != nil {
			t.Errorf("Error getting feedback: %v", err)
		}
		dbFeedback, ok := dbObj.(*feedback.Feedback)
		if !ok {
			t.Errorf("Error invalid type")
		}
		if dbFeedback.ReplyReview().Cardinality() != 10 {
			t.Errorf("Error creating feedback")
		}
	})

	t.Run(`answer a feedback`, func(t *testing.T) {
		err := feedbackService.AnswerAFeedback(
			ctx,
			complaintID.String(),
			reviewedID,
			"reviewedIMG.jpg",
			"reviewedName",
			"answer body",
			stringDate,
			false,
			stringDate,
			stringDate,
		)
		if err != nil {
			t.Errorf("Error answering feedback: %v", err)
		}
		dbAnswers, err := answerRepository.FindByFeedbackID(ctx, complaintID.String())
		if err != nil {
			t.Errorf("Error getting feedback: %v", err)
		}
		if len(dbAnswers) != 1 {
			t.Errorf("Error answering feedback")
		}
	})

	t.Run(`Mark an answer as read`, func(t *testing.T) {
		dbAnswers, err := answerRepository.FindByFeedbackID(ctx, complaintID.String())
		if err != nil {
			t.Errorf("Error getting feedback: %v", err)
		}
		if len(dbAnswers) == 0 {
			t.Errorf("Error marking answer as read")
		}
		err = feedbackService.MarkAnswerAsRead(ctx, dbAnswers[0].ID().String())
		if err != nil {
			t.Errorf("Error marking answer as read: %v", err)
		}
		dbObj, err := answerRepository.Get(ctx, dbAnswers[0].ID().String())
		if err != nil {
			t.Errorf("Error getting feedback: %v", err)
		}
		dbAnswer, ok := dbObj.(*feedback.Answer)
		if !ok {
			t.Errorf("Error invalid type")
		}
		if !dbAnswer.Read() {
			t.Errorf("Error marking answer as read")
		}
	})

	t.Cleanup(
		func() {
			err := feedbackRepository.Remove(ctx, complaintID.String())
			if err != nil {
				t.Errorf("Error deleting feedback: %v", err)
			}
			dbAnswers, err := answerRepository.FindByFeedbackID(ctx, complaintID.String())
			if err != nil {
				t.Errorf("Error getting feedback: %v", err)
			}
			for _, dbAnswer := range dbAnswers {
				err := answerRepository.Remove(ctx, dbAnswer.ID().String())
				if err != nil {
					t.Errorf("Error deleting feedback: %v", err)
				}
			}
		},
	)
}
