package enterprise

import (
	"errors"
	"fmt"
)

var ErrWrongTypeAssertion = errors.New("wrong type assertion")
var ErrUnknownColleagueType = errors.New("unknown colleague type")
var ErrPositionNotExists = errors.New("position not exists")
var ErrForbidden = errors.New("either you don't belong to this enterprise or you don't have the necessary permissions")
var ErrInvalidChatID = errors.New("invalid chat ID")
var ErrReplyNotFound = errors.New("reply not found in enterprise chat")

type ValidationError struct {
	Message string
}

func (e ValidationError) Error() string {
	return fmt.Errorf("validation error: %s", e.Message).Error()
}
