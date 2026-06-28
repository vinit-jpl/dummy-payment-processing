package service

import (
	"context"
	"dummy-payment-processing/internal/dto"
	"dummy-payment-processing/internal/models"
	"dummy-payment-processing/internal/utils"
	"errors"
)

func (s *transactionService) processTransaction(txnId string) error {
	status := utils.ProcessTransaction(txnId)

	s.mu.Lock()
	defer s.mu.Unlock()

	txn, exists := s.transactions[txnId]
	if !exists {
		return errors.New("transaction not found")
	}
	txn.Status = status
	s.transactions[txnId] = txn

	return nil
}

func (s *transactionService) CreateTransaction(ctx context.Context, req *dto.CreateTransactionRequest) (*dto.CreateTransactionResponse, error) {

	s.mu.Lock()
	defer s.mu.Unlock()
	// generate unique transaction id
	txnId := utils.GenerateTransactionIdString()

	// store transaction in memory and set initial state as processing
	transaction := models.Transaction{
		TransactionID: txnId,
		Amount:        req.Amount,
		ReferenceID:   req.ReferenceID,
		Status:        "processing",
	}

	s.transactions[txnId] = transaction

	//  start async processing with dealy of 2-5 seconds
	go s.processTransaction(txnId)

	return &dto.CreateTransactionResponse{
		TransactionID: txnId,
		Status:        "processing",
	}, nil

}
