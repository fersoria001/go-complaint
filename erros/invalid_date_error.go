package erros

//Package erros
type InvalidDateError struct {
}

func (e *InvalidDateError) Error() string {
	return "The date is invalid because it is after the current date"
}
