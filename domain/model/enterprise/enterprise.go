package enterprise

import (
	"context"
	"go-complaint/domain"
	"go-complaint/domain/model/common"
	"go-complaint/erros"
	"net/mail"
	"time"
)

// Package enterprise
// <<Aggregate>> Enterprise
// Enterprise is a struct that represent the enterprise entity
// in the domain of the company
// id is the owner id and is not unique
// Its name is unique and its the pk
type Enterprise struct {
	name           string
	owner          *Owner
	logoIMG        string
	bannerIMG      string
	website        string
	email          string
	phone          string
	address        common.Address
	industry       Industry
	registerAt     common.Date
	foundationDate common.Date
}

func (e *Enterprise) ChangeAddress(ctx context.Context, country, county, city string) error {

	if country != "" || county != "" || city != "" {
		publisher := domain.DomainEventPublisherInstance()
		oldValue := e.Address()
		if country == "" {
			country = oldValue.Country()
		}
		if county == "" {
			county = oldValue.County()
		}
		if city == "" {
			city = oldValue.City()
		}
		address, err := common.NewAddress(country, county, city)
		if err != nil {
			return err
		}
		err = e.setAddress(address)
		if err != nil {
			return err
		}
		event, err := NewEnterpriseUpdated(e.Name(), e.Industry().Name(),
			map[string]string{"country": oldValue.Country(), "county": oldValue.County(), "city": oldValue.City()},
			map[string]string{"country": address.Country(), "county": address.County(), "city": address.City()})
		if err != nil {
			return err
		}
		err = publisher.Publish(ctx, event)
		if err != nil {
			return err
		}
	} else {
		return &erros.NullValueError{}
	}
	return nil
}

func (e *Enterprise) ChangePhone(ctx context.Context, phone string) error {
	oldValue := e.Phone()
	err := e.setPhone(phone)
	if err != nil {
		return err
	}
	publisher := domain.DomainEventPublisherInstance()
	event, err := NewEnterpriseUpdated(e.Name(), e.Industry().Name(),
		map[string]string{"phone": oldValue},
		map[string]string{"phone": phone})
	if err != nil {
		return err
	}
	err = publisher.Publish(ctx, event)
	if err != nil {
		return err
	}
	return nil
}

func (e *Enterprise) ChangeWebsite(ctx context.Context, website string) error {
	oldValue := e.Website()
	err := e.setWebsite(website)
	if err != nil {
		return err
	}
	publisher := domain.DomainEventPublisherInstance()
	event, err := NewEnterpriseUpdated(e.Name(), e.Industry().Name(),
		map[string]string{"website": oldValue},
		map[string]string{"website": website})
	if err != nil {
		return err
	}
	err = publisher.Publish(ctx, event)
	if err != nil {
		return err
	}
	return nil
}

func (e *Enterprise) ChangeEmail(ctx context.Context, email string) error {
	oldValue := e.Email()
	err := e.setEmail(email)
	if err != nil {
		return err
	}
	publisher := domain.DomainEventPublisherInstance()
	event, err := NewEnterpriseUpdated(e.Name(), e.Industry().Name(),
		map[string]string{"email": oldValue},
		map[string]string{"email": email})
	if err != nil {
		return err
	}
	err = publisher.Publish(ctx, event)
	if err != nil {
		return err
	}
	return nil
}

func (e *Enterprise) ChangeLogoIMG(ctx context.Context, logoIMG string) error {
	oldValue := e.LogoIMG()
	err := e.setLogoIMG(logoIMG)
	if err != nil {
		return err
	}
	publisher := domain.DomainEventPublisherInstance()
	event, err := NewEnterpriseUpdated(e.Name(), e.Industry().Name(),
		map[string]string{"logoIMG": oldValue},
		map[string]string{"logoIMG": logoIMG})
	if err != nil {
		return err
	}
	err = publisher.Publish(ctx, event)
	if err != nil {
		return err
	}
	return nil
}

func (e *Enterprise) ChangeBannerIMG(ctx context.Context, bannerIMG string) error {
	oldValue := e.BannerIMG()
	err := e.setBannerIMG(bannerIMG)
	if err != nil {
		return err
	}
	publisher := domain.DomainEventPublisherInstance()
	event, err := NewEnterpriseUpdated(e.Name(), e.Industry().Name(),
		map[string]string{"logoIMG": oldValue},
		map[string]string{"logoIMG": bannerIMG})
	if err != nil {
		return err
	}
	err = publisher.Publish(ctx, event)
	if err != nil {
		return err
	}
	return nil
}

// factory method and publisher
func CreateEnterprise(ctx context.Context,
	ownerID string,
	name,
	logoIMG,
	bannerIMG,
	website, email, phone, country, county, city, industryName, foundationDate string) (*Enterprise, error) {
	var (
		address   common.Address
		industry  Industry
		fDate     common.Date
		regat     common.Date
		e         *Enterprise
		publisher = domain.DomainEventPublisherInstance()
		event     *EnterpriseCreated
		err       error
	)
	address, err = common.NewAddress(country, county, city)
	if err != nil {
		return nil, err
	}
	industry, err = NewIndustry(0, industryName)
	if err != nil {
		return nil, err
	}
	fDate, err = common.NewDateFromString(foundationDate)
	if err != nil {
		return nil, err
	}
	regat = common.NewDate(time.Now())

	e, err = NewEnterprise(ownerID, name, logoIMG, bannerIMG, website, email, phone, address, industry, regat, fDate)
	if err != nil {
		return nil, err
	}
	event, err = NewEnterpriseCreated(e.Name(), e.Industry().Name(), e.RegisterAt().Date())
	if err != nil {
		return nil, err
	}
	err = publisher.Publish(ctx, event)
	if err != nil {
		return nil, err
	}
	return e, err
}

// constructor
func NewEnterprise(
	ownerID string,
	name, logoIMG, bannerIMG, website, email, phone string,
	address common.Address, industry Industry,
	registerAt, foundationDate common.Date) (*Enterprise, error) {
	var e = new(Enterprise)
	var err error
	err = e.setName(name)
	if err != nil {
		return nil, err
	}
	_, err = mail.ParseAddress(ownerID)
	if err != nil {
		return nil, &erros.ValidationError{
			Expected: "valid email as ownerID",
		}
	}
	err = e.setOwner(NewOwner(ownerID))
	if err != nil {
		return nil, err
	}
	err = e.setLogoIMG(logoIMG)
	if err != nil {
		return nil, err
	}
	err = e.setWebsite(website)
	if err != nil {
		return nil, err
	}
	err = e.setEmail(email)
	if err != nil {
		return nil, err
	}
	err = e.setPhone(phone)
	if err != nil {
		return nil, err
	}
	err = e.setAddress(address)
	if err != nil {
		return nil, err
	}
	err = e.setIndustry(industry)
	if err != nil {
		return nil, err
	}
	err = e.setRegisterAt(registerAt)
	if err != nil {
		return nil, err
	}
	err = e.setFoundationDate(foundationDate)
	if err != nil {
		return nil, err
	}
	return e, nil
}

func (e *Enterprise) setOwner(owner *Owner) error {
	if owner == nil {
		return &erros.NullValueError{}
	}
	e.owner = owner
	return nil
}

func (e *Enterprise) setLogoIMG(logoIMG string) error {
	if logoIMG == "" {
		return &erros.NullValueError{}
	}
	e.logoIMG = logoIMG
	return nil
}

func (e *Enterprise) setBannerIMG(bannerIMG string) error {
	if bannerIMG == "" {
		return &erros.NullValueError{}
	}
	e.bannerIMG = bannerIMG
	return nil
}

func (e *Enterprise) setName(name string) error {
	if name == "" {
		return &erros.NullValueError{}
	}
	if len(name) < 3 {
		return &erros.InvalidLengthError{AttributeName: "name", MinLength: 3, MaxLength: 120, CurrentLength: len(name)}
	}

	if len(name) > 120 {
		return &erros.InvalidLengthError{AttributeName: "name", MinLength: 3, MaxLength: 120, CurrentLength: len(name)}
	}

	e.name = name
	return nil
}

func (e *Enterprise) setWebsite(website string) error {
	if website == "" {
		return &erros.NullValueError{}
	}
	e.website = website
	return nil
}

func (e *Enterprise) setEmail(email string) error {
	if email == "" {
		return &erros.NullValueError{}
	}
	var valid, err = mail.ParseAddress(email)
	if err != nil {
		return &erros.InvalidEmailError{}
	}
	if valid == nil {
		return &erros.NullValueError{}
	}
	e.email = email
	return nil
}

func (e *Enterprise) setPhone(phone string) error {

	if phone == "" {
		return &erros.NullValueError{}
	}
	if len(phone) < 10 {
		return &erros.InvalidLengthError{AttributeName: "phone", MinLength: 10, MaxLength: 21, CurrentLength: len(phone)}
	}

	if len(phone) > 21 {
		return &erros.InvalidLengthError{AttributeName: "phone", MinLength: 10, MaxLength: 21, CurrentLength: len(phone)}
	}
	e.phone = phone
	return nil
}

func (e *Enterprise) setAddress(address common.Address) error {
	if address == (common.Address{}) {
		return &erros.NullValueError{}
	}
	e.address = address
	return nil
}

func (e *Enterprise) setIndustry(industry Industry) error {
	if industry == (Industry{}) {
		return &erros.EmptyStructError{}
	}
	e.industry = industry
	return nil
}

func (e *Enterprise) setRegisterAt(registerAt common.Date) error {
	if registerAt == (common.Date{}) {
		return &erros.EmptyStructError{}
	}
	e.registerAt = registerAt
	return nil
}

func (e *Enterprise) setFoundationDate(foundationDate common.Date) error {
	if foundationDate == (common.Date{}) {
		return &erros.EmptyStructError{}
	}
	e.foundationDate = foundationDate
	return nil
}

func (e *Enterprise) LogoIMG() string {
	return e.logoIMG
}

func (e *Enterprise) BannerIMG() string {
	return e.bannerIMG
}

func (e *Enterprise) Name() string {
	return e.name
}

func (e *Enterprise) Website() string {
	return e.website
}

func (e *Enterprise) Email() string {
	return e.email
}

func (e *Enterprise) Phone() string {
	return e.phone
}

func (e *Enterprise) Address() common.Address {
	return e.address
}

func (e *Enterprise) Industry() Industry {
	return e.industry
}

func (e *Enterprise) RegisterAt() common.Date {
	return e.registerAt
}

func (e *Enterprise) FoundationDate() common.Date {
	return e.foundationDate
}
func (e *Enterprise) OwnerID() string {
	return e.owner.id
}

func (e *Enterprise) Owner() *Owner {
	return e.owner
}
