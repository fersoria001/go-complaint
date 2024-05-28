package dto

type Message struct {
	Content string   `json:"content"`
	Reply   ReplyDTO `json:"reply"`
}
