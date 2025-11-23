package model

type Transaction struct {
	Timestamp   int64  `json:"timestamp"`
	Name        string `json:"name"`
	Type        string `json:"type"` // "DEBIT" or "CREDIT"
	Amount      int64  `json:"amount"`
	Status      string `json:"status"` // "SUCCESS", "FAILED", "PENDING", etc.
	Description string `json:"description"`
}
