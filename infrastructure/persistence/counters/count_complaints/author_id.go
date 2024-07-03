package count_complaints

type AuthorID struct {
	query string
	args  []interface{}
}

func WhereAuthorID(authorID string) *AuthorID {
	return &AuthorID{
		query: string(`
	SELECT
 		COUNT(*)
	FROM 
	complaint
	WHERE author_id = $1`),
		args: []interface{}{authorID},
	}
}

func (e *AuthorID) Query() string {
	return e.query
}

func (e *AuthorID) Args() []interface{} {
	return e.args
}
