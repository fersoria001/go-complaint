package commands

import (
	"context"
	"go-complaint/application"
	"go-complaint/dto"
	"go-complaint/infrastructure/persistence/repositories"

	"github.com/google/uuid"
)

type MarkNotificationAsReadCommand struct {
	Id string `json:"id"`
}

func NewMarkNotificationAsReadCommand(id string) *MarkNotificationAsReadCommand {
	return &MarkNotificationAsReadCommand{
		Id: id,
	}
}

func (c MarkNotificationAsReadCommand) Execute(ctx context.Context) error {
	id, err := uuid.Parse(c.Id)
	if err != nil {
		return err
	}
	reg := repositories.MapperRegistryInstance()
	repository, ok := reg.Get("Notification").(repositories.NotificationRepository)
	if !ok {
		return ErrWrongTypeAssertion
	}
	n, err := repository.Get(ctx, id)
	if err != nil {
		return err
	}
	n.MarkAsRead()
	err = repository.Update(ctx, *n)
	if err != nil {
		return err
	}
	pub := application.ApplicationMessagePublisherInstance()
	pub.Publish(application.NewApplicationMessage(
		n.Owner().Id().String(),
		"notification",
		*dto.NewNotification(*n),
	))
	return nil
}
