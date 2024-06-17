package repositories

import (
	"context"
	"go-complaint/domain/model/common"
	"go-complaint/domain/model/feedback"
	"go-complaint/infrastructure/persistence/datasource"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type FeedbackAnswerRepository struct {
	schema datasource.Schema
}

func NewFeedbackAnswerRepository(
	feedbackSchema datasource.Schema,
) FeedbackAnswerRepository {
	return FeedbackAnswerRepository{schema: feedbackSchema}
}

func (fr FeedbackAnswerRepository) DeleteAll(
	ctx context.Context,
	feedbackID uuid.UUID,
) error {
	conn, err := fr.schema.Acquire(ctx)
	if err != nil {
		return err
	}
	deleteCommand := string(`
	DELETE FROM feedback_answers
	WHERE feedback_id = $1`)
	_, err = conn.Exec(ctx, deleteCommand, &feedbackID)
	if err != nil {
		return err
	}
	defer conn.Release()
	return nil
}

func (fr FeedbackAnswerRepository) Save(
	ctx context.Context,
	answers mapset.Set[feedback.Answer],
) error {
	conn, err := fr.schema.Acquire(ctx)
	if err != nil {
		return err
	}
	insertCommand := string(`
		INSERT INTO 
		feedback_answers (
			id,
			feedback_id,
			sender_id,
			answer_body,
			created_at,
			read,
			read_at,
			updated_at,
			is_enterprise,
			enterprise_id
		  )
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`)
	tx, err := conn.Begin(ctx)
	if err != nil {
		return err
	}
	for answer := range answers.Iter() {
		var (
			ID           = answer.ID()
			FeedbackID   = answer.FeedbackID()
			SenderID     = answer.SenderID()
			Body         = answer.Body()
			CreatedAt    = answer.CreatedAt().StringRepresentation()
			Read         = answer.Read()
			ReadAt       = answer.ReadAt().StringRepresentation()
			UpdatedAt    = answer.UpdatedAt().StringRepresentation()
			IsEnterprise = answer.IsEnterprise()
			EnterpriseID = answer.EnterpriseID()
		)
		_, err = tx.Exec(
			ctx,
			insertCommand,
			&ID,
			&FeedbackID,
			&SenderID,
			&Body,
			&CreatedAt,
			&Read,
			&ReadAt,
			&UpdatedAt,
			&IsEnterprise,
			&EnterpriseID,
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

// blocking connection for more time in this approach
func (fr FeedbackAnswerRepository) FindAll(
	ctx context.Context,
	statementSource StatementSource,
) (mapset.Set[*feedback.Answer], error) {
	conn, err := fr.schema.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	rows, err := conn.Query(ctx, statementSource.Query(), statementSource.Args()...)
	if err != nil {
		return nil, err
	}
	answers, err := fr.loadAll(ctx, rows)
	if err != nil {
		return nil, err
	}
	defer func() {
		rows.Close()
		conn.Release()
	}()
	return answers, nil
}

func (fr FeedbackAnswerRepository) loadAll(
	ctx context.Context,
	rows pgx.Rows,
) (mapset.Set[*feedback.Answer], error) {
	answers := mapset.NewSet[*feedback.Answer]()
	for rows.Next() {
		answer, err := fr.load(ctx, rows)
		if err != nil {
			return nil, err
		}
		answers.Add(answer)
	}
	return answers, nil
}

func (fr FeedbackAnswerRepository) load(
	ctx context.Context,
	row pgx.Row,
) (*feedback.Answer, error) {
	mapper := MapperRegistryInstance().Get("User")
	if mapper == nil {
		return nil, ErrMapperNotRegistered
	}
	userRepository, ok := mapper.(UserRepository)
	if !ok {
		return nil, ErrWrongTypeAssertion
	}
	enterpriseRepository := NewEnterpriseRepository(
		fr.schema,
	)
	var (
		id           uuid.UUID
		feedbackId   uuid.UUID
		senderID     string
		senderIMG    string
		senderName   string
		body         string
		createdAt    string
		read         bool
		readAt       string
		updatedAt    string
		isEnterprise bool
		enterpriseID string
	)

	err := row.Scan(
		&id,
		&feedbackId,
		&senderID,
		&body,
		&createdAt,
		&read,
		&readAt,
		&updatedAt,
		&isEnterprise,
		&enterpriseID,
	)
	if err != nil {
		return nil, err
	}
	sender, err := userRepository.Get(ctx, senderID)
	if err != nil {
		return nil, err
	}
	senderName = sender.FullName()
	if isEnterprise {
		enterprise, err := enterpriseRepository.Get(
			ctx, enterpriseID)
		if err != nil {
			return nil, err
		}
		senderIMG = enterprise.LogoIMG()
		enterpriseID = enterprise.Name()
	} else {
		senderIMG = sender.ProfileIMG()
		enterpriseID = ""
	}
	createdAtDate, err := common.NewDateFromString(createdAt)
	if err != nil {
		return nil, err
	}
	readAtDate, err := common.NewDateFromString(readAt)
	if err != nil {
		return nil, err
	}
	updatedAtDate, err := common.NewDateFromString(updatedAt)
	if err != nil {
		return nil, err
	}
	return feedback.NewAnswer(
		id,
		feedbackId,
		senderID,
		senderIMG,
		senderName,
		body,
		createdAtDate,
		read,
		readAtDate,
		updatedAtDate,
		isEnterprise,
		enterpriseID,
	)
}
