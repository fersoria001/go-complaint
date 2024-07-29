package count_complaints

import "github.com/google/uuid"

type AuthorOrReceiverId struct {
	query string
	args  []interface{}
}

func WhereAuthorOrReceiverId(authorId uuid.UUID, status []string) *AuthorOrReceiverId {
	return &AuthorOrReceiverId{
		query: string(`
	SELECT
 		COUNT(*)
	FROM complaint WHERE author_id = $1 OR receiver_id = $1
	AND status = any($2)`),
		args: []interface{}{authorId, status},
	}
}

func (e *AuthorOrReceiverId) Query() string {
	return e.query
}

func (e *AuthorOrReceiverId) Args() []interface{} {
	return e.args
}
