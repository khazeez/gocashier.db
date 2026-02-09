package services

import (
	"time"

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


func (t *TransactionService) GetReportToday() (*models.TransactionReport, error) {
	return t.transactionRepo.GetReportToday()
}

func (t *TransactionService) GetReportWithRange(startDate time.Time, endDate time.Time) (*models.TransactionReport, error) {
	return t.transactionRepo.GetReportWithRange(startDate, endDate)
}