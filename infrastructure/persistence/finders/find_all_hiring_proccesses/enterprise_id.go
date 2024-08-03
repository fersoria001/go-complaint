package find_all_hiring_proccesses

import "github.com/google/uuid"

type EnterpriseId struct {
	query string
	args  []interface{}
}

func ByEnterpriseId(enterpriseId uuid.UUID) *EnterpriseId {
	return &EnterpriseId{
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
	UPDATED_BY_ID
	FROM HIRING_PROCCESSES WHERE ENTERPRISE_ID = $1 `),
		args: []interface{}{enterpriseId},
	}
}

func (e *EnterpriseId) Query() string {
	return e.query
}

func (e *EnterpriseId) Args() []interface{} {
	return e.args
}
