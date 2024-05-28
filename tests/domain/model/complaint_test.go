package model_test

import (
	"context"
	"errors"
	"go-complaint/domain/model/complaint"
	"go-complaint/erros"
	"go-complaint/tests"
	"testing"
)

func TestComplaint(t *testing.T) {
	ctx := context.Background()
	authorID := "author@gmail.com"
	receiverID := "receiver@hotmail.com"

	title := tests.Repeat("a", 10)
	description := tests.Repeat("a", 30)
	content := tests.Repeat("a", 50)
	t.Run("A complaint is created and sent to the event bus, it returns the event to be saved",
		func(t *testing.T) {
			sent, err := complaint.SendComplaint(
				ctx,
				authorID,
				receiverID,
				title,
				description,
				content)
			if err != nil {
				t.Errorf("Expected no error, got %v", err)
			}
			if sent.Status() != complaint.OPEN {
				t.Errorf("Expected status to be OPEN, got %v", sent.Status())
			}

		})
	t.Run("A new complaint reply is created and sent to the event bus, it returns the reply",
		func(t *testing.T) {
			complaint, err := complaint.SendComplaint(
				ctx,
				authorID,
				receiverID,
				title,
				description,
				content)
			if err != nil {
				t.Errorf("Expected no error, got %v", err)
			}
			count := 0
			reply, err := complaint.ReplyComplaint(
				ctx,
				count,
				receiverID,
				"senderIMG",
				"senderName",
				"body")
			if err != nil {
				t.Errorf("Expected no error, got %v", err)
			}
			if reply.Read() != false {
				t.Errorf("Expected read to be false, got %v", reply.Read())
			}

		})

	t.Run("When you try to rate a not closed complaint, it returns an error",
		func(t *testing.T) {
			complaint, err := complaint.SendComplaint(
				ctx,
				authorID,
				receiverID,
				title,
				description,
				content)
			if err != nil {
				t.Errorf("Expected no error, got %v", err)
			}
			expectedErr := &erros.ValidationError{}
			err = complaint.Rate(ctx, authorID, 5, "this is a great solution to my problem")
			if err == nil {
				t.Errorf("Expected no error, got %v", err)
			}
			if !errors.As(err, &expectedErr) {
				t.Errorf("Expected %v, got %v", expectedErr, err)
			}
		})
	t.Run("Rate a closed complaint, it modifies the complaint pointer",
		func(t *testing.T) {
			complaintt, err := complaint.SendComplaint(
				ctx,
				authorID,
				receiverID,
				title,
				description,
				content)
			if err != nil {
				t.Errorf("Expected no error, got %v", err)
			}
			_, err = complaintt.ReplyComplaint(ctx,
				0,
				receiverID,
				"senderIMG",
				"senderName",
				"body")
			if err != nil {
				t.Errorf("Expected no error, got %v", err)
			}
			if complaintt.Status() != complaint.STARTED {
				t.Errorf("Expected status to be STARTED, got %v", complaintt.Status())
			}
			_, err = complaintt.ReplyComplaint(ctx,
				1,
				authorID,
				"senderIMG",
				"senderName",
				"body")
			if err != nil {
				t.Errorf("Expected no error, got %v", err)
			}
			if complaintt.Status() != complaint.IN_DISCUSSION {
				t.Errorf("Expected status to be IN_DISCUSSION, got %v", complaintt.Status())
			}

			err = complaintt.Close(ctx, authorID)
			if err != nil {
				t.Errorf("Expected no error, got %v", err)
			}
			err = complaintt.Rate(ctx, authorID, 5, "this is a great solution to my problem")
			if err != nil {
				t.Errorf("Expected no error, got %v", err)
			}
			if complaintt.Rating().Rate() != 5 {
				t.Errorf("Expected rate to be 5, got %v", complaintt.Rating().Rate())
			}
		})
}
