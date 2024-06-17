package queries

import (
	"context"
	"go-complaint/dto"
	notificationsfindall "go-complaint/infrastructure/persistence/finders/notifications_findall"
	"go-complaint/infrastructure/persistence/repositories"
)

type NotificationQuery struct {
	OwnerID string `json:"owner_id"`
}

func (notificationQuery NotificationQuery) Notifications(
	ctx context.Context,
) ([]dto.Notification, error) {
	domainObj, err := repositories.MapperRegistryInstance().Get("Notification").(repositories.NotificationRepository).FindAll(
		ctx,
		notificationsfindall.NewByOwnerID(notificationQuery.OwnerID),
	)
	if err != nil {
		return nil, err
	}
	notifications := make([]dto.Notification, 0)
	for notification := range domainObj.Iter() {
		notifications = append(notifications, dto.Notification{
			ID:        notification.ID().String(),
			OwnerID:   notification.OwnerID(),
			Thumbnail: notification.Thumbnail(),
			Title:     notification.Title(),
			Content:   notification.Content(),
			Link:      notification.Link(),
			Seen:      notification.Seen(),
		})
	}
	return notifications, nil
}
