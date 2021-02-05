package owners

import (
	"errors"

	"github.com/deepvalue-network/software/diamonds/domain/bills"
	"github.com/deepvalue-network/software/diamonds/domain/owners"
	"github.com/deepvalue-network/software/libs/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	owner       owners.Owner
	bill        bills.ViewBill
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		owner:       nil,
		bill:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter)
}

// WithOwner adds an owner to the builder
func (app *builder) WithOwner(owner owners.Owner) Builder {
	app.owner = owner
	return app
}

// WithViewBill adds a view bill to the builder
func (app *builder) WithViewBill(viewBill bills.ViewBill) Builder {
	app.bill = viewBill
	return app
}

// Now builds a new Bill instance
func (app *builder) Now() (Bill, error) {
	if app.owner == nil {
		return nil, errors.New("the owner is mandatory in order to build an owned Bill instance")
	}

	if app.bill == nil {
		return nil, errors.New("the view bill is mandatory in order to build an owned Bill instance")
	}

	hash, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.owner.Hash().Bytes(),
		app.bill.Hash().Bytes(),
	})

	if err != nil {
		return nil, err
	}

	return createBill(*hash, app.owner, app.bill), nil
}
