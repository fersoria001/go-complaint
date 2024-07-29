package mock_data

import (
	"go-complaint/domain/model/common"
	"go-complaint/domain/model/identity"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/google/uuid"
)

type RoleMock struct {
	Role identity.RolesEnum
}

type UserRoleMock struct {
	UserId       uuid.UUID
	EnterpriseId uuid.UUID
	Role         *RoleMock
}

type PersonMock struct {
	Id         uuid.UUID
	Genre      string
	Pronoun    string
	ProfileImg string
	Email      string
	FirstName  string
	LastName   string
	BirthDate  common.Date
	Phone      string
	Address    common.Address
}

type UserMock struct {
	Id           uuid.UUID
	UserName     string
	Password     string
	RegisterDate common.Date
	Person       *PersonMock
	IsConfirmed  bool
	UserRoles    mapset.Set[*identity.UserRole]
	RoleToAdd    *identity.UserRole
}

var (
	AssistantRole, _ = identity.NewRole(identity.ASSISTANT.String())
	ManagerRole, _   = identity.NewRole(identity.MANAGER.String())
	OwnerRole, _     = identity.NewRole(identity.OWNER.String())
	UserRole         = identity.NewUserRole(
		AssistantRole,
		uuid.MustParse("0d3baf1e-421c-448b-a784-78b210f42e1b"),
		uuid.MustParse("2ab7c1d3-f0e0-4d19-b12e-ac7821b4a302"),
	)
)

var NewUsers = map[string]*UserMock{
	"valid": {
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
				uuid.MustParse("0d3baf1e-421c-448b-a784-78b210f42e1b"),
				Country,
				CountryState,
				City,
			),
		},
		IsConfirmed: true,
		UserRoles:   mapset.NewSet[*identity.UserRole](),
		RoleToAdd:   UserRole,
	},
	"valid1": {
		Id:           uuid.MustParse("0d3baf1e-421c-448b-a784-78b210f42e2b"),
		UserName:     "email@gmail.com",
		Password:     "Password1",
		RegisterDate: CommonDate,
		Person: &PersonMock{
			Id:         uuid.MustParse("0d3baf1e-421c-448b-a784-78b210f42e2b"),
			Genre:      "male",
			Pronoun:    "he",
			ProfileImg: "/default.jpg",
			Email:      "email@gmail.com",
			FirstName:  "Fernando Agustin",
			LastName:   "Soria",
			BirthDate:  CommonDate,
			Phone:      "012345678910",
			Address: common.NewAddress(
				uuid.MustParse("0d3baf1e-421c-448b-a784-78b210f42e2b"),
				Country,
				CountryState,
				City,
			),
		},
		IsConfirmed: true,
		UserRoles:   mapset.NewSet[*identity.UserRole](),
		RoleToAdd:   UserRole,
	},
}
