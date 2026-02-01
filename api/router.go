package api

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"gocashier.db/api/handler"
	"gocashier.db/internal/repository"
	"gocashier.db/internal/services"
)


func Router(db *sql.DB) *gin.Engine{
	categoryRepo := repository.NewcategoryRepository(db)
	categoryService:= services.NewcategoryService(categoryRepo)
	categoryHandler := handler.NewCategoryHandler(categoryService)

	r := gin.Default()

	categoryRouter := r.Group("/categories")
	{
		categoryRouter.POST("/", categoryHandler.Create)
		categoryRouter.GET("/", categoryHandler.GetAll)
		categoryRouter.PUT("/:ID", categoryHandler.UpdateById)
		categoryRouter.GET("/:ID", categoryHandler.GetById)
		categoryRouter.DELETE("/:ID", categoryHandler.DeleteById)
	}

	return r

}
