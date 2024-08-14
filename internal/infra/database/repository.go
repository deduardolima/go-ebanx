package database

import (
	"go-ebanx/internal/entity/account"
	"sync"
)

type InMemoryAccountRepository struct {
	accounts map[string]*account.Account
	mu       sync.Mutex
}

func NewInMemoryAccountRepository() *InMemoryAccountRepository {
	return &InMemoryAccountRepository{
		accounts: make(map[string]*account.Account),
	}
}

func (r *InMemoryAccountRepository) Save(acc *account.Account) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.accounts[acc.ID] = acc
	return nil
}

func (r *InMemoryAccountRepository) FindByID(id string) (*account.Account, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	acc, exists := r.accounts[id]
	if !exists {
		return nil, account.ErrAccountNotFound
	}
	return acc, nil
}

func (r *InMemoryAccountRepository) Reset() {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.accounts = make(map[string]*account.Account)
}
