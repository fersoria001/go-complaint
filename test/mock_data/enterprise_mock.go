package mock_data

import (
	"go-complaint/domain/model/common"
	"go-complaint/domain/model/enterprise"
	"go-complaint/domain/model/identity"
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

var NewRegisterEnterprises = map[string]*RegisterEnterpriseCommandMock{
	"valid": {
		Id: uuid.MustParse("2ab7c1d3-f0e0-4d19-b12e-ac7821b4a302"),
		Owner: &UserMock{
			Id:           uuid.MustParse("0d3baf1e-421c-448b-a784-78b210f42e1b"),
			UserName:     "bercho001@gmail.com",
			Password:     "Password1",
			RegisterDate: CommonDate,
			Person: &PersonMock{
				Id:         uuid.MustParse("0d3baf1e-421c-448b-a784-78b210f42e1b"),
				Genre:      "male",
				Pronoun:    "he",
				ProfileImg: "/default.jpg",
				Email:      "bercho001@gmail.com",
				FirstName:  "Fernando Agustin",
				LastName:   "Soria",
				BirthDate:  CommonDate,
				Phone:      "012345678910",
				Address: common.NewAddress(
					uuid.New(),
					Country,
					CountryState,
					City,
				),
			},
			IsConfirmed: true,
			UserRoles:   mapset.NewSet[*identity.UserRole](),
		},
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
	},
}
