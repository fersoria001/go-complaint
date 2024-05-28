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
		eventStore = &EventStore{
			repository: repositories.NewEventRepository(datasource.EventSchema()),
		}
	})
	return eventStore, err
}

// this can be used with a key-value store like redis
// for faster read and write operations
func (es *EventStore) Save(ctx context.Context, event dto.StoredEvent) error {
	err := es.repository.Save(ctx, event)
	if err != nil {
		return err
	}
	return nil
}
