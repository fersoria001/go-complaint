package find_all_complaint_reviews

import "github.com/google/uuid"

type RatedById struct {
	query string
	args  []interface{}
}

func ByRatedById(ratedById uuid.UUID) *RatedById {
	return &RatedById{
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
	WHERE RATED_BY_ID=$1
	`),
		args: []interface{}{ratedById},
	}
}

func (e *RatedById) Query() string {
	return e.query
}

func (e *RatedById) Args() []interface{} {
	return e.args
}
