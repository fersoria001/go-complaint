package find_all_complaints

import "github.com/google/uuid"

type AuthorOrReceiver struct {
	query string
	args  []interface{}
}

func ByAuthorOrReceiver(id uuid.UUID, status []string) *AuthorOrReceiver {
	return &AuthorOrReceiver{
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
	FROM complaint WHERE author_id = $1 OR receiver_id = $1
	AND status = any($2)
	ORDER BY CREATED_AT
	`),
		args: []interface{}{id, status},
	}
}

func (e *AuthorOrReceiver) Query() string {
	return e.query
}

func (e *AuthorOrReceiver) Args() []interface{} {
	return e.args
}
