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
			`SELECT USER.EMAIL,
					USER.PASSWORD,
					USER.REGISTER_DATE,
					USER.IS_CONFIRMED
			FROM USER
			JOIN PERSON ON
					USER.EMAIL = PERSON.EMAIL
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
