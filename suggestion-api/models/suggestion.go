package models

type Suggestion struct {
	Name      string `json:"name"`
	Promotion string `json:"promotion"`
}

type SuggestionRequest struct {
	CustomerCode string `json:"customer_code"`
}
