package externals

import (
	"errors"
)

type builder struct {
	name string
	path string
}

func createBuilder() Builder {
	out := builder{
		name: "",
		path: "",
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithName adds a name to the builder
func (app *builder) WithName(name string) Builder {
	app.name = name
	return app
}

// WithPath adds a path to the builder
func (app *builder) WithPath(path string) Builder {
	app.path = path
	return app
}

// Now builds a new Import instance
func (app *builder) Now() (External, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build an Import instance")
	}

	if app.path == "" {
		return nil, errors.New("the path is mandatory in order to build an Import instance")
	}

	return createExternal(app.name, app.path), nil
}
