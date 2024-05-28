package erros

type ValueNotFoundError struct {
}

func (e *ValueNotFoundError) Error() string {
	return "Not found"
}
