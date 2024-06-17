package find_all_feedback

type ReviewedID struct {
	query string
	args  []interface{}
}

func ByReviewedID(reviewedID string) *ReviewedID {
	return &ReviewedID{
		query: string(`
		SELECT 
			id,
			complaint_id,
			reviewed_id
		FROM feedback
		WHERE reviewed_id = $1`),
		args: []interface{}{reviewedID},
	}
}

func (e *ReviewedID) Query() string {
	return e.query
}

func (e *ReviewedID) Args() []interface{} {
	return e.args
}
