package employeefindall

type ByEnterpriseID struct {
	query string
	args  []interface{}
}

func NewByEnterpriseID(enterpriseID string) *ByEnterpriseID {
	return &ByEnterpriseID{
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
		WHERE enterprise_id = $1`,
		),
		args: []interface{}{enterpriseID},
	}
}

func (e *ByEnterpriseID) Query() string {
	return e.query
}

func (e *ByEnterpriseID) Args() []interface{} {
	return e.args
}
