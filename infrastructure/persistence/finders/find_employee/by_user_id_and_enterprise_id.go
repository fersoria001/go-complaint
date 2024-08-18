package find_employee

import "github.com/google/uuid"

type UserIdAndEnterpriseId struct {
	query string
	args  []interface{}
}

func ByUserIdAndEnterpriseId(userId, enterpriseId uuid.UUID) *UserIdAndEnterpriseId {
	return &UserIdAndEnterpriseId{
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
		WHERE user_id = $1 AND enterprise_id = $2`,
		),
		args: []interface{}{enterpriseId},
	}
}

func (e *UserIdAndEnterpriseId) Query() string {
	return e.query
}

func (e *UserIdAndEnterpriseId) Args() []interface{} {
	return e.args
}
