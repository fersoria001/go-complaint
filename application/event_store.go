package application

import (
	"context"
	"go-complaint/dto"
	"go-complaint/infrastructure/persistence/datasource"
	"go-complaint/infrastructure/persistence/repositories"
	"sync"
)

var eventStore *EventStore
var once sync.Once

type EventStore struct {
	repository *repositories.EventRepository
}

func EventStoreInstance() (*EventStore, error) {
	var err error = nil
	once.Do(func() {
		repository := repositories.NewEventRepository(datasource.PublicSchema())
		eventStore = &EventStore{
			repository: &repository,
		}
	})
	return eventStore, err
}

func (es *EventStore) Save(ctx context.Context, event *dto.StoredEvent) error {
	err := es.repository.Save(ctx, event)
	if err != nil {
		return err
	}
	return nil
}
