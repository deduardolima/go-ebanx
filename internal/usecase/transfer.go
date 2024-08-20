package usecase

import "go-ebanx/internal/entity/account"

type TransferUseCase struct {
	AccountRepo account.AccountRepositoryInterface
}

func NewTransferUseCase(repo account.AccountRepositoryInterface) *TransferUseCase {
	return &TransferUseCase{AccountRepo: repo}
}

func (uc *TransferUseCase) Execute(originID, destinationID string, amount int) (*account.Account, *account.Account, error) {
	originAccount, err := uc.AccountRepo.FindByID(originID)
	if err != nil {
		return nil, nil, err
	}

	destinationAccount, err := uc.AccountRepo.FindByID(destinationID)
	if err != nil {
		destinationAccount, err = account.NewAccount(destinationID, 0)
		if err != nil {
			return nil, nil, err
		}
	}

	err = originAccount.Withdraw(amount)
	if err != nil {
		return nil, nil, err
	}

	err = destinationAccount.Deposit(amount)
	if err != nil {
		return nil, nil, err
	}

	err = uc.AccountRepo.Save(originAccount)
	if err != nil {
		return nil, nil, err
	}

	err = uc.AccountRepo.Save(destinationAccount)
	if err != nil {
		return nil, nil, err
	}

	return originAccount, destinationAccount, nil
}
