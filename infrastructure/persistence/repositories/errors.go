package repositories

import "fmt"

var ErrMapperNotRegistered = fmt.Errorf("mapper not registered")
var ErrWrongTypeAssertion = fmt.Errorf("wrong type assertion")
var ErrFeedbackAlreadyExists = fmt.Errorf("duplicate key value violates unique constraint \"feedback_complaint_id_enterprise_id_key\" (SQLSTATE 23505)")

type EmailError struct {
	StatusCode   int
	ResponseBody string
}

func (e *EmailError) Error() string {
	return fmt.Sprintf("email error: %d", e.StatusCode)
}
