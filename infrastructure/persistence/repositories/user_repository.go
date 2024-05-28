package repositories

import (
	"context"
	"fmt"
	"go-complaint/domain/model/common"
	"go-complaint/domain/model/identity"
	"go-complaint/dto"
	"go-complaint/infrastructure/persistence/datasource"
	"go-complaint/infrastructure/persistence/models"
	"log"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Package persistence
// implements the repository interface
type UserRepository struct {
	schema *datasource.Schema
}

func NewUserRepository(schema *datasource.Schema) *UserRepository {
	return &UserRepository{schema: schema}
}

func (r *UserRepository) Save(ctx context.Context, user *identity.User) error {
	var conn, err = r.schema.Pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()
	var userModel = models.NewUser(user)
	var personModel = models.NewPerson(user.Person())
	insertUser := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)",
		userModel.Table(), models.StringColumns(userModel.Columns()), userModel.Args())
	insertPerson := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)",
		personModel.Table(), models.StringColumns(personModel.Columns()), personModel.Args())
	tx, err := conn.Begin(ctx)
	if err != nil {
		return err
	}
	_, err = tx.Exec(ctx, insertUser,
		userModel.Values()...)
	if err != nil {
		return err
	}
	_, err = tx.Exec(ctx, insertPerson,
		personModel.Values()...)
	if err != nil {
		tx.Rollback(ctx)
		return err
	}
	err = tx.Commit(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) GetAll(ctx context.Context) ([]*identity.User, error) {
	var conn, err = r.schema.Pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	var (
		userModel            = &models.User{}
		personModel          = &models.Person{}
		aliasedUserColumns   = models.Alias(userModel.Columns(), userModel.Table())
		aliasedPersonColumns = models.Alias(personModel.Columns(), personModel.Table())
		joinUserPerson       = fmt.Sprintf("SELECT %s, %s FROM %s JOIN %s ON %s.email = %s.email",
			models.StringColumns(aliasedUserColumns), models.StringColumns(aliasedPersonColumns),
			userModel.Table(), personModel.Table(), userModel.Table(), personModel.Table(),
		)
	)
	rows, err := conn.Query(ctx, joinUserPerson)
	if err != nil {
		return nil, err
	}
	users := make([]*identity.User, 0)
	for rows.Next() {
		err = rows.Scan(append(userModel.Values(), personModel.Values()...)...)
		if err != nil {
			return nil, err
		}
		birthDate, err := common.NewDateFromString(personModel.BirthDate)
		if err != nil {
			return nil, err
		}
		address, err := common.NewAddress(personModel.Country, personModel.County, personModel.City)
		if err != nil {
			return nil, err
		}
		person, err := identity.NewPerson(
			personModel.Email,
			personModel.FirstName,
			personModel.LastName,
			personModel.Phone,
			birthDate,
			address,
		)
		if err != nil {
			return nil, err
		}
		registerDate, err := common.NewDateFromString(userModel.RegisterDate)
		if err != nil {
			return nil, err
		}
		user, err := identity.NewUser(
			userModel.ProfileIMG,
			registerDate,
			userModel.Email,
			userModel.Password,
			person,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *UserRepository) Get(ctx context.Context, id string) (*identity.User, error) {
	var conn, err = r.schema.Pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	var userModel = &models.User{}
	var personModel = &models.Person{}
	var userRole = &models.UserRole{}
	var selectUser = fmt.Sprintf("SELECT %s FROM %s WHERE email = $1", models.StringColumns(userModel.Columns()), userModel.Table())
	var selectPerson = fmt.Sprintf("SELECT %s FROM %s WHERE email = $1", models.StringColumns(personModel.Columns()), personModel.Table())
	var selectUserRoles = fmt.Sprintf("SELECT %s FROM %s WHERE user_id = $1", models.StringColumns(userRole.Columns()), userRole.Table())
	row := conn.QueryRow(ctx, selectUser, id)
	err = row.Scan(userModel.Values()...)
	if err != nil {
		return nil, err
	}
	row = conn.QueryRow(ctx, selectPerson, id)
	err = row.Scan(personModel.Values()...)
	if err != nil {
		return nil, err
	}
	birthDate, err := common.NewDateFromString(personModel.BirthDate)
	if err != nil {
		return nil, err
	}
	address, err := common.NewAddress(personModel.Country, personModel.County, personModel.City)
	if err != nil {
		return nil, err
	}
	person, err := identity.NewPerson(
		personModel.Email,
		personModel.FirstName,
		personModel.LastName,
		personModel.Phone,
		birthDate,
		address,
	)
	if err != nil {
		return nil, err
	}
	registerDate, err := common.NewDateFromString(userModel.RegisterDate)
	if err != nil {
		return nil, err
	}
	user, err := identity.NewUser(
		userModel.ProfileIMG,
		registerDate,
		userModel.Email,
		userModel.Password,
		person,
	)
	if err != nil {
		return nil, err
	}
	userRoles := make(map[string][]*identity.Role, 0)
	rows, err := conn.Query(ctx, selectUserRoles, id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {

		err = rows.Scan(userRole.Values()...)
		if err != nil {
			return nil, err
		}
		parsedRole, err := identity.ParseRole(userRole.RoleID)
		if err != nil {
			return nil, err
		}
		var role = identity.NewRole(parsedRole)
		if _, ok := userRoles[userRole.EnterpriseID]; !ok {
			userRoles[userRole.EnterpriseID] = make([]*identity.Role, 0)
		}
		userRoles[userRole.EnterpriseID] = append(userRoles[userRole.EnterpriseID], role)
	}
	err = user.AddRoles(ctx, userRoles)
	if err != nil {
		return nil, err
	}

	return user, nil

}

func (r *UserRepository) Remove(ctx context.Context, id string) error {
	var conn, err = r.schema.Pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()
	var personModel = &models.Person{}
	var userModel = &models.User{}
	var userRoleModel = &models.UserRole{}
	var deletePerson = fmt.Sprintf(`DELETE FROM %s WHERE email = $1`, personModel.Table())
	var deleteUser = fmt.Sprintf(`DELETE FROM %s WHERE email = $1`, userModel.Table())
	var deleteUserRoles = fmt.Sprintf(`DELETE FROM %s WHERE user_id = $1`, userRoleModel.Table())
	tx, err := conn.Begin(ctx)
	if err != nil {
		return err
	}
	_, err = tx.Exec(ctx, deletePerson, id)
	if err != nil {
		return err
	}
	_, err = tx.Exec(ctx, deleteUser, id)
	if err != nil {
		tx.Rollback(ctx)
		return err
	}
	_, err = tx.Exec(ctx, deleteUserRoles, id)
	if err != nil {
		tx.Rollback(ctx)
		return err
	}
	err = tx.Commit(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) Update(ctx context.Context, user *identity.User) error {
	var (
		conn           *pgxpool.Conn
		err            error
		userModel      = models.NewUser(user)
		personModel    = models.NewPerson(user.Person())
		userRolesModel = models.NewUserRoles(user.UserRoles())
		userRoleModel  = &models.UserRole{}
		updateUser     = fmt.Sprintf("UPDATE %s SET %s WHERE email = $3",
			userModel.Table(), models.StringKeyArgs(userModel.Columns()))
		updatePerson = fmt.Sprintf("UPDATE %s SET %s WHERE email = $1",
			personModel.Table(), models.StringKeyArgs(personModel.Columns()))
		deleteUserRoles = fmt.Sprintf("DELETE FROM %s WHERE user_id = $1", userRoleModel.Table())
		insertUserRole  = fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", userRoleModel.Table(),
			models.StringColumns(userRoleModel.Columns()), userRoleModel.Args())
	)
	conn, err = r.schema.Pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()
	tx, err := conn.Begin(ctx)
	if err != nil {
		return err
	}
	_, err = tx.Exec(ctx, updateUser, userModel.Values()...)

	if err != nil {

		tx.Rollback(ctx)
		return err
	}
	_, err = tx.Exec(ctx, updatePerson, personModel.Values()...)
	if err != nil {

		tx.Rollback(ctx)
		return err
	}
	_, err = tx.Exec(ctx, deleteUserRoles, user.Email())

	if err != nil {

		tx.Rollback(ctx)
		return err
	}
	for _, ur := range userRolesModel {
		_, err = tx.Exec(ctx, insertUserRole, ur.Values()...)

		if err != nil {
			tx.Rollback(ctx)
			return err
		}
	}
	err = tx.Commit(ctx)
	if err != nil {
		return err
	}

	return nil

}

//Addittional behaviour

func (r *UserRepository) FindByName(ctx context.Context, term string) ([]*dto.Receiver, error) {
	conn, err := r.schema.Pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	var findPerson = `
	SELECT P.FIRST_NAME,
       P.LAST_NAME,
       P.EMAIL,
	   U.PROFILE_IMG
FROM PERSONS P
JOIN USERS U ON
U.EMAIL = P.EMAIL
WHERE LOWER(P.FIRST_NAME) LIKE $1 OR
LOWER(P.LAST_NAME) LIKE $1;
`
	wildCardTerm := "%" + strings.ToLower(term) + "%"
	rows, err := conn.Query(ctx, findPerson, wildCardTerm)
	if err != nil {
		return nil, err
	}
	receivers := make([]*dto.Receiver, 0)
	var firstName string
	var lastName string
	for rows.Next() {
		var receiver = &dto.Receiver{}
		err = rows.Scan(&firstName, &lastName, &receiver.ID, &receiver.IMG)
		if err != nil {
			return nil, err
		}
		receiver.FullName = fmt.Sprintf("%s %s", firstName, lastName)
		receivers = append(receivers, receiver)
	}

	return receivers, nil
}
func arrayParams(ids []string) (string, int) {
	var paramrefs string
	var lastParamIndex int
	for i := range ids {
		paramrefs += `$` + strconv.Itoa(i+1) + `,`
		lastParamIndex = i + 1
	}
	paramrefs = paramrefs[:len(paramrefs)-1]
	return paramrefs, lastParamIndex
}
func castToAny(ids []string) []interface{} {
	emailInterfaces := make([]interface{}, len(ids))
	for i, v := range ids {
		emailInterfaces[i] = v
	}
	return emailInterfaces
}
func (r *UserRepository) FindByNotEmailsAndName(ctx context.Context, name string, emails []string, limit, offset int) ([]*identity.User, int, error) {
	var conn, err = r.schema.Pool.Acquire(ctx)

	if err != nil {
		return nil, 0, err
	}
	defer conn.Release()
	arrayString, lastParamIndex := arrayParams(emails)
	var (
		userModel            = &models.User{}
		personModel          = &models.Person{}
		aliasedUserColumns   = models.Alias(userModel.Columns(), userModel.Table())
		aliasedPersonColumns = models.Alias(personModel.Columns(), personModel.Table())
		joinUserPerson       = fmt.Sprintf(`SELECT %s, %s FROM %s  JOIN %s  ON %s.email = %s.email
		WHERE LOWER(CONCAT(%s.first_name,' ', %s.last_name)) LIKE $%d AND %s.email NOT in (%s) LIMIT $%d OFFSET $%d`,
			models.StringColumns(aliasedUserColumns), models.StringColumns(aliasedPersonColumns),
			userModel.Table(), personModel.Table(), userModel.Table(), personModel.Table(),
			personModel.Table(), personModel.Table(), lastParamIndex+1,
			userModel.Table(), arrayString, lastParamIndex+2, lastParamIndex+3,
		)
		countQuery = fmt.Sprintf(`SELECT COUNT(*) FROM %s WHERE LOWER(CONCAT(%s.first_name,' ', %s.last_name))  LIKE $%d  AND %s.email NOT in (%s)`,
			personModel.Table(), personModel.Table(), personModel.Table(), lastParamIndex+1, personModel.Table(), arrayString)
	)
	var count int
	name = "%" + strings.ToLower(name) + "%"
	arrayAndName := append(castToAny(emails), name)
	err = conn.QueryRow(ctx, countQuery, arrayAndName...).Scan(&count)
	if err != nil {
		log.Printf("Query %s, array %v, err %v", countQuery, arrayAndName, err)
		return nil, 0, err
	}
	if count == 0 {
		return nil, 0, nil
	}
	allArgs := append(arrayAndName, limit, offset)
	rows, err := conn.Query(ctx, joinUserPerson, allArgs...)
	if err != nil {
		log.Printf("Query %s, array %v", joinUserPerson, allArgs)
		return nil, count, err
	}
	users := make([]*identity.User, 0)
	for rows.Next() {
		err = rows.Scan(append(userModel.Values(), personModel.Values()...)...)
		if err != nil {
			return nil, count, err
		}
		birthDate, err := common.NewDateFromString(personModel.BirthDate)
		if err != nil {
			return nil, count, err
		}
		address, err := common.NewAddress(personModel.Country, personModel.County, personModel.City)
		if err != nil {
			return nil, count, err
		}
		person, err := identity.NewPerson(
			personModel.Email,
			personModel.FirstName,
			personModel.LastName,
			personModel.Phone,
			birthDate,
			address,
		)
		if err != nil {
			return nil, count, err
		}
		registerDate, err := common.NewDateFromString(userModel.RegisterDate)
		if err != nil {
			return nil, count, err
		}
		user, err := identity.NewUser(
			userModel.ProfileIMG,
			registerDate,
			userModel.Email,
			userModel.Password,
			person,
		)
		if err != nil {
			return nil, count, err
		}
		users = append(users, user)
	}

	return users, count, nil
}
