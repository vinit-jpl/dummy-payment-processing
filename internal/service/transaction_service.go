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

func (s *transactionService) GetTransactionStatus(ctx context.Context, txnId string) (*dto.GetTransactionStatusResponse, error) {

	s.mu.RLock()
	defer s.mu.RUnlock()
	// check if txnID exists or not
	txn, exists := s.transactions[txnId]
	if !exists {
		return &dto.GetTransactionStatusResponse{}, errors.New("transaction id not found")
	}

	resp := &dto.GetTransactionStatusResponse{
		TransactionID: txn.TransactionID,
		Status:        txn.Status,
		Amount:        txn.Amount,
		ReferenceID:   txn.ReferenceID,
	}

	// if present return the response. nil
	return resp, nil
}

func (s *transactionService) GetTransactionStats(ctx context.Context) (*dto.GetTransactionStatsResponse, error) {

	// check if any transactions are present
	s.mu.RLock()
	defer s.mu.RUnlock()
	size := len(s.transactions)

	resp := &dto.GetTransactionStatsResponse{}

	if size == 0 {
		return &dto.GetTransactionStatsResponse{}, errors.New("no transactions found")
	}

	for _, txn := range s.transactions {
		switch txn.Status {
		case "success":
			resp.Success++
		case "failed":
			resp.Failed++
		case "processing":
			resp.Processing++
		}
	}

	return resp, nil
}

func (s *transactionService) GetAllTransactions(ctx context.Context) *dto.GetAllTransactionsResponse {

	s.mu.RLock()
	defer s.mu.RUnlock()

	txns := make([]models.Transaction, 0, len(s.transactions))

	for _, txn := range s.transactions {
		txns = append(txns, txn)
	}

	resp := &dto.GetAllTransactionsResponse{
		Count:        len(txns),
		Transactions: txns,
	}

	return resp

}
