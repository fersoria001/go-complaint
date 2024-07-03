package find_all_complaints

type ByAuthorIDWithLimitAndOffset struct {
	query string
	args  []interface{}
}

func NewByAuthorIDWithLimitAndOffset(authorID string, limit, offset int) *ByAuthorIDWithLimitAndOffset {
	return &ByAuthorIDWithLimitAndOffset{
		query: string(`
	SELECT
	id,
	author_id,
	receiver_id,
	complaint_status,
	title,
	complaint_description,
	body,
	rating_rate,
	rating_comment,
	created_at,
	updated_at
	FROM 
	complaint
	WHERE author_id = $1
	LIMIT $2
	OFFSET $3
	`),
		args: []interface{}{authorID, limit, offset},
	}
}

func (e *ByAuthorIDWithLimitAndOffset) Query() string {
	return e.query
}

func (e *ByAuthorIDWithLimitAndOffset) Args() []interface{} {
	return e.args
}
