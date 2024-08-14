package usecase

import "go-ebanx/internal/entity/account"

type TransferUseCase struct {
	AccountRepo account.AccountRepositoryInterface
}

func NewTransferUseCase(repo account.AccountRepositoryInterface) *TransferUseCase {
	return &TransferUseCase{AccountRepo: repo}
}

func (uc *TransferUseCase) Execute(originID, destinationID string, amount int) (*account.Account, *account.Account, error) {
	// Busca a conta de origem pelo ID
	originAccount, err := uc.AccountRepo.FindByID(originID)
	if err != nil {
		return nil, nil, err
	}

	// Busca a conta de destino pelo ID, ou cria uma nova se não existir
	destinationAccount, err := uc.AccountRepo.FindByID(destinationID)
	if err != nil {
		destinationAccount, err = account.NewAccount(destinationID, 0)
		if err != nil {
			return nil, nil, err
		}
	}

	// Realiza o saque na conta de origem
	err = originAccount.Withdraw(amount)
	if err != nil {
		return nil, nil, err
	}

	// Realiza o depósito na conta de destino
	err = destinationAccount.Deposit(amount)
	if err != nil {
		return nil, nil, err
	}

	// Salva as alterações na conta de origem
	err = uc.AccountRepo.Save(originAccount)
	if err != nil {
		return nil, nil, err
	}

	// Salva as alterações na conta de destino
	err = uc.AccountRepo.Save(destinationAccount)
	if err != nil {
		return nil, nil, err
	}

	return originAccount, destinationAccount, nil
}
