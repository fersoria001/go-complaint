package repositories

import (
	"context"
	"go-complaint/domain/model/identity"
	"go-complaint/infrastructure/persistence/datasource"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/jackc/pgx/v5"
)

type UserRoleRepository struct {
	schema datasource.Schema
}

func NewUserRoleRepository(schema datasource.Schema) UserRoleRepository {
	return UserRoleRepository{schema: schema}
}
func (uur UserRoleRepository) RemoveAll(
	ctx context.Context,
	userID string,
) error {
	conn, err := uur.schema.Acquire(ctx)
	if err != nil {
		return err
	}
	deleteCommand := string(`
	DELETE FROM user_role
	WHERE user_id = $1
	  `)
	_, err = conn.Exec(ctx, deleteCommand, userID)
	if err != nil {
		return err
	}
	defer conn.Release()
	return nil
}
func (uur UserRoleRepository) SaveAll(
	ctx context.Context,
	userRole mapset.Set[identity.UserRole],
) error {
	conn, err := uur.schema.Acquire(ctx)
	if err != nil {
		return err
	}
	insertCommand := string(`
	INSERT INTO user_role
	(user_id, role_id, enterprise_id) VALUES ($1, $2, $3)`)
	tx, err := conn.Begin(ctx)
	if err != nil {
		return err
	}
	for userRole := range userRole.Iter() {
		_, err = tx.Exec(
			ctx,
			insertCommand,
			userRole.UserID(),
			userRole.GetRole().String(),
			userRole.EnterpriseID(),
		)
		if err != nil {
			tx.Rollback(ctx)
			return err
		}
	}
	err = tx.Commit(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()
	return nil
}

func (urr UserRoleRepository) FindAll(
	ctx context.Context,
	source StatementSource,
) (mapset.Set[*identity.UserRole], error) {
	conn, err := urr.schema.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	rows, err := conn.Query(
		ctx,
		source.Query(),
		source.Args()...,
	)
	if err != nil {
		return nil, err
	}
	userRoles, err := urr.loadAll(ctx, rows)
	if err != nil {
		return nil, err
	}
	defer func() {
		rows.Close()
		conn.Release()
	}()
	return userRoles, nil
}

func (uur UserRoleRepository) loadAll(
	ctx context.Context,
	rows pgx.Rows,
) (mapset.Set[*identity.UserRole], error) {
	userRoles := mapset.NewSet[*identity.UserRole]()
	for rows.Next() {
		userRole, err := uur.load(ctx, rows)
		if err != nil {
			return nil, err
		}
		userRoles.Add(userRole)
	}
	return userRoles, nil
}

func (uur UserRoleRepository) load(
	_ context.Context,
	row pgx.Row,
) (*identity.UserRole, error) {
	var (
		userID       string
		roleID       string
		enterpriseID string
	)
	err := row.Scan(
		&userID,
		&roleID,
		&enterpriseID,
	)
	if err != nil {
		return nil, err
	}
	role, err := identity.NewRole(roleID)
	if err != nil {
		return nil, err
	}
	userRole := identity.NewUserRole(role, userID, enterpriseID)
	return userRole, nil
}
