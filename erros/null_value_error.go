package erros

//Package erros
type NullValueError struct {
}


func (e *NullValueError) Error() string {
	return "The value is null"
}

