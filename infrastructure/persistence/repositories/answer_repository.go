package repositories

import (
	"context"
	"fmt"
	"go-complaint/domain/model/common"
	"go-complaint/domain/model/feedback"
	"go-complaint/erros"
	"go-complaint/infrastructure/persistence/datasource"
	"go-complaint/infrastructure/persistence/models"

	"github.com/google/uuid"
)

type AnswerRepository struct {
	schema *datasource.Schema
}

func NewAnswerRepository(answerSchema *datasource.Schema) *AnswerRepository {
	return &AnswerRepository{schema: answerSchema}
}

func (answerRepository *AnswerRepository) Save(ctx context.Context, obj interface{}) error {
	aggregate, ok := obj.(*feedback.Answer)
	if !ok {
		return &erros.InvalidTypeError{}
	}
	conn, err := answerRepository.schema.Pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()
	model := models.NewAnswer(aggregate)
	saveCommand := fmt.Sprintf(
		`INSERT INTO %s (%s) VALUES (%s)`,
		model.Table(),
		models.StringColumns(model.Columns()),
		model.Args(),
	)
	_, err = conn.Exec(ctx, saveCommand, model.Values()...)
	if err != nil {
		return err
	}
	return nil
}

func (answerRepository *AnswerRepository) Get(
	ctx context.Context,
	id string,
) (interface{}, error) {
	conn, err := answerRepository.schema.Pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	model := models.NewAnswer(&feedback.Answer{})
	parsedID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	selectQuery := fmt.Sprintf(
		`SELECT %s FROM %s WHERE id = $1`,
		models.StringColumns(model.Columns()),
		model.Table(),
	)
	row := conn.QueryRow(ctx, selectQuery, parsedID)
	err = row.Scan(model.Values()...)
	if err != nil {
		return nil, err
	}
	createdAt, err := common.NewDateFromString(model.CreatedAt)
	if err != nil {
		return nil, err
	}
	readAt, err := common.NewDateFromString(model.ReadAt)
	if err != nil {
		return nil, err
	}
	updatedAt, err := common.NewDateFromString(model.UpdatedAt)
	if err != nil {
		return nil, err
	}
	value, err := feedback.NewAnswer(
		model.ID,
		model.FeedbackID,
		model.SenderID,
		model.SenderIMG,
		model.SenderName,
		model.Body,
		createdAt,
		model.Read,
		readAt,
		updatedAt,
	)
	if err != nil {
		return nil, err
	}
	return value, nil
}

func (answerRepository *AnswerRepository) Update(
	ctx context.Context,
	obj interface{},
) error {
	aggregate, ok := obj.(*feedback.Answer)
	if !ok {
		return &erros.InvalidTypeError{}
	}
	conn, err := answerRepository.schema.Pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()
	model := models.NewAnswer(aggregate)
	updateCommand := fmt.Sprintf(
		`UPDATE %s SET %s WHERE id = $1`,
		model.Table(),
		models.StringKeyArgs(model.Columns()),
	)
	_, err = conn.Exec(ctx, updateCommand, model.Values()...)
	if err != nil {
		return err
	}
	return nil
}

func (answerRepository *AnswerRepository) Remove(
	ctx context.Context,
	id string,
) error {
	conn, err := answerRepository.schema.Pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()
	model := models.NewAnswer(&feedback.Answer{})
	parsedID, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	deleteCommand := fmt.Sprintf(
		`DELETE FROM %s WHERE id = $1`,
		model.Table(),
	)
	_, err = conn.Exec(ctx, deleteCommand, parsedID)
	if err != nil {
		return err
	}
	return nil
}

func (answerRepository *AnswerRepository) FindByFeedbackID(
	ctx context.Context,
	id string,
) ([]*feedback.Answer, error) {
	conn, err := answerRepository.schema.Pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	model := models.NewAnswer(&feedback.Answer{})
	parsedID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	selectQuery := fmt.Sprintf(
		`SELECT %s FROM %s WHERE feedback_id = $1`,
		models.StringColumns(model.Columns()),
		model.Table(),
	)
	rows, err := conn.Query(ctx, selectQuery, parsedID)
	if err != nil {
		return nil, err
	}
	var answers []*feedback.Answer
	for rows.Next() {
		err = rows.Scan(model.Values()...)
		if err != nil {
			return nil, err
		}
		createdAt, err := common.NewDateFromString(model.CreatedAt)
		if err != nil {
			return nil, err
		}
		readAt, err := common.NewDateFromString(model.ReadAt)
		if err != nil {
			return nil, err
		}
		updatedAt, err := common.NewDateFromString(model.UpdatedAt)
		if err != nil {
			return nil, err
		}
		value, err := feedback.NewAnswer(
			model.ID,
			model.FeedbackID,
			model.SenderID,
			model.SenderIMG,
			model.SenderName,
			model.Body,
			createdAt,
			model.Read,
			readAt,
			updatedAt,
		)
		if err != nil {
			return nil, err
		}
		answers = append(answers, value)
	}
	return answers, nil
}
