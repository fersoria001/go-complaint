package infrastructure

import "fmt"

type EmailError struct {
	StatusCode   int
	ResponseBody string
}

func (e *EmailError) Error() string {
	return fmt.Sprintf("email error: %d", e.StatusCode)
}

var ErrConfirmationRetryLimit = fmt.Errorf("confirmation retry limit reached")
var ErrUserRoleNotFound = fmt.Errorf("user role not found")
var ErrFileAlreadyExists = fmt.Errorf("file already exists")
var ErrFileTooBig = fmt.Errorf("file too big")
