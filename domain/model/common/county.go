package common

type CountryState struct {
	id   int
	name string
}

func NewCountryState(id int, name string) CountryState {
	return CountryState{id: id, name: name}
}

func (c CountryState) ID() int {
	return c.id
}

func (c CountryState) Name() string {
	return c.name
}
