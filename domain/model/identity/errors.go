package identity

import "fmt"

var ErrUserRoleNotFound = fmt.Errorf("user role not found")
var ErrNilPtr = fmt.Errorf("the pointer is nil")
var ErrNilValue = fmt.Errorf("the value is nil")
var ErrUserRoleAlreadyAdded = fmt.Errorf("the user already have this role")
