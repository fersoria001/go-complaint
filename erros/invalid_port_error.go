package erros

import "fmt"

type InvalidPortError struct {
	Port      int
}

func (e *InvalidPortError) Error() string {
	return fmt.Sprintf("port %d is invalid, out of range(0-65535)", e.Port)
}