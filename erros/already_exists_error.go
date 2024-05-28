package erros

//Package erros
type AlreadyExistsError struct {

}

func (e *AlreadyExistsError) Error() string {
	return "The value already exists in the collection"
}
