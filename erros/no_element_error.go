package erros

type NoElementError struct {
}

func (n *NoElementError) Error() string {
	return "No element"
}
