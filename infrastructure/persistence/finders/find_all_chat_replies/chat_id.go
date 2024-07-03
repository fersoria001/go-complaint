package find_all_chat_replies

import (
	"go-complaint/domain/model/enterprise"
)

type ChatID struct {
	query string
	args  []interface{}
}

func ByChatID(chatID enterprise.ChatID) *ChatID {
	one := "%" + chatID.String() + "%"
	two := "%" + chatID.Reverse().String() + "%"

	return &ChatID{
		query: string(`
		SELECT 
			ID,
			CHAT_ID,
			USER_ID,
			CONTENT,
			SEEN,
			CREATED_AT,
			UPDATED_AT
		FROM chat_reply
		WHERE chat_id LIKE $1
 		OR chat_id LIKE $2
		`),
		args: []interface{}{one, two},
	}
}

func (e *ChatID) Query() string {
	return e.query
}

func (e *ChatID) Args() []interface{} {
	return e.args
}
