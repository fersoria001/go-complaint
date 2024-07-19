package find_all_enterprises

type OwnerId struct {
	query string
	args  []interface{}
}

func ByOwnerId(term string) *OwnerId {
	return &OwnerId{
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
		WHERE owner_user_id = $1`,
		),
		args: []interface{}{term},
	}
}

func (e *OwnerId) Query() string {
	return e.query
}

func (e *OwnerId) Args() []interface{} {
	return e.args
}
