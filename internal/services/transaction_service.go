package services

import (
	"gocashier.db/internal/models"
	"gocashier.db/internal/repository"
)

type TransactionService struct {
	transactionRepo repository.TransactionRepository
}

func NewTransactionService(transactionRepo repository.TransactionRepository) TransactionService {
	return TransactionService{
		transactionRepo: transactionRepo,
	}
}

func (t *TransactionService) CreateTransaction(items []models.CheckoutItem) (*models.Transaction, error) {
	return t.transactionRepo.CreateTransaction(items)
}
