package find_all_complaint_data

import "github.com/google/uuid"

type OwnerIdAndNotDataOwnership struct {
	query string
	args  []interface{}
}

func ByOwnerIdAndNotDataOwnership(ownerId uuid.UUID) *OwnerIdAndNotDataOwnership {
	return &OwnerIdAndNotDataOwnership{
		query: string(`SELECT
		ID,
		OWNER_ID,
		AUTHOR_ID,
		RECEIVER_ID,
		COMPLAINT_ID,
		OCCURRED_ON,
		DATA_TYPE
		FROM COMPLAINT_DATA
		WHERE 
		OWNER_ID=$1 AND RECEIVER_ID != $1 AND AUTHOR_ID != $1`),
		args: []interface{}{ownerId},
	}
}

func (e *OwnerIdAndNotDataOwnership) Query() string {
	return e.query
}

func (e *OwnerIdAndNotDataOwnership) Args() []interface{} {
	return e.args
}
