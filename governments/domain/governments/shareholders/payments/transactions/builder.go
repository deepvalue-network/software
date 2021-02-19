package transactions

import (
	"errors"

	"github.com/deepvalue-network/software/governments/domain/governments/shareholders/payments"
	"github.com/deepvalue-network/software/libs/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	payment     payments.Payment
	note        string
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		payment:     nil,
		note:        "",
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter)
}

// WithPayment adds a payment to the builder
func (app *builder) WithPayment(payment payments.Payment) Builder {
	app.payment = payment
	return app
}

// WithNote adds a note to the builder
func (app *builder) WithNote(note string) Builder {
	app.note = note
	return app
}

// Now builds a new Transaction instance
func (app *builder) Now() (Transaction, error) {
	if app.payment == nil {
		return nil, errors.New("the payment is mandatory in order to build a Transaction instance")
	}

	hash, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.payment.Hash().Bytes(),
		[]byte(app.note),
	})

	if err != nil {
		return nil, err
	}

	return createTransaction(*hash, app.payment, app.note), nil
}
