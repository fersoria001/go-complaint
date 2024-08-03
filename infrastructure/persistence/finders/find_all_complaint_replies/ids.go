package find_all_complaint_replies

import "github.com/google/uuid"

type IDs struct {
	query string
	args  []interface{}
}

func ByAnyIDs(ids []uuid.UUID) *IDs {
	return &IDs{
		query: string(`
	SELECT
	id,
	complaint_id,
	author_id,
	body,
	is_read,
	read_at,
	created_at,
	updated_at
	FROM complaint_replies
	WHERE id = any($1)
	`),
		args: []interface{}{ids},
	}
}

func (e *IDs) Query() string {
	return e.query
}

func (e *IDs) Args() []interface{} {
	return e.args
}
