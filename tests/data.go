package tests

import (
	"go-complaint/application/commands"
	"go-complaint/domain/model/common"
	"go-complaint/domain/model/complaint"
	"go-complaint/domain/model/employee"
	"go-complaint/domain/model/enterprise"
	"go-complaint/domain/model/feedback"
	"go-complaint/domain/model/identity"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/google/uuid"
)

var (
	msBirthDate          = common.NewDate(fixedDate).StringRepresentation()
	UserRegisterCommands = map[string]commands.UserCommand{
		"1": {
			Email:          "bercho001@gmail.com",
			Password:       "Password1",
			FirstName:      "Fernando Agustin",
			LastName:       "Soria",
			Gender:         "MALE",
			Pronoun:        "He",
			BirthDate:      msBirthDate,
			PhoneCode:      "54",
			Phone:          "12345678910",
			CountryID:      11,
			CountryStateID: 3634,
			CityID:         644,
		},
		"2": {
			Email:          "fernandosoria1379@gmail.com",
			Password:       "Password1",
			FirstName:      "Fernando Agustin",
			LastName:       "Soria",
			Gender:         "MALE",
			Pronoun:        "He",
			BirthDate:      msBirthDate,
			Phone:          "12345678910",
			CountryID:      11,
			CountryStateID: 3634,
			CityID:         644,
		},
	}
	UserRegisterAndVerifyCommands = map[string]commands.UserCommand{
		"1": {
			Email:          "fernandosoria@gocomplaint.com",
			Password:       "Password1",
			FirstName:      "Fernando Agustin",
			LastName:       "Soria",
			Gender:         "MALE",
			Pronoun:        "He",
			BirthDate:      msBirthDate,
			Phone:          "12345678910",
			CountryID:      11,
			CountryStateID: 3634,
			CityID:         644,
		},
		"2": {
			Email:          "soriafernando@hotmail.com",
			Password:       "Password1",
			FirstName:      "Fernando Agustin",
			LastName:       "Soria",
			Gender:         "MALE",
			Pronoun:        "He",
			BirthDate:      msBirthDate,
			Phone:          "12345678910",
			CountryID:      11,
			CountryStateID: 3634,
			CityID:         644,
		},
		"3": {
			Email:          "soriafernando@hlive.com",
			Password:       "Password1",
			FirstName:      "Fernando Agustin",
			LastName:       "Soria",
			Gender:         "MALE",
			Pronoun:        "He",
			BirthDate:      msBirthDate,
			Phone:          "12345678910",
			CountryID:      11,
			CountryStateID: 3634,
			CityID:         644,
		},
		"4": {
			Email:          "soriaf@hlive.com",
			Password:       "Password1",
			FirstName:      "Fernando Agustin",
			LastName:       "Soria",
			Gender:         "MALE",
			Pronoun:        "He",
			BirthDate:      msBirthDate,
			Phone:          "12345678910",
			CountryID:      11,
			CountryStateID: 3634,
			CityID:         644,
		},
	}
	CreateEnterpriseCommands = map[string]commands.EnterpriseCommand{
		"Spoon company": {
			OwnerID:        UserRegisterAndVerifyCommands["2"].Email,
			Name:           "Spoon company",
			Website:        "www.spooncompany.com",
			Email:          "spoon-company@enterprise.com",
			Phone:          "12345678910",
			CountryID:      11,
			CountryStateID: 3634,
			CityID:         644,
			IndustryID:     97,
			FoundationDate: msBirthDate,
		},
	}
)

var (
	emptyUserRoles  = mapset.NewSet[*identity.UserRole]()
	fixedDate       = time.Date(2020, time.April, 11, 21, 34, 01, 0, time.UTC)
	CommonDate      = common.NewDate(fixedDate)
	country         = common.NewCountry(11, "Argentina", "54")
	county          = common.NewCountryState(3634, "San Juan")
	city            = common.NewCity(644, "Albardón", "AR", -31.43722, -68.52556)
	Address1ID      = uuid.MustParse("0d3baf1e-421c-448b-a784-78b210f42e1b")
	Address2ID      = uuid.MustParse("2ab7c1d3-f0e0-4d19-b12e-ac7821b4a302")
	Address3ID      = uuid.MustParse("f41298ea-7c1b-441a-8a0b-32d01ca27e2b")
	Address4ID      = uuid.MustParse("8c910eaf-b942-430f-b01d-21a38e271d0f")
	Address5ID      = uuid.MustParse("5e4df321-64b9-42ab-a0cd-b18c07e12a14")
	Address6ID      = uuid.MustParse("107a2be9-38e4-4429-87b1-c08d3b1f4823")
	Address1        = common.NewAddress(Address1ID, country, county, city)
	Address2        = common.NewAddress(Address2ID, country, county, city)
	Address3        = common.NewAddress(Address3ID, country, county, city)
	Address4        = common.NewAddress(Address4ID, country, county, city)
	Address5        = common.NewAddress(Address5ID, country, county, city)
	Address6        = common.NewAddress(Address6ID, country, county, city)
	PersonClient, _ = identity.NewPerson(
		"mock-client@user.com",
		"/default.jpg",
		"male",
		"him",
		"mock",
		"user",
		"012345678910",
		CommonDate,
		Address1,
	)
	UserClient, _ = identity.NewUser(
		"mock-client@user.com",
		"Password1",
		common.NewDate(fixedDate),
		PersonClient,
		true,
		emptyUserRoles,
	)
	PersonAssistant, _ = identity.NewPerson(
		"mock-assistant@employee.com",
		"/default.jpg",
		"female",
		"her",
		"mocka",
		"assistant",
		"012345678910",
		CommonDate,
		Address2,
	)
	UserAssistant, _ = identity.NewUser(
		"mock-assistant@employee.com",
		"Password1",
		CommonDate,
		PersonAssistant,
		true,
		emptyUserRoles,
	)
	PersonManager, _ = identity.NewPerson(
		"mock-manager@employee.com",
		"/default.jpg",
		"female",
		"her",
		"mocka",
		"manager",
		"012345678910",
		CommonDate,
		Address3,
	)
	UserManager, _ = identity.NewUser(
		"mock-manager@employee.com",
		"Password1",
		CommonDate,
		PersonManager,
		true,
		emptyUserRoles,
	)
	PersonOwner, _ = identity.NewPerson(
		"mock-owner@owner.com",
		"/default.jpg",
		"female",
		"her",
		"mocka",
		"owner",
		"012345678910",
		CommonDate,
		Address4,
	)
	UserOwner, _ = identity.NewUser(
		"mock-owner@owner.com",
		"Password1",
		CommonDate,
		PersonOwner,
		true,
		emptyUserRoles,
	)
	Industry, _  = enterprise.NewIndustry(190, "Oil & Gas")
	AssistantID  = uuid.MustParse("f7b4e29f-488e-4f1b-88e5-a4a9eebf7c5d")
	ManagerID    = uuid.MustParse("2c1a8e12-7d1d-4d1d-b02e-e2c4f5f4272e")
	EnterpriseID = "FloatingPoint. Ltd"
	Manager, _   = employee.NewEmployee(
		ManagerID,
		EnterpriseID,
		UserManager,
		enterprise.MANAGER,
		CommonDate,
		false,
		CommonDate,
	)
	Asisstant, _ = employee.NewEmployee(
		AssistantID,
		EnterpriseID,
		UserAssistant,
		enterprise.ASSISTANT,
		CommonDate,
		false,
		CommonDate,
	)
	Enterprise1, _ = enterprise.NewEnterprise(
		UserOwner.Email(),
		EnterpriseID,
		"/logo.jpg",
		"/banner.jpg",
		"www.floatingpoint.com",
		"floating-point@enterprise.com",
		"012345678910",
		Address5,
		Industry,
		CommonDate,
		CommonDate,
		CommonDate,
		mapset.NewSet[enterprise.Employee](),
	)
	ComplaintID = uuid.MustParse("31b6617f-578d-4d4e-83f7-02dc399d5738")
	status      = complaint.OPEN
	msg, _      = complaint.NewMessage(
		RepeatString("t", 11),
		RepeatString("d", 31),
		RepeatString("b", 51),
	)
	rating, _     = complaint.NewRating(5, "good")
	emptySet      = mapset.NewSet[*complaint.Reply]()
	Complaint1, _ = complaint.NewComplaint(
		ComplaintID,
		UserClient.Email(),
		UserClient.FullName(),
		UserClient.ProfileIMG(),
		Enterprise1.Name(),
		Enterprise1.Name(),
		Enterprise1.LogoIMG(),
		status,
		msg,
		CommonDate,
		CommonDate,
		rating,
		emptySet,
	)
	ReceiverReply1ID  = uuid.MustParse("b21ea78c-d4f2-4a0e-984b-17f238b12e0a")
	ReceiverReply2ID  = uuid.MustParse("b21ea78c-d4f2-4a0e-984b-17f238b12e0b")
	AuthorReply1ID    = uuid.MustParse("1ca0f32b-87d9-4f1e-a9b2-02c47e8f1b21")
	AuthorReply2ID    = uuid.MustParse("1ca0f32b-87d9-4f1e-a9b2-02c47e8f1b22")
	ReceiverReply1, _ = complaint.NewReply(
		ReceiverReply1ID,
		ComplaintID,
		Asisstant.Email(),
		Asisstant.ProfileIMG(),
		Asisstant.FullName(),
		"reply body",
		false,
		CommonDate,
		CommonDate,
		CommonDate,
		true,
		EnterpriseID,
	)
	ReceiverReply2, _ = complaint.NewReply(
		ReceiverReply2ID,
		ComplaintID,
		Asisstant.Email(),
		Asisstant.ProfileIMG(),
		Asisstant.FullName(),
		"reply body",
		false,
		CommonDate,
		CommonDate,
		CommonDate,
		true,
		EnterpriseID,
	)
	AuthorReply1, _ = complaint.NewReply(
		AuthorReply1ID,
		ComplaintID,
		UserClient.Email(),
		UserClient.ProfileIMG(),
		UserClient.FullName(),
		"reply body",
		false,
		CommonDate,
		CommonDate,
		CommonDate,
		false,
		"",
	)
	AuthorReply2, _ = complaint.NewReply(
		AuthorReply2ID,
		ComplaintID,
		UserClient.Email(),
		UserClient.ProfileIMG(),
		UserClient.FullName(),
		"reply body",
		false,
		CommonDate,
		CommonDate,
		CommonDate,
		false,
		"",
	)
	ReplyReview1ID      = uuid.MustParse("e8742109-a23b-4d42-827e-b98d1c07f1a3")
	ReplyReview1Replies = mapset.NewSet[complaint.Reply](*ReceiverReply1)
	Review1             = feedback.NewReview(
		ComplaintID,
		ReplyReview1ID,
		"nice reply",
	)
	ReplyReview1, _ = feedback.NewReplyReview(
		ReplyReview1ID,
		ComplaintID,
		ReplyReview1Replies,
		*Asisstant.User,
		Review1,
		"#375774",
		time.Now(),
	)
	replyReviews    = mapset.NewSet[*feedback.ReplyReview](ReplyReview1)
	feedbackAnswers = mapset.NewSet[*feedback.Answer]()
	Feedback1ID     = uuid.MustParse("5445ef98-23d0-46a4-8fb4-a55a43b234b7")
	Feedback1, _    = feedback.NewFeedback(
		Feedback1ID,
		ComplaintID,
		EnterpriseID,
		replyReviews,
		feedbackAnswers,
		time.Now(),
		time.Now(),
		false,
	)
)
