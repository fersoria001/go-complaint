package count_complaints_where

type StatusAndReceiverIDOrAuthorID struct {
	query string
	args  []interface{}
}

func StatusAndReceiverIDOrAuthorIDAre(complaintStatusAndReceiverIDOrAuthorID, ownerID string) *StatusAndReceiverIDOrAuthorID {
	return &StatusAndReceiverIDOrAuthorID{
		query: string(`
	SELECT
 		COUNT(*)
	FROM 
	public."complaint"
	WHERE 
	"complaint".complaint_status = $1 AND "complaint".receiver_id = $2
	OR
	"complaint".complaint_status = $1 AND "complaint".author_id = $2
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
