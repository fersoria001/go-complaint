package find_chats

import "github.com/google/uuid"

type RecipientOneTwoAndEnterpriseId struct {
	query string
	args  []interface{}
}

func ByRecipientOneTwoAndEnterpriseId(recipientOne, recipientTwo, enterpriseId uuid.UUID) *RecipientOneTwoAndEnterpriseId {
	return &RecipientOneTwoAndEnterpriseId{
		query: string(`
		SELECT ID,
		ENTERPRISE_ID,
		RECIPIENT_ONE_ID,
		RECIPIENT_TWO_ID
		FROM chats
			WHERE recipient_one_id = $1 AND recipient_two_id = $2 and enterprise_id = $3 OR
			recipient_one_id = $2 AND recipient_two_id = $1 and enterprise_id = $3
	`),
		args: []interface{}{recipientOne, recipientTwo, enterpriseId},
	}
}

func (e *RecipientOneTwoAndEnterpriseId) Query() string {
	return e.query
}

func (e *RecipientOneTwoAndEnterpriseId) Args() []interface{} {
	return e.args
}
