package find_all_complaints

type Receiver struct {
	query string
	args  []interface{}
}

func ByReceiver(receiverID string) *Receiver {
	return &Receiver{
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
	`),
		args: []interface{}{receiverID},
	}
}

func (e *Receiver) Query() string {
	return e.query
}

func (e *Receiver) Args() []interface{} {
	return e.args
}
