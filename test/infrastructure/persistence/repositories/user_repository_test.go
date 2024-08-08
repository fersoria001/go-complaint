package repositories_test

import (
	"context"
	"go-complaint/domain/model/identity"
	"go-complaint/infrastructure/persistence/finders/find_user"
	"go-complaint/infrastructure/persistence/repositories"
	"go-complaint/test/mock_data"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Save(t *testing.T) {
	ctx := context.Background()
	repository, ok := repositories.MapperRegistryInstance().Get("User").(repositories.UserRepository)
	assert.True(t, ok)
	for _, v := range mock_data.NewUsers {
		person, err := identity.NewPerson(
			v.Person.Id,
			v.Person.Email,
			v.Person.ProfileImg,
			v.Person.Genre,
			v.Person.Pronoun,
			v.Person.FirstName,
			v.Person.LastName,
			v.Person.Phone,
			v.Person.BirthDate,
			v.Person.Address,
		)
		assert.Nil(t, err)
		user, err := identity.NewUser(
			v.Id,
			v.UserName,
			v.Password,
			v.RegisterDate,
			person,
			v.IsConfirmed,
			v.UserRoles,
		)
		assert.Nil(t, err)
		err = repository.Save(ctx, user)
		assert.Nil(t, err)
		t.Cleanup(func() {
			err := repository.Remove(ctx, v.Id)
			assert.Nil(t, err)
		})
	}
}

func TestUserRepository_Get(t *testing.T) {
	ctx := context.Background()
	repository, ok := repositories.MapperRegistryInstance().Get("User").(repositories.UserRepository)
	assert.True(t, ok)
	for _, v := range mock_data.NewUsers {
		person, err := identity.NewPerson(
			v.Person.Id,
			v.Person.Email,
			v.Person.ProfileImg,
			v.Person.Genre,
			v.Person.Pronoun,
			v.Person.FirstName,
			v.Person.LastName,
			v.Person.Phone,
			v.Person.BirthDate,
			v.Person.Address,
		)
		assert.Nil(t, err)
		user, err := identity.NewUser(
			v.Id,
			v.UserName,
			v.Password,
			v.RegisterDate,
			person,
			v.IsConfirmed,
			v.UserRoles,
		)
		assert.Nil(t, err)
		err = repository.Save(ctx, user)
		assert.Nil(t, err)
		dbUser, err := repository.Get(ctx, user.Id())
		assert.Nil(t, err)
		assert.Equal(t, v.Id, dbUser.Id())
		assert.Equal(t, v.UserName, dbUser.UserName())
		assert.Equal(t, v.Password, dbUser.Password())
		assert.Equal(t, v.RegisterDate.StringRepresentation(), dbUser.RegisterDate().StringRepresentation())
		assert.Equal(t, v.IsConfirmed, dbUser.IsConfirmed())
		assert.Equal(t, v.UserRoles.Cardinality(), dbUser.UserRoles().Cardinality())
		assert.NotNil(t, dbUser.Person)
		t.Cleanup(func() {
			err := repository.Remove(ctx, v.Id)
			assert.Nil(t, err)
		})
	}
}

func TestUserRepository_Find_ByUserName(t *testing.T) {
	ctx := context.Background()
	repository, ok := repositories.MapperRegistryInstance().Get("User").(repositories.UserRepository)
	assert.True(t, ok)
	for _, v := range mock_data.NewUsers {
		person, err := identity.NewPerson(
			v.Person.Id,
			v.Person.Email,
			v.Person.ProfileImg,
			v.Person.Genre,
			v.Person.Pronoun,
			v.Person.FirstName,
			v.Person.LastName,
			v.Person.Phone,
			v.Person.BirthDate,
			v.Person.Address,
		)
		assert.Nil(t, err)
		user, err := identity.NewUser(
			v.Id,
			v.UserName,
			v.Password,
			v.RegisterDate,
			person,
			v.IsConfirmed,
			v.UserRoles,
		)
		assert.Nil(t, err)
		err = repository.Save(ctx, user)
		assert.Nil(t, err)
		dbUser, err := repository.Find(ctx, find_user.ByUsername(user.UserName()))
		assert.Nil(t, err)
		assert.Equal(t, v.Id, dbUser.Id())
		assert.Equal(t, v.UserName, dbUser.UserName())
		assert.Equal(t, v.Password, dbUser.Password())
		assert.Equal(t, v.RegisterDate.StringRepresentation(), dbUser.RegisterDate().StringRepresentation())
		assert.Equal(t, v.IsConfirmed, dbUser.IsConfirmed())
		assert.Equal(t, v.UserRoles.Cardinality(), dbUser.UserRoles().Cardinality())
		assert.NotNil(t, dbUser.Person)
		t.Cleanup(func() {
			err := repository.Remove(ctx, v.Id)
			assert.Nil(t, err)
		})
	}
}

func TestUserRepository_Update(t *testing.T) {
	ctx := context.Background()
	repository, ok := repositories.MapperRegistryInstance().Get("User").(repositories.UserRepository)
	assert.True(t, ok)
	for _, v := range mock_data.NewUsers {
		person, err := identity.NewPerson(
			v.Person.Id,
			v.Person.Email,
			v.Person.ProfileImg,
			v.Person.Genre,
			v.Person.Pronoun,
			v.Person.FirstName,
			v.Person.LastName,
			v.Person.Phone,
			v.Person.BirthDate,
			v.Person.Address,
		)
		assert.Nil(t, err)
		user, err := identity.NewUser(
			v.Id,
			v.UserName,
			v.Password,
			v.RegisterDate,
			person,
			v.IsConfirmed,
			v.UserRoles,
		)
		assert.Nil(t, err)
		err = repository.Save(ctx, user)
		assert.Nil(t, err)
		dbUser, err := repository.Get(ctx, user.Id())
		assert.Nil(t, err)
		assert.Equal(t, v.UserRoles.Cardinality(), dbUser.UserRoles().Cardinality())
		err = dbUser.AddRole(ctx, v.RoleToAdd.GetRole(), v.RoleToAdd.EnterpriseId(), v.RoleToAdd.EnterpriseName())
		assert.Nil(t, err)
		err = repository.Update(ctx, dbUser)
		assert.Nil(t, err)
		updatedUser, err := repository.Get(ctx, dbUser.Id())
		assert.Nil(t, err)
		assert.Equal(t, 1, updatedUser.UserRoles().Cardinality())
		t.Cleanup(func() {
			err := repository.Remove(ctx, v.Id)
			assert.Nil(t, err)
		})
	}
}
