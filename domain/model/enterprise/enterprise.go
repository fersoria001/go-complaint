package enterprise

import (
	"context"
	"fmt"
	"go-complaint/domain"
	"go-complaint/domain/model/common"
	"go-complaint/domain/model/identity"
	"go-complaint/erros"
	"log"
	"net/mail"
	"slices"
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
	id             uuid.UUID
	name           string
	ownerId        uuid.UUID
	logoImg        string
	bannerImg      string
	website        string
	email          string
	phone          string
	address        common.Address
	industry       Industry
	registerAt     common.Date
	updatedAt      common.Date
	foundationDate common.Date
	employees      mapset.Set[*Employee]
}

func (e *Enterprise) PromoteEmployee(
	ctx context.Context,
	promotedBy string,
	employeeId uuid.UUID,
	newPosition Position) (*identity.User, error) {
	slice := e.employees.ToSlice()
	index, ok := slices.BinarySearchFunc(slice, employeeId, func(i *Employee, j uuid.UUID) int {
		if i.Id() == j {
			return 0
		}
		return -1
	})
	if !ok {
		return nil, fmt.Errorf("employee with id %s not found", employeeId)
	}
	emp := slice[index]
	if !emp.ApprovedHiring() {
		return nil, &erros.ValidationError{Expected: "employee needs to be approved first"}
	}
	err := emp.SetPosition(newPosition)
	if err != nil {
		return nil, err
	}
	user := emp.GetUser()
	role, err := identity.ParseRole(newPosition.String())
	if err != nil {
		return nil, err
	}
	urSlice := user.UserRoles().ToSlice()
	for _, r := range urSlice {
		if r.EnterpriseId() == e.id {
			err = user.RemoveUserRole(ctx, r.GetRole(), e.id)
			if err != nil {
				return nil, err
			}
		}
	}
	err = user.AddRole(ctx, role, e.id)
	if err != nil {
		return nil, err
	}
	event := NewEmployeePromoted(e.name, promotedBy, emp.Email(), newPosition)
	err = domain.DomainEventPublisherInstance().Publish(ctx, event)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (e *Enterprise) EmployeeLeave(ctx context.Context, employeeId uuid.UUID) (*identity.User, error) {
	var (
		publisher = domain.DomainEventPublisherInstance()
		err       error
	)
	var employee *Employee
	var user *identity.User
	for emp := range e.employees.Iter() {
		if employeeId == emp.Id() {
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
				e.Id(),
			)
			if err != nil {
				return nil, err
			}
			employee = emp
			e.employees.Remove(emp)
		}
	}

	err = publisher.Publish(ctx, NewEmployeeLeaved(employee))
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (e *Enterprise) FireEmployee(
	ctx context.Context,
	emitedBy uuid.UUID,
	employeeId uuid.UUID) (*identity.User, error) {
	var (
		publisher = domain.DomainEventPublisherInstance()
		err       error
	)
	slice := e.employees.ToSlice()
	index, ok := slices.BinarySearchFunc(slice, employeeId, func(i *Employee, j uuid.UUID) int {
		if i.Id() == j {
			return 0
		}
		return -1
	})
	if !ok {
		return nil, fmt.Errorf("employee with id %s not found", employeeId)
	}
	emp := slice[index]

	if !emp.ApprovedHiring() {
		return nil, &erros.ValidationError{Expected: "employee needs to be approved first"}
	}
	role, err := identity.ParseRole(emp.Position().String())
	if err != nil {
		return nil, err
	}
	user := emp.GetUser()
	err = emp.GetUser().RemoveUserRole(ctx,
		role,
		e.Id(),
	)
	if err != nil {
		return nil, err
	}
	log.Println("befor remove employee from enterprise", len(slice))
	slice = slices.DeleteFunc(slice, func(i *Employee) bool {
		return i.Id() == employeeId
	})
	log.Println("after remove employee from enterprise", len(slice))
	e.employees = mapset.NewSet(slice...)
	err = publisher.Publish(ctx, NewEmployeeFired(emitedBy, emp))
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (e *Enterprise) CancelHiringProccess(
	ctx context.Context,
	candidateId,
	emitedBy uuid.UUID,
	reason string,
	position Position,
) error {
	return domain.DomainEventPublisherInstance().Publish(
		ctx,
		NewHiringProccessCanceled(
			e.id,
			candidateId,
			emitedBy,
			reason,
			position,
		),
	)
}

func (e *Enterprise) HireEmployee(
	ctx context.Context,
	emitedBy uuid.UUID,
	employee *Employee,
) error {
	// role, err := identity.ParseRole(Position().String())

	// if err != nil {
	// 	return err
	// }
	// err = GetUser().AddRole(
	// 	ctx,
	// 	role,
	// 	e.Id(),
	// )
	// if err != nil {
	// 	return err
	// }
	employee.SetApprovedHiring(true)
	err := e.AddEmployee(employee)
	if err != nil {
		return err
	}
	domainEventPublisher := domain.DomainEventPublisherInstance()
	ev := NewEmployeeHired(
		e.Id(),
		emitedBy,
		employee.Id(),
		employee.Email(),
		employee.Position(),
	)
	err = domainEventPublisher.Publish(
		ctx,
		ev,
	)
	if err != nil {
		return err
	}

	return nil
}

func (e *Enterprise) InviteToProject(
	ctx context.Context,
	userId uuid.UUID,
	proposedTo uuid.UUID,
	proposalPosition Position,
) error {
	event := NewHiringInvitationSent(
		e.id,
		userId,
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
	id,
	ownerId uuid.UUID,
	name,
	logoImg,
	bannerImg,
	website,
	email,
	phone string,
	foundationDate common.Date,
	industry Industry,
	address common.Address,
) (*Enterprise, error) {
	regat := common.NewDate(time.Now())
	emptySet := mapset.NewSet[*Employee]()
	e, err := NewEnterprise(
		id,
		ownerId,
		name,
		logoImg,
		bannerImg,
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
		e.Id(),
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
	return e, err
}

// constructor
func NewEnterprise(
	id,
	ownerId uuid.UUID,
	name,
	logoImg,
	bannerImg,
	website,
	email,
	phone string,
	address common.Address,
	industry Industry,
	registerAt,
	updatedAt,
	foundationDate common.Date,
	employees mapset.Set[*Employee],
) (*Enterprise, error) {
	e := new(Enterprise)
	err := e.setName(name)
	if err != nil {
		return nil, err
	}
	e.id = id
	e.ownerId = ownerId
	err = e.setLogoIMG(logoImg)
	if err != nil {
		return nil, err
	}
	err = e.setBannerIMG(bannerImg)
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

func (e *Enterprise) setEmployees(employees mapset.Set[*Employee]) error {
	if employees == nil {
		return &erros.NullValueError{}
	}
	e.employees = employees
	return nil
}

func (e *Enterprise) setLogoIMG(logoIMG string) error {
	if logoIMG == "" {
		return &erros.NullValueError{}
	}
	e.logoImg = logoIMG
	return nil
}

func (e *Enterprise) setBannerIMG(bannerIMG string) error {
	if bannerIMG == "" {
		return &erros.NullValueError{}
	}
	e.bannerImg = bannerIMG
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

func (e Enterprise) Id() uuid.UUID {
	return e.id
}

func (e Enterprise) LogoIMG() string {
	return e.logoImg
}

func (e Enterprise) BannerIMG() string {
	return e.bannerImg
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
func (e Enterprise) OwnerId() uuid.UUID {
	return e.ownerId
}
func (e Enterprise) Employees() mapset.Set[Employee] {
	valueCopy := mapset.NewSet[Employee]()
	for v := range e.employees.Iter() {
		valueCopy.Add(*v)
	}
	return valueCopy
}
func (e *Enterprise) AddEmployee(employee *Employee) error {
	if employee == nil {
		return ErrNilPointer
	}
	e.employees.Add(employee)
	return nil
}
