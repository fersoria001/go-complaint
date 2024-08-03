package repositories

import (
	"context"
	"fmt"
	"go-complaint/domain/model/common"
	"go-complaint/domain/model/complaint"
	"go-complaint/infrastructure/persistence/datasource"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type ComplaintDataRepository struct {
	schema datasource.Schema
}

func NewComplaintDataRepository(schema datasource.Schema) ComplaintDataRepository {
	return ComplaintDataRepository{
		schema: schema,
	}
}

func (cdr ComplaintDataRepository) Remove(ctx context.Context, id uuid.UUID) error {
	conn, err := cdr.schema.Acquire(ctx)
	defer conn.Release()
	if err != nil {
		return err
	}
	deleteCommand := string(`DELETE FROM COMPLAINT_DATA WHERE ID=$1`)
	_, err = conn.Exec(ctx, deleteCommand, &id)
	if err != nil {
		return err
	}
	return nil
}

func (cdr ComplaintDataRepository) Save(ctx context.Context, complaintData complaint.ComplaintData) error {
	conn, err := cdr.schema.Acquire(ctx)
	defer conn.Release()
	if err != nil {
		return err
	}
	insertCommand := string(`INSERT INTO COMPLAINT_DATA (
	ID,
	OWNER_ID,
	COMPLAINT_ID,
	OCCURRED_ON,
	DATA_TYPE) VALUES ($1, $2, $3, $4, $5)`)
	var (
		id          = complaintData.Id()
		ownerId     = complaintData.OwnerId()
		complaintId = complaintData.ComplaintId()
		occurredOn  = common.StringDate(complaintData.OccurredOn())
		dataType    = complaintData.DataType().String()
	)
	_, err = conn.Exec(ctx, insertCommand, &id, &ownerId, &complaintId, &occurredOn, &dataType)
	if err != nil {
		return err
	}
	return nil
}

func (cdr ComplaintDataRepository) Find(ctx context.Context, src StatementSource) (*complaint.ComplaintData, error) {
	conn, err := cdr.schema.Acquire(ctx)
	defer conn.Release()
	if err != nil {
		return nil, err
	}
	row := conn.QueryRow(ctx, src.Query(), src.Args()...)
	return cdr.load(ctx, row)
}

func (cdr ComplaintDataRepository) Get(ctx context.Context, id uuid.UUID) (*complaint.ComplaintData, error) {
	conn, err := cdr.schema.Acquire(ctx)
	defer conn.Release()
	if err != nil {
		return nil, err
	}
	selectQuery := string(`SELECT
	ID,
	OWNER_ID,
	COMPLAINT_ID,
	OCCURRED_ON,
	DATA_TYPE
	FROM COMPLAINT_DATA
	WHERE ID=$1`)
	row := conn.QueryRow(ctx, selectQuery, &id)
	return cdr.load(ctx, row)
}

func (cdr ComplaintDataRepository) load(_ context.Context, row pgx.Row) (*complaint.ComplaintData, error) {
	var (
		id          uuid.UUID
		ownerId     uuid.UUID
		complaintId uuid.UUID
		occurredOn  string
		dataType    string
	)
	err := row.Scan(&id, &ownerId, &complaintId, &occurredOn, &dataType)
	if err != nil {
		return nil, err
	}
	date, err := common.ParseDate(occurredOn)
	if err != nil {
		return nil, err
	}
	dType := complaint.ParseComplaintDataType(dataType)
	if dType < 1 {
		return nil, fmt.Errorf("dataType is unknown")
	}
	return complaint.NewComplaintData(id, ownerId, complaintId, date, dType), nil
}

func (cdr ComplaintDataRepository) FindAll(ctx context.Context, src StatementSource) ([]*complaint.ComplaintData, error) {
	conn, err := cdr.schema.Acquire(ctx)
	defer conn.Release()
	if err != nil {
		return nil, err
	}
	rows, err := conn.Query(ctx, src.Query(), src.Args()...)
	if err != nil {
		return nil, err
	}
	result, err := cdr.loadAll(ctx, rows)
	if err != nil {
		return nil, err
	}
	rows.Close()
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (cdr ComplaintDataRepository) loadAll(ctx context.Context, rows pgx.Rows) ([]*complaint.ComplaintData, error) {
	results := make([]*complaint.ComplaintData, 0)
	for rows.Next() {
		complaintData, err := cdr.load(ctx, rows)
		if err != nil {
			return nil, err
		}
		results = append(results, complaintData)
	}
	return results, nil
}
