package enterprise_test

import (
	"go-complaint/domain/model/enterprise"
	"go-complaint/domain/model/identity"
	"go-complaint/test/mock_data"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmployee_New(t *testing.T) {
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

	}
}
