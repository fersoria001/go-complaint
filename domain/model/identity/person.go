package identity

import (
	"fmt"
	"go-complaint/domain/model/common"
	"go-complaint/erros"
	"regexp"
)

// Package identityandaccess
// <<Entity>> Person
type Person struct {
	//Gender
	email     string
	firstName string
	lastName  string
	birthDate common.Date
	phone     string
	address   common.Address
}

// Resolved in User model
func (p *Person) ChangePhone(phone string) error                   { return nil }
func (p *Person) ChangeAddress(country, county, city string) error { return nil }
func (p *Person) ChangeName(firstName, lastName string) error      { return nil }

func NewPerson(email string, firstName string, lastName string, phone string,
	birthDate common.Date, address common.Address) (*Person, error) {
	var person = new(Person)
	var err error
	err = person.setEmail(email)
	if err != nil {
		return nil, err
	}
	err = person.setFirstName(firstName)
	if err != nil {
		return nil, err
	}
	err = person.setLastName(lastName)
	if err != nil {
		return nil, err
	}
	err = person.setBirthDate(birthDate)
	if err != nil {
		return nil, err
	}
	err = person.setPhone(phone)
	if err != nil {
		return nil, err
	}
	err = person.setAddress(address)
	if err != nil {
		return nil, err
	}
	return person, nil
}

func (p *Person) setEmail(email string) error {
	if email == "" {
		return &erros.NullValueError{}
	}
	p.email = email
	return nil
}

func (p *Person) setAddress(address common.Address) error {
	if address == (common.Address{}) {
		return &erros.EmptyStructError{}
	}
	p.address = address
	return nil
}

func (p *Person) setFirstName(firstName string) error {
	var len = len(firstName)
	if len < 2 || len > 50 {
		return &erros.InvalidLengthError{AttributeName: "person.firstName", CurrentLength: len, MinLength: 2, MaxLength: 50}
	}
	regex, err := regexp.Compile(`^[^\d]*$`)
	if err != nil {
		return err
	}
	if !regex.MatchString(firstName) {
		return &erros.InvalidNameError{}
	}
	p.firstName = firstName
	return nil
}

func (p *Person) setLastName(lastName string) error {
	var len = len(lastName)
	if len < 2 || len > 50 {
		return &erros.InvalidLengthError{AttributeName: "person.lastName", CurrentLength: len, MinLength: 2, MaxLength: 50}
	}
	regex, err := regexp.Compile(`^[^\d]*$`)
	if err != nil {
		return err
	}
	if !regex.MatchString(lastName) {
		return &erros.InvalidNameError{}
	}

	p.lastName = lastName
	return nil
}

func (p *Person) setBirthDate(birthDate common.Date) error {
	if birthDate == (common.Date{}) {
		return &erros.EmptyStructError{}
	}
	p.birthDate = birthDate
	return nil
}

func (p *Person) setPhone(phone string) error {
	var len = len(phone)
	if len < 9 || len > 21 {
		return &erros.InvalidLengthError{AttributeName: "person.phone", CurrentLength: len, MinLength: 9, MaxLength: 21}
	}
	p.phone = phone
	return nil
}

func (p *Person) FullName() string {
	return fmt.Sprintf(`%s %s`, p.firstName, p.lastName)
}

func (p *Person) Email() string {
	return p.email
}

func (p *Person) FirstName() string {
	return p.firstName
}

func (p *Person) LastName() string {
	return p.lastName
}

func (p *Person) BirthDate() common.Date {
	return p.birthDate
}

func (p *Person) Phone() string {
	return p.phone
}

func (p *Person) Address() common.Address {
	return p.address
}

func (p *Person) Age() int {
	return p.birthDate.Age()
}
