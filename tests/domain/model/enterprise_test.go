package model_test

import (
	"context"
	"errors"
	"go-complaint/domain/model/common"
	"go-complaint/domain/model/enterprise"
	"go-complaint/erros"
	"go-complaint/tests"
	"testing"
	"time"
)

func TestEnterprise_VALID(t *testing.T) {
	id := "owner@gmail.com"
	address, err := common.NewAddress("country", "county", "city")
	if err != nil {
		t.Errorf("expected nil error, got %v", err)
	}
	sameDate := common.NewDate(time.Now())
	industry, err := enterprise.NewIndustry(1, "industry")
	if err != nil {
		t.Errorf("expected nil error, got %v", err)
	}
	valueMap := map[string]struct {
		id             string
		name           string
		website        string
		bannerIMG      string
		logoIMG        string
		email          string
		phone          string
		address        common.Address
		industry       enterprise.Industry
		registerAt     common.Date
		foundationDate common.Date
	}{
		"basic": {
			id:             id,
			name:           "name",
			logoIMG:        "logo.png",
			bannerIMG:      "/banner.jpg",
			website:        "website",
			email:          "email@gmail.com",
			phone:          "0123456789",
			address:        address,
			industry:       industry,
			registerAt:     sameDate,
			foundationDate: sameDate,
		},
	}
	for _, data := range valueMap {
		_, err := enterprise.NewEnterprise(
			data.id,
			data.name,
			data.logoIMG,
			data.bannerIMG,
			data.website,
			data.email,
			data.phone,
			data.address,
			data.industry,
			data.registerAt,
			data.foundationDate,
		)
		if err != nil {
			t.Errorf("expected nil error, got %v", err)
		}
	}
}

func TestEnterprise_INVALID(t *testing.T) {
	id := "owner@gmail.com"
	address, err := common.NewAddress("country", "county", "city")
	if err != nil {
		t.Errorf("expected nil error, got %v", err)
	}
	sameDate := common.NewDate(time.Now())
	industry, err := enterprise.NewIndustry(1, "industry")
	if err != nil {
		t.Errorf("expected nil error, got %v", err)
	}
	valueMap := map[string]struct {
		id             string
		name           string
		logoIMG        string
		bannerIMG      string
		website        string
		email          string
		phone          string
		address        common.Address
		industry       enterprise.Industry
		registerAt     common.Date
		foundationDate common.Date
		expected       interface{}
	}{
		"empty name": {
			id:             id,
			name:           "",
			logoIMG:        "logo.png",
			bannerIMG:      "/banner.jpg",
			website:        "website",
			email:          "website@gmail.com",
			phone:          "0123456789",
			address:        address,
			industry:       industry,
			registerAt:     sameDate,
			foundationDate: sameDate,
			expected:       &erros.NullValueError{},
		},
		"empty website": {
			id:             id,
			name:           "name",
			logoIMG:        "logo.png",
			bannerIMG:      "/banner.jpg",
			website:        "",
			email:          "website@gmail.com",
			phone:          "0123456789",
			address:        address,
			industry:       industry,
			registerAt:     sameDate,
			foundationDate: sameDate,
			expected:       &erros.NullValueError{},
		},
		"empty email": {
			id:             id,
			name:           "name",
			logoIMG:        "logo.png",
			bannerIMG:      "/banner.jpg",
			website:        "website",
			email:          "",
			phone:          "0123456789",
			address:        address,
			industry:       industry,
			registerAt:     sameDate,
			foundationDate: sameDate,
			expected:       &erros.NullValueError{},
		},
		"empty phone": {
			id:             id,
			name:           "name",
			logoIMG:        "logo.png",
			bannerIMG:      "/banner.jpg",
			website:        "website",
			email:          "website@gmail.com",
			phone:          "",
			address:        address,
			industry:       industry,
			registerAt:     sameDate,
			foundationDate: sameDate,
			expected:       &erros.NullValueError{},
		},
		"shortName": {
			id:             id,
			name:           "n",
			logoIMG:        "logo.png",
			bannerIMG:      "/banner.jpg",
			website:        "website",
			email:          "website@gmail.com",
			phone:          "0123456789",
			address:        address,
			industry:       industry,
			registerAt:     sameDate,
			foundationDate: sameDate,
			expected:       &erros.InvalidLengthError{},
		},
		"largeName": {
			id:             id,
			name:           tests.Repeat("a", 121),
			logoIMG:        "logo.png",
			bannerIMG:      "/banner.jpg",
			website:        "website",
			email:          "website@gmail.com",
			phone:          "0123456789",
			address:        address,
			industry:       industry,
			registerAt:     sameDate,
			foundationDate: sameDate,
			expected:       &erros.InvalidLengthError{},
		},
		"invalidEmail": {
			id:             id,
			name:           "name",
			logoIMG:        "logo.png",
			bannerIMG:      "/banner.jpg",
			website:        "website",
			email:          "website",
			phone:          "0123456789",
			address:        address,
			industry:       industry,
			registerAt:     sameDate,
			foundationDate: sameDate,
			expected:       &erros.InvalidEmailError{},
		},
		"shortPhone": {
			id:             id,
			name:           "name",
			logoIMG:        "logo.png",
			bannerIMG:      "/banner.jpg",
			website:        "website",
			email:          "website@gmail.com",
			phone:          "012345678",
			address:        address,
			industry:       industry,
			registerAt:     sameDate,
			foundationDate: sameDate,
			expected:       &erros.InvalidLengthError{},
		},
		"largePhone": {
			id:             id,
			name:           "name",
			logoIMG:        "logo.png",
			bannerIMG:      "/banner.jpg",
			website:        "website",
			email:          "website@gmail.com",
			phone:          tests.Repeat("1", 22),
			address:        address,
			industry:       industry,
			registerAt:     sameDate,
			foundationDate: sameDate,
			expected:       &erros.InvalidLengthError{},
		},
	}
	for _, data := range valueMap {
		_, err := enterprise.NewEnterprise(
			data.id,
			data.name,
			data.logoIMG,
			data.bannerIMG,
			data.website,
			data.email,
			data.phone,
			data.address,
			data.industry,
			data.registerAt,
			data.foundationDate,
		)
		if err == nil {
			t.Errorf("expected %v error, got nil", data.expected)
		}
		if data.expected != nil {
			if !errors.As(err, &data.expected) {
				t.Errorf("expected %v, got %v", data.expected, err)
			}
		}

	}
}

func TestEnterprise(t *testing.T) {

	ctx := context.Background()
	ownerID := "owner@gmail.com"
	address, err := common.NewAddress("country", "county", "city")
	if err != nil {
		t.Errorf("expected nil error, got %v", err)
	}
	sameDate := common.NewDate(time.Now())
	industry, err := enterprise.NewIndustry(1, "industry")
	if err != nil {
		t.Errorf("expected nil error, got %v", err)
	}
	newEnterprise, err := enterprise.NewEnterprise(
		ownerID,
		"KirkShouwnbedm",
		"logo.png",
		"banner.jpg",
		"website.com",
		"admin@kirkshownbedm.com",
		"0123456789",
		address,
		industry,
		sameDate,
		sameDate,
	)
	t.Run("The logoIMG is changed and an event is sent to the event bus",
		func(t *testing.T) {
			newLogoIMG := "newLogo.png"
			err = newEnterprise.ChangeLogoIMG(ctx, newLogoIMG)
			if err != nil {
				t.Errorf("expected nil error, got %v", err)
			}
			if newEnterprise.LogoIMG() != newLogoIMG {
				t.Errorf("expected %v, got %v", newLogoIMG, newEnterprise.LogoIMG())
			}

		})
	t.Run("The logoIMG is changed to an invalid value it returns and error, don't publish an event",
		func(t *testing.T) {
			newLogoIMG := ""
			err = newEnterprise.ChangeLogoIMG(ctx, newLogoIMG)
			if err == nil {
				t.Errorf("expected error, got nil")
			}
			expectedError := &erros.NullValueError{}
			if !errors.As(err, &expectedError) {
				t.Errorf("expected %v, got %v", &erros.NullValueError{}, err)
			}

		})
	t.Run("Change the email and publish the update event",
		func(t *testing.T) {
			newEmail := "Kualtum@kualtum.com"
			err = newEnterprise.ChangeEmail(ctx, newEmail)
			if err != nil {
				t.Errorf("expected nil error, got %v", err)
			}
			if newEnterprise.Email() != newEmail {
				t.Errorf("expected %v, got %v", newEmail, newEnterprise.Email())
			}

		})
	t.Run("Change to an invalid email and return an error don't publish an event",
		func(t *testing.T) {
			newInvalidEmail := "kualtum"
			err = newEnterprise.ChangeEmail(ctx, newInvalidEmail)
			if err == nil {
				t.Errorf("expected error, got nil")
			}
			expectedError := &erros.InvalidEmailError{}
			if !errors.As(err, &expectedError) {
				t.Errorf("expected %v, got %v", &erros.InvalidEmailError{}, err)
			}

		})
	t.Run("Change the website and publish the update event",
		func(t *testing.T) {
			newWebsite := "kualtum.com"
			err = newEnterprise.ChangeWebsite(ctx, newWebsite)
			if err != nil {
				t.Errorf("expected nil error, got %v", err)
			}
			if newEnterprise.Website() != newWebsite {
				t.Errorf("expected %v, got %v", newWebsite, newEnterprise.Website())
			}

		})
	t.Run("Change to an invalid website and return an error don't publish an event",
		func(t *testing.T) {
			newInvalidWebsite := ""
			err = newEnterprise.ChangeWebsite(ctx, newInvalidWebsite)
			if err == nil {
				t.Errorf("expected error, got nil")
			}
			expectedError := &erros.NullValueError{}
			if !errors.As(err, &expectedError) {
				t.Errorf("expected %v, got %v", &erros.NullValueError{}, err)
			}

		})
	t.Run("Change the phone and publish the update event",
		func(t *testing.T) {
			newPhone := "9876543210"
			err = newEnterprise.ChangePhone(ctx, newPhone)
			if err != nil {
				t.Errorf("expected nil error, got %v", err)
			}
			if newEnterprise.Phone() != newPhone {
				t.Errorf("expected %v, got %v", newPhone, newEnterprise.Phone())
			}

		})
	t.Run("Change to an invalid phone and return an error don't publish an event",
		func(t *testing.T) {
			newInvalidPhone := ""
			err = newEnterprise.ChangePhone(ctx, newInvalidPhone)
			if err == nil {
				t.Errorf("expected error, got nil")
			}
			expectedError := &erros.NullValueError{}
			if !errors.As(err, &expectedError) {
				t.Errorf("expected %v, got %v", &erros.NullValueError{}, err)
			}

		})

	t.Run("Change the address and publish the update event",
		func(t *testing.T) {
			newAddress, err := common.NewAddress("newCountry", "newCounty", "newCity")
			if err != nil {
				t.Errorf("expected nil error, got %v", err)
			}
			err = newEnterprise.ChangeAddress(ctx, "newCountry", "newCounty", "newCity")
			if err != nil {
				t.Errorf("expected nil error, got %v", err)
			}
			if newEnterprise.Address() != newAddress {
				t.Errorf("expected %v, got %v", newAddress, newEnterprise.Address())
			}

		})
	t.Run("Change to an invalid address and return an error don't publish an event",
		func(t *testing.T) {
			err = newEnterprise.ChangeAddress(ctx, "", "", "")
			if err == nil {
				t.Errorf("expected error, got nil")
			}
			expectedError := &erros.NullValueError{}
			if !errors.As(err, &expectedError) {
				t.Errorf("expected %v, got %v", &erros.NullValueError{}, err)
			}

		})
}
