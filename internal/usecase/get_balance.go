package usecase

import "go-ebanx/internal/entity/account"

type GetBalanceUseCase struct {
	AccountRepo account.AccountRepositoryInterface
}

func NewGetBalanceUseCase(repo account.AccountRepositoryInterface) *GetBalanceUseCase {
	return &GetBalanceUseCase{AccountRepo: repo}
}

func (uc *GetBalanceUseCase) Execute(accountID string) (int, error) {
	account, err := uc.AccountRepo.FindByID(accountID)
	if err != nil {
		return 0, err
	}
	return account.Balance, nil
}
