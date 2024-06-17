package find_all_feedback_review

type ReviewerID struct {
	query string
	args  []interface{}
}

func ByReviewerID(reviewerID string) *ReviewerID {
	return &ReviewerID{
		query: string(`
		SELECT 
		id,
		reviewer_id,
		reviewed_at,
		review_comment
		FROM feedback_reviews
		WHERE reviewer_id = $1
		`),
		args: []interface{}{reviewerID},
	}
}

func (e *ReviewerID) Query() string {
	return e.query
}

func (e *ReviewerID) Args() []interface{} {
	return e.args
}
