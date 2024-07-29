package find_all_enterprises

import "github.com/google/uuid"

type AnyId struct {
	query string
	args  []interface{}
}

func ByAnyId(idSlice []uuid.UUID) *AnyId {
	return &AnyId{
		query: string(
			`SELECT
		enterprise_id,
		enterprise_name,
		owner_user_id,
		logo_img,
		banner_img,
		website,
		email,
		phone,
		industry_id,
		created_at,
		updated_at,
		foundation_date
		FROM enterprise
		WHERE enterprise_id = any($1)`,
		),
		args: []interface{}{idSlice},
	}
}

func (e *AnyId) Query() string {
	return e.query
}

func (e *AnyId) Args() []interface{} {
	return e.args
}
