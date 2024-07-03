package countryfindall

type Countries struct {
	query string
	args  []interface{}
}

func NewCountries() *Countries {
	return &Countries{
		query: string(`
		SELECT
		id,
		name,
		phonecode
		FROM country
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
