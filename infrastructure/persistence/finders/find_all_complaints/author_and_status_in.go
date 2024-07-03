package find_all_complaints

type AuthorAndStatusIn struct {
	query string
	args  []interface{}
}

func ByAuthorAndStatusIn(authorID string, status []string) *AuthorAndStatusIn {
	return &AuthorAndStatusIn{
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
	WHERE author_id = $1 AND
	complaint_status = any ($2)
	`),
		args: []interface{}{authorID, status},
	}
}

func (e *AuthorAndStatusIn) Query() string {
	return e.query
}

func (e *AuthorAndStatusIn) Args() []interface{} {
	return e.args
}
