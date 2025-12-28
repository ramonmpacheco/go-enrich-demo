package models

type Address struct {
	Street  string `json:"street,omitempty"`
	Number  string `json:"number,omitempty"`
	City    string `json:"city,omitempty"`
	State   string `json:"state,omitempty"`
	Country string `json:"country,omitempty"`
}
