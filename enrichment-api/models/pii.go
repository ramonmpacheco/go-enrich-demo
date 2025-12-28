package models

type Pii struct {
	Name  string `json:"name,omitempty"`
	Cpf   string `json:"cpf,omitempty"`
	Email string `json:"email,omitempty"`
}
