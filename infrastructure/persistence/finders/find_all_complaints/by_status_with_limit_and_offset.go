package find_all_complaints

type StatusStatusAndReceiverIDOrAuthorIDWithLimitAndOffset struct {
	query string
	args  []interface{}
}

func ByStatusStatusAndReceiverIDOrAuthorIDWithLimitAndOffset(complaintStatus, ownerID string, limit, offset int) *StatusStatusAndReceiverIDOrAuthorIDWithLimitAndOffset {
	return &StatusStatusAndReceiverIDOrAuthorIDWithLimitAndOffset{
		query: `
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
	WHERE 
	complaint_status = $1 AND receiver_id = $2
	OR
	complaint_status = $1 AND author_id = $2
	LIMIT $3
	OFFSET $4
	`,
		args: []interface{}{complaintStatus, ownerID, limit, offset},
	}
}

func (e *StatusStatusAndReceiverIDOrAuthorIDWithLimitAndOffset) Query() string {
	return e.query
}

func (e *StatusStatusAndReceiverIDOrAuthorIDWithLimitAndOffset) Args() []interface{} {
	return e.args
}
