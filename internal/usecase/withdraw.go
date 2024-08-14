package usecase

import "go-ebanx/internal/entity/account"

type WithdrawUseCase struct {
	AccountRepo account.AccountRepositoryInterface
}

func NewWithdrawUseCase(repo account.AccountRepositoryInterface) *WithdrawUseCase {
	return &WithdrawUseCase{AccountRepo: repo}
}

func (uc *WithdrawUseCase) Execute(originID string, amount int) (*account.Account, error) {
	acc, err := uc.AccountRepo.FindByID(originID)
	if err != nil {
		return nil, err
	}

	err = acc.Withdraw(amount)
	if err != nil {
		return nil, err
	}

	err = uc.AccountRepo.Save(acc)
	if err != nil {
		return nil, err
	}

	return acc, nil
}
