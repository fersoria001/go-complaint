package application

import (
	"context"
	"go-complaint/dto"
	"go-complaint/infrastructure/persistence/repositories"
)

type EventStore struct {
}

func (es *EventStore) Save(ctx context.Context, event *dto.StoredEvent) error {
	repository := repositories.MapperRegistryInstance().Get("Event").(repositories.EventRepository)
	err := repository.Save(ctx, event)
	if err != nil {
		return err
	}
	return nil
}
