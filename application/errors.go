package application

import "fmt"

var ErrConfirmationRetryLimit = fmt.Errorf("confirmation retry limit reached")
var ErrUserRoleNotFound = fmt.Errorf("user role not found")
