package queries_test

import (
	"context"
	"go-complaint/application/queries"
	"go-complaint/domain/model/identity"
	"go-complaint/infrastructure/persistence/repositories"
	"go-complaint/test/mock_data"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserByUsernameQuery_Execute(t *testing.T) {
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
		q := queries.NewUserByUsernameQuery(user.UserName())
		dbUser, err := q.Execute(ctx)
		assert.Nil(t, err)
		assert.NotNil(t, dbUser)
		assert.Equal(t, v.Id.String(), dbUser.Id)
		assert.Equal(t, v.UserName, dbUser.Username)
		assert.Equal(t, v.Person.ProfileImg, dbUser.Person.ProfileImg)
		assert.Equal(t, v.Person.Email, dbUser.Person.Email)
		assert.Equal(t, v.Person.FirstName, dbUser.Person.FirstName)
		assert.Equal(t, v.Person.LastName, dbUser.Person.LastName)
		assert.Equal(t, v.Person.Genre, dbUser.Person.Genre)
		assert.Equal(t, v.Person.Pronoun, dbUser.Person.Pronoun)
		assert.Equal(t, v.Person.BirthDate.Age(), dbUser.Person.Age)
		assert.Equal(t, v.Person.Phone, dbUser.Person.Phone)
		assert.Equal(t, v.Person.Address.Country().Name(), dbUser.Person.Address.Country)
		assert.Equal(t, v.Person.Address.CountryState().Name(), dbUser.Person.Address.County)
		assert.Equal(t, v.Person.Address.City().Name(), dbUser.Person.Address.City)
		t.Cleanup(func() {
			err := repository.Remove(ctx, v.Id)
			assert.Nil(t, err)
		})
	}
}
