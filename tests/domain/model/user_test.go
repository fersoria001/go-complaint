package model_test

import (
	"errors"
	"go-complaint/domain/model/common"
	"go-complaint/domain/model/identity"
	"go-complaint/erros"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestNewUser_VALID(t *testing.T) {
	email := "email@gmail.com"
	address, err := common.NewAddress("country", "county", "city")
	if err != nil {
		t.Errorf("NewAddress() got = %v, want nil", err)
	}
	birthDate := common.NewDate(time.Now())
	registerDate := common.NewDate(time.Now())
	person, err := identity.NewPerson(email, "firstName", "lastName", "01234567890", birthDate, address)
	if err != nil {
		t.Errorf("expected nil error, got %v", err)
	}
	pass := "Password1"

	data := map[string]struct {
		personID     uuid.UUID
		profileIMG   string
		registerDate common.Date
		email        string
		password     string
		person       *identity.Person
	}{
		"basic": {
			profileIMG:   "profileIMG",
			registerDate: registerDate,
			email:        "email@gmail.com",
			password:     pass,
			person:       person,
		},
		"emailAlternative1": {
			profileIMG:   "profileIMG",
			registerDate: registerDate,
			email:        "email@gmail.com.ar",
			password:     pass,
			person:       person,
		},
		"emailAlternative2": {
			profileIMG:   "profileIMG",
			registerDate: registerDate,
			email:        "email.em@hotmail.live.ar",
			password:     pass,
			person:       person,
		},
		"emailAlternative3": {
			profileIMG:   "profileIMG",
			registerDate: registerDate,
			email:        "email123@hotmail.live",
			password:     pass,
			person:       person,
		},
	}

	for name, testData := range data {
		t.Run(name, func(t *testing.T) {
			_, err := identity.NewUser(testData.profileIMG, testData.registerDate, testData.email, testData.password, testData.person)
			if err != nil {
				t.Errorf("%s test expected nil error, got %v", name, err)
			}
		})
	}
}

func TestNewUser_INVALID(t *testing.T) {
	email := "email@gmail.com"
	address, err := common.NewAddress("country", "county", "city")
	if err != nil {
		t.Errorf("NewAddress() got = %v, want nil", err)
	}
	birthDate := common.NewDate(time.Now())
	registerDate := common.NewDate(time.Now())
	person, err := identity.NewPerson(email, "firstName", "lastName", "01234567890", birthDate, address)
	if err != nil {
		t.Errorf("expected nil error, got %v", err)
	}
	pass := "Password1"

	data := map[string]struct {
		profileIMG   string
		registerDate common.Date
		email        string
		password     string
		person       *identity.Person
		expected     interface{}
	}{
		"nullProfileIMG": {
			profileIMG:   "",
			registerDate: registerDate,
			email:        "email@gmail.com",
			password:     pass,
			person:       person,
			expected:     &erros.NullValueError{},
		},
		"nullRegisterDate": {
			profileIMG:   "profileIMG",
			registerDate: common.Date{},
			email:        "email@gmail.com",
			password:     pass,
			person:       person,
			expected:     &erros.EmptyStructError{},
		},
		"nullEmail": {
			profileIMG:   "profileIMG",
			registerDate: registerDate,
			email:        "",
			password:     pass,
			person:       person,
			expected:     &erros.NullValueError{},
		},
		"nullPassword": {
			profileIMG:   "profileIMG",
			registerDate: registerDate,
			email:        "email@gmail.com",
			password:     "",
			person:       person,
			expected:     &erros.NullValueError{},
		},
		"invalidEmail1": {
			profileIMG:   "profileIMG",
			registerDate: registerDate,
			email:        "email",
			password:     pass,
			person:       person,
			expected:     &erros.InvalidEmailError{},
		},
		"invalidEmail2": {
			profileIMG:   "profileIMG",
			registerDate: registerDate,
			email:        "email@",
			password:     pass,
			person:       person,
			expected:     &erros.InvalidEmailError{},
		},
		"invalidEmail3": {
			profileIMG:   "profileIMG",
			registerDate: registerDate,
			email:        "email@.com",
			password:     pass,
			person:       person,
			expected:     &erros.InvalidEmailError{},
		},
		"invalidEmail4": {
			profileIMG:   "profileIMG",
			registerDate: registerDate,
			email:        "email@.com.ar",
			password:     pass,
			person:       person,
			expected:     &erros.InvalidEmailError{},
		},
	}
	for name, testData := range data {
		t.Run(name, func(t *testing.T) {
			_, err := identity.NewUser(testData.profileIMG, testData.registerDate, testData.email, testData.password, testData.person)
			if err == nil {
				t.Errorf("expected error, got nil")
			}
			if testData.expected != nil {
				if !errors.As(err, &testData.expected) {
					t.Errorf("%v expected %v, got %v", name, testData.expected, err)
				}
			}

		})
	}

}
