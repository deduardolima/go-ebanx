package main

import (
	"go-ebanx/internal/infra/database"
	"go-ebanx/internal/infra/web"
	"go-ebanx/internal/infra/web/webserver"
	"go-ebanx/internal/usecase"
)

func main() {
	accountRepo := database.NewInMemoryAccountRepository()

	getBalanceUseCase := usecase.NewGetBalanceUseCase(accountRepo)
	depositUseCase := usecase.NewDepositUseCase(accountRepo)
	withdrawUseCase := usecase.NewWithdrawUseCase(accountRepo)
	transferUseCase := usecase.NewTransferUseCase(accountRepo)

	handler := web.NewHandler(getBalanceUseCase, depositUseCase, withdrawUseCase, transferUseCase)

	server := webserver.NewWebServer(handler)
	server.Start()
}
