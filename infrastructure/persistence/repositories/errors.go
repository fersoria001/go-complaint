package repositories

import "fmt"

var ErrMapperNotRegistered = fmt.Errorf("mapper not registered")
var ErrWrongTypeAssertion = fmt.Errorf("wrong type assertion")

type EmailError struct {
	StatusCode   int
	ResponseBody string
}

func (e *EmailError) Error() string {
	return fmt.Sprintf("email error: %d", e.StatusCode)
}
