package commands

import (
	"context"
	"go-complaint/domain"
	"go-complaint/infrastructure/persistence/repositories"
	"time"

	"github.com/google/uuid"
)

type NotificationCommand struct {
	ID        string `json:"id"`
	OwnerID   string `json:"owner_id"`
	Thumbnail string `json:"thumbnail"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Link      string `json:"link"`
}

func (notificationCommand NotificationCommand) SaveNew(
	ctx context.Context,
) error {
	newNotification, err := domain.NewNotification(
		uuid.New(),
		notificationCommand.OwnerID,
		notificationCommand.Thumbnail,
		notificationCommand.Title,
		notificationCommand.Content,
		notificationCommand.Link,
		time.Now(),
		false,
	)
	if err != nil {
		return err
	}
	return repositories.MapperRegistryInstance().Get("Notification").(repositories.NotificationRepository).Save(ctx, newNotification)
}

func (notificationCommand NotificationCommand) MarkAsRead(
	ctx context.Context,
) error {
	if notificationCommand.ID == "" {
		return nil
	}
	parsedID, err := uuid.Parse(notificationCommand.ID)
	if err != nil {
		return err
	}
	n, err := repositories.MapperRegistryInstance().Get("Notification").(repositories.NotificationRepository).Get(
		ctx,
		parsedID,
	)
	if err != nil {
		return err
	}
	n.MarkAsRead()
	return repositories.MapperRegistryInstance().Get("Notification").(repositories.NotificationRepository).Update(ctx, *n)
}
