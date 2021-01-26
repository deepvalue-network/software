package parsers

import "errors"

type importSingleBuilder struct {
	name string
	path RelativePath
}

func createImportSingleBuilder() ImportSingleBuilder {
	out := importSingleBuilder{
		name: "",
		path: nil,
	}

	return &out
}

// Create initializes the builder
func (app *importSingleBuilder) Create() ImportSingleBuilder {
	return createImportSingleBuilder()
}

// WithName adds a name to the builder
func (app *importSingleBuilder) WithName(name string) ImportSingleBuilder {
	app.name = name
	return app
}

// WithPath adds a path to the builder
func (app *importSingleBuilder) WithPath(path RelativePath) ImportSingleBuilder {
	app.path = path
	return app
}

// Now builds a new ImportSingle instance
func (app *importSingleBuilder) Now() (ImportSingle, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build an ImportSingle instance")
	}

	if app.path == nil {
		return nil, errors.New("the RelativePath is mandatory in order to build an ImportSingle instance")
	}

	return createImportSingle(app.name, app.path), nil
}
