package complaint

import "fmt"

var ErrComplaintClosed = fmt.Errorf("the complaint is closed")
var ErrReplyNotAdded = fmt.Errorf("the complaint reply couldn't be added to the replies set")

type ValidationError struct {
	Message string
}

func (e ValidationError) Error() string {
	return fmt.Errorf("validation error: %s", e.Message).Error()
}
