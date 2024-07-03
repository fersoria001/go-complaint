package find_all_complaints

type ByReceiverIDWithLimitAndOffset struct {
	query string
	args  []interface{}
}

func NewByReceiverIDWithLimitAndOffset(receiverID string, limit, offset int) *ByReceiverIDWithLimitAndOffset {
	return &ByReceiverIDWithLimitAndOffset{
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
	WHERE receiver_id = $1
	LIMIT $2
	OFFSET $3
	`),
		args: []interface{}{receiverID, limit, offset},
	}
}

func (e *ByReceiverIDWithLimitAndOffset) Query() string {
	return e.query
}

func (e *ByReceiverIDWithLimitAndOffset) Args() []interface{} {
	return e.args
}
