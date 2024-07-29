package repositories

import (
	"context"
	"go-complaint/domain/model/recipient"
	"go-complaint/infrastructure/persistence/datasource"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type RecipientRepository struct {
	schema datasource.Schema
}

func NewRecipientRepository(schema datasource.Schema) RecipientRepository {
	return RecipientRepository{schema: schema}
}

func (r RecipientRepository) Remove(ctx context.Context, id uuid.UUID) error {
	conn, err := r.schema.Acquire(ctx)
	defer conn.Release()
	if err != nil {
		return err
	}
	deleteCommand := string(`DELETE FROM RECIPIENTS WHERE ID = $1`)
	_, err = conn.Exec(ctx, deleteCommand, &id)
	if err != nil {
		return err
	}
	return nil
}

func (r RecipientRepository) Update(ctx context.Context, recipient recipient.Recipient) error {
	conn, err := r.schema.Acquire(ctx)
	defer conn.Release()
	if err != nil {
		return err
	}
	insertCommand := string(`
	UPDATE RECIPIENTS 
	SET SUBJECT_NAME=$2,SUBJECT_THUMBNAIL=$3, SUBJECT_EMAIL=$4
	WHERE ID=$1`)
	var (
		id               = recipient.Id()
		subjectName      = recipient.SubjectName()
		subjectThumbnail = recipient.SubjectThumbnail()
		subjectEmail     = recipient.SubjectEmail()
	)
	_, err = conn.Exec(ctx, insertCommand, &id, &subjectName, &subjectThumbnail, &subjectEmail)
	if err != nil {
		return err
	}
	return nil
}

func (r RecipientRepository) Save(ctx context.Context, recipient recipient.Recipient) error {
	conn, err := r.schema.Acquire(ctx)
	defer conn.Release()
	if err != nil {
		return err
	}
	insertCommand := string(`
	INSERT INTO RECIPIENTS (
	ID, IS_ENTERPRISE, SUBJECT_NAME, SUBJECT_THUMBNAIL,SUBJECT_EMAIL) 
	VALUES ($1, $2, $3, $4, $5)`)
	var (
		id               = recipient.Id()
		isEnterprise     = recipient.IsEnterprise()
		subjectName      = recipient.SubjectName()
		subjectThumbnail = recipient.SubjectThumbnail()
		subjectEmail     = recipient.SubjectEmail()
	)
	_, err = conn.Exec(ctx, insertCommand, &id, &isEnterprise, &subjectName, &subjectThumbnail, &subjectEmail)
	if err != nil {
		return err
	}
	return nil
}

func (r RecipientRepository) Find(ctx context.Context, src StatementSource) (*recipient.Recipient, error) {
	conn, err := r.schema.Acquire(ctx)
	defer conn.Release()
	if err != nil {
		return nil, err
	}
	row := conn.QueryRow(ctx, src.Query(), src.Args()...)
	return r.load(ctx, row)
}

func (r RecipientRepository) FindAll(ctx context.Context, src StatementSource) ([]*recipient.Recipient, error) {
	conn, err := r.schema.Acquire(ctx)
	defer conn.Release()
	if err != nil {
		return nil, err
	}
	rows, err := conn.Query(ctx, src.Query(), src.Args()...)
	if err != nil {
		return nil, err
	}
	result, err := r.loadAll(ctx, rows)
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

func (r RecipientRepository) loadAll(ctx context.Context, rows pgx.Rows) ([]*recipient.Recipient, error) {
	results := make([]*recipient.Recipient, 0)
	for rows.Next() {
		recipient, err := r.load(ctx, rows)
		if err != nil {
			return nil, err
		}
		results = append(results, recipient)
	}
	return results, nil
}

func (r RecipientRepository) Get(ctx context.Context, id uuid.UUID) (*recipient.Recipient, error) {
	conn, err := r.schema.Acquire(ctx)
	defer conn.Release()
	if err != nil {
		return nil, err
	}
	selectQuery := string(`SELECT ID, IS_ENTERPRISE,SUBJECT_NAME,SUBJECT_THUMBNAIL,SUBJECT_EMAIL FROM RECIPIENTS WHERE ID=$1`)
	row := conn.QueryRow(ctx, selectQuery, &id)
	return r.load(ctx, row)
}

func (r RecipientRepository) load(_ context.Context, row pgx.Row) (*recipient.Recipient, error) {
	var (
		id               uuid.UUID
		isEnterprise     bool
		subjectName      string
		subjectThumbnail string
		subjectEmail     string
	)
	err := row.Scan(&id, &isEnterprise, &subjectName, &subjectThumbnail, &subjectEmail)
	if err != nil {
		return nil, err
	}
	return recipient.NewRecipient(id, subjectName, subjectThumbnail, subjectEmail, isEnterprise), nil
}
