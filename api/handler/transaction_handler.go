package handler

import (
	"net/http"
	"time"

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
		"data":    data,
	})

}

func (t *transactionHandler) GetReportToday(h *gin.Context) {
	data, err := t.transactionService.GetReportToday()
	if err != nil {
		h.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	h.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    data,
	})
}

func (t *transactionHandler) GetReportWithRange(h *gin.Context) {
	startDateStr := h.Query("start_date")
	endDateStr   := h.Query("end_date")

	if startDateStr == "" || endDateStr == "" {
		h.JSON(400, gin.H{"error": "start_date and end_date required"})
		return
	}

	layout := "2006-01-02"

	startDate, err := time.Parse(layout, startDateStr)
	if err != nil {
		h.JSON(400, gin.H{"error": "invalid start_date format"})
		return
	}

	endDate, err := time.Parse(layout, endDateStr)
	if err != nil {
		h.JSON(400, gin.H{"error": "invalid end_date format"})
		return
	}

	endDate = endDate.AddDate(0, 0, 1)

	report, err := t.transactionService.GetReportWithRange(startDate, endDate)
	if err != nil {
		h.JSON(500, gin.H{"error": err.Error()})
		return
	}

	h.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    report,
	})


}