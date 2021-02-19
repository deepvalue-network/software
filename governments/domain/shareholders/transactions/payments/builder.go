package payments

import (
	"errors"

	"github.com/deepvalue-network/software/governments/domain/governments/shareholders/payments"
	"github.com/deepvalue-network/software/libs/hash"
)

type paymentBuilder struct {
	hashAdapter hash.Adapter
	payment     payments.Payment
	note        string
}

func createPaymentBuilder(
	hashAdapter hash.Adapter,
) PaymentBuilder {
	out := paymentBuilder{
		hashAdapter: hashAdapter,
		payment:     nil,
		note:        "",
	}

	return &out
}

// Create initializes the paymentBuilder
func (app *paymentBuilder) Create() PaymentBuilder {
	return createPaymentBuilder(app.hashAdapter)
}

// WithPayment adds a payment to the paymentBuilder
func (app *paymentBuilder) WithPayment(payment payments.Payment) PaymentBuilder {
	app.payment = payment
	return app
}

// WithNote adds a note to the paymentBuilder
func (app *paymentBuilder) WithNote(note string) PaymentBuilder {
	app.note = note
	return app
}

// Now builds a new Payment instance
func (app *paymentBuilder) Now() (Payment, error) {
	if app.payment == nil {
		return nil, errors.New("the payment is mandatory in order to build a Payment instance")
	}

	hash, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.payment.Hash().Bytes(),
		[]byte(app.note),
	})

	if err != nil {
		return nil, err
	}

	return createPayment(*hash, app.payment, app.note), nil
}
