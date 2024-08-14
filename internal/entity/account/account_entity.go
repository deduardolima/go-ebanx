package account

import "errors"

var ErrAccountNotFound = errors.New("account not found")

type Account struct {
	ID      string
	Balance int
}

func NewAccount(id string, initialBalance int) (*Account, error) {
	if id == "" {
		return nil, errors.New("invalid account id")
	}
	if initialBalance < 0 {
		return nil, errors.New("initial balance cannot be negative")
	}
	return &Account{
		ID:      id,
		Balance: initialBalance,
	}, nil
}

func (a *Account) Deposit(amount int) error {
	if amount <= 0 {
		return errors.New("deposit amount must be positive")
	}
	a.Balance += amount
	return nil
}

func (a *Account) Withdraw(amount int) error {
	if amount <= 0 {
		return errors.New("withdraw amount must be positive")
	}
	if a.Balance < amount {
		return errors.New("insufficient funds")
	}
	a.Balance -= amount
	return nil
}
