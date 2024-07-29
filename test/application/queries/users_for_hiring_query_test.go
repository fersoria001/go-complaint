package queries_test

import (
	"context"
	"go-complaint/application/queries"
	"go-complaint/domain/model/enterprise"
	"go-complaint/domain/model/identity"
	"go-complaint/infrastructure/persistence/repositories"
	"go-complaint/test/mock_data"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUsersForHiringQuery_Setup(t *testing.T) {
	ctx := context.Background()
	mockOwner := mock_data.NewUsers["valid"]
	mockEnterprise := mock_data.NewEnterprises["valid"]
	userRepository, ok := repositories.MapperRegistryInstance().Get("User").(repositories.UserRepository)
	assert.True(t, ok)
	enterpriseRepository, ok := repositories.MapperRegistryInstance().Get("Enterprise").(repositories.EnterpriseRepository)
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
		err = userRepository.Save(ctx, user)
		assert.Nil(t, err)
	}
	industry, err := enterprise.NewIndustry(mockEnterprise.Industry.Id, mockEnterprise.Industry.Name)
	assert.Nil(t, err)
	e, err := enterprise.NewEnterprise(
		mockEnterprise.Id,
		mockOwner.Id,
		mockEnterprise.Name,
		mockEnterprise.LogoImg,
		mockEnterprise.BannerImg,
		mockEnterprise.Website,
		mockEnterprise.Email,
		mockEnterprise.Phone,
		mockEnterprise.Address,
		industry,
		mockEnterprise.RegisterAt,
		mockEnterprise.UpdatedAt,
		mockEnterprise.FoundationDate,
		mockEnterprise.Employees,
	)
	assert.Nil(t, err)
	err = enterpriseRepository.Save(ctx, e)
	assert.Nil(t, err)
}

func TestUsersForHiringQuery_Execute(t *testing.T) {
	TestUsersForHiringQuery_Setup(t)
	ctx := context.Background()
	mockOwner := mock_data.NewUsers["valid"]
	mockEnterprise := mock_data.NewEnterprises["valid"]
	term := ""
	q := queries.NewUsersForHiringQuery(
		mockEnterprise.Id.String(),
		term,
		10,
		0,
	)
	usersForHiring, err := q.Execute(ctx)
	assert.Nil(t, err)
	assert.Greater(t, len(usersForHiring.Users), 0)
	for _, v := range usersForHiring.Users {
		assert.NotEqual(t, mockOwner.Id.String(), v.Id)
	}
	t.Cleanup(func() {
		userRepository, ok := repositories.MapperRegistryInstance().Get("User").(repositories.UserRepository)
		assert.True(t, ok)
		enterpriseRepository, ok := repositories.MapperRegistryInstance().Get("Enterprise").(repositories.EnterpriseRepository)
		assert.True(t, ok)
		err := enterpriseRepository.Remove(ctx, mockEnterprise.Id)
		assert.Nil(t, err)
		for _, v := range mock_data.NewUsers {
			err := userRepository.Remove(ctx, v.Id)
			assert.Nil(t, err)
		}
	})
}

func TestUsersForHiringQuery_Execute_WithTerm_TRUE(t *testing.T) {
	TestUsersForHiringQuery_Setup(t)
	ctx := context.Background()
	mockUser := mock_data.NewUsers["valid1"]
	mockOwner := mock_data.NewUsers["valid"]
	mockEnterprise := mock_data.NewEnterprises["valid"]
	term := mockUser.Person.FirstName
	q := queries.NewUsersForHiringQuery(
		mockEnterprise.Id.String(),
		term,
		10,
		0,
	)
	usersForHiring, err := q.Execute(ctx)
	assert.Nil(t, err)
	assert.Greater(t, len(usersForHiring.Users), 0)
	for _, v := range usersForHiring.Users {
		assert.NotEqual(t, mockOwner.Id.String(), v.Id)
	}
	t.Cleanup(func() {
		userRepository, ok := repositories.MapperRegistryInstance().Get("User").(repositories.UserRepository)
		assert.True(t, ok)
		enterpriseRepository, ok := repositories.MapperRegistryInstance().Get("Enterprise").(repositories.EnterpriseRepository)
		assert.True(t, ok)
		err := enterpriseRepository.Remove(ctx, mockEnterprise.Id)
		assert.Nil(t, err)
		for _, v := range mock_data.NewUsers {
			err := userRepository.Remove(ctx, v.Id)
			assert.Nil(t, err)
		}
	})
}

func TestUsersForHiringQuery_Execute_WithTerm_FALSE(t *testing.T) {
	TestUsersForHiringQuery_Setup(t)
	ctx := context.Background()
	mockEnterprise := mock_data.NewEnterprises["valid"]
	term := "notIncludedWord"
	q := queries.NewUsersForHiringQuery(
		mockEnterprise.Id.String(),
		term,
		10,
		0,
	)
	usersForHiring, err := q.Execute(ctx)
	assert.Nil(t, err)
	assert.Equal(t, 0, len(usersForHiring.Users))
	t.Cleanup(func() {
		userRepository, ok := repositories.MapperRegistryInstance().Get("User").(repositories.UserRepository)
		assert.True(t, ok)
		enterpriseRepository, ok := repositories.MapperRegistryInstance().Get("Enterprise").(repositories.EnterpriseRepository)
		assert.True(t, ok)
		err := enterpriseRepository.Remove(ctx, mockEnterprise.Id)
		assert.Nil(t, err)
		for _, v := range mock_data.NewUsers {
			err := userRepository.Remove(ctx, v.Id)
			assert.Nil(t, err)
		}
	})
}
