package saves

import (
	"errors"

	"github.com/steve-care-software/products/bobby/domain/selectors"
	"github.com/steve-care-software/products/bobby/domain/transactions/bodies/containers/tables"
	"github.com/steve-care-software/products/libs/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	metaData    tables.Transaction
	parent      selectors.Selector
	graphbase   selectors.Selector
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		metaData:    nil,
		parent:      nil,
		graphbase:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter)
}

// WithMetaData adds a metaData transaction to the builder
func (app *builder) WithMetaData(metaData tables.Transaction) Builder {
	app.metaData = metaData
	return app
}

// WithGraphbase adds a graphbase to the builder
func (app *builder) WithGraphbase(graphbase selectors.Selector) Builder {
	app.graphbase = graphbase
	return app
}

// WithParent adds a parent to the builder
func (app *builder) WithParent(parent selectors.Selector) Builder {
	app.parent = parent
	return app
}

// Now builds a new Transaction instance
func (app *builder) Now() (Transaction, error) {
	if app.metaData == nil {
		return nil, errors.New("the metaData is mandatory in order to build a Transaction instance")
	}

	data := [][]byte{
		app.metaData.Hash().Bytes(),
	}

	if app.graphbase != nil {
		data = append(data, app.graphbase.Hash().Bytes())
	}

	if app.parent != nil {
		data = append(data, app.parent.Hash().Bytes())
	}

	hsh, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.graphbase != nil && app.parent != nil {
		return createTransactionWithParentAndGraphbase(*hsh, app.metaData, app.parent, app.graphbase), nil
	}

	if app.graphbase != nil {
		return createTransactionWithGraphbase(*hsh, app.metaData, app.graphbase), nil
	}

	if app.parent != nil {
		return createTransactionWithParent(*hsh, app.metaData, app.parent), nil
	}

	return createTransaction(*hsh, app.metaData), nil
}
