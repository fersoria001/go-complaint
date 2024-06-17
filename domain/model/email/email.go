package email

type Email struct {
	TemplateID string `json:"template_id"`
	From       struct {
		Email string `json:"email"`
		Name  string `json:"name"`
	} `json:"from"`
	To struct {
		Email string `json:"email"`
		Name  string `json:"name"`
	} `json:"to"`
	Subject         string                   `json:"subject"`
	Personalization []map[string]interface{} `json:"personalization"`
}
