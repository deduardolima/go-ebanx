package usecase_test

import (
	"go-ebanx/internal/entity/account"
	"go-ebanx/internal/infra/database"
	"go-ebanx/internal/usecase"
	"testing"
)

func TestDepositUseCase_Execute(t *testing.T) {
	repo := database.NewMockAccountRepository()
	depositUseCase := usecase.NewDepositUseCase(repo)

	// Teste de depósito para uma nova conta
	account, err := depositUseCase.Execute("123", 100)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if account.ID != "123" || account.Balance != 100 {
		t.Errorf("expected account ID '123' and balance 100, got ID '%s' and balance %d", account.ID, account.Balance)
	}

	// Teste de depósito em uma conta existente
	account, err = depositUseCase.Execute("123", 50)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if account.Balance != 150 {
		t.Errorf("expected balance to be 150, got %d", account.Balance)
	}

	// Teste de depósito com valor inválido
	_, err = depositUseCase.Execute("123", -50)
	if err == nil || err.Error() != "deposit amount must be positive" {
		t.Errorf("expected error 'deposit amount must be positive', got %v", err)
	}
}

func TestGetBalanceUseCase_Execute(t *testing.T) {
	repo := database.NewMockAccountRepository()
	balanceUseCase := usecase.NewGetBalanceUseCase(repo)

	// Configura uma conta no repositório
	repo.Save(&account.Account{ID: "123", Balance: 100})

	// Teste para obter o saldo de uma conta existente
	balance, err := balanceUseCase.Execute("123")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if balance != 100 {
		t.Errorf("expected balance to be 100, got %d", balance)
	}

	// Teste para conta não existente
	_, err = balanceUseCase.Execute("999")
	if err == nil || err.Error() != account.ErrAccountNotFound.Error() {
		t.Errorf("expected error 'account not found', got %v", err)
	}
}

func TestTransferUseCase_Execute(t *testing.T) {
	repo := database.NewMockAccountRepository()
	transferUseCase := usecase.NewTransferUseCase(repo)

	// Configura duas contas no repositório
	repo.Save(&account.Account{ID: "100", Balance: 200})
	repo.Save(&account.Account{ID: "200", Balance: 50})

	// Teste de transferência válida
	originAcc, destAcc, err := transferUseCase.Execute("100", "200", 100)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if originAcc.Balance != 100 || destAcc.Balance != 150 {
		t.Errorf("expected origin balance 100 and destination balance 150, got origin %d and destination %d", originAcc.Balance, destAcc.Balance)
	}

	// Teste de transferência com saldo insuficiente
	_, _, err = transferUseCase.Execute("100", "200", 200)
	if err == nil || err.Error() != "insufficient funds" {
		t.Errorf("expected error 'insufficient funds', got %v", err)
	}
}

func TestWithdrawUseCase_Execute(t *testing.T) {
	repo := database.NewMockAccountRepository()
	withdrawUseCase := usecase.NewWithdrawUseCase(repo)

	// Configura uma conta no repositório
	repo.Save(&account.Account{ID: "123", Balance: 100})

	// Teste de saque válido
	acc, err := withdrawUseCase.Execute("123", 50)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if acc.Balance != 50 {
		t.Errorf("expected balance to be 50, got %d", acc.Balance)
	}

	// Teste de saque com saldo insuficiente
	_, err = withdrawUseCase.Execute("123", 200)
	if err == nil || err.Error() != "insufficient funds" {
		t.Errorf("expected error 'insufficient funds', got %v", err)
	}
}
