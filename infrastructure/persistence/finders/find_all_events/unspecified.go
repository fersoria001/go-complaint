package find_all_events

type Unspecified struct {
	query string
	args  []interface{}
}

func By() *Unspecified {
	return &Unspecified{
		query: string(`
	SELECT
		event_id,
		event_body,
		occurred_on,
		type_name
	FROM events
	`),
		args: []interface{}{},
	}
}

func (e *Unspecified) Query() string {
	return e.query
}

func (e *Unspecified) Args() []interface{} {
	return e.args
}
