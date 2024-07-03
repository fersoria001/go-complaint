package find_all_complaints

type ReceiverAndStatusIn struct {
	query string
	args  []interface{}
}

func ByReceiverAndStatusIn(receiverID string, status []string) *ReceiverAndStatusIn {
	return &ReceiverAndStatusIn{
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
	WHERE receiver_id = $1 AND
	complaint_status = any ($2)
	`),
		args: []interface{}{receiverID, status},
	}
}

func (e *ReceiverAndStatusIn) Query() string {
	return e.query
}

func (e *ReceiverAndStatusIn) Args() []interface{} {
	return e.args
}
