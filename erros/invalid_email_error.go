package erros


//Package erros
type InvalidEmailError struct {
}

func (e *InvalidEmailError) Error() string {
	return "The email is invalid, it should be a valid email according to the RFC  5322  specification"
}