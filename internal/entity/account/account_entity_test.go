package account

import (
	"testing"
)

func TestNewAccount(t *testing.T) {
	// Testando a criação de uma conta válida
	account, err := NewAccount("123", 100)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if account.ID != "123" {
		t.Errorf("expected ID to be '123', got %v", account.ID)
	}
	if account.Balance != 100 {
		t.Errorf("expected balance to be 100, got %v", account.Balance)
	}

	_, err = NewAccount("", 100)
	if err == nil || err.Error() != "invalid account id" {
		t.Errorf("expected error 'invalid account id', got %v", err)
	}

	_, err = NewAccount("123", -100)
	if err == nil || err.Error() != "initial balance cannot be negative" {
		t.Errorf("expected error 'initial balance cannot be negative', got %v", err)
	}
}

func TestDeposit(t *testing.T) {
	account, _ := NewAccount("123", 100)

	err := account.Deposit(50)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if account.Balance != 150 {
		t.Errorf("expected balance to be 150, got %v", account.Balance)
	}

	err = account.Deposit(-50)
	if err == nil || err.Error() != "deposit amount must be positive" {
		t.Errorf("expected error 'deposit amount must be positive', got %v", err)
	}
}

func TestWithdraw(t *testing.T) {
	account, _ := NewAccount("123", 100)

	err := account.Withdraw(50)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if account.Balance != 50 {
		t.Errorf("expected balance to be 50, got %v", account.Balance)
	}

	err = account.Withdraw(-50)
	if err == nil || err.Error() != "withdraw amount must be positive" {
		t.Errorf("expected error 'withdraw amount must be positive', got %v", err)
	}

	err = account.Withdraw(100)
	if err == nil || err.Error() != "insufficient funds" {
		t.Errorf("expected error 'insufficient funds', got %v", err)
	}
}
