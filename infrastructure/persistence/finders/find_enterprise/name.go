package find_enterprise

type Name struct {
	query string
	args  []interface{}
}

func ByName(name string) *Name {
	return &Name{
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
		WHERE enterprise_name = $1`,
		),
		args: []interface{}{name},
	}
}

func (e *Name) Query() string {
	return e.query
}

func (e *Name) Args() []interface{} {
	return e.args
}
