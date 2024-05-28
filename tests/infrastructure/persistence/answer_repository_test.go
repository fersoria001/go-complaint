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

	"github.com/google/uuid"
)

func TestAnswerRepository(t *testing.T) {
	//setup
	ctx := context.Background()
	answerRepository := repositories.NewAnswerRepository(datasource.FeedbackSchema())
	commonDate := common.NewDate(time.Now())
	aComplaintUUID := uuid.New()
	anEmployeeID := tests.NewEmployeeID("enterpriseName", "manager@gmail.com", "g@gmail.com")
	newAnswer, err := feedback.NewAnswer(
		uuid.New(),
		aComplaintUUID,
		anEmployeeID,
		"SenderIMG",
		"SenderName Fullname",
		"Body of the answer",
		commonDate,
		false,
		commonDate,
		commonDate,
	)
	if err != nil {
		t.Errorf("Error creating answer: %v", err)
	}
	//test
	t.Run("Create", func(t *testing.T) {
		err := answerRepository.Save(ctx, newAnswer)
		if err != nil {
			t.Errorf("Error creating answer: %v", err)
		}
	})
	t.Run("Get", func(t *testing.T) {
		dbObj, err := answerRepository.Get(ctx, newAnswer.ID().String())
		if err != nil {
			t.Errorf("Error getting answer: %v", err)
		}
		_, ok := dbObj.(*feedback.Answer)
		if !ok {
			t.Errorf("Error getting answer: invalid type")
		}
	})

	t.Run("Update", func(t *testing.T) {
		newAnswer.MarkAsRead()
		err := answerRepository.Update(ctx, newAnswer)
		if err != nil {
			t.Errorf("Error updating answer: %v", err)
		}
		dbObj, err := answerRepository.Get(ctx, newAnswer.ID().String())
		if err != nil {
			t.Errorf("Error getting answer: %v", err)
		}
		dbAnswer, ok := dbObj.(*feedback.Answer)
		if !ok {
			t.Errorf("Error getting answer: invalid type")
		}
		if !dbAnswer.Read() {
			t.Errorf("Error updating answer: read field not updated")
		}
	})

	t.Run("Find by feedback id", func(t *testing.T) {
		dbObj, err := answerRepository.FindByFeedbackID(ctx, aComplaintUUID.String())
		if err != nil {
			t.Errorf("Error finding answer by feedback id: %v", err)
		}
		if len(dbObj) == 0 {
			t.Errorf("Error finding answer by feedback id: no answer found")
		}
	})

	t.Run("Delete", func(t *testing.T) {
		err := answerRepository.Remove(ctx, newAnswer.ID().String())
		if err != nil {
			t.Errorf("Error deleting answer: %v", err)
		}
		dbObj, err := answerRepository.Get(ctx, newAnswer.ID().String())
		if err == nil {
			t.Errorf("Error deleting answer: answer not deleted")
		}
		if dbObj != nil {
			t.Errorf("Error deleting answer: answer not deleted")
		}
	})
}
