package count_complaints

type ReceiverID struct {
	query string
	args  []interface{}
}

func WhereReceiverID(receiverID string) *ReceiverID {
	return &ReceiverID{
		query: string(`
	SELECT
 		COUNT(*)
	FROM 
	complaint
	WHERE receiver_id = $1`),
		args: []interface{}{receiverID},
	}
}

func (e *ReceiverID) Query() string {
	return e.query
}

func (e *ReceiverID) Args() []interface{} {
	return e.args
}
