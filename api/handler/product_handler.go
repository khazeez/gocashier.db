package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gocashier.db/internal/models"
	"gocashier.db/internal/services"
)

type productHandler struct {
	productService services.ProductService
}

func NewProductHandler(productService services.ProductService) *productHandler {
	return &productHandler{
		productService: productService,
	}

}

func (c *productHandler) Create(h *gin.Context) {
	var product models.Product
	err := h.ShouldBindJSON(&product)
	if err != nil {
		h.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request payload",
		})
		return
	}

	err = c.productService.Create(&product)
	if err != nil {
		h.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	h.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": product,
	})


}

func (c *productHandler) GetAll(h *gin.Context) {
	
	
}

func (c *productHandler) UpdateById(h *gin.Context) {
	
}

func (c *productHandler) DeleteById(h *gin.Context) {
	
}

func (c *productHandler) GetById(h *gin.Context) {
	
}

