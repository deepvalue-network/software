package rows

import (
	"github.com/deepvalue-network/software/libs/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	list        []Row
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		list:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter)
}

// WithRows add rows to the builder
func (app *builder) WithRows(rows []Row) Builder {
	app.list = rows
	return app
}

// Now builds a new Rows instance
func (app *builder) Now() (Rows, error) {
	if app.list == nil {
		app.list = []Row{}
	}

	data := [][]byte{
		[]byte("initial"),
	}

	for _, oneRow := range app.list {
		data = append(data, oneRow.Resource().Hash().Bytes())
	}

	hash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	return createRows(*hash, app.list), nil
}
