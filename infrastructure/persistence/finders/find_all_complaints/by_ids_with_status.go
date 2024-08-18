package find_all_complaints

import "github.com/google/uuid"

type IdsWithStatus struct {
	query string
	args  []interface{}
}

func ByIdsWithStatus(ids []uuid.UUID, status []string) *IdsWithStatus {
	return &IdsWithStatus{
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
	FROM complaint WHERE id = any($1) AND status = any($2)
	ORDER BY CREATED_AT
	`),
		args: []interface{}{ids, status},
	}
}

func (e *IdsWithStatus) Query() string {
	return e.query
}

func (e *IdsWithStatus) Args() []interface{} {
	return e.args
}
