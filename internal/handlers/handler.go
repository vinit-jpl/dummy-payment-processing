package handlers

import "dummy-payment-processing/internal/service"

type TransactionHandler struct {
	service service.TransactionService
}

func NewTransactionHandler(s service.TransactionService) *TransactionHandler {
	return &TransactionHandler{
		service: s,
	}
}
