package find_hiring_proccess

import "github.com/google/uuid"

type UserIdAndEnterpriseId struct {
	query string
	args  []interface{}
}

func ByUserIdAndEnterpriseId(userId, enterpriseId uuid.UUID) *UserIdAndEnterpriseId {
	return &UserIdAndEnterpriseId{
		query: string(`SELECT 
	ID,
	ENTERPRISE_ID,
	USER_ID,
	ROLE,
	STATUS,
	REASON,
	EMITED_BY_ID,
	OCCURRED_ON,
	LAST_UPDATE,
	UPDATED_BY_ID,
	INDUSTRY_ID
	FROM HIRING_PROCCESSES WHERE ENTERPRISE_ID = $1 AND USER_ID =$2`),
		args: []interface{}{enterpriseId, userId},
	}
}

func (e *UserIdAndEnterpriseId) Query() string {
	return e.query
}

func (e *UserIdAndEnterpriseId) Args() []interface{} {
	return e.args
}
