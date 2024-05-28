package model_test

import (
	"errors"
	"go-complaint/domain/model/common"
	"go-complaint/domain/model/complaint"
	"go-complaint/erros"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestNewReply_VALID(t *testing.T) {
	id := uuid.New()
	senderID := "UniqueSenderID"
	senderName := "UniqueSenderName"
	surrogateID := uuid.New()
	sameDate := common.NewDate(time.Now())
	valuesMap := map[string]struct {
		ID          uuid.UUID
		SurrogateID uuid.UUID
		SenderID    string
		SenderIMG   string
		SenderName  string
		Body        string
		CreatedAt   common.Date
		Read        bool
		ReadAt      common.Date
		UpdatedAt   common.Date
	}{
		"basic": {
			ID:          id,
			SurrogateID: surrogateID,
			SenderID:    senderID,
			SenderIMG:   "https://www.google.com",
			SenderName:  senderName,
			Body:        "This is a reply",
			CreatedAt:   sameDate,
			Read:        false,
			ReadAt:      sameDate,
			UpdatedAt:   sameDate,
		},
	}
	for name, v := range valuesMap {
		t.Run(name, func(t *testing.T) {
			_, err := complaint.NewReply(v.ID, v.SurrogateID, v.SenderID, v.SenderIMG, v.SenderName, v.Body, v.Read, v.CreatedAt, v.ReadAt, v.UpdatedAt)
			if err != nil {
				t.Errorf("NewReply() = %v, want nil", err)
			}
		})
	}
}

func TestNewReply_INVALID(t *testing.T) {
	id := uuid.New()
	senderID := "UniqueSenderID@gmail.com"
	senderName := "UniqueSenderName"
	surrogateID := uuid.New()
	sameDate := common.NewDate(time.Now())
	valuesMap := map[string]struct {
		ID          uuid.UUID
		SurrogateID uuid.UUID
		SenderID    string
		SenderIMG   string
		SenderName  string
		Body        string
		CreatedAt   common.Date
		Read        bool
		ReadAt      common.Date
		UpdatedAt   common.Date
		Expected    interface{}
	}{
		"nullID": {
			ID:          uuid.Nil,
			SurrogateID: surrogateID,
			SenderID:    senderID,
			SenderIMG:   "https://www.google.com",
			SenderName:  senderName,
			Body:        "This is a reply",
			CreatedAt:   sameDate,
			Read:        false,
			ReadAt:      sameDate,
			UpdatedAt:   sameDate,
			Expected:    &erros.NullValueError{},
		},
		"nullComplaintID": {
			ID:          id,
			SurrogateID: uuid.Nil,
			SenderID:    senderID,
			SenderIMG:   "https://www.google.com",
			SenderName:  senderName,
			Body:        "This is a reply",
			CreatedAt:   sameDate,
			Read:        false,
			ReadAt:      sameDate,
			UpdatedAt:   sameDate,
			Expected:    &erros.NullValueError{},
		},
		"nullSurrogateID": {
			ID:          id,
			SurrogateID: uuid.Nil,
			SenderID:    senderID,
			SenderIMG:   "https://www.google.com",
			SenderName:  senderName,
			Body:        "This is a reply",
			CreatedAt:   sameDate,
			Read:        false,
			ReadAt:      sameDate,
			UpdatedAt:   sameDate,
			Expected:    &erros.NullValueError{},
		},
		"nullSenderIMG": {
			ID:          id,
			SurrogateID: surrogateID,
			SenderID:    senderID,
			SenderIMG:   "",
			SenderName:  senderName,
			Body:        "This is a reply",
			CreatedAt:   sameDate,
			Read:        false,
			ReadAt:      sameDate,
			UpdatedAt:   sameDate,
			Expected:    &erros.NullValueError{},
		},
		"nullSenderName": {
			ID:          id,
			SurrogateID: surrogateID,
			SenderID:    senderID,
			SenderIMG:   "https://www.google.com",
			SenderName:  "",
			Body:        "This is a reply",
			CreatedAt:   sameDate,
			Read:        false,
			ReadAt:      sameDate,
			UpdatedAt:   sameDate,
			Expected:    &erros.NullValueError{},
		},
		"nullBody": {
			ID:          id,
			SurrogateID: surrogateID,
			SenderID:    senderID,
			SenderIMG:   "https://www.google.com",
			SenderName:  senderName,
			Body:        "",
			CreatedAt:   sameDate,
			Read:        false,
			ReadAt:      sameDate,
			UpdatedAt:   sameDate,
			Expected:    &erros.NullValueError{},
		},
		"nullCreatedAt": {
			ID:          id,
			SurrogateID: surrogateID,
			SenderID:    senderID,
			SenderIMG:   "https://www.google.com",
			SenderName:  senderName,
			Body:        "This is a reply",
			CreatedAt:   common.Date{},
			Read:        false,
			ReadAt:      sameDate,
			UpdatedAt:   sameDate,
			Expected:    &erros.EmptyStructError{},
		},
		"nullReadAt": {
			ID:          id,
			SurrogateID: surrogateID,
			SenderID:    senderID,
			SenderIMG:   "https://www.google.com",
			SenderName:  senderName,
			Body:        "This is a reply",
			CreatedAt:   sameDate,
			Read:        false,
			ReadAt:      common.Date{},
			UpdatedAt:   sameDate,
			Expected:    &erros.NullValueError{},
		},
		"nullUpdatedAt": {
			ID:          id,
			SurrogateID: surrogateID,
			SenderID:    senderID,
			SenderIMG:   "https://www.google.com",
			SenderName:  senderName,
			Body:        "This is a reply",
			CreatedAt:   sameDate,
			Read:        false,
			ReadAt:      sameDate,
			UpdatedAt:   common.Date{},
			Expected:    &erros.NullValueError{},
		},
	}
	for name, v := range valuesMap {
		t.Run(name, func(t *testing.T) {
			_, err := complaint.NewReply(v.ID, v.SurrogateID, v.SenderID, v.SenderIMG, v.SenderName, v.Body, v.Read, v.CreatedAt, v.ReadAt, v.UpdatedAt)
			if err == nil {
				t.Errorf("%s didn't throw error, want %v", name, v.Expected)
			}
			if v.Expected != nil {
				if !errors.As(err, &v.Expected) {
					t.Errorf("%s expected %v, got %v", name, v.Expected, err)
				}
			}
		})
	}
}
