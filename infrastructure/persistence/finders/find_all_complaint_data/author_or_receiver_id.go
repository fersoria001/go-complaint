package find_all_complaint_data

import "github.com/google/uuid"

type AuthorOrReceiverId struct {
	query string
	args  []interface{}
}

func ByAuthorOrReceiverId(id uuid.UUID) *AuthorOrReceiverId {
	return &AuthorOrReceiverId{
		query: string(`SELECT
		ID,
		OWNER_ID,
		AUTHOR_ID,
		RECEIVER_ID,
		COMPLAINT_ID,
		OCCURRED_ON,
		DATA_TYPE
		FROM COMPLAINT_DATA
		WHERE AUTHOR_ID = $1 OR RECEIVER_ID = $1`),
		args: []interface{}{id},
	}
}

func (e *AuthorOrReceiverId) Query() string {
	return e.query
}

func (e *AuthorOrReceiverId) Args() []interface{} {
	return e.args
}
