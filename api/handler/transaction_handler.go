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


// CreateTransaction godoc
// @Summary Checkout transaction
// @Description Create new transaction from items
// @Tags transactions
// @Accept json
// @Produce json
// @Param request body models.CheckoutRequest true "Checkout Data"
// @Success 200 {object} map[string]interface{}
// @Router /transaction/checkout [post]
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



// GetReportToday godoc
// @Summary Get today's transaction report
// @Description Get all transactions for today
// @Tags transactions
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /transaction/report/today [get]
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



// GetReportWithRange godoc
// @Summary Get transaction report by date range
// @Description Get report between start_date and end_date
// @Tags transactions
// @Produce json
// @Param start_date query string true "Start Date (YYYY-MM-DD)"
// @Param end_date query string true "End Date (YYYY-MM-DD)"
// @Success 200 {object} map[string]interface{}
// @Router /transaction/report [get]
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