package heads

import (
	"errors"

	"github.com/deepvalue-network/software/libs/hash"
)

type externalBuilder struct {
	hashAdapter hash.Adapter
	name        string
	path        string
}

func createExternalBuilder(
	hashAdapter hash.Adapter,
) ExternalBuilder {
	out := externalBuilder{
		hashAdapter: hashAdapter,
		name:        "",
		path:        "",
	}

	return &out
}

// Create initializes the builder
func (app *externalBuilder) Create() ExternalBuilder {
	return createExternalBuilder(app.hashAdapter)
}

// WithName adds a name to the builder
func (app *externalBuilder) WithName(name string) ExternalBuilder {
	app.name = name
	return app
}

// WithPath adds a path to the builder
func (app *externalBuilder) WithPath(path string) ExternalBuilder {
	app.path = path
	return app
}

// Now builds a new Import instance
func (app *externalBuilder) Now() (External, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build an Import instance")
	}

	if app.path == "" {
		return nil, errors.New("the path is mandatory in order to build an Import instance")
	}

	return createExternal(app.name, app.path), nil
}
