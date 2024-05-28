package dto

type Country struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type County struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type City struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type PhoneCode struct {
	ID   int    `json:"id"`
	Code string `json:"code"`
}
