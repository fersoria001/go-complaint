package erros
//Package erros
type InvalidPasswordError struct {
}



func (e *InvalidPasswordError) Error() string {
	return "The password is invalid, it should have a minimum of 8 characters and at least one digit character and at least one uppercase letter and at least one lowercase letter"
}