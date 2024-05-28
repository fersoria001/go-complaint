package erros

//Package erros
type InvalidNameError struct {
}

func (e *InvalidNameError) Error() string {
	return "The name is invalid, valid formats can be containing letters, hyphens, and apostrophes, including accented characters from various European languages and no trailing spaces"
}
