package persistence_test

import (
	"context"
	"go-complaint/domain/model/common"
	"go-complaint/domain/model/identity"
	"go-complaint/infrastructure"
	"go-complaint/infrastructure/persistence/datasource"
	"go-complaint/infrastructure/persistence/repositories"
	"log"
	"testing"
	"time"
)

func TestUserRepository(t *testing.T) {
	schema := datasource.IdentityAndAccessSchema()

	ctx := context.Background()
	err := schema.Connect(ctx)
	if err != nil {
		t.Errorf("error at connect schema, got %v", err)
	}
	repository := repositories.NewUserRepository(schema)
	//The aggregate has the identity of the root entity
	address, err := common.NewAddress("country", "county", "city")
	if err != nil {
		t.Errorf("NewAddress() got = %v, want nil", err)
	}
	email := "bgas@gmail.com"
	person, err := identity.NewPerson(email, "firstName", "lastName", "01234567890", common.NewDate(time.Now()), address)
	if err != nil {
		t.Errorf("expected nil error, got %v", err)
	}
	password := "Password1"
	encryptionService := infrastructure.NewEncryptionService()
	hash, err := encryptionService.Encrypt(password)
	if err != nil {
		t.Errorf("expected nil error, got %v", err)
	}
	user, err := identity.NewUser(
		"profileIMG",
		common.NewDate(time.Now()),
		email,
		string(hash),
		person,
	)
	if err != nil {
		t.Errorf("expected nil error, got %v", err)
	}
	ownership := identity.NewRole(identity.OWNER)
	firstRole := identity.NewRole(identity.ASSISTANT)
	secondrole := identity.NewRole(identity.MANAGER)

	t.Run("Remove", func(t *testing.T) {
		err = repository.Remove(ctx, email)
		if err != nil {
			t.Errorf("expected nil error, got %v", err)
		}
	})

	t.Run("Save", func(t *testing.T) {
		err = repository.Save(ctx, user)
		if err != nil {
			t.Errorf("expected nil error, got %v", err)
		}
		_, err = repository.Get(ctx, email)
		if err != nil {
			t.Errorf("expected nil error, got %v", err)
		}
		err = repository.Remove(ctx, email)
		if err != nil {
			t.Errorf("expected nil error, got %v", err)
		}
	})

	t.Run("Update", func(t *testing.T) {
		err = repository.Save(ctx, user)
		if err != nil {
			t.Errorf("expected nil error, got %v", err)
		}
		dbUser, err := repository.Get(ctx, email)
		if err != nil {
			t.Errorf("expected nil error, got %v", err)
		}
		log.Printf("OLDPASSWORD %s", dbUser.Password())
		newHash, err := infrastructure.NewEncryptionService().Encrypt("newPassword2")
		if err != nil {
			t.Errorf("expected nil error, got %v", err)
		}
		err = dbUser.ChangePassword(ctx, password, "newPassword2", string(newHash))
		if err != nil {
			t.Errorf("expected nil error, got %v", err)
		}
		err = dbUser.AddRole(ctx, ownership, "enterpriseID")
		if err != nil {
			t.Errorf("expected nil error, got %v", err)
		}
		err = dbUser.AddRole(ctx, firstRole, "enterpriseName")
		if err != nil {
			t.Errorf("expected nil error, got %v", err)
		}
		err = dbUser.ChangePersonalData(ctx,
			"NewImage",
			"newFirstName",
			"newLastName",
			"012345678901",
			"country",
			"province",
			"city")
		if err != nil {
			t.Errorf("expected nil error, got %v", err)
		}
		err = dbUser.AddRole(ctx, secondrole, "enterpriseID")
		if err != nil {
			t.Errorf("expected nil error, got %v", err)
		}
		err = repository.Update(ctx, dbUser)
		if err != nil {
			t.Errorf("expected nil error, got %v", err)
		}
		userDB, err := repository.Get(ctx, email)
		if err != nil {
			t.Errorf("expected nil error, got %v", err)
		}
		count := userDB.UserRoles().Cardinality()
		if count != 3 {
			t.Errorf("expected 3, got %v", count)
		}
		if userDB.ProfileIMG() != "NewImage" {
			t.Errorf("expected NewImage, got %v", userDB.ProfileIMG())
		}
		if userDB.Person().FirstName() != "newFirstName" {
			t.Errorf("expected newFirstName, got %v", userDB.Person().FirstName())
		}
		if userDB.Person().LastName() != "newLastName" {
			t.Errorf("expected newLastName, got %v", userDB.Person().LastName())
		}
		if userDB.Person().Phone() != "012345678901" {
			t.Errorf("expected 012345678901, got %v", userDB.Person().Phone())
		}
		if userDB.Person().Address().Country() != "country" {
			t.Errorf("expected country, got %v", userDB.Person().Address().Country())
		}
		if userDB.Person().Address().County() != "province" {
			t.Errorf("expected province, got %v", userDB.Person().Address().County())
		}
		if userDB.Person().Address().City() != "city" {
			t.Errorf("expected city, got %v", userDB.Person().Address().City())
		}
		log.Printf("USER DB PASSWORD %s", userDB.Password())
		err = repository.Remove(ctx, email)
		if err != nil {
			t.Errorf("expected nil error, got %v", err)
		}

	})
	t.Run("Get with value objects", func(t *testing.T) {
		err = repository.Save(ctx, user)
		if err != nil {
			t.Errorf("expected nil error, got %v", err)
		}
		dbUser, err := repository.Get(ctx, email)
		if err != nil {
			t.Errorf("expected nil error, got %v", err)
		}
		dbUser.AddRole(ctx, ownership, "enterpriseID")
		dbUser.AddRole(ctx, firstRole, "enterpriseName")
		dbUser.AddRole(ctx, secondrole, "enterpriseID")
		err = repository.Update(ctx, dbUser)
		if err != nil {
			t.Errorf("expected nil error, got %v", err)
		}
		dbUser, err = repository.Get(ctx, email)
		if err != nil {
			t.Errorf("expected nil error, got %v", err)
		}
		count := dbUser.UserRoles().Cardinality()
		log.Printf("count: %v", dbUser.UserRoles())
		if count != 3 {
			t.Errorf("expected 3, got %v", count)
		}
		err = repository.Remove(ctx, email)
		if err != nil {
			t.Errorf("expected nil error, got %v", err)
		}
	})

}

func TestUserRepository_MULTIPLE(t *testing.T) {
	schema := datasource.IdentityAndAccessSchema()

	ctx := context.Background()
	err := schema.Connect(ctx)
	if err != nil {
		t.Errorf("error at connect schema, got %v", err)
	}
	repository := repositories.NewUserRepository(schema)
	address, err := common.NewAddress("country", "county", "city")
	if err != nil {
		t.Errorf("NewAddress() got = %v, want nil", err)
	}
	email := "bgas@gmail.com"
	email2 := "vderdmon@gmail.com"
	person1, err := identity.NewPerson(email, "firstName", "lastName", "01234567890", common.NewDate(time.Now()), address)
	if err != nil {
		t.Errorf("expected nil error, got %v", err)
	}
	person2, err := identity.NewPerson(email2, "firstName", "lastName", "01234567890", common.NewDate(time.Now()), address)
	if err != nil {
		t.Errorf("expected nil error, got %v", err)
	}
	password := "Password1"
	encryptionService := infrastructure.NewEncryptionService()
	hash, err := encryptionService.Encrypt(password)
	if err != nil {
		t.Errorf("expected nil error, got %v", err)
	}
	user1, err := identity.NewUser(
		"profileIMG",
		common.NewDate(time.Now()),
		email,
		string(hash),
		person1,
	)
	if err != nil {
		t.Errorf("expected nil error, got %v", err)
	}
	user2, err := identity.NewUser(
		"profileIMG",
		common.NewDate(time.Now()),
		email2,
		string(hash),
		person2,
	)
	if err != nil {
		t.Errorf("expected nil error, got %v", err)
	}
	t.Run("Save and recover trough GetAll without user roles",
		func(t *testing.T) {
			err = repository.Save(ctx, user1)
			if err != nil {
				t.Errorf("expected nil error, got %v", err)
			}
			err = repository.Save(ctx, user2)
			if err != nil {
				t.Errorf("expected nil error, got %v", err)
			}
			users, err := repository.GetAll(ctx)
			if err != nil {
				t.Errorf("expected nil error, got %v", err)
			}
			if len(users) < 2 {
				t.Errorf("expected 2, got %v", len(users))
			}
		})
	t.Cleanup(func() {
		err = repository.Remove(ctx, email)
		if err != nil {
			t.Errorf("expected nil error, got %v", err)
		}
		err = repository.Remove(ctx, email2)
		if err != nil {
			t.Errorf("expected nil error, got %v", err)
		}
	})
}
