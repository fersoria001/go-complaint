package erros

type UnauthorizedError struct {
}

func (e *UnauthorizedError) Error() string {
	return "Unauthorized: User not found in context"
}
