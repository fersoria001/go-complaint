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

type HiringProccessRepository struct {
	schema datasource.Schema
}

func NewHiringProccessRepository(schema datasource.Schema) HiringProccessRepository {
	return HiringProccessRepository{
		schema: schema,
	}
}

func (r HiringProccessRepository) Remove(ctx context.Context, id uuid.UUID) error {
	conn, err := r.schema.Acquire(ctx)
	defer conn.Release()
	if err != nil {
		return err
	}
	deleteCommand := string(`DELETE FROM HIRING_PROCCESSES WHERE ID = $1`)
	_, err = conn.Exec(ctx, deleteCommand, &id)
	if err != nil {
		return err
	}
	return nil
}

func (r HiringProccessRepository) Update(ctx context.Context, hiringProccess enterprise.HiringProccess) error {
	conn, err := r.schema.Acquire(ctx)
	defer conn.Release()
	if err != nil {
		return err
	}
	updateCommand := string(`
	UPDATE HIRING_PROCCESSES 
	SET
	STATUS = $2,
	REASON = $3,
	LAST_UPDATE = $4,
	UPDATED_BY_ID= $5
	WHERE ID = $1`)
	var (
		id          uuid.UUID = hiringProccess.Id()
		status      string    = hiringProccess.Status().String()
		reason      string    = hiringProccess.Reason()
		lastUpdate  string    = common.StringDate(hiringProccess.LastUpdate())
		updatedById uuid.UUID = hiringProccess.UpdatedBy().Id()
	)
	_, err = conn.Exec(ctx, updateCommand, &id, &status, &reason, &lastUpdate, &updatedById)
	if err != nil {
		return err
	}
	return nil
}

func (r HiringProccessRepository) Save(ctx context.Context, hiringProccess enterprise.HiringProccess) error {
	conn, err := r.schema.Acquire(ctx)
	defer conn.Release()
	if err != nil {
		return err
	}
	insertCommand := string(`
	INSERT INTO HIRING_PROCCESSES (
	ID,
	ENTERPRISE_ID,
	USER_ID,
	ROLE,
	STATUS,
	REASON,
	EMITED_BY_ID,
	OCCURRED_ON,
	LAST_UPDATE,
	UPDATED_BY_ID) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`)
	var (
		id           uuid.UUID = hiringProccess.Id()
		enterpriseId uuid.UUID = hiringProccess.Enterprise().Id()
		userId       uuid.UUID = hiringProccess.User().Id()
		role         string    = hiringProccess.Role().String()
		status       string    = hiringProccess.Status().String()
		reason       string    = hiringProccess.Reason()
		emitedBy     uuid.UUID = hiringProccess.EmitedBy().Id()
		occurredOn   string    = common.StringDate(hiringProccess.OccurredOn())
		lastUpdate   string    = common.StringDate(hiringProccess.LastUpdate())
		updatedById  uuid.UUID = hiringProccess.UpdatedBy().Id()
	)
	_, err = conn.Exec(ctx, insertCommand, &id, &enterpriseId, &userId, &role, &status,
		&reason, &emitedBy, &occurredOn, &lastUpdate, &updatedById)
	if err != nil {
		return err
	}
	return nil
}

func (r HiringProccessRepository) FindAll(ctx context.Context, src StatementSource) ([]*enterprise.HiringProccess, error) {
	conn, err := r.schema.Acquire(ctx)
	defer conn.Release()
	if err != nil {
		return nil, err
	}
	rows, err := conn.Query(ctx, src.Query(), src.Args()...)
	if err != nil {
		return nil, err
	}
	results, err := r.loadAll(ctx, rows)
	if err != nil {
		return nil, err
	}
	if rows.Err() != nil {
		return nil, err
	}
	rows.Close()
	return results, nil
}

func (r HiringProccessRepository) loadAll(ctx context.Context, rows pgx.Rows) ([]*enterprise.HiringProccess, error) {
	result := make([]*enterprise.HiringProccess, 0)
	for rows.Next() {
		h, err := r.load(ctx, rows)
		if err != nil {
			return nil, err
		}
		result = append(result, h)
	}
	return result, nil
}
func (r HiringProccessRepository) Find(ctx context.Context, src StatementSource) (*enterprise.HiringProccess, error) {
	conn, err := r.schema.Acquire(ctx)
	defer conn.Release()
	if err != nil {
		return nil, err
	}
	row := conn.QueryRow(ctx, src.Query(), src.Args()...)
	return r.load(ctx, row)
}

func (r HiringProccessRepository) Get(ctx context.Context, id uuid.UUID) (*enterprise.HiringProccess, error) {
	conn, err := r.schema.Acquire(ctx)
	defer conn.Release()
	if err != nil {
		return nil, err
	}
	selectQuery := string(`SELECT 
	ID,
	ENTERPRISE_ID,
	USER_ID,
	ROLE,
	STATUS,
	REASON,
	EMITED_BY_ID,
	OCCURRED_ON,
	LAST_UPDATE,
	UPDATED_BY_ID
	FROM HIRING_PROCCESSES WHERE ID = $1 `)
	row := conn.QueryRow(ctx, selectQuery, &id)
	return r.load(ctx, row)
}

func (r HiringProccessRepository) load(ctx context.Context, row pgx.Row) (*enterprise.HiringProccess, error) {
	var (
		id           uuid.UUID
		enterpriseId uuid.UUID
		userId       uuid.UUID
		role         string
		status       string
		reason       string
		emitedById   uuid.UUID
		occurredOn   string
		lastUpdate   string
		updatedById  uuid.UUID
	)
	err := row.Scan(
		&id,
		&enterpriseId,
		&userId,
		&role,
		&status,
		&reason,
		&emitedById,
		&occurredOn,
		&lastUpdate,
		&updatedById)
	if err != nil {
		return nil, err
	}
	reg := MapperRegistryInstance()
	recipientRepository, ok := reg.Get("Recipient").(RecipientRepository)
	if !ok {
		return nil, ErrWrongTypeAssertion
	}
	dbEnterprise, err := recipientRepository.Get(ctx, enterpriseId)
	if err != nil {
		return nil, err
	}
	user, err := recipientRepository.Get(ctx, userId)
	if err != nil {
		return nil, err
	}
	emitedBy, err := recipientRepository.Get(ctx, emitedById)
	if err != nil {
		return nil, err
	}
	updatedBy, err := recipientRepository.Get(ctx, updatedById)
	if err != nil {
		return nil, err
	}
	parsedRole := enterprise.ParsePosition(role)
	if parsedRole < 0 {
		return nil, fmt.Errorf("enterprise role isn't valid")
	}
	parsedStatus := enterprise.ParseHiringProccessStatus(status)
	if parsedStatus < 0 {
		return nil, fmt.Errorf("hiring proccess status is not valid")
	}
	occurredOnTime, err := common.ParseDate(occurredOn)
	if err != nil {
		return nil, err
	}
	lastUpdateTime, err := common.ParseDate(lastUpdate)
	if err != nil {
		return nil, err
	}
	return enterprise.NewHiringProccess(
		id,
		*dbEnterprise,
		*user,
		parsedRole,
		parsedStatus,
		reason,
		*emitedBy,
		occurredOnTime,
		lastUpdateTime,
		*updatedBy,
	), nil
}
