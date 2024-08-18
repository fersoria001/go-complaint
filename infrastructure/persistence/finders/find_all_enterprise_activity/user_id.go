package find_all_enterprise_activity

import "github.com/google/uuid"

type UserId struct {
	query string
	args  []interface{}
}

func ByUserId(userId uuid.UUID) *UserId {
	return &UserId{
		query: string(`
	SELECT
		ID,
		USER_ID,
		ACTIVITY_ID,
		ENTERPRISE_ID,
		ENTERPRISE_NAME,
		OCCURRED_ON,
		ACTIVITY_TYPE
	FROM ENTERPRISE_ACTIVITY
	WHERE USER_ID = $1
	`),
		args: []interface{}{userId},
	}
}

func (e *UserId) Query() string {
	return e.query
}

func (e *UserId) Args() []interface{} {
	return e.args
}
