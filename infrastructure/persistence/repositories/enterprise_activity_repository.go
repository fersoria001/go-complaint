package repositories

import (
	"context"
	"fmt"
	"go-complaint/domain/model/common"
	"go-complaint/domain/model/enterprise"
	"go-complaint/infrastructure/persistence/datasource"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type EnterpriseActivityRepository struct {
	schema datasource.Schema
}

func NewEnterpriseActivityRepository(schema datasource.Schema) EnterpriseActivityRepository {
	return EnterpriseActivityRepository{
		schema: schema,
	}
}

func (ear EnterpriseActivityRepository) Save(ctx context.Context, enterpriseActivity enterprise.EnterpriseActivity) error {
	conn, err := ear.schema.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()
	insertCommand := string(`
	INSERT INTO ENTERPRISE_ACTIVITY 
	(ID, USER_ID, ACTIVITY_ID, ENTERPRISE_ID, ENTERPRISE_NAME, OCCURRED_ON, ACTIVITY_TYPE) VALUES 
	($1, $2, $3, $4, $5, $6, $7)`)
	var (
		id             = enterpriseActivity.Id()
		userId         = enterpriseActivity.User().Id()
		activityId     = enterpriseActivity.ActivityId()
		enterpriseId   = enterpriseActivity.EnterpriseId()
		enterpriseName = enterpriseActivity.EnterpriseName()
		occurredOn     = common.StringDate(enterpriseActivity.OccurredOn())
		activityType   = enterpriseActivity.ActivityType().String()
	)
	_, err = conn.Exec(ctx, insertCommand, &id, &userId, &activityId, &enterpriseId, &enterpriseName, &occurredOn, &activityType)
	if err != nil {
		return err
	}
	return nil
}

func (ear EnterpriseActivityRepository) Get(ctx context.Context, id uuid.UUID) (*enterprise.EnterpriseActivity, error) {
	conn, err := ear.schema.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	selectQuery := string(`
	SELECT
		ID,
		USER_ID,
		ACTIVITY_ID,
		ENTERPRISE_ID,
		ENTERPRISE_NAME,
		OCCURRED_ON,
		ACTIVITY_TYPE
	FROM ENTERPRISE_ACTIVITY
	WHERE ID = $1
	`)
	row := conn.QueryRow(ctx, selectQuery, &id)
	return ear.load(ctx, row)
}

func (ear EnterpriseActivityRepository) Find(ctx context.Context, src StatementSource) (*enterprise.EnterpriseActivity, error) {
	conn, err := ear.schema.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	row := conn.QueryRow(ctx, src.Query(), src.Args()...)
	return ear.load(ctx, row)
}

func (ear EnterpriseActivityRepository) Remove(ctx context.Context, id uuid.UUID) error {
	conn, err := ear.schema.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()
	removeCommand := string(`DELETE FROM ENTERPRISE_ACTIVITY WHERE ID = $1`)
	_, err = conn.Exec(ctx, removeCommand, &id)
	if err != nil {
		return err
	}
	return nil
}

func (ear EnterpriseActivityRepository) FindAll(ctx context.Context, src StatementSource) ([]*enterprise.EnterpriseActivity, error) {
	conn, err := ear.schema.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	rows, err := conn.Query(ctx, src.Query(), src.Args()...)
	if err != nil {
		return nil, err
	}
	results, err := ear.loadAll(ctx, rows)
	if err != nil {
		return nil, err
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}
	rows.Close()
	return results, nil
}

func (ear EnterpriseActivityRepository) load(ctx context.Context, row pgx.Row) (*enterprise.EnterpriseActivity, error) {
	var (
		id             uuid.UUID
		userId         uuid.UUID
		activityId     uuid.UUID
		enterpriseId   uuid.UUID
		enterpriseName string
		occurredOn     string
		activityType   string
	)
	err := row.Scan(&id, &userId, &activityId, &enterpriseId, &enterpriseName, &occurredOn, &activityType)
	if err != nil {
		return nil, err
	}
	date, err := common.ParseDate(occurredOn)
	if err != nil {
		return nil, err
	}
	activity := enterprise.ParseActivityType(activityType)
	if activity < 0 {
		return nil, fmt.Errorf("unexpected enterprise activity string in enterprise activity data mapper")
	}
	reg := MapperRegistryInstance()
	recipientRepository, ok := reg.Get("Recipient").(RecipientRepository)
	if !ok {
		return nil, ErrWrongTypeAssertion
	}
	user, err := recipientRepository.Get(ctx, userId)
	if err != nil {
		return nil, err
	}
	return enterprise.NewEnterpriseActivity(id, activityId, enterpriseId, *user, enterpriseName, date, activity), nil
}

func (ear EnterpriseActivityRepository) loadAll(ctx context.Context, rows pgx.Rows) ([]*enterprise.EnterpriseActivity, error) {
	results := make([]*enterprise.EnterpriseActivity, 0)
	for rows.Next() {
		ea, err := ear.load(ctx, rows)
		if err != nil {
			return nil, err
		}
		results = append(results, ea)
	}
	return results, nil
}
