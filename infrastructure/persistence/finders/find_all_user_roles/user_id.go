package find_all_user_roles

import "github.com/google/uuid"

type UserId struct {
	query string
	args  []interface{}
}

func ByUserId(userId uuid.UUID) *UserId {
	return &UserId{
		query: string(`
			SELECT
				user_id,
				role_id,
				enterprise_id
			FROM
				user_roles
			WHERE
				user_id = $1`,
		),
		args: []interface{}{userId},
	}
}

func (e *UserId) Query() string {
	return e.query
}

func (e *UserId) Args() []interface{} {
	return e.args
}
