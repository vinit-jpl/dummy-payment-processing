package service

import (
	"context"
	"dummy-payment-processing/internal/dto"
	"dummy-payment-processing/internal/models"
	"sync"
)

type TransactionService interface {
	CreateTransaction(ctx context.Context, req *dto.CreateTransactionRequest) (*dto.CreateTransactionResponse, error)
	GetTransactionStatus(ctx context.Context, txnId string) (*dto.GetTransactionStatusResponse, error)
	GetTransactionStats(ctx context.Context) (*dto.GetTransactionStatsResponse, error)
}
type transactionService struct {
	transactions map[string]models.Transaction
	mu           sync.RWMutex
}

func NewTransactionService() TransactionService {
	return &transactionService{
		transactions: make(map[string]models.Transaction),
	}
}
