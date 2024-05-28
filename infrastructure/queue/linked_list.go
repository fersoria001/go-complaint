package queue

import (
	"go-complaint/erros"
)

/*
It cant be instantiated with an interface
wrap it into a struct,
tail is always nil
append and insert doesn't move the cursor
*/
type LList[E any] struct {
	head *Link[E]
	tail *Link[E]
	curr *Link[E]
	cnt  int
}

// void init() == new(LList[E])
// remove all garbage collector
func NewLList[E any]() *LList[E] {
	nilNewLink := NewLink[E]()
	return &LList[E]{
		head: nilNewLink,
		tail: nilNewLink,
		curr: nilNewLink,
		cnt:  0,
	}
}
func (l *LList[E]) Clear() {
	for l.head.Next != nil {
		l.curr = l.head
		l.head = l.head.Next
		l.curr.Next = nil
		l.cnt--
	}
}

// Insert "it" at current position
// if tail is equal to current position,
// then current position is the new tail
func (l *LList[E]) Insert(it *E) {
	l.curr.Next = NewLink[E](WithValue(it), WithNext(l.curr.Next))
	if l.tail == l.curr {
		l.tail = l.curr.Next
	}
	l.cnt++
}

// Append "it" to list
// tail is always the last element
func (l *LList[E]) Append(it *E) {
	l.tail.Next = NewLink[E](WithValue(it))
	l.tail = l.tail.Next
	l.cnt++
}

func (l *LList[E]) Remove() error {
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
	return nil
}
func (l *LList[E]) MoveToStart() {
	l.curr = l.head
}
func (l *LList[E]) MoveToEnd() {
	l.curr = l.tail
}
func (l *LList[E]) Prev() {
	if l.curr == l.head {
		return
	}
	temp := l.head
	for temp.Next != l.curr {
		temp = temp.Next
	}
	l.curr = temp
}
func (l *LList[E]) Next() {
	if l.curr != l.tail {
		l.curr = l.curr.Next
	}
}
func (l *LList[E]) Length() int {
	return l.cnt
}
func (l *LList[E]) CurrPos() int {
	temp := l.head
	var i int
	for l.curr != temp {
		i++
		temp = temp.Next
	}
	return i
}
func (l *LList[E]) MoveToPos(pos int) error {
	if pos < 0 || pos > l.cnt {
		return &erros.OutOfRangeError{}
	}
	l.curr = l.head
	for i := 0; i < pos; i++ {
		l.curr = l.curr.Next
	}
	return nil
}
func (l *LList[E]) GetValue() (*E, error) {
	if l.curr.Next == nil {
		return nil, &erros.NoElementError{}
	}
	return l.curr.Next.Element, nil
}

//LinkedQueue
