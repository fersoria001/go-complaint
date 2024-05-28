package persistence_test

import (
	"context"
	"go-complaint/dto"
	"go-complaint/infrastructure/persistence/datasource"
	"go-complaint/infrastructure/persistence/repositories"
	"go-complaint/tests"
	"testing"

	mapset "github.com/deckarep/golang-set/v2"
)

func TestEventRepository_VALID(t *testing.T) {
	ctx := context.Background()
	eventSchema := datasource.EventSchema()
	err := eventSchema.Connect(ctx)
	if err != nil {
		t.Error(err)
	}
	eventRepository := repositories.NewEventRepository(eventSchema)
	storedSet := mapset.NewSet[*dto.StoredEvent]()
	for i := 0; i < 10; i++ {
		stored, err := dto.NewStoredEvent(tests.NewFakeEvent())
		if err != nil {
			t.Error(err)
		}
		storedSet.Add(stored)
	}

	t.Run(`Save a new event`, func(t *testing.T) {
		oneEvent, ok := storedSet.Pop()
		if !ok {
			t.Error(err)
		}
		err := eventRepository.Save(ctx, *oneEvent)
		if err != nil {
			t.Error(err)
		}
	})
	t.Run(`Save all events`, func(t *testing.T) {
		err := eventRepository.SaveAll(ctx, storedSet)
		if err != nil {
			t.Error(err)
		}
	})

	t.Run(`Get all events`, func(t *testing.T) {
		set, err := eventRepository.GetAll(ctx)
		if err != nil {
			t.Error(err)
		}
		if set.Cardinality() == 0 {
			t.Error(`No events found`)
		}
	})
	t.Run(`Get by ID`, func(t *testing.T) {
		set, err := eventRepository.GetAll(ctx)
		if err != nil {
			t.Error(err)
		}
		if set.Cardinality() == 0 {
			t.Error(`No events found`)
		}
		oneEvent, ok := set.Pop()
		if !ok {
			t.Error(err)
		}
		_, err = eventRepository.Get(ctx, oneEvent.EventId.String())
		if err != nil {
			t.Error(err)
		}
	})
	t.Run(`Delete by ID`, func(t *testing.T) {
		set, err := eventRepository.GetAll(ctx)
		if err != nil {
			t.Error(err)
		}
		if set.Cardinality() == 0 {
			t.Error(`No events found`)
		}
		oneEvent, ok := set.Pop()
		if !ok {
			t.Error(err)
		}
		err = eventRepository.Remove(ctx, oneEvent.EventId.String())
		if err != nil {
			t.Error(err)
		}
	})
	t.Run(`Delete all events of an id slice`, func(t *testing.T) {
		set, err := eventRepository.GetAll(ctx)
		if err != nil {
			t.Error(err)
		}
		if set.Cardinality() == 0 {
			t.Error(`No events found`)
		}
		ids := make([]string, 0)
		for event := range set.Iter() {
			ids = append(ids, event.EventId.String())
		}
		err = eventRepository.RemoveAll(ctx, ids)
		if err != nil {
			t.Error(err)
		}
	})

}
