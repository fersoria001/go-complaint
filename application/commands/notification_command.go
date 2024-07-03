package commands

import (
	"context"
	"fmt"
	"go-complaint/domain"
	"go-complaint/dto"
	"go-complaint/infrastructure/cache"
	"go-complaint/infrastructure/persistence/repositories"
	"log"
	"time"

	"github.com/google/uuid"
)

type NotificationCommand struct {
	ID          string `json:"id"`
	OwnerID     string `json:"owner_id"`
	ThumbnailID string `json:"thumbnail_id"`
	Thumbnail   string `json:"thumbnail"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	Link        string `json:"link"`
}

func (notificationCommand NotificationCommand) SaveNew(
	ctx context.Context,
) error {
	newNotification, err := domain.NewNotification(
		uuid.New(),
		notificationCommand.OwnerID,
		notificationCommand.ThumbnailID,
		notificationCommand.Title,
		notificationCommand.Content,
		notificationCommand.Link,
		time.Now(),
		false,
	)
	if err != nil {
		return err
	}
	notificationDto := dto.NewNotification(*newNotification)
	notificationDto.Thumbnail = notificationCommand.Thumbnail
	log.Println("notificationDto", notificationDto)
	cache.RequestChannel <- cache.Request{
		Key:     fmt.Sprintf("notifications:%s", notificationCommand.OwnerID),
		Payload: notificationDto,
	}

	return repositories.MapperRegistryInstance().Get("Notification").(repositories.NotificationRepository).Save(ctx, newNotification)
}

func (notificationCommand NotificationCommand) MarkAsRead(
	ctx context.Context,
) error {
	if notificationCommand.ID == "" || notificationCommand.OwnerID == "" {
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
	if n.Seen() {
		return nil
	}
	n.MarkAsRead()
	cache.RequestChannel <- cache.Request{
		Key:     fmt.Sprintf("notifications:%s", notificationCommand.OwnerID),
		Payload: dto.NewNotification(*n),
	}
	log.Println("notificationCommand.OwnerID", notificationCommand.OwnerID)
	return repositories.MapperRegistryInstance().Get("Notification").(repositories.NotificationRepository).Update(ctx, *n)
}
