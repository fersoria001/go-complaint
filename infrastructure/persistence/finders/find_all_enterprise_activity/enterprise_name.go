package find_all_enterprise_activity

type EnterpriseName struct {
	query string
	args  []interface{}
}

func ByEnterpriseName(enterpriseName string) *EnterpriseName {
	return &EnterpriseName{
		query: string(`
	SELECT
		ID,
		USER_ID,
		ACTIVITY_ID,
		ENTERPRISE_ID,
		ENTERPRISE_NAME,
		OCCURRED_ON,
		ACTIVITY_TYPE
	FROM ENTERPRISE_ACTIVITY
	WHERE ENTERPRISE_NAME = $1
	`),
		args: []interface{}{enterpriseName},
	}
}

func (e *EnterpriseName) Query() string {
	return e.query
}

func (e *EnterpriseName) Args() []interface{} {
	return e.args
}
