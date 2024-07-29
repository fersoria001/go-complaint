package dto

type Message struct {
	Content       string `json:"content"`
	Reply         Reply  `json:"reply"`
	StatusChanged bool   `json:"status_changed"`
}
