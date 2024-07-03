package statecitiesfindall

type ByStateID struct {
	query string
	args  []interface{}
}

func NewByStateID(stateID int) *ByStateID {
	return &ByStateID{
		query: string(`
		SELECT
		id,
		name,
		country_code,
	 	latitude,
	  	longitude
		FROM state_cities
		WHERE state_id = $1
		`),
		args: []interface{}{stateID},
	}
}

func (e *ByStateID) Query() string {
	return e.query
}

func (e *ByStateID) Args() []interface{} {
	return e.args
}
