package services

import (
	"gocashier.db/internal/models"
	"gocashier.db/internal/repository"
)

type transactionService struct {
	transactionRepo repository.TransactionRepository
}

func NewTransactionRepository(transactionRepo repository.TransactionRepository) transactionService {
	return transactionService{
		transactionRepo: transactionRepo,
	}
}


func (t *transactionService) CreateTransaction(items []models.CheckoutItem) (*models.Transaction, error) {
		return t.transactionRepo.CreateTransaction(items)
}
