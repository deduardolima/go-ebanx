package database

import (
	"go-ebanx/internal/entity/account"
)

type MockAccountRepository struct {
	Accounts map[string]*account.Account
}

func NewMockAccountRepository() *MockAccountRepository {
	return &MockAccountRepository{
		Accounts: make(map[string]*account.Account),
	}
}

func (r *MockAccountRepository) FindByID(id string) (*account.Account, error) {
	if acc, exists := r.Accounts[id]; exists {
		return acc, nil
	}
	return nil, account.ErrAccountNotFound
}

func (r *MockAccountRepository) Save(acc *account.Account) error {
	r.Accounts[acc.ID] = acc
	return nil
}

func (r *MockAccountRepository) Reset() {
	r.Accounts = make(map[string]*account.Account)
}
