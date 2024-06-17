package queue

import (
	"go-complaint/erros"
	"sync"
)

type LinkedQueue[E any] struct {
	front *Link[E]
	rear  *Link[E]
	size  int
	mu    sync.Mutex
}

func NewLinkedQueue[E any]() *LinkedQueue[E] {
	nilNewLink := NewLink[E]()
	return &LinkedQueue[E]{
		front: nilNewLink,
		rear:  nilNewLink,
		size:  0,
	}
}

func (q *LinkedQueue[E]) Clear() {
	q.mu.Lock()
	for q.front.Next == nil {
		q.rear = q.front
		q.rear = nil
	}
	q.rear = q.front
	q.size = 0
	q.mu.Unlock()
}

func (q *LinkedQueue[E]) Enqueue(e E) {
	q.mu.Lock()
	q.rear.Next = NewLink[E](WithValue(e))
	q.rear = q.rear.Next
	q.size++
	q.mu.Unlock()
}

func (q *LinkedQueue[E]) Dequeue() (E, error) {
	q.mu.Lock()
	var nilE E
	if q.size == 0 {
		return nilE, &erros.QueueIsEmptyError{}
	}
	it := q.front.Next.Element
	ltemp := q.front.Next
	q.front.Next = ltemp.Next
	if q.rear == ltemp {
		q.rear = q.front
	}
	ltemp = nil
	q.size--
	q.mu.Unlock()
	return it, nil
}

func (q *LinkedQueue[E]) FrontValue() (E, error) {
	var nilE E
	if q.size == 0 {
		return nilE, &erros.QueueIsEmptyError{}
	}
	return q.front.Next.Element, nil
}

func (q *LinkedQueue[E]) Length() int {
	return q.size
}
