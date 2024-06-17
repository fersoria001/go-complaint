package identity

import (
	"context"
	"fmt"
	"go-complaint/domain"
	"go-complaint/domain/model/common"
	"go-complaint/erros"
	"regexp"
)

// Package identityandaccess
// <<Entity>> Person
type Person struct {
	gender     string
	pronoun    string
	profileIMG string
	email      string
	firstName  string
	lastName   string
	birthDate  common.Date
	phone      string
	address    common.Address
}

func (p *Person) ChangePronoun(ctx context.Context, pronoun string) error {
	var oldValue = p.pronoun
	if oldValue == pronoun {
		return nil
	}
	err := p.setPronoun(pronoun)
	if err != nil {
		return err
	}
	eventt, err := NewPersonPronounChanged(p.email, oldValue, pronoun)
	if err != nil {
		return err
	}
	domain.DomainEventPublisherInstance().Publish(
		ctx,
		eventt,
	)
	return nil
}

func (p *Person) ChangeGender(ctx context.Context, gender string) error {
	var oldValue = p.profileIMG
	if oldValue == gender {
		return nil
	}
	err := p.setGender(gender)
	if err != nil {
		return err
	}
	eventt, err := NewPersonGenderChanged(p.email, oldValue, gender)
	if err != nil {
		return err
	}
	domain.DomainEventPublisherInstance().Publish(
		ctx,
		eventt,
	)
	return nil
}
func (p *Person) ChangeProfileIMG(ctx context.Context, profileIMG string) error {
	var oldValue = p.profileIMG
	if oldValue == profileIMG {
		return nil
	}
	err := p.setProfileIMG(profileIMG)
	if err != nil {
		return err
	}
	eventt, err := NewPersonProfileIMGChanged(p.email, oldValue, profileIMG)
	if err != nil {
		return err
	}
	domain.DomainEventPublisherInstance().Publish(
		ctx,
		eventt,
	)
	return nil
}

// Resolved in User model
func (p *Person) ChangePhone(
	ctx context.Context,
	phone string) error {
	var oldValue = p.phone
	if oldValue == phone {
		return nil
	}
	err := p.setPhone(phone)
	if err != nil {
		return err
	}
	eventt, err := NewPersonPhoneChanged(p.email, oldValue, phone)
	if err != nil {
		return err
	}
	domain.DomainEventPublisherInstance().Publish(
		ctx,
		eventt,
	)
	return nil
}

func (p *Person) ChangeCity(
	ctx context.Context,
	cityID int,
) error {
	var oldValue = p.address.City().ID()
	if oldValue == cityID {
		return nil
	}
	newCity := common.NewCity(cityID, "", "", 0, 0)
	newAddress := common.NewAddress(
		p.address.ID(),
		p.address.Country(),
		p.address.CountryState(),
		newCity,
	)
	err := p.setAddress(newAddress)
	if err != nil {
		return err
	}
	eventt, err := NewPersonCityChanged(p.email, oldValue, cityID)
	if err != nil {
		return err
	}
	domain.DomainEventPublisherInstance().Publish(
		ctx,
		eventt,
	)
	return nil
}

func (p *Person) ChangeCountryState(
	ctx context.Context,
	countryStateID int,
) error {
	var oldValue = p.address.CountryState().ID()
	if oldValue == countryStateID {
		return nil
	}
	newCountryState := common.NewCountryState(countryStateID, "")
	newAddress := common.NewAddress(
		p.address.ID(),
		p.address.Country(),
		newCountryState,
		p.address.City(),
	)
	err := p.setAddress(newAddress)
	if err != nil {
		return err
	}
	eventt, err := NewPersonCountryStateChanged(p.email, oldValue, countryStateID)
	if err != nil {
		return err
	}
	domain.DomainEventPublisherInstance().Publish(
		ctx,
		eventt,
	)
	return nil
}

func (p *Person) ChangeCountry(ctx context.Context, countryID int) error {
	var oldValue = p.address.Country().ID()
	if oldValue == countryID {
		return nil
	}
	newCountry := common.NewCountry(countryID, "", "")
	newAddress := common.NewAddress(
		p.address.ID(),
		newCountry,
		p.address.CountryState(),
		p.address.City(),
	)
	err := p.setAddress(newAddress)
	if err != nil {
		return err
	}
	eventt, err := NewPersonCountryChanged(p.email, oldValue, countryID)
	if err != nil {
		return err
	}
	domain.DomainEventPublisherInstance().Publish(
		ctx,
		eventt,
	)
	return nil
}

func (p *Person) ChangeLastName(ctx context.Context, setLastName string) error {
	var oldValue = p.lastName
	if oldValue == setLastName {
		return nil
	}
	err := p.setLastName(setLastName)
	if err != nil {
		return err
	}
	eventt, err := NewPersonLastNameChanged(p.email, oldValue, setLastName)
	if err != nil {
		return err
	}
	domain.DomainEventPublisherInstance().Publish(
		ctx,
		eventt,
	)
	return nil
}

func (p *Person) ChangeFirstName(ctx context.Context, setFirstName string) error {
	var oldValue = p.firstName
	if oldValue == setFirstName {
		return nil
	}
	err := p.setFirstName(setFirstName)
	if err != nil {
		return err
	}
	eventt, err := NewPersonFirstNameChanged(p.email, oldValue, setFirstName)
	if err != nil {
		return err
	}
	domain.DomainEventPublisherInstance().Publish(
		ctx,
		eventt,
	)
	return nil
}

func NewPerson(
	email string,
	profileIMG string,
	gender string,
	pronoun string,
	firstName string,
	lastName string,
	phone string,
	birthDate common.Date,
	address common.Address,
) (*Person, error) {
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
	err = person.setGender(gender)
	if err != nil {
		return nil, err
	}
	err = person.setPronoun(pronoun)
	if err != nil {
		return nil, err
	}
	err = person.setProfileIMG(profileIMG)
	if err != nil {
		return nil, err
	}
	return person, nil
}

func (p *Person) setGender(gender string) error {
	if gender == "" {
		return &erros.NullValueError{}
	}
	p.gender = gender
	return nil
}

func (p *Person) setPronoun(pronoun string) error {
	if pronoun == "" {
		return &erros.NullValueError{}
	}
	p.pronoun = pronoun
	return nil
}

func (p *Person) setProfileIMG(profileIMG string) error {
	if profileIMG == "" {
		return &erros.NullValueError{}
	}
	p.profileIMG = profileIMG
	return nil
}

func (p *Person) setEmail(email string) error {
	if email == "" {
		return &erros.NullValueError{}
	}
	p.email = email
	return nil
}

func (p *Person) setAddress(address common.Address) error {
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

func (p Person) FullName() string {
	return fmt.Sprintf(`%s %s`, p.firstName, p.lastName)
}

func (p Person) Email() string {
	return p.email
}

func (p Person) FirstName() string {
	return p.firstName
}

func (p Person) LastName() string {
	return p.lastName
}

func (p Person) BirthDate() common.Date {
	return p.birthDate
}

func (p Person) Phone() string {
	return p.phone
}

func (p Person) Address() common.Address {
	return p.address
}

func (p Person) Age() int {
	return p.birthDate.Age()
}

func (p Person) Gender() string {
	return p.gender
}

func (p Person) Pronoun() string {
	return p.pronoun
}

func (p Person) ProfileIMG() string {
	return p.profileIMG
}
