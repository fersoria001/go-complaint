package count_complaints

type ReceiverIDAndStatus struct {
	query string
	args  []interface{}
}

func WhereReceiverIDAndStatusIN(authorID string, statusSlice []string) *ReceiverIDAndStatus {
	return &ReceiverIDAndStatus{
		query: string(`
	SELECT
 		COUNT(*)
	FROM 
	complaint
	WHERE receiver_id = $1 AND complaint_status = any ($2)`),
		args: []interface{}{authorID, statusSlice},
	}
}

func (e *ReceiverIDAndStatus) Query() string {
	return e.query
}

func (e *ReceiverIDAndStatus) Args() []interface{} {
	return e.args
}
