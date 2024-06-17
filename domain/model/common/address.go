package common

import (
	"go-complaint/domain"

	"github.com/google/uuid"
)

// First opinion
// Value objects must not be passed as pointers because
// they are immutable and side effect free
type Address struct {
	id       uuid.UUID
	country  Country
	county   CountryState
	city     City
	director domain.Director
}

func NewAddress(
	id uuid.UUID,
	country Country,
	county CountryState,
	city City,
) Address {
	a := Address{
		id:      id,
		country: country,
		county:  county,
		city:    city,
	}
	return a
}
func NewAddressWithDirector(director domain.Director) *Address {
	a := &Address{
		director: director,
	}
	return a
}

func (a *Address) Changed() {
	a.director.Changed(a)
}

func (a Address) ID() uuid.UUID {
	return a.id
}

func (a Address) Country() Country {
	return a.country
}

func (a Address) CountryState() CountryState {
	return a.county
}

func (a Address) City() City {
	return a.city
}
