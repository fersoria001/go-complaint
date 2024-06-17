package employeefindall

type ByUserIDAndEnterpriseID struct {
	query string
	args  []interface{}
}

func NewByUserIDAndEnterpriseID(userID, enterpriseID string) *ByUserIDAndEnterpriseID {
	return &ByUserIDAndEnterpriseID{
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
		WHERE enterprise_id = $1 AND user_id = $2`,
		),
		args: []interface{}{userID, enterpriseID},
	}
}

func (e *ByUserIDAndEnterpriseID) Query() string {
	return e.query
}

func (e *ByUserIDAndEnterpriseID) Args() []interface{} {
	return e.args
}
