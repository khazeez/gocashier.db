package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gocashier.db/internal/models"
	"gocashier.db/internal/services"
)

type categoryHandler struct {
	categoryService services.CategoryService
}

func NewCategoryHandler(categoryService services.CategoryService) *categoryHandler {
	return &categoryHandler{
		categoryService: categoryService,
	}

}


// Create godoc
// @Summary Create new category
// @Tags categories
// @Accept json
// @Produce json
// @Param category body models.Category true "Category Data"
// @Success 201 {object} models.Category
// @Router /categories [post]
func (c *categoryHandler) Create(h *gin.Context) {
	var category models.Category

	// Bind JSON ke struct
	if err := h.ShouldBindJSON(&category); err != nil {
		h.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request payload",
		})
		return
	}

	// Simpan ke database
	if err := c.categoryService.Create(&category); err != nil {
		h.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to create category",
		})
		return
	}

	// Kembalikan response data yang sudah tersimpan (termasuk ID)
	h.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    category,
	})
}



// GetAll godoc
// @Summary Get all categories
// @Tags categories
// @Produce json
// @Success 200 {array} models.Category
// @Router /categories [get]
func (c *categoryHandler) GetAll(h *gin.Context) {
	data, err := c.categoryService.GetAll()
	if err != nil {
		h.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	h.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    data,
	})
}

// UpdateById godoc
// @Summary Update category by ID
// @Tags categories
// @Accept json
// @Produce json
// @Param ID path int true "Category ID"
// @Param category body models.Category true "Category Data"
// @Success 200 {object} models.Category
// @Router /categories/{ID} [put]
func (c *categoryHandler) UpdateById(h *gin.Context) {
	id, err := strconv.Atoi(h.Param("ID"))
	if err != nil {
		h.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	var category models.Category
	if err := h.ShouldBindJSON(&category); err != nil {
		h.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request payload",
		})
		return

	}

	if err := c.categoryService.UpdateById(id, &category); err != nil {
		h.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	h.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    category,
	})

}

// DeleteById godoc
// @Summary Delete category by ID
// @Tags categories
// @Produce json
// @Param ID path int true "Category ID"
// @Success 200 {string} string "Success delete data"
// @Router /categories/{ID} [delete]
func (c *categoryHandler) DeleteById(h *gin.Context) {
	id, err := strconv.Atoi(h.Param("ID"))
	if err != nil {
		h.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	err = c.categoryService.DeleteById(id)
	if err != nil {
		h.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Succsess delete data",
	})

}


// GetById godoc
// @Summary Get category by ID
// @Tags categories
// @Produce json
// @Param ID path int true "Category ID"
// @Success 200 {object} models.Category
// @Router /categories/{ID} [get]
func (c *categoryHandler) GetById(h *gin.Context) {
	id, err := strconv.Atoi(h.Param("ID"))
	if err != nil {
		h.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	data, err := c.categoryService.GetById(id)
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
