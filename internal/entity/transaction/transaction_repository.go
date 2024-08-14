package transaction

type TransactionRepositoryInterface interface {
	Save(transaction *Transaction) error
	FindAll() ([]*Transaction, error)
}
