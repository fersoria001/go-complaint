package application_test

import (
	"context"
	"go-complaint/application"
	"go-complaint/domain/model/common"
	"go-complaint/infrastructure"
	"go-complaint/infrastructure/persistence/datasource"
	"go-complaint/infrastructure/persistence/repositories"
	"testing"
	"time"
)

var updateUserProfileMap = map[string]struct {
	ProfileIMG string
	FirstName  string
	LastName   string
	Phone      string
	Country    string
	County     string
	City       string
}{
	"only_profile_img": {
		ProfileIMG: "imageNewValuen.jpg",
		FirstName:  "",
		LastName:   "",
		Phone:      "",
		Country:    "",
		County:     "",
		City:       "",
	},
	"only_first_name": {
		ProfileIMG: "",
		FirstName:  "firstNameNewValue",
		LastName:   "",
		Phone:      "",
		Country:    "",
		County:     "",
		City:       "",
	},
	"only_last_name": {
		ProfileIMG: "",
		FirstName:  "",
		LastName:   "lastNameNewValue",
		Phone:      "",
		Country:    "",
		County:     "",
		City:       "",
	},
	"only_phone": {
		ProfileIMG: "",
		FirstName:  "",
		LastName:   "",
		Phone:      "123456789068",
		Country:    "",
		County:     "",
		City:       "",
	},
	"only_country": {
		ProfileIMG: "",
		FirstName:  "",
		LastName:   "",
		Phone:      "",
		Country:    "NewCountryValue",
		County:     "",
		City:       "",
	},
	"only_county": {
		ProfileIMG: "",
		FirstName:  "",
		LastName:   "",
		Phone:      "",
		Country:    "",
		County:     "NewCountyValue",
		City:       "",
	},
	"only_city": {
		ProfileIMG: "",
		FirstName:  "",
		LastName:   "",
		Phone:      "",
		Country:    "",
		County:     "",
		City:       "NewCityValue",
	},
	"all_fields": {
		ProfileIMG: "imageNewValue33.jpg",
		FirstName:  "firstNameNew",
		LastName:   "lastNameNew",
		Phone:      "123456789000",
		Country:    "NewCountry",
		County:     "NewCounty",
		City:       "NewCity",
	},
	"no_fields": {
		ProfileIMG: "",
		FirstName:  "",
		LastName:   "",
		Phone:      "",
		Country:    "",
		County:     "",
		City:       "",
	},
}

func TestIdentityService(t *testing.T) {
	ctx := context.Background()
	schema := datasource.IdentityAndAccessSchema()
	repository := repositories.NewUserRepository(schema)
	identityService := application.NewIdentityService(repository)
	birthDateStr := common.StringDate(time.Date(1993, 9, 7, 6, 5, 4, 3, time.UTC))
	email := "email@gmail.com"
	password := "Password1"
	encryptionService := infrastructure.NewEncryptionService()

	t.Run(`A user must be created and saved in the repository,
	 also an event is sent to the event bus`, func(t *testing.T) {
		err := identityService.RegisterUser(
			ctx,
			"image.jpg",
			email,
			password,
			"firstName",
			"lastName",
			birthDateStr,
			"1234567890",
			"country",
			"county",
			"city",
		)
		if err != nil {
			t.Error(err)
		}
		user, err := repository.Get(ctx, email)
		if err != nil {
			t.Error(err)
		}
		if user == nil {
			t.Error("user not found")
		}

	})
	t.Run(`An user personal information must be updated and saved in the repository,
	it sent an event to the event bus on each execution`, func(t *testing.T) {
		for name, fields := range updateUserProfileMap {
			err := identityService.UpdateUserProfile(
				ctx,
				email,
				fields.ProfileIMG,
				fields.FirstName,
				fields.LastName,
				fields.Phone,
				fields.Country,
				fields.County,
				fields.City,
			)
			if err != nil {
				t.Error(err)
			}
			user, err := repository.Get(ctx, email)
			if err != nil {
				t.Error(err)
			}
			if user == nil {
				t.Error("user not found")
			}
			if fields.ProfileIMG != "" && user.ProfileIMG() != fields.ProfileIMG {
				t.Errorf("profile image not updated, %s", name)
			}
			if fields.FirstName != "" && user.Person().FirstName() != fields.FirstName {
				t.Errorf("first name not updated, %s", name)
			}
			if fields.LastName != "" && user.Person().LastName() != fields.LastName {
				t.Errorf("last name not updated, %s", name)
			}
			if fields.Phone != "" && user.Person().Phone() != fields.Phone {
				t.Errorf("phone not updated, %s", name)
			}
			if fields.Country != "" && user.Person().Address().Country() != fields.Country {
				t.Errorf("country not updated, %s", name)
			}
			if fields.County != "" && user.Person().Address().County() != fields.County {
				t.Errorf("county not updated, %s", name)
			}
			if fields.City != "" && user.Person().Address().City() != fields.City {
				t.Errorf("city not updated, %s", name)
			}

		}
	})

	t.Run(`An user is authenticated, returns a JWT token`,
		func(t *testing.T) {
			token, err := identityService.Login(
				ctx,
				email,
				password,
				false)
			if err != nil {
				t.Error(err)
			}
			if token == "" {
				t.Error("token not generated")
			}

		})
	t.Run(`If the account doesn't exists or the password
	 is incorrect, Login must return an error`,
		func(t *testing.T) {
			_, err := identityService.Login(
				ctx,
				"nonexists@gmail.com",
				password,
				false)
			if err == nil {
				t.Error("error not returned")
			}
			_, err = identityService.Login(
				ctx,
				email,
				"wrongPassword",
				false)
			if err == nil {
				t.Error("error not returned")
			}

		})

	t.Run(`When an user change its password, the password is updated and a new event 
	is sent to the event bus`, func(t *testing.T) {
		newPassword := "newPassword"
		err := identityService.ChangePassword(ctx, email, password, newPassword)
		if err != nil {
			t.Error(err)
		}
		user, err := repository.Get(ctx, email)
		if err != nil {
			t.Error(err)
		}
		if user == nil {
			t.Error("user not found")
		}
		if encryptionService.Compare(user.Password(), newPassword) != nil {
			t.Error("password not updated")
		}

	})

	t.Run(`When a password recovery is triggered,
	the user password change to a random one, and an event is sent
	to the event bus`,
		func(t *testing.T) {
			err := identityService.RecoverPassword(ctx, email)
			if err != nil {
				t.Error(err)
			}
			user, err := repository.Get(ctx, email)
			if err != nil {
				t.Error(err)
			}
			if user == nil {
				t.Error("user not found")
			}
			if user.Password() == password {
				t.Error("password not changed")
			}

		})

	t.Run(`When a role is added to an user,
	 the role is added to the user roles set
	and an event is sent to the event bus`,
		func(t *testing.T) {

			err := identityService.AddNewRole(ctx, email, "OWNER", "enterpriseName")
			if err != nil {
				t.Error(err)
			}

			user, err := repository.Get(ctx, email)
			if err != nil {
				t.Error(err)
			}
			if user == nil {
				t.Error("user not found")
			}
			if user.UserRoles().Cardinality() < 1 {
				t.Error("role not added")
			}
			if len(user.GetAuthorities("enterpriseName")) < 1 {
				t.Error("authority not created")
			}

		})
	t.Run(`More than one role inside the same enterprise can be added to an user`,
		func(t *testing.T) {
			err := identityService.AddNewRole(ctx, email, "MANAGER", "enterpriseName")
			if err != nil {
				t.Error(err)
			}
			user, err := repository.Get(ctx, email)
			if err != nil {
				t.Error(err)
			}
			if user == nil {
				t.Error("user not found")
			}
			if user.UserRoles().Cardinality() < 2 {
				t.Error("role not added")
			}
			if len(user.GetAuthorities("enterpriseName")) < 2 {
				t.Error("authority not created")
			}

		})
	t.Run(`A user can has roles in different enterprises`,
		func(t *testing.T) {
			err := identityService.AddNewRole(ctx, email, "OWNER", "enterpriseName2")
			if err != nil {
				t.Error(err)
			}
			err = identityService.AddNewRole(ctx, email, "MANAGER", "enterpriseName2")
			if err != nil {
				t.Error(err)
			}
			user, err := repository.Get(ctx, email)
			if err != nil {
				t.Error(err)
			}
			if user == nil {
				t.Error("user not found")
			}
			if user.UserRoles().Cardinality() < 4 {
				t.Error("role not added")
			}
			if len(user.GetAuthorities("enterpriseName2")) < 2 {
				t.Error("authority not created")
			}

		})
	t.Run(`When a role is removed from an user,
	 the role is removed from the user roles set
	and an event is sent to the event bus`,
		func(t *testing.T) {
			err := identityService.RemoveRole(ctx, email, "OWNER", "enterpriseName")
			if err != nil {
				t.Error(err)
			}
			user, err := repository.Get(ctx, email)
			if err != nil {
				t.Error(err)
			}
			if user == nil {
				t.Error("user not found")
			}
			if user.UserRoles().Cardinality() > 3 {
				t.Error("role not removed")
			}
			if len(user.GetAuthorities("enterpriseName")) > 1 {
				t.Error("authority not removed")
			}

		})
	t.Run(`Change role must remove the old role and add the new one
	two events are sent to the event bus one for the role removed and
	another for the role added`,
		func(t *testing.T) {
			err := identityService.ChangeRole(ctx, email, "enterpriseName", "MANAGER", "OWNER")
			if err != nil {
				t.Error(err)
			}
			user, err := repository.Get(ctx, email)
			if err != nil {
				t.Error(err)
			}
			if user == nil {
				t.Error("user not found")
			}
			if user.UserRoles().Cardinality() != 3 {
				t.Error("role cardinality is not the same as before")
			}
			if len(user.GetAuthorities("enterpriseName")) != 1 {
				t.Error("authority cardinality is not the same as before")
			}
			if user.GetAuthorities("enterpriseName")[0].Authority() != "OWNER" {
				t.Error("role change not reflect in the authorities")
			}

		})

	t.Cleanup(func() {
		err := repository.Remove(ctx, email)
		if err != nil {
			t.Error(err)
		}
	})
}
