package enterprise_test

import (
	"go-complaint/domain/model/enterprise"
	"go-complaint/tests"
	"testing"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/stretchr/testify/assert"
)

func TestEnteprise(t *testing.T) {
	industry, err := enterprise.NewIndustry(190, "Oil & Gas")
	assert.Nil(t, err)
	emptySet := mapset.NewSet[enterprise.Employee]()
	newEnterprise, err := enterprise.NewEnterprise(
		tests.UserOwner.Email(),
		"FloatingPoint. Ltd",
		"/logo.jpg",
		"/banner.jpg",
		"www.floatingpoint.com",
		"floating-point@enterprise.com",
		"012345678910",
		tests.Address5,
		industry,
		tests.CommonDate,
		tests.CommonDate,
		tests.CommonDate,
		emptySet,
	)
	assert.Nil(t, err)
	assert.NotNil(t, newEnterprise)
	assert.Equal(t, "FloatingPoint. Ltd", newEnterprise.Name())

}
