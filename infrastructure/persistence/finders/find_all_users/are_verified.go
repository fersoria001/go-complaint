package find_all_users

type AreVerified struct {
	query string
	args  []interface{}
}

func ThatAreVerified(limit, offset int) *AreVerified {
	return &AreVerified{
		query: string(
			`SELECT PUBLIC.USER.EMAIL,
					PUBLIC.USER.PASSWORD,
					PUBLIC.USER.REGISTER_DATE,
					PUBLIC.USER.IS_CONFIRMED
			FROM PUBLIC.USER
			JOIN PUBLIC.PERSON ON
					PUBLIC.USER.EMAIL = PUBLIC.PERSON.EMAIL
			WHERE 		PUBLIC.USER.IS_CONFIRMED = TRUE
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
