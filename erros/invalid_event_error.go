package erros


type InvalidEventError struct {
}

func (e *InvalidEventError) Error() string {
	return "the data []byte doesn't correspond to any domain event type"
}