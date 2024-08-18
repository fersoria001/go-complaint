package repositories

import (
	"context"
	"go-complaint/domain"
	"go-complaint/domain/model/common"
	"go-complaint/infrastructure/persistence/datasource"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type NotificationRepository struct {
	schema datasource.Schema
}

func NewNotificationRepository(schema datasource.Schema) NotificationRepository {
	return NotificationRepository{schema: schema}
}

func (pr NotificationRepository) Remove(ctx context.Context, id uuid.UUID) error {
	conn, err := pr.schema.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()
	deleteCommand := string(`DELETE FROM NOTIFICATIONS WHERE ID=$1`)
	_, err = conn.Exec(ctx, deleteCommand, &id)
	if err != nil {
		return err
	}
	return nil
}

func (pr NotificationRepository) Get(
	ctx context.Context,
	notificationID uuid.UUID,
) (*domain.Notification, error) {
	conn, err := pr.schema.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	selectQuery := string(`
		SELECT
			id,
			owner_id,
			sender_id,
			title,
			content,
			link,
			occurred_on,
			seen
		FROM notifications
		WHERE id = $1
		`)
	row := conn.QueryRow(
		ctx,
		selectQuery,
		&notificationID,
	)
	notification, err := pr.load(ctx, row)
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	return notification, nil
}

func (pr NotificationRepository) FindAll(
	ctx context.Context,
	source StatementSource,
) ([]*domain.Notification, error) {
	conn, err := pr.schema.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	rows, err := conn.Query(
		ctx,
		source.Query(),
		source.Args()...,
	)
	if err != nil {
		return nil, err
	}
	notifications, err := pr.loadAll(ctx, rows)
	if err != nil {
		return nil, err
	}
	if rows.Err() != nil {
		return nil, err
	}
	rows.Close()
	return notifications, nil
}

func (pr NotificationRepository) loadAll(
	ctx context.Context,
	rows pgx.Rows,
) ([]*domain.Notification, error) {
	notifications := make([]*domain.Notification, 0)
	for rows.Next() {
		notification, err := pr.load(ctx, rows)
		if err != nil {
			return nil, err
		}
		notifications = append(notifications, notification)
	}
	return notifications, nil
}

func (pr NotificationRepository) load(ctx context.Context, row pgx.Row) (*domain.Notification, error) {
	var (
		id         uuid.UUID
		ownerId    uuid.UUID
		senderId   uuid.UUID
		title      string
		content    string
		link       string
		occurredOn string
		seen       bool
	)
	err := row.Scan(
		&id,
		&ownerId,
		&senderId,
		&title,
		&content,
		&link,
		&occurredOn,
		&seen,
	)
	if err != nil {
		return nil, err
	}
	reg := MapperRegistryInstance()
	recipientRepository, ok := reg.Get("Recipient").(RecipientRepository)
	if !ok {
		return nil, ErrWrongTypeAssertion
	}
	owner, err := recipientRepository.Get(ctx, ownerId)
	if err != nil {
		return nil, err
	}
	sender, err := recipientRepository.Get(ctx, senderId)
	if err != nil {
		return nil, err
	}
	occurredOnTime, err := common.NewDateFromString(occurredOn)
	if err != nil {
		return nil, err
	}
	notification := domain.NewNotification(
		id,
		*owner,
		*sender,
		title,
		content,
		link,
		occurredOnTime.Date(),
		seen,
	)
	return notification, nil
}

func (pr NotificationRepository) Save(
	ctx context.Context,
	notification domain.Notification,
) error {
	conn, err := pr.schema.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()
	insertCommand := string(`
		INSERT INTO 
		notifications (
			id,
			owner_id,
			sender_id,
			title,
			content,
			link,
			occurred_on,
			seen
			)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`)
	var (
		id         uuid.UUID = notification.ID()
		ownerId    uuid.UUID = notification.Owner().Id()
		senderId   uuid.UUID = notification.Sender().Id()
		title      string    = notification.Title()
		content    string    = notification.Content()
		link       string    = notification.Link()
		occurredOn string    = common.StringDate(notification.OccurredOn())
		seen       bool      = notification.Seen()
	)
	_, err = conn.Exec(
		ctx,
		insertCommand,
		&id,
		&ownerId,
		&senderId,
		&title,
		&content,
		&link,
		&occurredOn,
		&seen,
	)
	if err != nil {
		return err
	}
	return nil
}

func (pr NotificationRepository) Update(
	ctx context.Context,
	notification domain.Notification,
) error {
	conn, err := pr.schema.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()
	insertCommand := string(`
		UPDATE notifications
		SET
			seen = $1
		WHERE id = $2`)
	var (
		seen bool      = notification.Seen()
		id   uuid.UUID = notification.ID()
	)
	_, err = conn.Exec(
		ctx,
		insertCommand,
		&seen,
		&id,
	)
	if err != nil {
		return err
	}
	return nil
}
