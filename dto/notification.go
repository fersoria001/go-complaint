package dto

type Notification struct {
	ID        string `json:"id"`
	OwnerID   string `json:"owner_id"`
	Thumbnail string `json:"thumbnail"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Link      string `json:"link"`
	Seen      bool   `json:"seen"`
}
