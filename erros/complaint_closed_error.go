package erros

type ComplaintClosedError struct {
}

func (e *ComplaintClosedError) Error() string {
	return "complaint is closed"
}
