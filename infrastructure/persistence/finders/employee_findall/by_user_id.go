package employeefindall

type ByUserID struct {
	query string
	args  []interface{}
}

func NewByUserID(email string) *ByUserID {
	return &ByUserID{
		query: string(
			`SELECT 
		employee_id,
		enterprise_id,
		user_id,
		hiring_date,
		approved_hiring,
		approved_hiring_at,
		job_position
		FROM employee
		WHERE user_id = $1`,
		),
		args: []interface{}{email},
	}
}

func (e *ByUserID) Query() string {
	return e.query
}

func (e *ByUserID) Args() []interface{} {
	return e.args
}
