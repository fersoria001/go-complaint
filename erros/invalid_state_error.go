package erros

type InvalidStateError struct {
}

func (e *InvalidStateError) Error() string {
	return "invalid state"
}
