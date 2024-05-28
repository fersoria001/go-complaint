package infrastructure_test

import (
	"errors"
	"go-complaint/erros"
	"go-complaint/infrastructure/queue"
	"go-complaint/tests"
	"testing"
)

func TestLinkedQueue(t *testing.T) {
	fakeEvent := tests.NewFakeEvent()
	queuedEvent1 := &QueuedEvent{
		typeName: "FakeEvent",
		event:    fakeEvent,
	}
	queuedEvent2 := &QueuedEvent{
		typeName: "FakeEvent1",
		event:    fakeEvent,
	}
	queuedEvent3 := &QueuedEvent{
		typeName: "FakeEvent2",
		event:    fakeEvent,
	}
	expectedErr := &erros.QueueIsEmptyError{}
	linkedQueue := queue.NewLinkedQueue[QueuedEvent]()
	t.Run("Test NewLinkedQueue", func(t *testing.T) {
		if linkedQueue == nil {
			t.Error("NewLinkedQueue() should not return nil")
		}
	})
	t.Run("Test Enqueue it will insert all elements at rear, pushing the first to the front - 1", func(t *testing.T) {
		linkedQueue.Enqueue(queuedEvent1)
		linkedQueue.Enqueue(queuedEvent2)
		linkedQueue.Enqueue(queuedEvent3)
		if linkedQueue.Length() != 3 {
			t.Error("Length() should return 3")
		}
		frontValue, err := linkedQueue.FrontValue()
		if err != nil {
			t.Error("FrontValue() should not return error")
		}
		if frontValue != queuedEvent1 {
			t.Error("FrontValue() should return queuedEvent1")
		}
	})
	t.Run("Test Dequeue it will remove the front element and return it", func(t *testing.T) {
		linkedQueue.Enqueue(queuedEvent1)
		linkedQueue.Enqueue(queuedEvent2)
		linkedQueue.Enqueue(queuedEvent3)
		dequeuedEvent, err := linkedQueue.Dequeue()
		if err != nil {
			t.Error("Dequeue() should not return error")
		}
		if dequeuedEvent != queuedEvent1 {
			t.Error("Dequeue() should return queuedEvent1")
		}
		if linkedQueue.Length() != 2 {
			t.Error("Length() should return 2")
		}
		dequeuedEvent, err = linkedQueue.Dequeue()
		if err != nil {
			t.Error("Dequeue() should not return error")
		}
		if dequeuedEvent != queuedEvent2 {
			t.Error("Dequeue() should return queuedEvent2")
		}
		if linkedQueue.Length() != 1 {
			t.Error("Length() should return 1")
		}
		dequeuedEvent, err = linkedQueue.Dequeue()
		if err != nil {
			t.Error("Dequeue() should not return error")
		}
		if dequeuedEvent != queuedEvent3 {
			t.Error("Dequeue() should return queuedEvent3")
		}
		if linkedQueue.Length() != 0 {
			t.Error("Length() should return 0")
		}

		_, err = linkedQueue.Dequeue()
		if err == nil {
			t.Error("Dequeue() should return error")
		}
		if linkedQueue.Length() != 0 {
			t.Error("Length() should return 0")
		}
		if !errors.As(err, &expectedErr) {
			t.Error("Dequeue() should return QueueIsEmptyError")
		}
	})
	t.Run("Test Clear it will remove all elements", func(t *testing.T) {
		linkedQueue.Enqueue(queuedEvent1)
		linkedQueue.Enqueue(queuedEvent2)
		linkedQueue.Enqueue(queuedEvent3)
		linkedQueue.Clear()
		if linkedQueue.Length() != 0 {
			t.Error("Length() should return 0")
		}
	})
}
