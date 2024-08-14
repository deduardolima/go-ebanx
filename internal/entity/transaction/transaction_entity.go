package transaction

import "errors"

type Transaction struct {
	ID            string
	OriginID      string
	DestinationID string
	Amount        int
}

func NewTransaction(id, originID, destinationID string, amount int) (*Transaction, error) {
	if id == "" {
		return nil, errors.New("invalid transaction id")
	}
	if amount <= 0 {
		return nil, errors.New("transaction amount must be positive")
	}
	return &Transaction{
		ID:            id,
		OriginID:      originID,
		DestinationID: destinationID,
		Amount:        amount,
	}, nil
}
