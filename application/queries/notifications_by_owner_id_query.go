package queries

import (
	"context"
	"errors"
	"go-complaint/dto"
	"go-complaint/infrastructure/persistence/finders/find_all_notifications"
	"go-complaint/infrastructure/persistence/repositories"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type NotificationsByOwnerIdQuery struct {
	OwnerId string `json:"ownerId"`
}

func NewNotificationsByOwnerIdQuery(ownerId string) *NotificationsByOwnerIdQuery {
	return &NotificationsByOwnerIdQuery{
		OwnerId: ownerId,
	}
}

func (q NotificationsByOwnerIdQuery) Execute(ctx context.Context) ([]*dto.Notification, error) {
	id, err := uuid.Parse(q.OwnerId)
	if err != nil {
		return nil, err
	}
	result := make([]*dto.Notification, 0)
	repository := repositories.MapperRegistryInstance().Get("Notification").(repositories.NotificationRepository)
	notifications, err := repository.FindAll(ctx, find_all_notifications.ByOwnerId(id))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return result, nil
		}
		return nil, err
	}
	for i := range notifications {
		result = append(result, dto.NewNotification(*notifications[i]))
	}
	return result, nil
}
