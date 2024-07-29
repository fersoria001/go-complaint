package employeefindall

import "github.com/google/uuid"

type EnterpriseId struct {
	query string
	args  []interface{}
}

func ByEnterpriseId(enterpriseId uuid.UUID) *EnterpriseId {
	return &EnterpriseId{
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
		args: []interface{}{enterpriseId},
	}
}

func (e *EnterpriseId) Query() string {
	return e.query
}

func (e *EnterpriseId) Args() []interface{} {
	return e.args
}
