package account

type AccountRepositoryInterface interface {
	Save(account *Account) error
	FindByID(id string) (*Account, error)
	Reset()
}
