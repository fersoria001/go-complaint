package countrystatefindall

type ByCountryID struct {
	query string
	args  []interface{}
}

func NewByCountryID(countryID int) *ByCountryID {
	return &ByCountryID{
		query: string(`
		SELECT
		id,
		name
		FROM country_states
		WHERE country_id = $1
		`),
		args: []interface{}{countryID},
	}
}

func (e *ByCountryID) Query() string {
	return e.query
}

func (e *ByCountryID) Args() []interface{} {
	return e.args
}
