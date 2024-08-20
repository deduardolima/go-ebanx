package transaction

import (
	"testing"
)

func TestNewTransaction(t *testing.T) {
	transaction, err := NewTransaction("1", "123", "456", 100)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if transaction.ID != "1" {
		t.Errorf("expected ID to be '1', got %v", transaction.ID)
	}
	if transaction.Amount != 100 {
		t.Errorf("expected amount to be 100, got %v", transaction.Amount)
	}

	// Testando a criação de uma transação com ID vazio
	_, err = NewTransaction("", "123", "456", 100)
	if err == nil || err.Error() != "invalid transaction id" {
		t.Errorf("expected error 'invalid transaction id', got %v", err)
	}

	// Testando a criação de uma transação com valor negativo
	_, err = NewTransaction("1", "123", "456", -100)
	if err == nil || err.Error() != "transaction amount must be positive" {
		t.Errorf("expected error 'transaction amount must be positive', got %v", err)
	}
}
