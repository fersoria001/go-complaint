package model_test

import (
	"errors"
	"go-complaint/domain/model/complaint"
	"go-complaint/erros"
	"testing"
)

func TestNewMessage_VALID(t *testing.T) {
	valuesMap := map[string]struct {
		title       string //more than 10 less than 80
		description string //more than 30 less than 120
		body        string //more than 50 less than 250
	}{
		"basic": {
			title:       "This is a title",
			description: "This is a description that has more than 30 characters",
			body:        "This is a body that has more than 50 characters and is a valid body",
		},
		"minTitle": {
			title:       "This is a ",
			description: "This is a description that has more than 30 characters",
			body:        "This is a body that has more than 50 characters and is a valid body",
		},
		"maxTitle": {
			title:       "This is a title that has no more than eighty characters and is a valid max title",
			description: "This is a description that has more than 30 characters",
			body:        "This is a body that has more than 50 characters and is a valid body",
		},
		"minDescription": {
			title:       "This is a title",
			description: "This is a description that has",
			body:        "This is a body that has more than 50 characters and is a valid body",
		},
		"maxDescription": {
			title:       "This is a title",
			description: "This is a description that has more than 30 characters and is a valid max description because it has less than 120 chars",
			body:        "This is a body that has more than 50 characters and is a valid body",
		},
		"minBody": {
			title:       "This is a title",
			description: "This is a description that has more than 30 characters",
			body:        "This is a body that has more than 50 characters an",
		},
		"maxBody": {
			title:       "This is a title",
			description: "This is a description that has more than 30 characters",
			body:        "this is a body that has more than 50 characters and is a valid body that has more than 250 characters and is a valid max body this needs to be a long string to test the max body length and actually be more than 250 characters long please write more t",
		},
	}
	for name, v := range valuesMap {
		t.Run(name, func(t *testing.T) {
			_, err := complaint.NewMessage(v.title, v.description, v.body)
			if err != nil {
				t.Errorf("Expected no error, got %v at %s", err, name)
			}
		})
	}
}

func TestNewMessage_INVALID(t *testing.T) {
	valuesMap := map[string]struct {
		title       string //more than 10 less than 80
		description string //more than 30 less than 120
		body        string //more than 50 less than 250
		expected    interface{}
	}{
		"emptyTitle": {
			title:       "",
			description: "This is a description that has more than 30 characters",
			body:        "This is a body that has more than 50 characters and is a valid body",
			expected:    &erros.NullValueError{},
		},
		"lessThanMinTitle": {
			title:       "This is a",
			description: "This is a description that has more than 30 characters",
			body:        "This is a body that has more than 50 characters and is a valid body",
			expected:    &erros.InvalidLengthError{},
		},
		"moreThanMaxTitle": {
			title:       "This is a title that has more than eighty characters and is not a valid max title",
			description: "This is a description that has more than 30 characters",
			body:        "This is a body that has more than 50 characters and is a valid body",
			expected:    &erros.InvalidLengthError{},
		},
		"emptyDescription": {
			title:       "This is a title",
			description: "",
			body:        "This is a body that has more than 50 characters and is a valid body",
			expected:    &erros.NullValueError{},
		},
		"lessThanMinDescription": {
			title:       "This is a title",
			description: "This is a description",
			body:        "This is a body that has more than 50 characters and is a valid body",
			expected:    &erros.InvalidLengthError{},
		},
		"moreThanMaxDescription": {
			title:       "This is a title",
			description: "This is a description that has more than 30 characters and is a valid max description because it has more than 120 charss",
			body:        "This is a body that has more than 50 characters and is a valid body",
			expected:    &erros.InvalidLengthError{},
		},
		"emptyBody": {
			title:       "This is a title",
			description: "This is a description that has more than 30 characters",
			body:        "",
			expected:    &erros.NullValueError{},
		},
		"lessThanMinBody": {
			title:       "This is a title",
			description: "This is a description that has more than 30 characters",
			body:        "This is a body that has more than 50 characters a",
			expected:    &erros.InvalidLengthError{},
		},
		"moreThanMaxBody": {
			title:       "This is a title",
			description: "This is a description that has more than 30 characters",
			body:        "this is a body that has more than 50 characters and is a valid body that has more than 250 characters and is a valid max body this needs to be a long string to test the max body length and actually be more than 250 characters long please write more th",
			expected:    &erros.InvalidLengthError{},
		},
	}
	for name, v := range valuesMap {
		t.Run(name, func(t *testing.T) {
			_, err := complaint.NewMessage(v.title, v.description, v.body)
			if err == nil {
				t.Errorf("NewReply() = nil, want %v", v.expected)
			}
			if v.expected != nil {
				if !errors.As(err, &v.expected) {
					t.Errorf("expected %v, got %v", v.expected, err)
				}
			}
		})
	}
}
