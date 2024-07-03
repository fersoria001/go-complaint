package queue

import (
	"go-complaint/erros"
	"sync"
)

/*
It cant be instantiated with an interface
wrap it into a struct,
tail is always nil
append and insert doesn't move the cursor
*/
type LinkedList[E any] struct {
	head *Link[E]
	tail *Link[E]
	curr *Link[E]
	cnt  int
	mu   sync.Mutex
}

// void init() == new(LinkedList[E])
// remove all garbage collector
func NewLinkedList[E any]() *LinkedList[E] {
	nilNewLink := NewLink[E]()
	return &LinkedList[E]{
		head: nilNewLink,
		tail: nilNewLink,
		curr: nilNewLink,
		cnt:  0,
	}
}
func (l *LinkedList[E]) Clear() {
	l.mu.Lock()
	for l.head.Next != nil {
		l.curr = l.head
		l.head = l.head.Next
		l.curr.Next = nil
		l.cnt--
	}
	l.mu.Unlock()
}

// Insert "it" at current position
// if tail is equal to current position,
// then current position is the new tail
func (l *LinkedList[E]) Insert(it E) {
	l.mu.Lock()
	l.curr.Next = NewLink[E](WithValue(it), WithNext(l.curr.Next))
	if l.tail == l.curr {
		l.tail = l.curr.Next
	}
	l.cnt++
	l.mu.Unlock()
}

// Append "it" to list
// tail is always the last element
func (l *LinkedList[E]) Append(it E) {
	l.mu.Lock()
	l.tail.Next = NewLink[E](WithValue(it))
	l.tail = l.tail.Next
	l.cnt++
	l.mu.Unlock()
}

func (l *LinkedList[E]) Remove() error {
	l.mu.Lock()
	if l.curr.Next == nil {
		return &erros.NoElementError{}
	}
	if l.tail == l.curr.Next {
		l.tail = l.curr
	}
	ltemp := l.curr.Next.Next
	l.curr.Next = nil
	l.curr.Next = ltemp
	l.cnt--
	l.mu.Unlock()
	return nil
}
func (l *LinkedList[E]) MoveToStart() {
	l.mu.Lock()
	l.curr = l.head
	l.mu.Unlock()
}
func (l *LinkedList[E]) MoveToEnd() {
	l.mu.Lock()
	l.curr = l.tail
	l.mu.Unlock()
}
func (l *LinkedList[E]) Prev() {
	l.mu.Lock()
	if l.curr == l.head {
		return
	}
	temp := l.head
	for temp.Next != l.curr {
		temp = temp.Next
	}
	l.curr = temp
	l.mu.Unlock()
}
func (l *LinkedList[E]) Next() {
	l.mu.Lock()
	if l.curr != l.tail {
		l.curr = l.curr.Next
	}
	l.mu.Unlock()
}
func (l *LinkedList[E]) Length() int {
	return l.cnt
}
func (l *LinkedList[E]) CurrPos() int {
	l.mu.Lock()
	temp := l.head
	var i int
	for l.curr != temp {
		i++
		temp = temp.Next
	}
	l.mu.Unlock()
	return i
}
func (l *LinkedList[E]) MoveToPos(pos int) error {
	if pos < 0 || pos > l.cnt {
		return &erros.OutOfRangeError{}
	}
	l.mu.Lock()
	l.curr = l.head
	for i := 0; i < pos; i++ {
		l.curr = l.curr.Next
	}
	l.mu.Unlock()
	return nil
}
func (l *LinkedList[E]) GetValue() (E, error) {
	var nilE E
	if l.curr.Next == nil {
		return nilE, &erros.NoElementError{}
	}
	return l.curr.Next.Element, nil
}

//LinkedQueue
