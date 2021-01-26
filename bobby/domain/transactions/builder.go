package transactions

import (
	"errors"

	"github.com/steve-care-software/products/libs/hash"
)

type builder struct {
	hashAdapter  hash.Adapter
	transactions []Transaction
	isAtomic     bool
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter:  hashAdapter,
		transactions: nil,
		isAtomic:     false,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter)
}

// WithTransactions add transactions to the builder
func (app *builder) WithTransactions(list []Transaction) Builder {
	app.transactions = list
	return app
}

// IsAtomic flags the builder as atomic
func (app *builder) IsAtomic() Builder {
	app.isAtomic = true
	return app
}

// Now builds a new Transactions instance
func (app *builder) Now() (Transactions, error) {
	if app.transactions != nil && len(app.transactions) <= 0 {
		app.transactions = nil
	}

	if app.transactions == nil {
		return nil, errors.New("the []Transaction are mandatory in order to build a Transactions instance")
	}

	isAtomicStr := "false"
	if app.isAtomic {
		isAtomicStr = "true"
	}

	data := [][]byte{
		[]byte(isAtomicStr),
	}

	for _, oneTrx := range app.transactions {
		data = append(data, oneTrx.Hash().Bytes())
	}

	hsh, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.isAtomic {
		return createTransactionsWithAtomic(*hsh, app.transactions), nil
	}

	return createTransactions(*hsh, app.transactions), nil
}
