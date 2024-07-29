package find_all_complaint_reviews

import "github.com/google/uuid"

type TriggeredById struct {
	query string
	args  []interface{}
}

func ByTriggeredById(triggeredById uuid.UUID) *TriggeredById {
	return &TriggeredById{
		query: string(`
	SELECT 	
	ID,
	ATTENTION_BY_ID,
	TRIGGERED_BY_ID,
	RATED_BY_ID,
	COMPLAINT_ID,
	COMPLAINT_TITLE,
	STATUS,
	OCCURRED_ON
	FROM COMPLAINT_REVIEWS
	WHERE TRIGGERED_BY_ID=$1
	`),
		args: []interface{}{triggeredById},
	}
}

func (e *TriggeredById) Query() string {
	return e.query
}

func (e *TriggeredById) Args() []interface{} {
	return e.args
}
