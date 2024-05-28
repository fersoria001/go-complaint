package identity

import (
	"context"
	"go-complaint/domain"
	"go-complaint/domain/model/common"
	"go-complaint/erros"
	"net/mail"
	"regexp"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
)

// Package identity
// <<Aggregate>> User
type User struct {
	profileIMG   string
	registerDate common.Date
	email        string
	password     string
	person       *Person
	userRoles    mapset.Set[*UserRole]
}

func (u *User) GetAuthorities(enterprise string) []GrantedAuthority {
	authorities := make([]GrantedAuthority, 0)
	for userRole := range u.userRoles.Iter() {
		if userRole.enterprise == enterprise {
			authorities = append(authorities, NewAuthority(userRole.role.ID()))
		}
	}
	return authorities
}

func CreateUser(ctx context.Context, profileIMG string, registerDate common.Date, email string, password string,
	person *Person) (*User, error) {
	u, err := NewUser(profileIMG, registerDate, email, password, person)
	if err != nil {
		return nil, err
	}

	publisher := domain.DomainEventPublisherInstance()
	event, err := NewUserCreated(u.Email(), time.Now())
	if err != nil {
		return nil, err
	}
	err = publisher.Publish(ctx, event)
	if err != nil {
		return nil, err
	}
	return u, err
}

func NewUser(profileIMG string, registerDate common.Date, email string, password string,
	person *Person) (*User, error) {
	var user = new(User)
	var err error
	err = user.setProfileIMG(profileIMG)
	if err != nil {
		return nil, err
	}
	err = user.setRegisterDate(registerDate)
	if err != nil {
		return nil, err
	}
	err = user.setEmail(email)
	if err != nil {
		return nil, err
	}
	err = user.setPassword(password)
	if err != nil {
		return nil, err
	}
	err = user.setPerson(person)
	if err != nil {
		return nil, err
	}
	user.userRoles = mapset.NewSet[*UserRole]()
	return user, nil
}

/*
Intended to use to construct a User object, it's different from
the domain method AddRole that has more business logic
*/
func (u *User) AddRoles(ctx context.Context, roles map[string][]*Role) error {
	if roles == nil {
		return &erros.NullValueError{}
	}
	for enterprise, role := range roles {
		for _, r := range role {
			if !u.userRoles.Add(NewUserRole(r, u, enterprise)) {
				return &erros.AlreadyExistsError{}
			}
		}
	}

	return nil
}
func (u *User) RemoveUserRole(
	ctx context.Context,
	userRole *UserRole) error {
	publisher := domain.DomainEventPublisherInstance()
	var targetRole *UserRole
	for item := range u.userRoles.Iter() {
		if item.Equals(userRole) {
			targetRole = item
		}
	}
	if targetRole == nil {
		return &erros.NoElementError{}
	}
	u.userRoles.Remove(targetRole)
	publisher.Publish(
		ctx,
		NewRoleRemoved(u.Email(), userRole.Enterprise(), userRole.Role().ID()))

	return nil
}

func (u *User) AddRole(
	ctx context.Context,
	role *Role,
	enterprise string) error {
	if role == nil {
		return &erros.NullValueError{}
	}
	publisher := domain.DomainEventPublisherInstance()
	userRole := NewUserRole(role, u, enterprise)
	if u.userRoles.Add(userRole) {
		err := publisher.Publish(
			ctx,
			NewRoleAdded(u.Email(), enterprise, role.ID()),
		)
		if err != nil {
			return err
		}
	} else {
		return &erros.AlreadyExistsError{}
	}
	return nil
}

func (u *User) ResetPassword(ctx context.Context, randomGeneratedPassword, hash string) error {
	publisher := domain.DomainEventPublisherInstance()
	err := u.setPassword(hash)
	if err != nil {
		return err
	}
	event, err := NewPasswordReset(u.Email(), randomGeneratedPassword)
	if err != nil {
		return err
	}
	return publisher.Publish(ctx, event)
}
func (u *User) ChangePassword(ctx context.Context, oldPassword, newPassword, newHash string) error {
	publisher := domain.DomainEventPublisherInstance()
	if oldPassword == newPassword {
		return &erros.ValidationError{
			Expected: "old password and new password must be different",
		}
	}
	err := u.setPassword(newHash)
	if err != nil {
		return err
	}
	return publisher.Publish(ctx, NewPasswordChanged(u.Email(), oldPassword))
}

func (u *User) ChangePersonalData(ctx context.Context, profileIMG, firstName, lastName, phone, country, county, city string) error {
	var (
		publisher = domain.DomainEventPublisherInstance()
		err       error
		address   common.Address
		oldValues map[string]string = make(map[string]string)
		newValues map[string]string = make(map[string]string)
	)
	if profileIMG != u.ProfileIMG() && profileIMG != "" {
		oldValues["profileIMG"] = u.ProfileIMG()
		newValues["profileIMG"] = profileIMG
		err = u.setProfileIMG(profileIMG)
		if err != nil {
			return err
		}
	}
	if firstName != u.Person().FirstName() && firstName != "" {
		oldValues["firstName"] = u.Person().FirstName()
		newValues["firstName"] = firstName
		err = u.Person().setFirstName(firstName)
		if err != nil {
			return err
		}
	}
	if lastName != u.Person().LastName() && lastName != "" {
		oldValues["lastName"] = u.Person().LastName()
		newValues["lastName"] = lastName
		err = u.Person().setLastName(lastName)
		if err != nil {
			return err
		}
	}
	if phone != u.Person().Phone() && phone != "" {
		oldValues["phone"] = u.Person().Phone()
		newValues["phone"] = phone
		err = u.Person().setPhone(phone)
		if err != nil {
			return err
		}
	}
	if country != "" || county != "" || city != "" {
		if country == "" {
			country = u.Person().Address().Country()
		}
		if county == "" {
			county = u.Person().Address().County()
		}
		if city == "" {
			city = u.Person().Address().City()
		}
		oldValues["country"] = u.Person().Address().Country()
		oldValues["county"] = u.Person().Address().County()
		oldValues["city"] = u.Person().Address().City()
		address, err = common.NewAddress(country, county, city)
		newValues["country"] = address.Country()
		newValues["county"] = address.County()
		newValues["city"] = address.City()
		if err != nil {
			return err
		}
		err = u.Person().setAddress(address)
		if err != nil {
			return err
		}
	}
	if len(oldValues) != 0 || len(newValues) != 0 {
		publisher.Publish(context.Background(), NewUserUpdated(u.Email(), oldValues, newValues))
	}
	return nil
}

func (u *User) setPerson(person *Person) error {
	if person == nil {
		return &erros.NullValueError{}
	}
	u.person = person
	return nil
}
func (u *User) setPassword(password string) error {
	regex := regexp.MustCompile(`^(.{0,7}|[^0-9]*|[^A-Z]*|[^a-z]*)$`)
	if regex.MatchString(password) {
		return &erros.InvalidPasswordError{}
	}
	u.password = password
	return nil
}

func (u *User) setProfileIMG(profileIMG string) error {
	if profileIMG == "" {
		return &erros.NullValueError{}
	}
	u.profileIMG = profileIMG
	return nil
}

func (u *User) setRegisterDate(registerDate common.Date) error {
	if registerDate == (common.Date{}) {
		return &erros.NullValueError{}
	}
	u.registerDate = registerDate
	return nil
}

func (u *User) setEmail(email string) error {
	if email == "" {
		return &erros.NullValueError{}
	}
	valid, err := mail.ParseAddress(email)
	if err != nil {
		return &erros.InvalidEmailError{}
	}
	if valid == nil {
		return &erros.NullValueError{}
	}
	u.email = email
	return nil
}

func (u *User) ProfileIMG() string {
	return u.profileIMG
}

func (u *User) RegisterDate() common.Date {
	return u.registerDate
}

func (u User) Email() string {
	return u.email
}

func (u *User) Password() string {
	return u.password
}

func (u *User) Person() *Person {
	return u.person
}

func (u *User) UserRoles() mapset.Set[UserRole] {
	value := mapset.NewSet[UserRole]()
	for userRole := range u.userRoles.Iter() {
		value.Add(*userRole)
	}
	return value
}
