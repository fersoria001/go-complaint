package count_complaints

type AuthorIDAndStatus struct {
	query string
	args  []interface{}
}

func WhereAuthorIDAndStatus(authorID string, statusSlice []string) *AuthorIDAndStatus {
	return &AuthorIDAndStatus{
		query: string(`
	SELECT
 		COUNT(*)
	FROM 
	complaint
	WHERE author_id = $1 AND complaint_status  = any ($2)`),
		args: []interface{}{authorID, statusSlice},
	}
}

func (e *AuthorIDAndStatus) Query() string {
	return e.query
}

func (e *AuthorIDAndStatus) Args() []interface{} {
	return e.args
}
