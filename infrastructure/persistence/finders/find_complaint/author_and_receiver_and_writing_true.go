package find_complaint

import "github.com/google/uuid"

type AuthorAndReceiverAndWritingTrue struct {
	query string
	args  []interface{}
}

func ByAuthorAndReceiverAndWritingTrue(authorId, receiverId uuid.UUID) *AuthorAndReceiverAndWritingTrue {
	return &AuthorAndReceiverAndWritingTrue{
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
	FROM complaint WHERE author_id = $1 AND receiver_id = $2
	AND status = 'WRITING'
	`),
		args: []interface{}{authorId, receiverId},
	}
}

func (e *AuthorAndReceiverAndWritingTrue) Query() string {
	return e.query
}

func (e *AuthorAndReceiverAndWritingTrue) Args() []interface{} {
	return e.args
}
