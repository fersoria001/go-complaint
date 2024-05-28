package infrastructure_test

import (
	"go-complaint/domain"
	"go-complaint/infrastructure/queue"
	"go-complaint/tests"
	"testing"
)

type QueuedEvent struct {
	typeName string
	event    domain.DomainEvent
}

func TestLinkedList_INT(t *testing.T) {
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

	l := queue.NewLList[QueuedEvent]()
	t.Run("Test NewLList", func(t *testing.T) {
		if l == nil {
			t.Error("NewLList() should not return nil")
		}
	})
	t.Run("Test Insert it will insert all elements at head, pushing the first to the tail - 1", func(t *testing.T) {
		l.Insert(queuedEvent1)
		l.Insert(queuedEvent2)
		l.Insert(queuedEvent3)
		if l.Length() != 3 {
			t.Error("Size() should return 3")
		}
		l.MoveToStart()
		currentElement, err := l.GetValue()
		if err != nil {
			t.Error("GetValue() should not return error")
		}
		if currentElement != queuedEvent3 {
			t.Error("GetValue() should return queuedEvent1")
		}
		l.MoveToEnd()
		_, err = l.GetValue()
		if err == nil {
			t.Error("GetValue() tail should be nil")
		}
		l.Prev()
		currentElement, err = l.GetValue()
		if err != nil {
			t.Error("GetValue() should not return error")
		}
		if currentElement != queuedEvent1 {
			t.Error("GetValue() should return queuedEvent1")
		}
	})
	t.Run("Test Remove it should remove 1 element", func(t *testing.T) {
		l.Insert(queuedEvent1)
		l.Insert(queuedEvent2)
		err := l.Remove()
		if err != nil {
			t.Errorf("Remove() should not return error %v", err)
		}
		if l.Length() != 1 {
			t.Errorf("Size() should return 0, length: %v", l.Length())
		}
	})
	t.Run("Test clear, it should remove all nodes", func(t *testing.T) {
		l.Insert(queuedEvent1)
		l.Insert(queuedEvent2)
		l.Insert(queuedEvent3)
		l.Clear()
		if l.Length() != 0 {
			t.Error("Size() should return 0")
		}
	})
	t.Run("Test Append value 1 must be the head and value 3 the tail - 1", func(t *testing.T) {
		l.Append(queuedEvent1)
		l.Append(queuedEvent2)
		l.Append(queuedEvent3)
		if l.Length() != 3 {
			t.Error("Size() should return 3")
		}
		l.MoveToStart()
		currentElement, err := l.GetValue()
		if err != nil {
			t.Error("GetValue() should not return error")
		}
		if currentElement != queuedEvent1 {
			t.Error("GetValue() should return queuedEvent1")
		}
		l.MoveToEnd()
		l.Prev()
		currentElement, err = l.GetValue()
		if err != nil {
			t.Error("GetValue() should not return error")
		}
		if currentElement != queuedEvent3 {
			t.Error("GetValue() should return queuedEvent3")
		}
	})

	t.Run("Test traversal with append", func(t *testing.T) {
		l.Append(queuedEvent1)
		l.Append(queuedEvent2)
		l.Append(queuedEvent3)
		currPos := l.CurrPos()
		if currPos != 0 {
			t.Errorf("CurrPos() should return 0, got %v", currPos)
		}
		l.MoveToStart()
		currPos = l.CurrPos()
		if currPos != 0 {
			t.Errorf("CurrPos() should return 0, got %v", currPos)
		}
		l.MoveToEnd()
		currPos = l.CurrPos()
		if currPos != 3 {
			t.Errorf("CurrPos() should return 3, got %v", currPos)
		}
		l.MoveToPos(1)
		currPos = l.CurrPos()
		if currPos != 1 {
			t.Errorf("CurrPos() should return 1, got %v", currPos)
		}
		l.MoveToPos(2)
		currPos = l.CurrPos()
		if currPos != 2 {
			t.Errorf("CurrPos() should return 3, got %v", currPos)
		}
		currElem, err := l.GetValue()
		if err != nil {
			t.Error("GetValue() should not return error")
		}
		if currElem != queuedEvent3 {
			t.Errorf("GetValue() should return queuedEvent3, got %v", currElem)
		}
	})
}
