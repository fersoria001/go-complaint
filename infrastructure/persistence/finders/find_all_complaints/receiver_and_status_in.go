package find_all_complaints

import "github.com/google/uuid"

type ReceiverAndStatusIn struct {
	query string
	args  []interface{}
}

func ByReceiverAndStatusIn(receiverId uuid.UUID, status []string) *ReceiverAndStatusIn {
	return &ReceiverAndStatusIn{
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
	FROM complaint WHERE receiver_id = $1
	AND status = any($2)
	ORDER BY CREATED_AT
	`),
		args: []interface{}{receiverId, status},
	}
}

func (e *ReceiverAndStatusIn) Query() string {
	return e.query
}

func (e *ReceiverAndStatusIn) Args() []interface{} {
	return e.args
}
