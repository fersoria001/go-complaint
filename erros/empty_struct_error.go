package erros


//Package identityandaccesserrors
type EmptyStructError struct {
}


func (e *EmptyStructError) Error() string {
	return "The struct is empty"
}

