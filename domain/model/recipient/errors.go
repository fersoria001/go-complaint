package recipient

import "fmt"

type ValidationError struct {
	Message string
}

func (e ValidationError) Error() string {
	return fmt.Errorf("validation error: %s", e.Message).Error()
}
