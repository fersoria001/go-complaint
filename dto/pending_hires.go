package dto

type PendingHires struct {
	EventID    string `json:"event_id"`
	User       User   `json:"user"`
	Position   string `json:"position"`
	OccurredOn string `json:"occurred_on"`
}
