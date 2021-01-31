package overviews

import (
	"errors"

	domain_errors "github.com/steve-care-software/products/bobby/domain/errors"
	"github.com/steve-care-software/products/bobby/domain/transactions"
)

type invalidTransactionBuilder struct {
	trx   transactions.Transaction
	error domain_errors.Error
}

func createInvalidTransactionBuilder() InvalidTransactionBuilder {
	out := invalidTransactionBuilder{
		trx:   nil,
		error: nil,
	}

	return &out
}

// Create initializes the builder
func (app *invalidTransactionBuilder) Create() InvalidTransactionBuilder {
	return createInvalidTransactionBuilder()
}

// WithTransaction adds a transaction to the builder
func (app *invalidTransactionBuilder) WithTransaction(trx transactions.Transaction) InvalidTransactionBuilder {
	app.trx = trx
	return app
}

// WithError adds an error to the builder
func (app *invalidTransactionBuilder) WithError(err domain_errors.Error) InvalidTransactionBuilder {
	app.error = err
	return app
}

// Now builds a new InvalidTransaction instance
func (app *invalidTransactionBuilder) Now() (InvalidTransaction, error) {
	if app.trx == nil {
		return nil, errors.New("the transaction is mandatory in order to build an InvalidTransaction instance")
	}

	if app.error == nil {
		return nil, errors.New("the error is mandatory in order to build an InvalidTransaction instance")
	}

	return createInvalidTransaction(app.trx, app.error), nil
}
