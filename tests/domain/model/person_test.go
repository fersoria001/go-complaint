package model_test

import (
	"errors"
	"go-complaint/domain/model/common"
	"go-complaint/domain/model/identity"
	"go-complaint/erros"
	"testing"
	"time"
)

func TestNewPerson_VALID(t *testing.T) {
	email := "email@gmail.com"
	address, err := common.NewAddress("country", "county", "city")
	if err != nil {
		t.Errorf("NewAddress() got = %v, want nil", err)
	}
	birthDate := common.NewDate(time.Now())
	data := map[string]struct {
		email     string
		firstName string
		lastName  string
		birthDate common.Date
		phone     string
		address   common.Address
	}{
		"basic": {
			email:     email,
			firstName: "firstName",
			lastName:  "lastName",
			birthDate: birthDate,
			phone:     "1234567890",
			address:   address,
		},

		"nameAlternative1": {
			email:     email,
			firstName: "René",
			lastName:  "García",
			birthDate: birthDate,
			phone:     "1234567890",
			address:   address,
		},
		"nameAlternative2": {
			email:     email,
			firstName: "Jürgen",
			lastName:  "Müller",
			birthDate: birthDate,
			phone:     "1234567890",
			address:   address,
		},
		"nameAlternative3": {
			email:     email,
			firstName: "Hélène",
			lastName:  "Dupont",
			birthDate: birthDate,
			phone:     "1234567890",
			address:   address,
		},
		"nameAlternative4": {
			email:     email,
			firstName: "Andrés",
			lastName:  "Gómez",
			birthDate: birthDate,
			phone:     "1234567890",
			address:   address,
		},
		"nameAlternative5": {
			email:     email,
			firstName: "Chiara",
			lastName:  "Rossi",
			birthDate: birthDate,
			phone:     "1234567890",
			address:   address,
		},
		"nameAlternative6": {
			email:     email,
			firstName: "Søren",
			lastName:  "Jensen",
			birthDate: birthDate,
			phone:     "1234567890",
			address:   address,
		},
		"nameAlternative7": {
			email:     email,
			firstName: "François",
			lastName:  "Dupont",
			birthDate: birthDate,
			phone:     "1234567890",
			address:   address,
		},
		"nameAlternative8": {
			email:     email,
			firstName: "Günther",
			lastName:  "Schäfer",
			birthDate: birthDate,
			phone:     "1234567890",
			address:   address,
		},
		"nameAlternative9": {
			email:     email,
			firstName: "Björn",
			lastName:  "Andersson",
			birthDate: birthDate,
			phone:     "1234567890",
			address:   address,
		},
		"minLengthFirstName": {
			email:     email,
			firstName: "fi",
			lastName:  "lastName",
			birthDate: birthDate,
			phone:     "1234567890",
			address:   address,
		},
		"minLengthLastName": {
			email:     email,
			firstName: "firstName",
			lastName:  "la",
			birthDate: birthDate,
			phone:     "1234567890",
			address:   address,
		},
		"maxLengthFirstName": {
			email:     email,
			firstName: "firstNamesfirstNamesfirstNamesfirstNamesfirstNames",
			lastName:  "lastName",
			birthDate: birthDate,
			phone:     "1234567890",
			address:   address,
		},
		"maxLengthLastName": {
			email:     email,
			firstName: "firstName",
			lastName:  "lastNameslastNameslastNameslastNameslastNamesNames",
			birthDate: birthDate,
			phone:     "1234567890",
			address:   address,
		},
		"maxLengthPhone": {
			email:     email,
			firstName: "firstName",
			lastName:  "lastName",
			birthDate: birthDate,
			phone:     "1234567890 1234567890",
			address:   address,
		},
	}
	for _, tt := range data {
		t.Run(tt.firstName, func(t *testing.T) {
			_, err := identity.NewPerson(tt.email, tt.firstName, tt.lastName, tt.phone, tt.birthDate, tt.address)
			if err != nil {
				t.Errorf("NewPerson() got = %v, want nil", err)
			}
		})
	}
}

func TestNewPerson_INVALID(t *testing.T) {
	email := "email@gmail.com"
	address, err := common.NewAddress("country", "county", "city")
	if err != nil {
		t.Errorf("NewAddress() got = %v, want nil", err)
	}
	birthDate := common.NewDate(time.Now())
	data := map[string]struct {
		email     string
		firstName string
		lastName  string
		birthDate common.Date
		phone     string
		address   common.Address
		expected  interface{}
	}{
		"nullFirstName": {
			email:     email,
			firstName: "",
			lastName:  "lastName",
			birthDate: birthDate,
			phone:     "1234567890",
			address:   address,
			expected:  &erros.NullValueError{},
		},
		"nullLastName": {
			email:     email,
			firstName: "firstName",
			lastName:  "",
			birthDate: birthDate,
			phone:     "1234567890",
			address:   address,
			expected:  &erros.NullValueError{},
		},
		"nullPhone": {
			email:     email,
			firstName: "firstName",
			lastName:  "lastName",
			birthDate: birthDate,
			phone:     "",
			address:   address,
			expected:  &erros.NullValueError{},
		},
		"nullAddress": {
			email:     email,
			firstName: "firstName",
			lastName:  "lastName",
			birthDate: birthDate,
			phone:     "1234567890",
			address:   common.Address{},
			expected:  &erros.EmptyStructError{},
		},

		"minLengthPhone": {
			email:     email,
			firstName: "firstName",
			lastName:  "lastName",
			birthDate: birthDate,
			phone:     "12345678",
			address:   address,
			expected:  &erros.InvalidLengthError{},
		},
		"minLengthFirstName": {
			email:     email,
			firstName: "f",
			lastName:  "lastName",
			birthDate: birthDate,
			phone:     "1234567890",
			address:   address,
			expected:  &erros.InvalidLengthError{},
		},
		"minLengthLastName": {
			email:     email,
			firstName: "firstName",
			lastName:  "l",
			birthDate: birthDate,
			phone:     "1234567890",
			address:   address,
			expected:  &erros.InvalidLengthError{},
		},
		"maxLengthFirstName": {
			email:     email,
			firstName: "firstNamesfirstNamesfirstNamesfirstNamesfirstNames-",
			lastName:  "lastName",
			birthDate: birthDate,
			phone:     "1234567890",
			address:   address,
			expected:  &erros.InvalidLengthError{},
		},
		"maxLengthLastName": {
			email:     email,
			firstName: "firstName",
			lastName:  "lastNameslastNameslastNameslastNameslastNamesNames-",
			birthDate: birthDate,
			phone:     "1234567890",
			address:   address,
			expected:  &erros.InvalidLengthError{},
		},
		"maxLengthPhone": {
			email:     email,
			firstName: "firstName",
			lastName:  "lastName",
			birthDate: birthDate,
			phone:     "12345567890 1234567890",
			address:   address,
			expected:  &erros.InvalidLengthError{},
		},
		"nullBirthDate": {
			email:     email,
			firstName: "firstName",
			lastName:  "lastName",
			birthDate: common.Date{},
			phone:     "1234567890",
			address:   address,
			expected:  &erros.EmptyStructError{},
		},
		"invalidName": {
			email:     email,
			firstName: "firstName1",
			lastName:  "lastName",
			birthDate: birthDate,
			phone:     "1234567890",
			address:   address,
			expected:  &erros.InvalidNameError{},
		},
		// "invalidName2": {
		// 	email: email,
		// 	firstName: "firstName#",
		// 	lastName:  "lastName",
		// 	birthDate: birthDate,
		// 	phone:     "1234567890",
		// 	address:   address,
		// 	expected:  &erros.InvalidNameError{},
		// },
		"invalidLastName1": {
			email:     email,
			firstName: "firstName",
			lastName:  "lastName1",
			birthDate: birthDate,
			phone:     "1234567890",
			address:   address,
			expected:  &erros.InvalidNameError{},
		},
		// "invalidLastName2": {
		// 	email: email,
		// 	firstName: "firstName",
		// 	lastName:  "lastName#",
		// 	birthDate: birthDate,
		// 	phone:     "1234567890",
		// 	address:   address,
		// 	expected:  &erros.InvalidNameError{},
		// },
	}
	for k, tt := range data {
		t.Run(tt.firstName, func(t *testing.T) {
			_, err := identity.NewPerson(tt.email, tt.firstName, tt.lastName, tt.phone, tt.birthDate, tt.address)
			if err == nil {
				t.Errorf("expected error, got nil")
			}
			if tt.expected != nil {
				if !errors.As(err, &tt.expected) {
					t.Errorf("test %v expected %v, got %v", k, tt.expected, err)
				}
			}

		})
	}
}
