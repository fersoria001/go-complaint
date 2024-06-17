package notificationsfindall

type ByOwnerID struct {
	query string
	args  []interface{}
}

func NewByOwnerID(ownerID string) *ByOwnerID {
	return &ByOwnerID{
		query: string(`
		SELECT
			notification_id,
			owner_id,
			thumbnail,
			title,
			content,
			link,
			occurred_on,
			seen
		FROM public."notification"
		WHERE owner_id = $1
		`),
		args: []interface{}{ownerID},
	}
}

func (e *ByOwnerID) Query() string {
	return e.query
}

func (e *ByOwnerID) Args() []interface{} {
	return e.args
}
