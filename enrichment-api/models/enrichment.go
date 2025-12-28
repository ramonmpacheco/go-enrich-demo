package models

type EnrichmentResponse struct {
	Address     *Address     `json:"address"`
	Pii         *Pii         `json:"pii"`
	Suggestions []Suggestion `json:"suggestions"`
	Metadata    *Metadata    `json:"metadata"`
}

type Metadata struct {
	ProcessingTime string `json:"processing_time"`
}

type EnrichmentRequest struct {
	CustomerCode string `json:"customer_code"`
}
