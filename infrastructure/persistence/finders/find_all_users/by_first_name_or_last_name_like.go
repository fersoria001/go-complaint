package find_all_users

import "strings"

type ByFirstNameOrLastNameLike struct {
	query string
	args  []interface{}
}

func NewByFirstNameOrLastNameLike(term string) *ByFirstNameOrLastNameLike {
	wildCardTerm := "%" + strings.ToLower(term) + "%"
	return &ByFirstNameOrLastNameLike{
		query: string(
			`SELECT PUBLIC.USER.EMAIL,
					PUBLIC.USER.PASSWORD,
					PUBLIC.USER.REGISTER_DATE,
					PUBLIC.USER.IS_CONFIRMED
			FROM PUBLIC.USER
			JOIN PERSON ON
					PUBLIC.USER.EMAIL = PERSON.EMAIL
			WHERE 	LOWER(PERSON.FIRST_NAME) LIKE $1 OR
					LOWER(PERSON.LAST_NAME) LIKE $1;`,
		),
		args: []interface{}{wildCardTerm},
	}
}

func (e *ByFirstNameOrLastNameLike) Query() string {
	return e.query
}

func (e *ByFirstNameOrLastNameLike) Args() []interface{} {
	return e.args
}
