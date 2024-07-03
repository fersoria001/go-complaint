package find_all_users

type AreVerified struct {
	query string
	args  []interface{}
}

func ThatAreVerified(limit, offset int) *AreVerified {
	return &AreVerified{
		query: string(
			`SELECT USER.EMAIL,
					USER.PASSWORD,
					USER.REGISTER_DATE,
					USER.IS_CONFIRMED
			FROM USER
			JOIN PERSON ON
					USER.EMAIL = PERSON.EMAIL
			WHERE 		USER.IS_CONFIRMED = TRUE
			LIMIT $1 OFFSET $2;`,
		),
		args: []interface{}{limit, offset},
	}
}

func (e *AreVerified) Query() string {
	return e.query
}

func (e *AreVerified) Args() []interface{} {
	return e.args
}
