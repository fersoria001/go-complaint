package find_all_hiring_proccesses

import "github.com/google/uuid"

type UserId struct {
	query string
	args  []interface{}
}

func ByUserId(userId uuid.UUID) *UserId {
	return &UserId{
		query: string(`SELECT 
	ID,
	ENTERPRISE_ID,
	USER_ID,
	ROLE,
	STATUS,
	REASON,
	EMITED_BY_ID,
	OCCURRED_ON,
	LAST_UPDATE,
	UPDATED_BY_ID,
	INDUSTRY_ID
	FROM HIRING_PROCCESSES WHERE USER_ID = $1 `),
		args: []interface{}{userId},
	}
}

func (e *UserId) Query() string {
	return e.query
}

func (e *UserId) Args() []interface{} {
	return e.args
}
