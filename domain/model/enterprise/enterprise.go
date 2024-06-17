package enterprise

import (
	"context"
	"go-complaint/domain"
	"go-complaint/domain/model/common"
	"go-complaint/domain/model/identity"
	"go-complaint/erros"
	"net/mail"
	"strconv"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/google/uuid"
)

// Package enterprise
// <<Aggregate>> Enterprise
// Enterprise is a struct that represent the enterprise entity
// in the domain of the company
// id is the owner id and is not unique
// Its name is unique and its the pk
type Enterprise struct {
	name           string
	owner          string
	logoIMG        string
	bannerIMG      string
	website        string
	email          string
	phone          string
	address        common.Address
	industry       Industry
	registerAt     common.Date
	updatedAt      common.Date
	foundationDate common.Date
	employees      mapset.Set[Employee]
}

func (e *Enterprise) PromoteEmployee(
	ctx context.Context,
	promotedBy string,
	employeeID uuid.UUID,
	newPosition Position) (*identity.User, error) {
	var (
		err       error
		user      *identity.User
		publisher = domain.DomainEventPublisherInstance()
		event     *EmployeePromoted
	)
	for emp := range e.Employees().Iter() {
		if employeeID == emp.ID() {
			if emp.ApprovedHiring() {
				return nil, &erros.ValidationError{Expected: "employee needs to be approved first"}
			}
			err = emp.SetPosition(newPosition)
			if err != nil {
				return nil, err
			}
			user := emp.GetUser()
			role, err := identity.ParseRole(newPosition.String())
			if err != nil {
				return nil, err
			}
			for r := range user.UserRoles().Iter() {
				if r.EnterpriseID() == e.name {
					err = user.RemoveUserRole(ctx, r.GetRole(), e.name)
					if err != nil {
						return nil, err
					}
				}
			}
			err = user.AddRole(ctx, role, e.name)
			if err != nil {
				return nil, err
			}
			event = NewEmployeePromoted(e.name, promotedBy, emp.Email(), newPosition)
			break
		}
	}
	err = publisher.Publish(ctx, event)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (e *Enterprise) EmployeeLeave(ctx context.Context,
	employeeID uuid.UUID) (*identity.User, error) {
	var (
		publisher = domain.DomainEventPublisherInstance()
		err       error
	)
	var employee Employee
	var user *identity.User
	for emp := range e.Employees().Iter() {
		if employeeID == emp.ID() {
			if !emp.ApprovedHiring() {
				return nil, &erros.ValidationError{Expected: "employee needs to be approved first"}
			}
			role, err := identity.ParseRole(emp.Position().String())
			if err != nil {
				return nil, err
			}
			user = emp.GetUser()
			err = emp.GetUser().RemoveUserRole(ctx,
				role,
				e.Name(),
			)
			if err != nil {
				return nil, err
			}
			employee = emp
			e.Employees().Remove(emp)
		}
	}

	err = publisher.Publish(ctx, NewEmployeeLeaved(employee))
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (e *Enterprise) FireEmployee(ctx context.Context,
	employeeID uuid.UUID) (*identity.User, error) {
	var (
		publisher = domain.DomainEventPublisherInstance()
		err       error
	)
	var employee Employee
	var user *identity.User
	for emp := range e.Employees().Iter() {
		if employeeID == emp.ID() {
			if !emp.ApprovedHiring() {
				return nil, &erros.ValidationError{Expected: "employee needs to be approved first"}
			}
			role, err := identity.ParseRole(emp.Position().String())
			if err != nil {
				return nil, err
			}
			user = emp.GetUser()
			err = emp.GetUser().RemoveUserRole(ctx,
				role,
				e.Name(),
			)
			if err != nil {
				return nil, err
			}
			employee = emp
			e.Employees().Remove(emp)
		}
	}

	err = publisher.Publish(ctx, NewEmployeeFired(employee))
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (e *Enterprise) CancelHiringProccess(
	ctx context.Context,
	candidateID string,
	position Position,
) error {
	return domain.DomainEventPublisherInstance().Publish(
		ctx,
		NewHiringProccessCanceled(
			e.Name(),
			candidateID,
			position,
		),
	)
}

func (e *Enterprise) HireEmployee(
	ctx context.Context,
	user *identity.User,
	employee Employee,
) error {
	err := e.AddEmployee(employee)
	if err != nil {
		return err
	}
	role, err := identity.ParseRole(employee.Position().String())
	if err != nil {
		return err
	}
	err = user.AddRole(
		ctx,
		role,
		e.Name(),
	)
	if err != nil {
		return err
	}
	domainEventPublisher := domain.DomainEventPublisherInstance()
	err = domainEventPublisher.Publish(
		ctx,
		NewEmployeeHired(
			e.Name(),
			employee.ID(),
			employee.Email(),
			employee.Position(),
		),
	)
	if err != nil {
		return err
	}
	return nil
}

func (e *Enterprise) InviteToProject(
	ctx context.Context,
	userID string,
	proposedTo string,
	proposalPosition Position,
) error {
	ok := false
	if userID != e.owner {
		for employee := range e.employees.Iter() {
			if employee.Email() == userID &&
				employee.Position() == MANAGER {
				ok = true
				break
			}
		}
	} else {
		ok = true
	}
	if !ok {
		return ErrForbidden
	}
	event := NewHiringInvitationSent(
		e.name,
		userID,
		proposedTo,
		proposalPosition,
	)
	return domain.DomainEventPublisherInstance().Publish(ctx, event)
}

func (e *Enterprise) ChangeCity(
	ctx context.Context,
	cityID int,
) error {
	oldValue := e.Address().City().Name()
	city := common.NewCity(cityID, "", "", 0, 0)
	newAddress := common.NewAddress(e.address.ID(), e.Address().Country(), e.Address().CountryState(), city)
	err := e.setAddress(newAddress)
	if err != nil {
		return err
	}
	publisher := domain.DomainEventPublisherInstance()
	event, err := NewEnterpriseUpdated(e.Name(), e.Industry().Name(),
		map[string]string{"city": oldValue}, map[string]string{"city": city.Name()})
	if err != nil {
		return err
	}
	err = publisher.Publish(ctx, event)
	if err != nil {
		return err
	}
	return nil
}

func (e *Enterprise) ChangeCountryState(
	ctx context.Context,
	countryStateID int,
) error {
	oldValue := strconv.Itoa(e.Address().CountryState().ID())
	countryState := common.NewCountryState(countryStateID, "")
	newAddress := common.NewAddress(e.address.ID(), e.Address().Country(), countryState, e.Address().City())
	err := e.setAddress(newAddress)
	if err != nil {
		return err
	}
	publisher := domain.DomainEventPublisherInstance()
	event, err := NewEnterpriseUpdated(e.Name(), e.Industry().Name(),
		map[string]string{"countryState": oldValue}, map[string]string{"countryState": countryState.Name()})
	if err != nil {
		return err
	}
	err = publisher.Publish(ctx, event)
	if err != nil {
		return err
	}
	return nil
}

func (e *Enterprise) ChangeCountry(
	ctx context.Context,
	countryID int,
) error {
	oldValue := strconv.Itoa(e.Address().Country().ID())
	country := common.NewCountry(countryID, "", "")
	newAddress := common.NewAddress(e.address.ID(), country, e.Address().CountryState(), e.Address().City())
	err := e.setAddress(newAddress)
	if err != nil {
		return err
	}
	publisher := domain.DomainEventPublisherInstance()
	event, err := NewEnterpriseUpdated(e.Name(), e.Industry().Name(),
		map[string]string{"country": oldValue}, map[string]string{"country": country.Name()})
	if err != nil {
		return err
	}
	err = publisher.Publish(ctx, event)
	if err != nil {
		return err
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
func CreateEnterprise(
	ctx context.Context,
	owner *identity.User,
	name,
	website,
	email,
	phone string,
	foundationDate common.Date,
	industry Industry,
	address common.Address,
) (*Enterprise, error) {
	regat := common.NewDate(time.Now())
	emptySet := mapset.NewSet[Employee]()
	e, err := NewEnterprise(
		owner.Email(),
		name,
		"/banner.jpg",
		"/logo.jpg",
		website,
		email,
		phone,
		address,
		industry,
		regat,
		regat,
		foundationDate,
		emptySet,
	)
	if err != nil {
		return nil, err
	}
	event, err := NewEnterpriseCreated(
		e.Name(),
		e.Industry().ID(),
		e.CreatedAt().Date(),
	)
	if err != nil {
		return nil, err
	}
	err = domain.DomainEventPublisherInstance().Publish(ctx, event)
	if err != nil {
		return nil, err
	}
	err = owner.AddRole(
		ctx,
		identity.OWNER,
		name,
	)
	return e, err
}

// constructor
func NewEnterprise(
	ownerID,
	name,
	logoIMG,
	bannerIMG,
	website,
	email,
	phone string,
	address common.Address,
	industry Industry,
	registerAt,
	updatedAt,
	foundationDate common.Date,
	employees mapset.Set[Employee],
) (*Enterprise, error) {
	e := new(Enterprise)
	err := e.setName(name)
	if err != nil {
		return nil, err
	}
	_, err = mail.ParseAddress(ownerID)
	if err != nil {
		return nil, &erros.ValidationError{
			Expected: "valid email as ownerID",
		}
	}
	err = e.setOwner(ownerID)
	if err != nil {
		return nil, err
	}
	err = e.setLogoIMG(logoIMG)
	if err != nil {
		return nil, err
	}
	err = e.setBannerIMG(bannerIMG)
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
	err = e.setUpdatedAt(updatedAt)
	if err != nil {
		return nil, err
	}
	err = e.setFoundationDate(foundationDate)
	if err != nil {
		return nil, err
	}
	err = e.setEmployees(employees)
	if err != nil {
		return nil, err
	}
	return e, nil
}

// func (e *Enterprise) Changed() {
// 	e.director.Changed(e)
// }

func (e *Enterprise) setEmployees(employees mapset.Set[Employee]) error {
	if employees == nil {
		return &erros.NullValueError{}
	}
	e.employees = employees
	return nil
}

func (e *Enterprise) setOwner(owner string) error {
	if owner == "" {
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

func (e *Enterprise) setUpdatedAt(updatedAt common.Date) error {
	if updatedAt == (common.Date{}) {
		return &erros.EmptyStructError{}
	}
	e.updatedAt = updatedAt
	return nil
}

func (e *Enterprise) setFoundationDate(foundationDate common.Date) error {
	if foundationDate == (common.Date{}) {
		return &erros.EmptyStructError{}
	}
	e.foundationDate = foundationDate
	return nil
}

func (e Enterprise) LogoIMG() string {
	return e.logoIMG
}

func (e Enterprise) BannerIMG() string {
	return e.bannerIMG
}

func (e Enterprise) Name() string {
	return e.name
}

func (e Enterprise) Website() string {
	return e.website
}

func (e Enterprise) Email() string {
	return e.email
}

func (e Enterprise) Phone() string {
	return e.phone
}

func (e Enterprise) Address() common.Address {
	return e.address
}

func (e Enterprise) Industry() Industry {
	return e.industry
}

func (e Enterprise) CreatedAt() common.Date {
	return e.registerAt
}
func (e Enterprise) UpdatedAt() common.Date {
	return e.updatedAt
}
func (e Enterprise) FoundationDate() common.Date {
	return e.foundationDate
}
func (e Enterprise) Owner() string {
	return e.owner
}

func (e Enterprise) Employees() mapset.Set[Employee] {
	return e.employees
}

func (e *Enterprise) AddEmployee(employee Employee) error {
	if employee == nil {
		return &erros.NullValueError{}
	}
	e.employees.Add(employee)
	return nil
}
