package erros

type OutOfRangeError struct {
}

func (o *OutOfRangeError) Error() string {
	return "out of range"
}