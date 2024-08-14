package usecase

import "go-ebanx/internal/entity/account"

type DepositUseCase struct {
	AccountRepo account.AccountRepositoryInterface
}

func NewDepositUseCase(repo account.AccountRepositoryInterface) *DepositUseCase {
	return &DepositUseCase{AccountRepo: repo}
}

func (uc *DepositUseCase) Execute(destinationID string, amount int) (*account.Account, error) {
	acc, err := uc.AccountRepo.FindByID(destinationID)
	if err != nil {
		acc, err = account.NewAccount(destinationID, amount)
		if err != nil {
			return nil, err
		}
	} else {
		err = acc.Deposit(amount)
		if err != nil {
			return nil, err
		}
	}

	err = uc.AccountRepo.Save(acc)
	if err != nil {
		return nil, err
	}

	return acc, nil
}
