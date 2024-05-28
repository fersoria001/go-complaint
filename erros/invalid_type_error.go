package erros

type InvalidTypeError struct {
}

func (i *InvalidTypeError) Error() string {
	return "invalid type"
}
