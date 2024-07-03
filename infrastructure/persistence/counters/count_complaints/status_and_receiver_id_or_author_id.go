package count_complaints

type StatusAndReceiverIDOrAuthorID struct {
	query string
	args  []interface{}
}

func WhereStatusAndReceiverIDOrAuthorID(complaintStatusAndReceiverIDOrAuthorID, ownerID string) *StatusAndReceiverIDOrAuthorID {
	return &StatusAndReceiverIDOrAuthorID{
		query: string(`
	SELECT
 		COUNT(*)
	FROM 
	complaint
	WHERE 
	complaint_status = $1 AND receiver_id = $2
	OR
	complaint_status = $1 AND author_id = $2
	`),
		args: []interface{}{complaintStatusAndReceiverIDOrAuthorID, ownerID},
	}
}

func (e *StatusAndReceiverIDOrAuthorID) Query() string {
	return e.query
}

func (e *StatusAndReceiverIDOrAuthorID) Args() []interface{} {
	return e.args
}
