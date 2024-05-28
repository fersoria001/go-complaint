package erros

import "fmt"


type ColumnNotFoundError struct {
	ColumnName string
}

func (e *ColumnNotFoundError) Error() string {
	return fmt.Sprintf("column %s not found", e.ColumnName)
}