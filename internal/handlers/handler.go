package handlers

import "dummy-payment-processing/internal/service"

type TransactionHandler struct {
	service service.TransactionService
}

// accept interface return concrete types
func NewTransactionHandler(s service.TransactionService) *TransactionHandler {
	return &TransactionHandler{
		service: s,
	}
}
