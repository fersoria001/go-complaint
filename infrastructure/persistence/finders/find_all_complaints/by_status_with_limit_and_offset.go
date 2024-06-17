package find_all_complaints

type StatusStatusAndReceiverIDOrAuthorIDWithLimitAndOffset struct {
	query string
	args  []interface{}
}

func ByStatusStatusAndReceiverIDOrAuthorIDWithLimitAndOffset(complaintStatus, ownerID string, limit, offset int) *StatusStatusAndReceiverIDOrAuthorIDWithLimitAndOffset {
	return &StatusStatusAndReceiverIDOrAuthorIDWithLimitAndOffset{
		query: `
	SELECT
	"complaint".id,
	"complaint".author_id,
	"complaint".receiver_id,
	"complaint".complaint_status,
	"complaint".title,
	"complaint".complaint_description,
	"complaint".body,
	"complaint".rating_rate,
	"complaint".rating_comment,
	"complaint".created_at,
	"complaint".updated_at
	FROM 
	public."complaint"
	WHERE 
	"complaint".complaint_status = $1 AND "complaint".receiver_id = $2
	OR
	"complaint".complaint_status = $1 AND "complaint".author_id = $2
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
