package shareholders

import (
	"errors"

	"github.com/deepvalue-network/software/libs/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	list        []ShareHolder
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

// WithShareHolders add shareholders to the builder
func (app *builder) WithShareHolders(shareHolders []ShareHolder) Builder {
	app.list = shareHolders
	return app
}

// Now builds a new ShareHolders instance
func (app *builder) Now() (ShareHolders, error) {
	if app.list == nil {
		app.list = []ShareHolder{}
	}

	if len(app.list) <= 1 {
		return nil, errors.New("thee must be at least 1 ShareHolder instance in order to build a ShareHolders instance")
	}

	data := [][]byte{}
	for _, oneHolder := range app.list {
		data = append(data, oneHolder.Hash().Bytes())
	}

	hash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	return createShareHolders(*hash, app.list), nil
}
