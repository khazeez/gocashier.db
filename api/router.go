package api

import (
	"database/sql"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"

	_ "gocashier.db/docs"

	"gocashier.db/api/handler"
	"gocashier.db/api/middleware" // tambahkan import ini
	"gocashier.db/internal/repository"
	"gocashier.db/internal/services"
)

func Router(db *sql.DB) *gin.Engine {

	// Category
	categoryRepo := repository.NewcategoryRepository(db)
	categoryService := services.NewcategoryService(categoryRepo)
	categoryHandler := handler.NewCategoryHandler(categoryService)

	// Product
	productRepo := repository.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	productHandler := handler.NewProductHandler(productService)

	// Transaction
	transactionRepo := repository.NewTransactionRepository(db)
	transactionService := services.NewTransactionService(transactionRepo)
	transactionHandler := handler.NewTransactionHandler(transactionService)

	r := gin.Default()

	// Daftarkan middleware logger (sebelum routes)
	r.Use(middleware.LoggerMiddleware())

	// Swagger endpoint
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiRouter := r.Group("/api")
	{
		categoryRouter := apiRouter.Group("/categories")
		{
			categoryRouter.POST("/", categoryHandler.Create)
			categoryRouter.GET("/", categoryHandler.GetAll)
			categoryRouter.GET("/:ID", categoryHandler.GetById)
			categoryRouter.PUT("/:ID", categoryHandler.UpdateById)
			categoryRouter.DELETE("/:ID", categoryHandler.DeleteById)
		}

		productRouter := apiRouter.Group("/products")
		{
			productRouter.POST("/", productHandler.Create)
			productRouter.GET("/", productHandler.GetAll)
			productRouter.GET("/:ID", productHandler.GetById)
			productRouter.PUT("/:ID", productHandler.UpdateById)
			productRouter.DELETE("/:ID", productHandler.DeleteById)
			productRouter.GET("/:ID/detail", productHandler.GetDetailProductById)
		}

		transactionRouter := apiRouter.Group("/transaction")
		{
			transactionRouter.POST("/checkout", transactionHandler.CreateTransaction)
			transactionRouter.GET("/report/today", transactionHandler.GetReportToday)
			transactionRouter.GET("/report", transactionHandler.GetReportWithRange)
		}
	}

	return r
}