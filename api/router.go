package api

import (
	"database/sql"
	// _ "go-swagger/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
	"gocashier.db/api/handler"
	"gocashier.db/internal/repository"
	"gocashier.db/internal/services"
)

func Router(db *sql.DB) *gin.Engine {
	//category
	categoryRepo := repository.NewcategoryRepository(db)
	categoryService := services.NewcategoryService(categoryRepo)
	categoryHandler := handler.NewCategoryHandler(categoryService)

	//Product
	productRepo := repository.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	productHandler := handler.NewProductHandler(productService)

	//Transaction
	transactionRepo := repository.NewTransactionRepository(db)
	transactionService := services.NewTransactionService(transactionRepo)
	transactionHandler := handler.NewTransactionHandler(transactionService)

	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiRouter := r.Group("/api")
	{
		categoryRouter := apiRouter.Group("/categories")
		{
			categoryRouter.POST("/", categoryHandler.Create)
			categoryRouter.GET("/", categoryHandler.GetAll)
			categoryRouter.PUT("/:ID", categoryHandler.UpdateById)
			categoryRouter.GET("/:ID", categoryHandler.GetById)
			categoryRouter.DELETE("/:ID", categoryHandler.DeleteById)
		}

		productRouter := apiRouter.Group("/products")
		{
			productRouter.POST("/", productHandler.Create)
			productRouter.GET("/", productHandler.GetAll)
			productRouter.PUT("/:ID", productHandler.UpdateById)
			productRouter.GET("/:ID", productHandler.GetById)
			productRouter.DELETE("/:ID", productHandler.DeleteById)
			productRouter.GET("/:ID/detail", productHandler.GetDetailProductById)
		}

		transactionRouter := apiRouter.Group("transaction")
		{
			transactionRouter.POST("/checkout", transactionHandler.CreateTransaction)
			transactionRouter.GET("/report/today", transactionHandler.GetReportToday)
			transactionRouter.GET("/report", transactionHandler.GetReportWithRange)
		}
	}

	return r

}
