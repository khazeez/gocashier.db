package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gocashier.db/internal/models"
	"gocashier.db/internal/services"
)

type transactionHandler struct {
	transactionService services.TransactionService
}

func NewTransactionHandler(transaactionService services.TransactionService) *transactionHandler {
	return &transactionHandler{
		transactionService: transaactionService,
	}
}

func (t *transactionHandler) CreateTransaction(h *gin.Context) {
	var items models.CheckoutRequest
	err := h.ShouldBindJSON(&items)
	if err != nil {
		h.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request payload",
		})
		return
	}

	data, err := t.transactionService.CreateTransaction(items.Items)
	if err != nil {
		h.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	h.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": data,
	})

}