package feedback

import (
	"fmt"
)

var ErrColorKeyNotFound = fmt.Errorf("color key not found")
var ErrReplyNotFound = fmt.Errorf("reply not found")
var ErrNilValue = fmt.Errorf("nil value")
var ErrNilPointer = fmt.Errorf("nil pointer")
var ErrNotFound = fmt.Errorf("not found")
var ErrReplyReviewNotFound = fmt.Errorf("reply review not found")
var ErrReplyReviewAlreadyExists = fmt.Errorf("reply review already exists")
var ErrReplyAlreadyExists = fmt.Errorf("reply already exists")
var ErrFeedbackIsNotDone = fmt.Errorf("feedback is not done")

type ValidationError struct {
	Message string
}

func (v ValidationError) Error() string {
	return fmt.Errorf("%s", v.Message).Error()
}
