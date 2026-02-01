package handler

import (
	"net/http"
	"strconv"

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

	data, err := c.productService.GetAll()
	if err != nil {
		h.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	h.JSON(http.StatusOK, gin.H{
		"succsess": true,
		"data": data,
	})

	
}

func (c *productHandler) UpdateById(h *gin.Context) {
	id, err := strconv.Atoi(h.Param("ID"))
	if err != nil {
		h.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var product models.Product
	if err := h.ShouldBindJSON(&product); err != nil {
		h.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request payload",
		})
		return
		
	}

	if err = c.productService.UpdateById(id, &product); err != nil {
		h.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.JSON(http.StatusOK, gin.H{
		"succsess": true,
		"data": product,
	})
}

func (c *productHandler) DeleteById(h *gin.Context) {
	id, err := strconv.Atoi(h.Param("ID"))
	if err != nil {
		h.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err = c.productService.DeleteById(id); err != nil {
		h.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	h.JSON(http.StatusOK, gin.H{
		"succsess": true,
		"message": "Success delete data",
	})
}

func (c *productHandler) GetById(h *gin.Context) {
	id, err := strconv.Atoi(h.Param("ID"))
	if err != nil {
		h.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	data, err := c.productService.GetById(id)
	if err != nil {
		h.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.JSON(http.StatusOK, gin.H{
		"succsess": true,
		"data": data,
	})
}

