package identity_test

import (
	"context"
	"go-complaint/domain/model/common"
	"go-complaint/domain/model/identity"
	"go-complaint/tests"
	"testing"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/stretchr/testify/assert"
)

func TestUser_New(t *testing.T) {
	// Arrange
	// Act
	person, err := identity.NewPerson(
		"mock-user@user.com",
		"/default.jpg",
		"male",
		"him",
		"mock",
		"user",
		"012345678910",
		common.NewDate(time.Now()),
		tests.Address1,
	)
	emptyUserRoles := mapset.NewSet[*identity.UserRole]()
	user, err1 := identity.NewUser(
		"mock-user@user.com",
		"Password1",
		common.NewDate(time.Now()),
		person,
		true,
		emptyUserRoles,
	)
	// Assert
	assert.Nil(t, err)
	assert.Nil(t, err1)
	assert.NotNil(t, user)
}

func TestUser_AddRole(t *testing.T) {
	// Arrange
	ctx := context.Background()
	person, err := identity.NewPerson(
		"mock-user@user.com",
		"/default.jpg",
		"male",
		"him",
		"mock",
		"user",
		"012345678910",
		common.NewDate(time.Now()),
		tests.Address1,
	)
	emptyUserRoles := mapset.NewSet[*identity.UserRole]()
	user, err1 := identity.NewUser(
		"mock-user@user.com",
		"Password1",
		common.NewDate(time.Now()),
		person,
		true,
		emptyUserRoles,
	)
	// Act
	err2 := user.AddRole(
		ctx,
		identity.ASSISTANT,
		"mock-enterprise",
	)
	authorities := user.Authorities()
	// Assert
	assert.Nil(t, err)
	assert.Nil(t, err1)
	assert.NotNil(t, user)
	assert.Nil(t, err2)
	assert.Equal(t, 1, len(authorities))
	assert.Equal(t, 1, len(authorities["mock-enterprise"]))
	assert.Equal(t, identity.ASSISTANT.String(), authorities["mock-enterprise"][0].Authority())
}

func TestUser_RemoveRole(t *testing.T) {
	// Arrange
	ctx := context.Background()
	person, err := identity.NewPerson(
		"mock-user@user.com",
		"/default.jpg",
		"male",
		"him",
		"mock",
		"user",
		"012345678910",
		common.NewDate(time.Now()),
		tests.Address1,
	)
	emptyUserRoles := mapset.NewSet[*identity.UserRole]()
	user, err1 := identity.NewUser(
		"mock-user@user.com",
		"Password1",
		common.NewDate(time.Now()),
		person,
		true,
		emptyUserRoles,
	)
	// Act
	err2 := user.AddRole(
		ctx,
		identity.ASSISTANT,
		"mock-enterprise",
	)
	authorities := user.Authorities()
	err3 := user.RemoveUserRole(
		ctx,
		identity.ASSISTANT,
		"mock-enterprise",
	)
	authorities2 := user.Authorities()
	// Assert
	assert.Nil(t, err)
	assert.Nil(t, err1)
	assert.NotNil(t, user)
	assert.Nil(t, err2)
	assert.Nil(t, err3)
	assert.Equal(t, 1, len(authorities["mock-enterprise"]))
	assert.Equal(t, 0, len(authorities2["mock-enterprise"]))
}
