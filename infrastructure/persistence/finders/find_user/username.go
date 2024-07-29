package find_user

type Username struct {
	query string
	args  []interface{}
}

func ByUsername(username string) *Username {
	return &Username{
		query: string(`
		SELECT 
		id,
		username,
		password,
		register_date,
		is_confirmed
		FROM users
		WHERE username = $1
		`),
		args: []interface{}{username},
	}
}

func (e *Username) Query() string {
	return e.query
}

func (e *Username) Args() []interface{} {
	return e.args
}
