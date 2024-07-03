package find_all_users

import "strings"

type FullNameLike struct {
	query string
	args  []interface{}
}

func ByFullNameLike(term string) *FullNameLike {
	wildCardTerm := "%" + strings.ToLower(term) + "%"
	return &FullNameLike{
		query: string(
			`SELECT email, password, register_date, is_confirmed FROM (SELECT 				
		USER.EMAIL,
		USER.PASSWORD,
		USER.REGISTER_DATE,
		USER.IS_CONFIRMED,
		person.first_name,
		person.last_name,
		person.first_name || ' ' || person.last_name AS full_name
FROM USER
JOIN PERSON ON
		USER.EMAIL = PERSON.EMAIL
) as users
where LOWER(first_name) like $1 OR
	  LOWER(last_name) like $1 OR
	  LOWER(full_name) like $1;`,
		),
		args: []interface{}{wildCardTerm},
	}
}

func (e *FullNameLike) Query() string {
	return e.query
}

func (e *FullNameLike) Args() []interface{} {
	return e.args
}
