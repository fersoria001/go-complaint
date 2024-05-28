package erros

import "fmt"

//Package erros
type InvalidLengthError struct {
	AttributeName string
	MaxLength	 int
	MinLength	 int
	CurrentLength int
}

func (e *InvalidLengthError) Error() string {
	return fmt.Sprintf("The length of %s must be between %d and %d, but it is %d", e.AttributeName, e.MinLength, e.MaxLength, e.CurrentLength)
}