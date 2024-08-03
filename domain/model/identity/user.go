package identity

import (
	"context"
	"go-complaint/domain"
	"go-complaint/domain/model/common"
	"go-complaint/erros"
	"net/mail"
	"regexp"
	"slices"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/google/uuid"
)

// Package identity
// <<Aggregate>> User
type User struct {
	id           uuid.UUID
	registerDate common.Date
	userName     string
	password     string
	userRoles    mapset.Set[*UserRole]
	isConfirmed  bool
	*Person
}

func (user *User) RejectHiringInvitation(
	ctx context.Context,
	enterpriseId uuid.UUID,
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
			enterpriseId,
			user.Id(),
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
	enterpriseId uuid.UUID,
	proposedPosition RolesEnum,
) error {
	var (
		publisher = domain.DomainEventPublisherInstance()
		err       error
	)
	err = publisher.Publish(
		ctx,
		NewHiringInvitationAccepted(
			enterpriseId,
			user.Id(),
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

func (u User) Authorities() map[uuid.UUID][]GrantedAuthority {
	authorities := make(map[uuid.UUID][]GrantedAuthority, 0)
	for userRole := range u.userRoles.Iter() {
		if _, ok := authorities[userRole.EnterpriseId()]; !ok {
			authorities[userRole.EnterpriseId()] = make([]GrantedAuthority, 0)
		}
		authorities[userRole.EnterpriseId()] = append(
			authorities[userRole.EnterpriseId()],
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
	id uuid.UUID,
	userName string,
	password string,
	emailVerificationToken string,
	registerDate common.Date,
	person *Person,
) (*User, error) {
	emptyUserRoles := mapset.NewSet[*UserRole]()
	u, err := NewUser(
		id,
		userName,
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
	id uuid.UUID,
	userName string,
	password string,
	registerDate common.Date,
	person *Person,
	isConfirmed bool,
	userRoles mapset.Set[*UserRole],
) (*User, error) {
	var user = new(User)
	var err error
	user.id = id
	err = user.setRegisterDate(registerDate)
	if err != nil {
		return nil, err
	}
	err = user.setUserName(userName)
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
func (u *User) AddRoles(ctx context.Context, roles map[uuid.UUID][]Role) error {
	if roles == nil {
		return &erros.NullValueError{}
	}
	for enterpriseId, role := range roles {
		for _, r := range role {
			if !u.userRoles.Add(NewUserRole(r, u.id, enterpriseId)) {
				return &erros.AlreadyExistsError{}
			}
		}
	}
	return nil
}

func (u *User) RemoveUserRole(
	ctx context.Context,
	role RolesEnum,
	enterpriseId uuid.UUID) error {
	publisher := domain.DomainEventPublisherInstance()
	s := u.userRoles.ToSlice()
	s = slices.DeleteFunc(s, func(e *UserRole) bool {
		if e.enterpriseId == enterpriseId && e.Role.role == role {
			return true
		}
		return false
	})
	u.userRoles = mapset.NewSet(s...)
	publisher.Publish(
		ctx,
		NewRoleRemoved(u.Id(), enterpriseId, role.String()))
	return nil
}

func (u User) UserRoles() mapset.Set[UserRole] {
	value := mapset.NewSet[UserRole]()
	for userRole := range u.userRoles.Iter() {
		value.Add(*userRole)
	}
	return value
}

func (u *User) AddRole(
	ctx context.Context,
	role RolesEnum,
	enterpriseId uuid.UUID) error {
	publisher := domain.DomainEventPublisherInstance()
	newRole, err := NewRole(role.String())
	if err != nil {
		return err
	}
	userRole := NewUserRole(newRole, u.id, enterpriseId)
	if u.userRoles.Add(userRole) {
		err := publisher.Publish(
			ctx,
			NewRoleAdded(u.Id(), enterpriseId, userRole.GetRole().String()),
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
			u.Id(),
			enterpriseId,
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

func (u *User) setUserName(userName string) error {
	if userName == "" {
		return &erros.NullValueError{}
	}
	valid, err := mail.ParseAddress(userName)
	if err != nil {
		return &erros.InvalidEmailError{}
	}
	if valid == nil {
		return &erros.NullValueError{}
	}
	u.userName = userName
	return nil
}

func (u User) Id() uuid.UUID {
	return u.id
}

func (u User) RegisterDate() common.Date {
	return u.registerDate
}

func (u User) UserName() string {
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
