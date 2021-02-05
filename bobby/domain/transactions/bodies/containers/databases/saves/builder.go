package saves

import (
	"errors"

	"github.com/deepvalue-network/software/bobby/domain/selectors"
	"github.com/deepvalue-network/software/libs/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	graphbase   selectors.Selector
	name        string
	database    selectors.Selector
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		graphbase:   nil,
		name:        "",
		database:    nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter)
}

// WithGraphbase adds a graphbase to the builder
func (app *builder) WithGraphbase(graphbase selectors.Selector) Builder {
	app.graphbase = graphbase
	return app
}

// WithName adds a name to the builder
func (app *builder) WithName(name string) Builder {
	app.name = name
	return app
}

// WithDatabase adds a database to the builder
func (app *builder) WithDatabase(database selectors.Selector) Builder {
	app.database = database
	return app
}

// Now builds a new Transaction instance
func (app *builder) Now() (Transaction, error) {
	if app.graphbase == nil {
		return nil, errors.New("the graphbase is mandatory in order to build a Transaction instance")
	}

	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Transaction instance")
	}

	data := [][]byte{
		app.graphbase.Hash().Bytes(),
		[]byte(app.name),
	}

	if app.database != nil {
		data = append(data, app.database.Hash().Bytes())
	}

	hsh, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.database != nil {
		return createTransactionWithDatabase(*hsh, app.graphbase, app.name, app.database), nil
	}

	return createTransaction(*hsh, app.graphbase, app.name), nil
}
