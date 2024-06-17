package countryfindall

type Countries struct {
	query string
	args  []interface{}
}

func NewCountries() *Countries {
	return &Countries{
		query: string(`
		SELECT
		"country".id,
		"country".name,
		"country".phonecode
		FROM public."country"
		`),
		args: []interface{}{},
	}
}

func (e *Countries) Query() string {
	return e.query
}

func (e *Countries) Args() []interface{} {
	return e.args
}
