package transactions

import (
	"errors"
	"fmt"

	"github.com/deepvalue-network/software/bobby/domain/transactions/bodies"
	"github.com/deepvalue-network/software/libs/cryptography/pk/signature"
	"github.com/deepvalue-network/software/libs/hash"
)

type transactionBuilder struct {
	hashAdapter hash.Adapter
	body        bodies.Body
	sig         signature.RingSignature
}

func createTransactionBuilder(
	hashAdapter hash.Adapter,
) TransactionBuilder {
	out := transactionBuilder{
		hashAdapter: hashAdapter,
		body:        nil,
		sig:         nil,
	}

	return &out
}

// Create initializes the builder
func (app *transactionBuilder) Create() TransactionBuilder {
	return createTransactionBuilder(app.hashAdapter)
}

// WithBody adds a body to the builder
func (app *transactionBuilder) WithBody(body bodies.Body) TransactionBuilder {
	app.body = body
	return app
}

// WithSignature adds a signature to the builder
func (app *transactionBuilder) WithSignature(sig signature.RingSignature) TransactionBuilder {
	app.sig = sig
	return app
}

// Now builds a new Transaction instance
func (app *transactionBuilder) Now() (Transaction, error) {
	if app.body == nil {
		return nil, errors.New("the body is mandatory in order to build a Transaction instance")
	}

	if app.sig == nil {
		return nil, errors.New("the signature is mandatory in order to build a Transaction instance")
	}

	if !app.sig.Verify(app.body.Resource().Hash().String()) {
		str := fmt.Sprintf("the given signature cannot be verified against the transaction's body hash: %s", app.body.Resource().Hash().String())
		return nil, errors.New(str)
	}

	hsh, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.body.Resource().Hash().Bytes(),
		[]byte(app.sig.String()),
	})

	if err != nil {
		return nil, err
	}

	return createTransaction(*hsh, app.body, app.sig), nil
}
