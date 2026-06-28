package models

type Transaction struct {
	TransactionID string `json:"transaction_id"`
	Amount        int    `json:"amount"`
	ReferenceID   string `json:"reference_id"`
	Status        string `json:"status"`
}
