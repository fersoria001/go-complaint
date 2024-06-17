package common

type City struct {
	id          int
	name        string
	countryCode string
	latitude    float64
	longitude   float64
}

func NewCity(id int, name string, countryCode string, latitude, longitude float64) City {
	return City{
		id:          id,
		name:        name,
		countryCode: countryCode,
		latitude:    latitude,
		longitude:   longitude,
	}
}

func (c City) ID() int {
	return c.id
}

func (c City) Name() string {
	return c.name
}

func (c City) CountryCode() string {
	return c.countryCode
}

func (c City) Latitude() float64 {
	return c.latitude
}

func (c City) Longitude() float64 {
	return c.longitude
}
