package find_all_complaints

type Author struct {
	query string
	args  []interface{}
}

func ByAuthor(authorID string) *Author {
	return &Author{
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
	WHERE author_id = $1
	`),
		args: []interface{}{authorID},
	}
}

func (e *Author) Query() string {
	return e.query
}

func (e *Author) Args() []interface{} {
	return e.args
}
