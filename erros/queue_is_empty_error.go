package erros

type QueueIsEmptyError struct {
}

func (q *QueueIsEmptyError) Error() string {
	return "queue is empty"
}
