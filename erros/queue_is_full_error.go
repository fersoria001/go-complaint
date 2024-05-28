package erros

type QueueIsFullError struct {
}

func (q *QueueIsFullError) Error() string {
	return "queue is full"
}
