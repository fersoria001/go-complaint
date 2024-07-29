package find_all_complaint_reviews

import "github.com/google/uuid"

type AttentionById struct {
	query string
	args  []interface{}
}

func ByAttentionById(attentionById uuid.UUID) *AttentionById {
	return &AttentionById{
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
	WHERE ATTENTION_BY_ID=$1
	`),
		args: []interface{}{attentionById},
	}
}

func (e *AttentionById) Query() string {
	return e.query
}

func (e *AttentionById) Args() []interface{} {
	return e.args
}
