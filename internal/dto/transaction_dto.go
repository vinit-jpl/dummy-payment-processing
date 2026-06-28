package dto

import model "dummy-payment-processing/internal/models"

type (
	CreateTransactionRequest struct {
		Amount      int    `json:"amount"`
		ReferenceID string `json:"reference_id"`
	}

	CreateTransactionResponse struct {
		TransactionID string `json:"transaction_id"`
		Status        string `json:"status"`
	}
)

type GetTransactionStatusResponse struct {
	TransactionID string `json:"transaction_id"`
	Status        string `json:"status"`
	Amount        int    `json:"amount"`
	ReferenceID   string `json:"reference_id"`
}

type GetAllTransactionsResponse struct {
	Count        int                 `json:"count"`
	Transactions []model.Transaction `json:"transactions"`
}
