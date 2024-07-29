package repositories_test

import (
	"context"
	"go-complaint/domain/model/identity"
	"go-complaint/infrastructure/persistence/repositories"
	"go-complaint/test/mock_data"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPersonRepository_Save(t *testing.T) {
	ctx := context.Background()
	repository, ok := repositories.MapperRegistryInstance().Get("Person").(repositories.PersonRepository)
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
		err = repository.Save(ctx, *person)
		assert.Nil(t, err)
		t.Cleanup(func() {
			err := repository.Remove(ctx, v.Person.Id)
			assert.Nil(t, err)
		})
	}
}

func TestPersonRepository_Get(t *testing.T) {
	ctx := context.Background()
	repository, ok := repositories.MapperRegistryInstance().Get("Person").(repositories.PersonRepository)
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
		err = repository.Save(ctx, *person)
		assert.Nil(t, err)
		dbPerson, err := repository.Get(ctx, person.Id())
		assert.Nil(t, err)
		assert.NotNil(t, dbPerson)
		assert.Equal(t, v.Person.Id, dbPerson.Id())
		assert.Equal(t, v.Person.Email, dbPerson.Email())
		assert.Equal(t, v.Person.ProfileImg, dbPerson.ProfileIMG())
		assert.Equal(t, v.Person.Genre, dbPerson.Genre())
		assert.Equal(t, v.Person.Pronoun, dbPerson.Pronoun())
		assert.Equal(t, v.Person.FirstName, dbPerson.FirstName())
		assert.Equal(t, v.Person.LastName, dbPerson.LastName())
		assert.Equal(t, v.Person.Phone, dbPerson.Phone())
		assert.Equal(t, v.Person.BirthDate.StringRepresentation(), dbPerson.BirthDate().StringRepresentation())
		assert.Equal(t, v.Person.Address, dbPerson.Address())
		t.Cleanup(func() {
			err := repository.Remove(ctx, dbPerson.Id())
			assert.Nil(t, err)
		})
	}
}
