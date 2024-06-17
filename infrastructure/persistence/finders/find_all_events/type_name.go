package find_all_events

type TypeName struct {
	query string
	args  []interface{}
}

func ByTypeName(eventType string) *TypeName {
	return &TypeName{
		query: string(`
	SELECT
		event_id,
		event_body,
		occurred_on,
		type_name
	FROM events
	WHERE type_name = $1
	`),
		args: []interface{}{eventType},
	}
}

func (e *TypeName) Query() string {
	return e.query
}

func (e *TypeName) Args() []interface{} {
	return e.args
}
