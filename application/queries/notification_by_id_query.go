package queries

import (
	"context"
	"go-complaint/dto"
	"go-complaint/infrastructure/persistence/repositories"

	"github.com/google/uuid"
)

type NotificationByIdQuery struct {
	Id string `json:"id"`
}

func NewNotificationByIdQuery(id string) *NotificationByIdQuery {
	return &NotificationByIdQuery{
		Id: id,
	}
}

func (q NotificationByIdQuery) Execute(ctx context.Context) (*dto.Notification, error) {
	id, err := uuid.Parse(q.Id)
	if err != nil {
		return nil, err
	}
	repository := repositories.MapperRegistryInstance().Get("Notification").(repositories.NotificationRepository)
	notification, err := repository.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return dto.NewNotification(*notification), nil
}
