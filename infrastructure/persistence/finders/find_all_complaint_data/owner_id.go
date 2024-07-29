package find_all_complaint_data

import "github.com/google/uuid"

type OwnerId struct {
	query string
	args  []interface{}
}

func ByOwnerId(ownerId uuid.UUID) *OwnerId {
	return &OwnerId{
		query: string(`SELECT
		ID,
		OWNER_ID,
		COMPLAINT_ID,
		OCCURRED_ON,
		DATA_TYPE
		FROM COMPLAINT_DATA
		WHERE OWNER_ID=$1`),
		args: []interface{}{ownerId},
	}
}

func (e *OwnerId) Query() string {
	return e.query
}

func (e *OwnerId) Args() []interface{} {
	return e.args
}
