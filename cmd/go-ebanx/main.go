package main

import (
	"go-ebanx/internal/infra/database"
	"go-ebanx/internal/infra/web"
	"go-ebanx/internal/usecase"

	_ "go-ebanx/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Go Ebanx API
// @version 1.0
// @description API para gerenciar operações bancárias como depósitos, saques, transferências e consulta de saldo.
// @host localhost:8080
// @BasePath /api

func main() {
	accountRepo := database.NewInMemoryAccountRepository()

	getBalanceUseCase := usecase.NewGetBalanceUseCase(accountRepo)
	depositUseCase := usecase.NewDepositUseCase(accountRepo)
	withdrawUseCase := usecase.NewWithdrawUseCase(accountRepo)
	transferUseCase := usecase.NewTransferUseCase(accountRepo)

	handler := web.NewHandler(getBalanceUseCase, depositUseCase, withdrawUseCase, transferUseCase)

	r := gin.Default()

	r.GET("/api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := r.Group("/api")
	{
		api.POST("/reset", func(c *gin.Context) {
			handler.Reset(c.Writer, c.Request)
		})

		api.GET("/balance", func(c *gin.Context) {
			handler.GetBalance(c.Writer, c.Request)
		})

		api.POST("/event", func(c *gin.Context) {
			handler.Event(c.Writer, c.Request)
		})
	}
	r.Run(":8080")
}
