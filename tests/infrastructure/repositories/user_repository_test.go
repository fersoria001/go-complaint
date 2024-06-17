package repositories_test

import (
	"context"
	"go-complaint/domain/model/identity"
	"go-complaint/infrastructure/persistence/datasource"
	"go-complaint/infrastructure/persistence/repositories"
	"go-complaint/tests"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserSave(t *testing.T) {
	// Arrange
	ctx := context.Background()
	schema := datasource.PublicSchema()
	userRepository := repositories.NewUserRepository(schema)
	// Act
	err := userRepository.Save(ctx, tests.UserClient)
	err1 := userRepository.Save(ctx, tests.UserAssistant)
	err2 := userRepository.Save(ctx, tests.UserManager)
	err3 := userRepository.Save(ctx, tests.UserOwner)
	// Assert
	assert.Nil(t, err)
	assert.Nil(t, err1)
	assert.Nil(t, err2)
	assert.Nil(t, err3)
}

func TestUserGet(t *testing.T) {
	// Arrange
	ctx := context.Background()
	schema := datasource.PublicSchema()
	userRepository := repositories.NewUserRepository(schema)
	// Act
	dbUser, err := userRepository.Get(ctx, tests.UserClient.Email())
	_, err1 := userRepository.Get(ctx, tests.UserAssistant.Email())
	_, err2 := userRepository.Get(ctx, tests.UserManager.Email())
	_, err3 := userRepository.Get(ctx, tests.UserOwner.Email())
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, tests.UserClient.Email(), dbUser.Email())
	assert.Equal(t, tests.UserClient.FirstName(), dbUser.FirstName())
	assert.Equal(t, tests.UserClient.LastName(), dbUser.LastName())
	assert.Equal(t, tests.UserClient.Password(), dbUser.Password())
	assert.Equal(t, tests.UserClient.BirthDate().StringRepresentation(), dbUser.BirthDate().StringRepresentation())
	assert.Equal(t, tests.UserClient.IsConfirmed(), dbUser.IsConfirmed())
	assert.Equal(t, tests.UserClient.Phone(), dbUser.Phone())
	assert.Equal(t, tests.UserClient.Address(), dbUser.Address())
	assert.Equal(t, tests.UserClient.Gender(), dbUser.Gender())
	assert.Equal(t, tests.UserClient.Pronoun(), dbUser.Pronoun())
	assert.Equal(t, tests.UserClient.ProfileIMG(), dbUser.ProfileIMG())
	t.Log(dbUser.UserRoles())
	assert.Nil(t, err1)
	assert.Nil(t, err2)
	assert.Nil(t, err3)
}

func TestUserUpdate_AddRole(t *testing.T) {
	// Arrange
	ctx := context.Background()
	schema := datasource.PublicSchema()
	userRepository := repositories.NewUserRepository(schema)
	// Act
	user, err := userRepository.Get(ctx, tests.UserClient.Email())
	user.AddRole(
		ctx,
		identity.ASSISTANT,
		tests.EnterpriseID,
	)
	err1 := userRepository.Update(ctx, user)
	user1, err2 := userRepository.Get(ctx, tests.UserClient.Email())
	// Assert
	assert.Nil(t, err)
	assert.Nil(t, err1)
	assert.Nil(t, err2)
	assert.NotNil(t, user1)
	assert.Equal(t, 1, user1.UserRoles().Cardinality())
}

func TestUserUpdate_RemoveRole(t *testing.T) {
	// Arrange
	ctx := context.Background()
	schema := datasource.PublicSchema()
	userRepository := repositories.NewUserRepository(schema)
	// Act
	user, err := userRepository.Get(ctx, tests.UserClient.Email())
	user.RemoveUserRole(
		ctx,
		identity.ASSISTANT,
		tests.EnterpriseID,
	)
	err1 := userRepository.Update(ctx, user)
	user1, err2 := userRepository.Get(ctx, tests.UserClient.Email())
	// Assert
	assert.Nil(t, err)
	assert.Nil(t, err1)
	assert.Nil(t, err2)
	assert.NotNil(t, user1)
	assert.Equal(t, 0, user1.UserRoles().Cardinality())
}
