package find_all_enterprises

import "strings"

type ByNameLike struct {
	query string
	args  []interface{}
}

func NewByNameLike(term string) *ByNameLike {
	wildCardTerm := "%" + strings.ToLower(term) + "%"
	return &ByNameLike{
		query: string(
			`SELECT
		enterprise_id,
		owner_user_id,
		logo_img,
		banner_img,
		website,
		email,
		phone,
		address_id,
		industry_id,
		created_at,
		updated_at,
		foundation_date
		FROM enterprise
		WHERE LOWER(enterprise_id) LIKE $1`,
		),
		args: []interface{}{wildCardTerm},
	}
}

func (e *ByNameLike) Query() string {
	return e.query
}

func (e *ByNameLike) Args() []interface{} {
	return e.args
}
