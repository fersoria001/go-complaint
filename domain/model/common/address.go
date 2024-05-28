package common

import "go-complaint/erros"

// First opinion
//Value objects must not be passed as pointers because
//they are immutable and side effect free
type Address struct {
	country string
	county  string
	city    string
}

func NewEmptyAddress() Address {
	return Address{}
}

func NewAddress(country string, county string, city string) (Address, error) {
	var address Address = *new(Address)
	err := address.setCountry(country)
	if err != nil {
		return Address{}, err
	}
	err = address.setCounty(county)
	if err != nil {
		return Address{}, err
	}
	err = address.setCity(city)
	if err != nil {
		return Address{}, err
	}
	return address, nil
}

func (a *Address) setCountry(country string) error {
	if country == "" {
		return &erros.NullValueError{}
	}
	a.country = country
	return nil
}

func (a *Address) setCounty(county string) error {
	if county == "" {
		return &erros.NullValueError{}
	}
	a.county = county
	return nil
}

func (a *Address) setCity(city string) error {
	if city == "" {
		return &erros.NullValueError{}
	}
	a.city = city
	return nil
}

func (a Address) Country() string {
	return a.country
}

func (a Address) County() string {
	return a.county
}

func (a Address) City() string {
	return a.city
}
