package heads

import (
	"errors"

	"github.com/deepvalue-network/software/pangolin/domain/parsers"
)

type builder struct {
	name    string
	version string
	imports []parsers.ImportSingle
}

func createBuilder() Builder {
	out := builder{
		name:    "",
		version: "",
		imports: nil,
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

// WithVersion adds a version to the builder
func (app *builder) WithVersion(version string) Builder {
	app.version = version
	return app
}

// WithImports add imports to the builder
func (app *builder) WithImports(imports []parsers.ImportSingle) Builder {
	app.imports = imports
	return app
}

// Now builds a new Head instance
func (app *builder) Now() (Head, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Head instance")
	}

	if app.version == "" {
		return nil, errors.New("the version is mandatory in order to build a Head instance")
	}

	if app.imports != nil && len(app.imports) <= 0 {
		app.imports = nil
	}

	if app.imports != nil {
		return createHeadWithImports(app.name, app.version, app.imports), nil
	}

	return createHead(app.name, app.version), nil
}
