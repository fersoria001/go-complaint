package model_test

import (
	"go-complaint/domain/model/complaint"
	"go-complaint/erros"
	"go-complaint/tests"
	"testing"
)

func TestRating_VALID(t *testing.T) {
	valueMap := map[string]struct {
		rate    int
		comment string
	}{
		"basic":          {rate: 3, comment: "good service"},
		"minimumComment": {rate: 3, comment: "bad"},
		"maximumComment": {rate: 3, comment: tests.Repeat("a", 250)},
		"minimumRate":    {rate: 0, comment: "good service"},
		"maximumRate":    {rate: 5, comment: "good service"},
	}
	for name, value := range valueMap {
		t.Run(name, func(t *testing.T) {
			_, err := complaint.NewRating(value.rate, value.comment)
			if err != nil {
				t.Errorf("expecting no error, got %v", err)
			}
		})
	}
}

func TestRating_INVALID(t *testing.T) {
	valueMap := map[string]struct {
		rate     int
		comment  string
		expected interface{}
	}{
		"longComment":  {rate: 3, comment: tests.Repeat("a", 251), expected: &erros.ValidationError{}},
		"negativeRate": {rate: -1, comment: "good service", expected: &erros.ValidationError{}},
		"highRate":     {rate: 6, comment: "good service", expected: &erros.ValidationError{}},
	}
	for name, value := range valueMap {
		t.Run(name, func(t *testing.T) {
			_, err := complaint.NewRating(value.rate, value.comment)
			if err == nil {
				t.Errorf("expecting an error, got nil")
			}
		})
	}
}
