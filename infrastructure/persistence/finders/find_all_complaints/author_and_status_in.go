package find_all_complaints

import "github.com/google/uuid"

type AuthorAndStatusIn struct {
	query string
	args  []interface{}
}

func ByAuthorAndStatusIn(authorId uuid.UUID, status []string) *AuthorAndStatusIn {
	return &AuthorAndStatusIn{
		query: string(`
	SELECT
	id,
	author_id,
	receiver_id,
	status,
	title,
	description,
	created_at,
	updated_at
	FROM complaint WHERE author_id = $1
	AND status = any($2)
	ORDER BY CREATED_AT
	`),
		args: []interface{}{authorId, status},
	}
}

func (e *AuthorAndStatusIn) Query() string {
	return e.query
}

func (e *AuthorAndStatusIn) Args() []interface{} {
	return e.args
}
