package repositories_test

import (
	"context"
	"go-complaint/domain/model/enterprise"
	"go-complaint/domain/model/identity"
	"go-complaint/infrastructure/persistence/repositories"
	"go-complaint/test/mock_data"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmployeeRepository_Save(t *testing.T) {
	ctx := context.Background()
	reg := repositories.MapperRegistryInstance()
	employeeRepository, ok := reg.Get("Employee").(repositories.EmployeeRepository)
	assert.True(t, ok)
	for _, v := range mock_data.NewEmployees {
		person, err := identity.NewPerson(
			v.User.Person.Id,
			v.User.Person.Email,
			v.User.Person.ProfileImg,
			v.User.Person.Genre,
			v.User.Person.Pronoun,
			v.User.Person.FirstName,
			v.User.Person.LastName,
			v.User.Person.Phone,
			v.User.Person.BirthDate,
			v.User.Person.Address,
		)
		assert.Nil(t, err)
		user, err := identity.NewUser(
			v.User.Id,
			v.User.UserName,
			v.User.Password,
			v.User.RegisterDate,
			person,
			v.User.IsConfirmed,
			v.User.UserRoles,
		)
		assert.Nil(t, err)
		e, err := enterprise.NewEmployee(
			v.Id,
			v.EnterpriseId,
			user,
			v.Position,
			v.HiringDate,
			v.ApprovedHiring,
			v.ApprovedHiringAt,
		)
		assert.Nil(t, err)
		assert.NotNil(t, e)
		err = employeeRepository.Save(ctx, e)
		assert.Nil(t, err)
	}
	t.Cleanup(func() {
		for _, v := range mock_data.NewEmployees {
			err := employeeRepository.Remove(ctx, v.Id)
			assert.Nil(t, err)
		}
	})
}

func TestEmployeeRepository_Get(t *testing.T) {
	ctx := context.Background()
	reg := repositories.MapperRegistryInstance()
	employeeRepository, ok := reg.Get("Employee").(repositories.EmployeeRepository)
	assert.True(t, ok)
	userRepository, ok := reg.Get("User").(repositories.UserRepository)
	assert.True(t, ok)
	for _, v := range mock_data.NewEmployees {
		person, err := identity.NewPerson(
			v.User.Person.Id,
			v.User.Person.Email,
			v.User.Person.ProfileImg,
			v.User.Person.Genre,
			v.User.Person.Pronoun,
			v.User.Person.FirstName,
			v.User.Person.LastName,
			v.User.Person.Phone,
			v.User.Person.BirthDate,
			v.User.Person.Address,
		)
		assert.Nil(t, err)
		user, err := identity.NewUser(
			v.User.Id,
			v.User.UserName,
			v.User.Password,
			v.User.RegisterDate,
			person,
			v.User.IsConfirmed,
			v.User.UserRoles,
		)
		assert.Nil(t, err)
		err = userRepository.Save(ctx, user)
		assert.Nil(t, err)
		e, err := enterprise.NewEmployee(
			v.Id,
			v.EnterpriseId,
			user,
			v.Position,
			v.HiringDate,
			v.ApprovedHiring,
			v.ApprovedHiringAt,
		)
		assert.Nil(t, err)
		assert.NotNil(t, e)
		err = employeeRepository.Save(ctx, e)
		assert.Nil(t, err)
		dbE, err := employeeRepository.Get(ctx, v.Id)
		assert.Nil(t, err)
		assert.NotNil(t, dbE)
		assert.Equal(t, v.Id, dbE.ID())
		assert.Equal(t, v.EnterpriseId, dbE.EnterpriseId())
		assert.Equal(t, user.Id(), dbE.User.Id())
		assert.Equal(t, v.HiringDate.StringRepresentation(), dbE.HiringDate().StringRepresentation())
		assert.Equal(t, v.ApprovedHiring, dbE.ApprovedHiring())
		assert.Equal(t, v.ApprovedHiringAt.StringRepresentation(), dbE.ApprovedHiringAt().StringRepresentation())
		assert.Equal(t, v.Position, dbE.Position())
	}
	t.Cleanup(func() {
		for _, v := range mock_data.NewEmployees {
			err := employeeRepository.Remove(ctx, v.Id)
			assert.Nil(t, err)
			err = userRepository.Remove(ctx, v.Id)
			assert.Nil(t, err)
		}
	})
}
