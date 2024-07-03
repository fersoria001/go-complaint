package find_all_complaints

type ReceiverAndStatus struct {
	query string
	args  []interface{}
}

func ByReceiverAndStatus(receiverID string, status string) *ReceiverAndStatus {
	return &ReceiverAndStatus{
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
	complaint_status = $2
	`),
		args: []interface{}{receiverID, status},
	}
}

func (e *ReceiverAndStatus) Query() string {
	return e.query
}

func (e *ReceiverAndStatus) Args() []interface{} {
	return e.args
}
