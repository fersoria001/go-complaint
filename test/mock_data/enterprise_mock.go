package mock_data

import (
	"go-complaint/domain/model/common"
	"go-complaint/domain/model/enterprise"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/google/uuid"
)

type IndustryMock struct {
	Id   int
	Name string
}

type EnterpriseMock struct {
	Id             uuid.UUID
	OwnerId        uuid.UUID
	Name           string
	LogoImg        string
	BannerImg      string
	Website        string
	Email          string
	Phone          string
	Address        common.Address
	Industry       IndustryMock
	RegisterAt     common.Date
	UpdatedAt      common.Date
	FoundationDate common.Date
	Employees      mapset.Set[*enterprise.Employee]
}

type EmployeeMock struct {
	Id               uuid.UUID
	EnterpriseId     uuid.UUID
	User             *UserMock
	HiringDate       common.Date
	ApprovedHiring   bool
	ApprovedHiringAt common.Date
	Position         enterprise.Position
}

type RegisterEnterpriseCommandMock struct {
	Id             uuid.UUID
	Owner          *UserMock
	Name           string
	LogoImg        string
	BannerImg      string
	Website        string
	Email          string
	Phone          string
	Address        common.Address
	Industry       IndustryMock
	RegisterAt     common.Date
	UpdatedAt      common.Date
	FoundationDate common.Date
	Employees      mapset.Set[*enterprise.Employee]
}

type EnterpriseActivityMock struct {
	Id             uuid.UUID
	UserId         uuid.UUID
	ActivityId     uuid.UUID
	EnterpriseId   uuid.UUID
	EnterpriseName string
	OccurredOn     time.Time
	ActivityType   enterprise.EnterpriseActivityType
}

var NewEnterpriseActivity = map[string]*EnterpriseActivityMock{
	"valid": {
		Id:             uuid.MustParse("b2a3b2c7-f0e0-4d19-b12e-ac7821b4a302"),
		UserId:         NewRecipients["user"].Id,
		ActivityId:     NewFeedbacks["valid"].Id,
		EnterpriseId:   NewEnterprises["valid"].Id,
		EnterpriseName: NewEnterprises["valid"].Name,
		OccurredOn:     CommonDate.Date(),
		ActivityType:   enterprise.FeedbacksStarted,
	},
}

var NewEnterprises = map[string]*EnterpriseMock{
	"valid": {
		Id:        uuid.MustParse("2ab7c1d3-f0e0-4d19-b12e-ac7821b4a302"),
		OwnerId:   uuid.MustParse("0d3baf1e-421c-448b-a784-78b210f42e1b"),
		Name:      "EnterpriseName",
		LogoImg:   "/default.jpg",
		BannerImg: "/default.jpg",
		Website:   "https://www.website.com",
		Email:     "enterpriseEmail@gmail.com",
		Phone:     "12345678910",
		Address:   common.NewAddress(uuid.MustParse("2ab7c1d3-f0e0-4d19-b12e-ac7821b4a302"), Country, CountryState, City),
		Industry: IndustryMock{
			Id:   190,
			Name: "Oil & Gas",
		},
		RegisterAt:     CommonDate,
		UpdatedAt:      CommonDate,
		FoundationDate: CommonDate,
		Employees:      mapset.NewSet[*enterprise.Employee](),
	},
	"valid1": {
		Id:        uuid.MustParse("5e4df321-64b9-42ab-a0cd-b18c07e12a14"),
		OwnerId:   uuid.MustParse("0d3baf1e-421c-448b-a784-78b210f42e2b"),
		Name:      "NameOfTheEnterprise",
		LogoImg:   "/default.jpg",
		BannerImg: "/default.jpg",
		Website:   "https://www.website.com",
		Email:     "enterpriseEmail@gmail.com",
		Phone:     "12345678910",
		Address:   common.NewAddress(uuid.MustParse("5e4df321-64b9-42ab-a0cd-b18c07e12a14"), Country, CountryState, City),
		Industry: IndustryMock{
			Id:   190,
			Name: "Oil & Gas",
		},
		RegisterAt:     CommonDate,
		UpdatedAt:      CommonDate,
		FoundationDate: CommonDate,
		Employees:      mapset.NewSet[*enterprise.Employee](),
	},
}

var NewEmployees = []*EmployeeMock{
	{
		Id:               NewUsers["valid"].Id,
		EnterpriseId:     NewEnterprises["valid"].Id,
		User:             NewUsers["valid"],
		HiringDate:       CommonDate,
		ApprovedHiring:   true,
		ApprovedHiringAt: CommonDate,
		Position:         enterprise.ASSISTANT,
	},
}

var NewRegisterEnterprises = map[string]*RegisterEnterpriseCommandMock{
	"valid": {
		Id:        uuid.MustParse("2ab7c1d3-f0e0-4d19-b12e-ac7821b4a302"),
		Owner:     NewUsers["valid"],
		Name:      "EnterpriseName",
		LogoImg:   "/default.jpg",
		BannerImg: "/default.jpg",
		Website:   "https://www.website.com",
		Email:     "enterpriseEmail@gmail.com",
		Phone:     "12345678910",
		Address:   common.NewAddress(uuid.MustParse("2ab7c1d3-f0e0-4d19-b12e-ac7821b4a302"), Country, CountryState, City),
		Industry: IndustryMock{
			Id:   190,
			Name: "Oil & Gas",
		},
		RegisterAt:     CommonDate,
		UpdatedAt:      CommonDate,
		FoundationDate: CommonDate,
		Employees:      mapset.NewSet[*enterprise.Employee](),
	},
}

type HiringProccessMock struct {
	Id         uuid.UUID
	Enterprise RecipientMock
	User       RecipientMock
	Role       enterprise.Position
	Status     enterprise.HiringProccessStatus
	Reason     string
	EmitedBy   RecipientMock
	OccurredOn time.Time
	LastUpdate time.Time
	UpdatedBy  RecipientMock
	Industry   *IndustryMock
}

var NewHiringProccesses = map[string]HiringProccessMock{
	"valid": {
		Id:         uuid.MustParse("2ab7c1d3-f0e0-4d19-b12e-ac7821b4a301"),
		Enterprise: *NewRecipients["enterprise"],
		User:       *NewRecipients["user"],
		Role:       enterprise.ASSISTANT,
		Status:     enterprise.PENDING,
		Reason:     "",
		EmitedBy:   *NewRecipients["user1"],
		OccurredOn: time.Now(),
		LastUpdate: time.Now(),
		UpdatedBy:  *NewRecipients["user1"],
		Industry:   &NewEnterprises["valid"].Industry,
	},
}
