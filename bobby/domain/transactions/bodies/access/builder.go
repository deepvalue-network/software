package access

import (
	"errors"

	access "github.com/steve-care-software/products/bobby/domain/resources"
	"github.com/steve-care-software/products/bobby/domain/selectors"
	"github.com/steve-care-software/products/libs/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	resources   selectors.Selector
	add         access.Access
	isRemove    bool
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		resources:   nil,
		add:         nil,
		isRemove:    false,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter)
}

// WithResources add resources to the builder
func (app *builder) WithResources(resources selectors.Selector) Builder {
	app.resources = resources
	return app
}

// Add adds an access to the builder
func (app *builder) Add(add access.Access) Builder {
	app.add = add
	return app
}

// Remove flags the builder as remove
func (app *builder) Remove() Builder {
	app.isRemove = true
	return app
}

// Now builds a new Transaction instance
func (app *builder) Now() (Transaction, error) {
	if app.resources == nil {
		return nil, errors.New("the resources selector is mandatory in order to build an access Transaction instance")
	}

	data := [][]byte{
		app.resources.Hash().Bytes(),
	}

	var content Content
	if app.add != nil {
		data = append(data, app.add.Resource().Hash().Bytes())
		content = createContentWithAdd(app.add)
	}

	if app.isRemove {
		data = append(data, []byte("is_removed"))
		content = createContentWithRemove()
	}

	if content == nil {
		return nil, errors.New("the content (added acces / remove access flag) is mandatory in order to build an access Transaction instance")
	}

	hsh, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	return createTransaction(*hsh, app.resources, content), nil
}
