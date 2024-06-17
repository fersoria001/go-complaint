package count_complaints_where

type ReceiverID struct {
	query string
	args  []interface{}
}

func NewReceiverID(receiverID string) *ReceiverID {
	return &ReceiverID{
		query: string(`
	SELECT
 		COUNT(*)
	FROM 
	public."complaint"
	WHERE "complaint".receiver_id = $1`),
		args: []interface{}{receiverID},
	}
}

func (e *ReceiverID) Query() string {
	return e.query
}

func (e *ReceiverID) Args() []interface{} {
	return e.args
}
