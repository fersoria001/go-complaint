package erros

import (
	"fmt"
)
//Package erros
type ValidationError struct {
	Expected string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("Expected %s", e.Expected)
}

