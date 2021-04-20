package heads

import (
	"errors"

	"github.com/deepvalue-network/software/pangolin/domain/middle/heads"
)

type valueBuilder struct {
	name    string
	version string
	imports []heads.External
}

func createValueBuilder() ValueBuilder {
	out := valueBuilder{
		name:    "",
		version: "",
		imports: nil,
	}

	return &out
}

// Create initializes the builder
func (app *valueBuilder) Create() ValueBuilder {
	return createValueBuilder()
}

// WithName adds a name to the builder
func (app *valueBuilder) WithName(name string) ValueBuilder {
	app.name = name
	return app
}

// WithVersion adds a version to the builder
func (app *valueBuilder) WithVersion(version string) ValueBuilder {
	app.version = version
	return app
}

// WithImports add imports to the builder
func (app *valueBuilder) WithImports(imports []heads.External) ValueBuilder {
	app.imports = imports
	return app
}

// Now builds a new Value instance
func (app *valueBuilder) Now() (Value, error) {
	if app.name == "" {
		return createValueWithName(app.name), nil
	}

	if app.version == "" {
		return createValueWithVersion(app.version), nil
	}

	if app.imports != nil && len(app.imports) <= 0 {
		app.imports = nil
	}

	if app.imports == nil {
		return createValueWithImports(app.imports), nil
	}

	return nil, errors.New("the Value is invalid")
}
