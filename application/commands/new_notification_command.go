package commands

import (
	"context"
	"go-complaint/application"
	"go-complaint/domain"
	"go-complaint/dto"
	"go-complaint/infrastructure/persistence/repositories"
	"time"

	"github.com/google/uuid"
)

type SendNotificationCommand struct {
	OwnerId  string `json:"ownerId"`
	SenderId string `json:"senderId"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	Link     string `json:"link"`
}

func NewSendNotificationCommand(ownerId, senderId, title, content, link string) *SendNotificationCommand {
	return &SendNotificationCommand{
		OwnerId:  ownerId,
		SenderId: senderId,
		Title:    title,
		Content:  content,
		Link:     link,
	}
}

func (c SendNotificationCommand) Execute(ctx context.Context) error {
	reg := repositories.MapperRegistryInstance()
	repository, ok := reg.Get("Notification").(repositories.NotificationRepository)
	if !ok {
		return ErrWrongTypeAssertion
	}
	recipientRepository, ok := reg.Get("Recipient").(repositories.RecipientRepository)
	if !ok {
		return ErrWrongTypeAssertion
	}
	ownerId, err := uuid.Parse(c.OwnerId)
	if err != nil {
		return err
	}
	senderId, err := uuid.Parse(c.SenderId)
	if err != nil {
		return err
	}
	owner, err := recipientRepository.Get(ctx, ownerId)
	if err != nil {
		return err
	}
	sender, err := recipientRepository.Get(ctx, senderId)
	if err != nil {
		return err
	}
	pub := application.ApplicationMessagePublisherInstance()
	newNotification := domain.NewNotification(
		uuid.New(),
		*owner,
		*sender,
		c.Title,
		c.Content,
		c.Link,
		time.Now(),
		false,
	)
	nDto := dto.NewNotification(*newNotification)
	err = repository.Save(ctx, *newNotification)
	if err != nil {
		return err
	}
	pub.Publish(application.NewApplicationMessage(nDto.Owner.Id, "notification", nDto))
	return nil
}
