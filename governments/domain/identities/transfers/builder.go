package transfers

import (
	"errors"

	"github.com/deepvalue-network/software/governments/domain/governments/shareholders/transfers/views"
	"github.com/deepvalue-network/software/libs/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	transfer    views.Transfer
	note        string
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		transfer:    nil,
		note:        "",
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter)
}

// WithTransfer adds a transfer to the builder
func (app *builder) WithTransfer(transfer views.Transfer) Builder {
	app.transfer = transfer
	return app
}

// WithNote adds a note to the builder
func (app *builder) WithNote(note string) Builder {
	app.note = note
	return app
}

// Now builds a new Transfer instance
func (app *builder) Now() (Transfer, error) {
	if app.transfer == nil {
		return nil, errors.New("the view transfer is mandatory in order to build an identity Transfer instance")
	}

	hsh, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.transfer.Hash().Bytes(),
		[]byte(app.note),
	})

	if err != nil {
		return nil, err
	}

	return createTransfer(*hsh, app.transfer, app.note), nil
}
