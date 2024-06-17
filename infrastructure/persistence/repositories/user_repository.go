package repositories

import (
	"context"
	"errors"
	"go-complaint/domain/model/common"
	"go-complaint/domain/model/identity"
	"go-complaint/infrastructure/persistence/datasource"
	userrolefindall "go-complaint/infrastructure/persistence/finders/user_role_findall"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/jackc/pgx/v5"
)

// Package persistence
// implements the repository interface
type UserRepository struct {
	schema datasource.Schema
}

func NewUserRepository(schema datasource.Schema) UserRepository {
	return UserRepository{schema: schema}
}

func (ur UserRepository) Update(
	ctx context.Context,
	user *identity.User,
) error {
	conn, err := ur.schema.Acquire(ctx)
	if err != nil {
		return err
	}
	insertCommand := string(`
		UPDATE public."user"
		SET
			password = $2,
			register_date = $3,
			is_confirmed = $4
		WHERE email = $1
		`)
	var (
		email        string = user.Email()
		password     string = user.Password()
		registerDate string = user.RegisterDate().StringRepresentation()
		isConfirmed  bool   = user.IsConfirmed()
	)
	_, err = conn.Exec(
		ctx,
		insertCommand,
		&email,
		&password,
		&registerDate,
		&isConfirmed,
	)
	if err != nil {
		return err
	}
	mapper := MapperRegistryInstance().Get("Person")
	if mapper == nil {
		return ErrMapperNotRegistered
	}
	personRepository, ok := mapper.(PersonRepository)
	if !ok {
		return ErrWrongTypeAssertion
	}
	mapper = MapperRegistryInstance().Get("UserRole")
	if mapper == nil {
		return ErrMapperNotRegistered
	}
	userRoleRepository, ok := mapper.(UserRoleRepository)
	if !ok {
		return ErrWrongTypeAssertion
	}
	err = personRepository.Update(ctx, user.GetPerson())
	if err != nil {
		return err
	}
	err = userRoleRepository.RemoveAll(ctx, user.Email())
	if err != nil {
		return err
	}
	err = userRoleRepository.SaveAll(ctx, user.UserRoles())
	if err != nil {
		return err
	}

	return nil
}

func (ur UserRepository) Save(
	ctx context.Context,
	user *identity.User,
) error {
	conn, err := ur.schema.Acquire(ctx)
	if err != nil {
		return err
	}
	//
	insertCommand := string(`
		INSERT INTO 
		public."user" (
			email,
			password,
			register_date,
			is_confirmed
		)
		VALUES ($1, $2, $3, $4)`)
	var (
		email        string = user.Email()
		password     string = user.Password()
		registerDate string = user.RegisterDate().StringRepresentation()
		isConfirmed  bool   = user.IsConfirmed()
	)
	_, err = conn.Exec(
		ctx,
		insertCommand,
		&email,
		&password,
		&registerDate,
		&isConfirmed,
	)
	if err != nil {
		return err
	}
	mapper := MapperRegistryInstance().Get("Person")
	if mapper == nil {
		return ErrMapperNotRegistered
	}
	personRepository, ok := mapper.(PersonRepository)
	if !ok {
		return ErrWrongTypeAssertion
	}
	mapper = MapperRegistryInstance().Get("UserRole")
	if mapper == nil {
		return ErrMapperNotRegistered
	}
	userRoleRepository, ok := mapper.(UserRoleRepository)
	if !ok {
		return ErrWrongTypeAssertion
	}
	err = userRoleRepository.SaveAll(ctx, user.UserRoles())
	if err != nil {
		return err
	}
	err = personRepository.Save(ctx, user.GetPerson())
	if err != nil {
		return err
	}
	return nil
}

func (ur UserRepository) Get(
	ctx context.Context,
	userEmail string,
) (*identity.User, error) {
	conn, err := ur.schema.Acquire(ctx)
	if err != nil {

		return nil, err
	}
	//
	selectQuery := string(`
		SELECT 
		email,
		password,
		register_date,
		is_confirmed
		FROM public."user"
		WHERE email = $1
		`)

	row := conn.QueryRow(
		ctx,
		selectQuery,
		userEmail,
	)
	defer conn.Release()
	return ur.load(ctx, row)
}
func (ur UserRepository) FindAll(
	ctx context.Context,
	source StatementSource,
) (mapset.Set[identity.User], error) {
	conn, err := ur.schema.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	rows, err := conn.Query(ctx, source.Query(), source.Args()...)
	if err != nil {
		return nil, err
	}
	users, err := ur.loadAll(ctx, rows)
	if err != nil {
		return nil, err
	}
	defer func() {
		rows.Close()
		conn.Release()
	}()
	return users, nil
}
func (ur UserRepository) loadAll(
	ctx context.Context,
	rows pgx.Rows,
) (mapset.Set[identity.User], error) {
	users := mapset.NewSet[identity.User]()
	for rows.Next() {
		user, err := ur.load(ctx, rows)
		if err != nil {
			return nil, err
		}
		users.Add(*user)
	}
	return users, nil
}
func (ur UserRepository) load(
	ctx context.Context,
	row pgx.Row,
) (*identity.User, error) {
	var (
		email        string
		password     string
		registerDate string
		isConfirmed  bool
	)
	err := row.Scan(
		&email,
		&password,
		&registerDate,
		&isConfirmed,
	)
	if err != nil {

		return nil, err
	}
	mapper := MapperRegistryInstance().Get("Person")
	if mapper == nil {
		return nil, ErrMapperNotRegistered
	}
	personRepository, ok := mapper.(PersonRepository)
	if !ok {
		return nil, ErrWrongTypeAssertion
	}
	mapper = MapperRegistryInstance().Get("UserRole")
	if mapper == nil {

		return nil, ErrMapperNotRegistered
	}
	userRoleRepository, ok := mapper.(UserRoleRepository)
	if !ok {
		return nil, ErrWrongTypeAssertion
	}
	person, err := personRepository.Get(ctx, email)
	if err != nil {
		return nil, err
	}
	commonRegisterDate, err := common.NewDateFromString(registerDate)
	if err != nil {
		return nil, err
	}
	userRoles, err := userRoleRepository.FindAll(
		ctx,
		userrolefindall.NewByUserID(email),
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			userRoles = mapset.NewSet[*identity.UserRole]()
		} else {
			return nil, err
		}
	}
	return identity.NewUser(
		email,
		password,
		commonRegisterDate,
		person,
		isConfirmed,
		userRoles,
	)
}
