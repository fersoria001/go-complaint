package count_complaints_where

type AuthorID struct {
	query string
	args  []interface{}
}

func NewAuthorID(authorID string) *AuthorID {
	return &AuthorID{
		query: string(`
	SELECT
 		COUNT(*)
	FROM 
	public."complaint"
	WHERE "complaint".author_id = $1`),
		args: []interface{}{authorID},
	}
}

func (e *AuthorID) Query() string {
	return e.query
}

func (e *AuthorID) Args() []interface{} {
	return e.args
}
