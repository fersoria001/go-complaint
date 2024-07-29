package queries

import (
	"fmt"
)

var ErrUserNotConfirmed = fmt.Errorf("user not confirmed")
var ErrAuthenticationFailed = fmt.Errorf("authentication failed")
var ErrConfirmationNotFound = fmt.Errorf("the token doesn't reference a login in proccess")
var ErrWrongTypeAssertion = fmt.Errorf("wrong type assertion")
var ErrConfirmationAlreadyDone = fmt.Errorf("confirmation already done")
var ErrConfirmationCodeNotMatch = fmt.Errorf("confirmation code doesn't match")
var ErrNilValue = fmt.Errorf("nil value")
var ErrBadRequest = fmt.Errorf("bad request")
var ErrForbidden = fmt.Errorf("forbidden")
var ErrUnauthorized = fmt.Errorf("unauthorized")
var ErrNotFound = fmt.Errorf("not found")
var ErrCursorOutOfRange = fmt.Errorf("cursor is out of range")
