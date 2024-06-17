package find_all_complaints

type ByAuthorIDWithLimitAndOffset struct {
	query string
	args  []interface{}
}

func NewByAuthorIDWithLimitAndOffset(authorID string, limit, offset int) *ByAuthorIDWithLimitAndOffset {
	return &ByAuthorIDWithLimitAndOffset{
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
	WHERE "complaint".author_id = $1
	LIMIT $2
	OFFSET $3
	`),
		args: []interface{}{authorID, limit, offset},
	}
}

func (e *ByAuthorIDWithLimitAndOffset) Query() string {
	return e.query
}

func (e *ByAuthorIDWithLimitAndOffset) Args() []interface{} {
	return e.args
}
