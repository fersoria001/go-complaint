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
	registerDate common.Date
	email        string
	password     string
	userRoles    mapset.Set[*UserRole]
	isConfirmed  bool
	*Person
}

func (user *User) RejectHiringInvitation(
	ctx context.Context,
	enterpriseID string,
	rejectionReason string,
	proposedPosition RolesEnum,
) error {
	var (
		publisher = domain.DomainEventPublisherInstance()
		err       error
	)
	err = publisher.Publish(
		ctx,
		NewHiringInvitationRejected(
			enterpriseID,
			user.Email(),
			rejectionReason,
			proposedPosition,
		),
	)
	if err != nil {
		return err
	}
	return nil
}

func (user *User) AcceptHiringInvitation(
	ctx context.Context,
	enterpriseID string,
	proposedPosition RolesEnum,
) error {
	var (
		publisher = domain.DomainEventPublisherInstance()
		err       error
	)
	err = publisher.Publish(
		ctx,
		NewHiringInvitationAccepted(
			enterpriseID,
			user.Email(),
			proposedPosition,
		),
	)
	if err != nil {
		return err
	}
	return nil
}

func (user *User) SignIn(ctx context.Context, code int, ip string,
	latitude, longitude float64, device string) {
	domain.DomainEventPublisherInstance().Publish(
		ctx,
		NewUserSignedIn(
			user.Email(),
			code,
			ip,
			device,
			latitude,
			longitude,
			time.Now(),
		),
	)
}

// factory method and publisher

func (u User) Authorities() map[string][]GrantedAuthority {
	authorities := make(map[string][]GrantedAuthority, 0)
	for userRole := range u.userRoles.Iter() {
		if _, ok := authorities[userRole.EnterpriseID()]; !ok {
			authorities[userRole.EnterpriseID()] = make([]GrantedAuthority, 0)
		}
		authorities[userRole.EnterpriseID()] = append(
			authorities[userRole.EnterpriseID()],
			NewAuthority(
				userRole.GetRole().String(),
			),
		)
	}
	return authorities
}

func (u *User) VerifyEmail(ctx context.Context) error {
	publisher := domain.DomainEventPublisherInstance()
	if u.isConfirmed {
		return &erros.AlreadyExistsError{}
	}
	u.isConfirmed = true
	err := publisher.Publish(ctx, NewUserEmailVerified(u.Email()))
	if err != nil {
		return err
	}
	return nil
}

func CreateUser(
	ctx context.Context,
	email string,
	password string,
	emailVerificationToken string,
	registerDate common.Date,
	person *Person,
) (*User, error) {
	emptyUserRoles := mapset.NewSet[*UserRole]()
	u, err := NewUser(
		email,
		password,
		registerDate,
		person,
		true,
		emptyUserRoles,
	)
	if err != nil {
		return nil, err
	}
	publisher := domain.DomainEventPublisherInstance()
	event, err := NewUserCreated(u.Email(), emailVerificationToken, time.Now())
	if err != nil {
		return nil, err
	}
	err = publisher.Publish(ctx, event)
	if err != nil {
		return nil, err
	}
	return u, err
}

func NewUser(
	email string,
	password string,
	registerDate common.Date,
	person *Person,
	isConfirmed bool,
	userRoles mapset.Set[*UserRole],
) (*User, error) {
	var user = new(User)
	var err error
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
	user.isConfirmed = isConfirmed
	user.userRoles = userRoles
	return user, nil
}

/*
Intended to use to construct a User object, it's different from
the domain method AddRole that has more business logic
*/
func (u *User) AddRoles(ctx context.Context, roles map[string][]Role) error {
	if roles == nil {
		return &erros.NullValueError{}
	}
	for enterprise, role := range roles {
		for _, r := range role {
			if !u.userRoles.Add(NewUserRole(r, u.email, enterprise)) {
				return &erros.AlreadyExistsError{}
			}
		}
	}

	return nil
}
func (u *User) RemoveUserRole(
	ctx context.Context,
	role RolesEnum,
	enterpriseID string) error {
	publisher := domain.DomainEventPublisherInstance()
	target, err := u.findUserRole(role, enterpriseID)
	if err != nil {
		return err
	}
	u.userRoles.Remove(target)
	publisher.Publish(
		ctx,
		NewRoleRemoved(u.Email(), enterpriseID, role.String()))
	return nil
}

func (u User) UserRoles() mapset.Set[UserRole] {
	value := mapset.NewSet[UserRole]()
	for userRole := range u.userRoles.Iter() {
		value.Add(*userRole)
	}
	return value
}

func (u *User) findUserRole(
	role RolesEnum,
	enterpriseID string) (*UserRole, error) {
	for userRole := range u.userRoles.Iter() {
		if userRole.GetRole().String() == role.String() && userRole.EnterpriseID() == enterpriseID {
			return userRole, nil
		}
	}
	return nil, ErrUserRoleNotFound
}

func (u *User) AddRole(
	ctx context.Context,
	role RolesEnum,
	enterpriseID string) error {
	publisher := domain.DomainEventPublisherInstance()
	newRole, err := NewRole(role.String())
	if err != nil {
		return err
	}
	userRole := NewUserRole(newRole, u.email, enterpriseID)
	if u.userRoles.Add(userRole) {
		err := publisher.Publish(
			ctx,
			NewRoleAdded(u.Email(), enterpriseID, userRole.GetRole().String()),
		)
		if err != nil {
			return err
		}
	} else {
		return &erros.AlreadyExistsError{}
	}
	domain.DomainEventPublisherInstance().Publish(
		ctx,
		NewRoleAdded(
			u.Email(),
			enterpriseID,
			role.String(),
		),
	)
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

func (u *User) setPerson(person *Person) error {
	if person == nil {
		return &erros.NullValueError{}
	}
	u.Person = person
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

func (u User) ProfileIMG() string {
	return u.profileIMG
}

func (u User) RegisterDate() common.Date {
	return u.registerDate
}

func (u User) Email() string {
	return u.email
}

func (u User) Password() string {
	return u.password
}

func (u User) GetPerson() Person {
	person := *u.Person
	return person
}

func (u User) IsConfirmed() bool {
	return u.isConfirmed
}
