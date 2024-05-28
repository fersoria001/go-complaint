package erros
import "fmt"

type NotImplementedError struct {
	Method string
}

func (e *NotImplementedError) Error() string {
	return fmt.Sprintf("Method %s is not implemented yet", e.Method)
}