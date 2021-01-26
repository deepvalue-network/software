package tables

import (
	"errors"

	"github.com/steve-care-software/products/bobby/domain/transactions/bodies/containers/tables/deletes"
	"github.com/steve-care-software/products/bobby/domain/transactions/bodies/containers/tables/saves"
)

type builder struct {
	del  deletes.Transaction
	save saves.Transaction
}

func createBuilder() Builder {
	out := builder{
		del:  nil,
		save: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithDelete adds a delete transaction to the builder
func (app *builder) WithDelete(del deletes.Transaction) Builder {
	app.del = del
	return app
}

// WithSave adds a save transaction to the builder
func (app *builder) WithSave(save saves.Transaction) Builder {
	app.save = save
	return app
}

// Now builds a new Transaction instance
func (app *builder) Now() (Transaction, error) {
	if app.del != nil {
		return createTransactionWithDelete(app.del), nil
	}

	if app.save != nil {
		return createTransactionWithSave(app.save), nil
	}

	return nil, errors.New("the database Transaction is invalid")
}
