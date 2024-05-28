package queue

type Queue interface {
	Clear()
	Enqueue(e interface{}) error
	Dequeue() (interface{}, error)
	FrontValue() (interface{}, error)
	length() int
}
