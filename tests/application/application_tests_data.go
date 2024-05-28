package application_test

import (
	"go-complaint/domain/model/common"
	"go-complaint/tests"
	"time"
)

var complaintInfo = map[string]struct {
	authorID    string
	receiverID  string
	title       string
	description string
	content     string
}{
	"userEnterprise": {
		authorID:    "worschwitz@live.com",
		receiverID:  "enterprise",
		title:       tests.Repeat("s", 30),
		description: tests.Repeat("b", 50),
		content:     tests.Repeat("c", 100),
	},
	"userUser": {
		authorID:    "worschwitz@live.com",
		receiverID:  "enterprise@hotmail.com",
		title:       tests.Repeat("s", 30),
		description: tests.Repeat("b", 50),
		content:     tests.Repeat("c", 100),
	},
	"enterpriseUser": {
		authorID:    "enterprise",
		receiverID:  "worschwitz@live.com",
		title:       tests.Repeat("s", 30),
		description: tests.Repeat("b", 50),
		content:     tests.Repeat("c", 100),
	},
	"enterpriseEnterprise": {
		authorID:    "enterprise",
		receiverID:  "OtherEnterprise",
		title:       tests.Repeat("s", 30),
		description: tests.Repeat("b", 50),
		content:     tests.Repeat("c", 100),
	},
}
var usersInfo = map[string]struct {
	profileImage string
	email        string
	password     string
	firstName    string
	lastName     string
	birthDate    string
	phone        string
	country      string
	county       string
	city         string
}{
	"owner": {
		profileImage: "owner.jpg",
		email:        "owner23@gmail.com",
		password:     "Password1",
		firstName:    "Owner",
		lastName:     "Ownership",
		birthDate:    common.StringDate(time.Now()),
		phone:        "01234567890",
		country:      "Country",
		county:       "County",
		city:         "City",
	},
	"user": {
		profileImage: "user.jpg",
		email:        "user35@gmail.com",
		password:     "Password1",
		firstName:    "User",
		lastName:     "User",
		birthDate:    common.StringDate(time.Now()),
		phone:        "01234567890",
		country:      "Country",
		county:       "County",
		city:         "City",
	},
	"employee": {
		profileImage: "employee.jpg",
		email:        "employee35@gmail.com",
		password:     "Password1",
		firstName:    "Employee",
		lastName:     "Employee",
		birthDate:    common.StringDate(time.Now()),
		phone:        "01234567890",
		country:      "Country",
		county:       "County",
		city:         "City",
	},
}
var userInfo = map[string]struct {
	ProfileIMG string
	Email      string
	Password   string
	FirstName  string
	LastName   string
	BirthDate  time.Time
	Phone      string
	Country    string
	County     string
	City       string
}{
	"user1": {
		ProfileIMG: "profile.jpg",
		Email:      "email@gmail.com",
		Password:   "Password1",
		FirstName:  "John",
		LastName:   "Doe",
		BirthDate:  time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
		Phone:      "12345678900",
		Country:    "Country",
		County:     "County",
		City:       "City",
	},
	"user2": {
		ProfileIMG: "profile2.jpg",
		Email:      "user2@gmail.com",
		Password:   "Password2",
		FirstName:  "Jane",
		LastName:   "Doe",
		BirthDate:  time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
		Phone:      "12345678900",
		Country:    "Country",
		County:     "County",
		City:       "City",
	},
	"user3": {
		ProfileIMG: "profile3.jpg",
		Email:      "user543@gmail.com",
		Password:   "Password3",
		FirstName:  "John",
		LastName:   "Doe",
		BirthDate:  time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
		Phone:      "12345678900",
		Country:    "Country",
		County:     "County",
		City:       "City",
	},
}

var enterprisesInfo = map[string]struct {
	ownerID        string
	enterpriseName string
	website        string
	email          string
	phone          string
	country        string
	county         string
	city           string
	industryName   string
	foundationDate string
}{
	"enterprise1": {
		ownerID:        "ownerID@gmail.com",
		enterpriseName: "enterprise1",
		website:        "www.enterprise1.com",
		email:          "enterpriseEmail@gmail.com",
		phone:          "12345678900",
		country:        "Country",
		county:         "County",
		city:           "City",
		industryName:   "Industry",
		foundationDate: common.StringDate(time.Now()),
	},
}
