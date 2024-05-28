package queue

import "go-complaint/erros"

type LinkedQueue[E any] struct {
	front *Link[E]
	rear  *Link[E]
	size  int
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
	for q.front.Next == nil {
		q.rear = q.front
		q.rear = nil
	}
	q.rear = q.front
	q.size = 0
}

func (q *LinkedQueue[E]) Enqueue(e *E) {
	q.rear.Next = NewLink[E](WithValue(e))
	q.rear = q.rear.Next
	q.size++
}

func (q *LinkedQueue[E]) Dequeue() (*E, error) {
	if q.size == 0 {
		return nil, &erros.QueueIsEmptyError{}
	}
	it := q.front.Next.Element
	ltemp := q.front.Next
	q.front.Next = ltemp.Next
	if q.rear == ltemp {
		q.rear = q.front
	}
	ltemp = nil
	q.size--
	return it, nil
}

func (q *LinkedQueue[E]) FrontValue() (*E, error) {
	if q.size == 0 {
		return nil, &erros.QueueIsEmptyError{}
	}
	return q.front.Next.Element, nil
}

func (q *LinkedQueue[E]) Length() int {
	return q.size
}
