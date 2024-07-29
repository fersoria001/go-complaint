package find_all_notifications

import "github.com/google/uuid"

type OwnerId struct {
	query string
	args  []interface{}
}

func ByOwnerId(ownerId uuid.UUID) *OwnerId {
	return &OwnerId{
		query: string(`
		SELECT
			id,
			owner_id,
			sender_id,
			title,
			content,
			link,
			occurred_on,
			seen
		FROM notifications
		WHERE owner_id = $1
		ORDER BY OCCURRED_ON
		`),
		args: []interface{}{ownerId},
	}
}

func (e *OwnerId) Query() string {
	return e.query
}

func (e *OwnerId) Args() []interface{} {
	return e.args
}
