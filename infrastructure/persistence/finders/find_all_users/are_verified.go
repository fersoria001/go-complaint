package find_all_users

type AreVerified struct {
	query string
	args  []interface{}
}

func ThatAreVerified(limit, offset int) *AreVerified {
	return &AreVerified{
		query: string(`
		SELECT 
		id,
		username,
		password,
		register_date,
		is_confirmed
		FROM users
		WHERE is_confirmed = TRUE
		LIMIT $1 OFFSET $2;
		`),
		args: []interface{}{limit, offset},
	}
}

func (e *AreVerified) Query() string {
	return e.query
}

func (e *AreVerified) Args() []interface{} {
	return e.args
}
