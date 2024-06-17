package userrolefindall

type ByUserID struct {
	query string
	args  []interface{}
}

func NewByUserID(userID string) *ByUserID {
	return &ByUserID{
		query: string(`
			SELECT
				user_id,
				role_id,
				enterprise_id
			FROM
				public."user_role"
			WHERE
				user_id = $1`,
		),
		args: []interface{}{userID},
	}
}

func (e *ByUserID) Query() string {
	return e.query
}

func (e *ByUserID) Args() []interface{} {
	return e.args
}
