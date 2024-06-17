package find_all_complaints

type ByReceiverIDWithLimitAndOffset struct {
	query string
	args  []interface{}
}

func NewByReceiverIDWithLimitAndOffset(receiverID string, limit, offset int) *ByReceiverIDWithLimitAndOffset {
	return &ByReceiverIDWithLimitAndOffset{
		query: string(`
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
	WHERE "complaint".receiver_id = $1
	LIMIT $2
	OFFSET $3
	`),
		args: []interface{}{receiverID, limit, offset},
	}
}

func (e *ByReceiverIDWithLimitAndOffset) Query() string {
	return e.query
}

func (e *ByReceiverIDWithLimitAndOffset) Args() []interface{} {
	return e.args
}
