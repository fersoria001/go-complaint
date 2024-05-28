package model_test

import (
	"context"
	"errors"
	"go-complaint/domain/model/common"
	"go-complaint/domain/model/enterprise"
	"go-complaint/erros"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestEmployee(t *testing.T) {

	ctx := context.Background()
	enterpriseName := "NukaCola"
	ownerEmail := "Roger@nukacola.com"
	ddmmyy := time.Now().Format("02/01/2006")
	uuidSegment := strings.Split(uuid.New().String(), "-")[4]
	ceoID := enterpriseName + "-" + ownerEmail + "-" + ddmmyy + "-" + uuidSegment
	hiringDate := common.NewDate(time.Now())
	//If a fakeInvitationID is used in the application it will fail, here is not a problem
	fakeInvitationID := uuid.New()
	ceo, err := enterprise.NewEmployee(
		ceoID,
		"iamge.jpg",
		"Roger",
		"Smith",
		30,
		ownerEmail,
		"555-555-555",
		enterprise.MANAGER,
		hiringDate,
		true,
		hiringDate,
	)
	if err != nil {
		t.Error(err)
	}
	manager, err := enterprise.NewEmployee(
		ceoID,
		"iamge.jpg",
		"Roger",
		"Smith",
		30,
		ownerEmail,
		"555-555-555",
		enterprise.MANAGER,
		hiringDate,
		true,
		hiringDate,
	)
	if err != nil {
		t.Error(err)
	}
	assistant, err := enterprise.NewEmployee(
		ceoID,
		"iamge.jpg",
		"Roger",
		"Smith",
		30,
		ownerEmail,
		"555-555-555",
		enterprise.ASSISTANT,
		hiringDate,
		true,
		hiringDate)
	if err != nil {
		t.Error(err)
	}
	if ceo.ID() != ceoID {
		t.Errorf("expected %s but got %s", ceoID, ceo.ID())
	}
	t.Run("A ceo can hire an employee,it publish an event to the event bus",
		func(t *testing.T) {

			newEmployee, err := ceo.HireEmployee(
				ctx,
				fakeInvitationID,
				ceoID,
				"iamge.jpg",
				"John",
				"Doe",
				"johndoe@gmail.com",
				"555-555-552",
				25,
			)
			if err != nil {
				t.Error(err)
			}
			if newEmployee.Position() != enterprise.NOT_ASSIGNED {
				t.Errorf("expected %s but got % s", enterprise.NOT_ASSIGNED, newEmployee.Position())
			}
			if newEmployee.ApprovedHiring() != false {
				t.Errorf("expected %t but got %t", false, newEmployee.ApprovedHiring())
			}

		})
	t.Run("A manager can hire an employee, it  publish an event to the event bus",
		func(t *testing.T) {
			newEmployee, err := manager.HireEmployee(
				ctx,
				fakeInvitationID,
				ceoID,
				"iamge.jpg",
				"John",
				"Doe",
				"johndoe@gmail.com",
				"555-555-552",
				25,
			)
			if err != nil {
				t.Error(err)
			}
			if newEmployee.Position() != enterprise.NOT_ASSIGNED {
				t.Errorf("expected %s but got %s", enterprise.NOT_ASSIGNED, newEmployee.Position())
			}
			if newEmployee.ApprovedHiring() != false {
				t.Errorf("expected %t but got %t", false, newEmployee.ApprovedHiring())
			}

		})
	t.Run("An assistant can not hire an employee, it returns an error and don't publish an event",
		func(t *testing.T) {
			_, err := assistant.HireEmployee(
				ctx,
				fakeInvitationID,
				ceoID,
				"iamge.jpg",
				"John",
				"Doe",
				"johndoe@gmail.com",
				"555-555-552",
				25,
			)
			expectedError := &erros.ValidationError{}
			if err == nil {
				t.Errorf("expected an error but got nil")
			}
			if !errors.As(err, &expectedError) {
				t.Errorf("expected %T but got %T", expectedError, err)
			}

		})

}
