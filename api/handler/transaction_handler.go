package handler

import (
	"github.com/gin-gonic/gin"
	"gocashier.db/internal/services"
)

type transactionHandler struct {
	transactionService services.TransactionService
}

func NewTransactionHandler(transaactionService services.TransactionService) *transaactionHandler {
	return &transactionHandler{
		transactionService: transaactionService,
	}
}

func (t *transactionHandler) CreateTransaction(h *gin.Context) {
	
}