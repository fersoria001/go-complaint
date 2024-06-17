package repositories

import (
	"context"
	"go-complaint/domain"
	"go-complaint/domain/model/common"
	"go-complaint/infrastructure/persistence/datasource"
	"net/mail"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type NotificationRepository struct {
	schema datasource.Schema
}

func NewNotificationRepository(schema datasource.Schema) NotificationRepository {
	return NotificationRepository{schema: schema}
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
			notification_id,
			owner_id,
			thumbnail,
			title,
			content,
			link,
			occurred_on,
			seen
		FROM public."notification"
		WHERE notification_id = $1
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
) (mapset.Set[*domain.Notification], error) {
	conn, err := pr.schema.Acquire(ctx)
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
	notifications, err := pr.loadAll(ctx, rows)
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	return notifications, nil
}

func (pr NotificationRepository) loadAll(
	ctx context.Context,
	rows pgx.Rows,
) (mapset.Set[*domain.Notification], error) {
	notifications := mapset.NewSet[*domain.Notification]()
	for rows.Next() {
		notification, err := pr.load(ctx, rows)
		if err != nil {
			return nil, err
		}
		notifications.Add(notification)
	}
	return notifications, nil
}

func (pr NotificationRepository) load(ctx context.Context, row pgx.Row) (*domain.Notification, error) {
	var (
		id         uuid.UUID
		ownerID    string
		thumbnail  string
		title      string
		content    string
		link       string
		occurredOn string
		seen       bool
	)
	err := row.Scan(
		&id,
		&ownerID,
		&thumbnail,
		&title,
		&content,
		&link,
		&occurredOn,
		&seen,
	)
	if err != nil {
		return nil, err
	}
	if _, err := mail.ParseAddress(thumbnail); err != nil {
		ep, err := MapperRegistryInstance().Get("Enterprise").(EnterpriseRepository).Get(ctx, thumbnail)
		if err != nil {
			return nil, err
		}
		thumbnail = ep.LogoIMG()
	} else {
		user, err := MapperRegistryInstance().Get("User").(UserRepository).Get(ctx, thumbnail)
		if err != nil {
			return nil, err
		}
		thumbnail = user.ProfileIMG()
	}
	occurredOnTime, err := common.NewDateFromString(occurredOn)
	if err != nil {
		return nil, err
	}
	notification, err := domain.NewNotification(
		id,
		ownerID,
		thumbnail,
		title,
		content,
		link,
		occurredOnTime.Date(),
		seen,
	)
	if err != nil {
		return nil, err
	}
	return notification, nil
}

func (pr NotificationRepository) Save(
	ctx context.Context,
	notification *domain.Notification,
) error {
	conn, err := pr.schema.Acquire(ctx)
	if err != nil {
		return err
	}
	insertCommand := string(`
		INSERT INTO 
		notification (
			notification_id,
			owner_id,
			thumbnail,
			title,
			content,
			link,
			occurred_on,
			seen
			)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`)
	var (
		id         uuid.UUID = notification.ID()
		ownerID    string    = notification.OwnerID()
		thumbnail  string    = notification.Thumbnail()
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
		&ownerID,
		&thumbnail,
		&title,
		&content,
		&link,
		&occurredOn,
		&seen,
	)
	if err != nil {
		return err
	}
	defer conn.Release()
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
	insertCommand := string(`
		UPDATE public."notification"
		SET
			seen = $1
		WHERE notification_id = $2`)
	var (
		seen bool = notification.Seen()
	)
	_, err = conn.Exec(
		ctx,
		insertCommand,
		&seen,
	)
	if err != nil {
		return err
	}

	defer conn.Release()
	return nil
}
