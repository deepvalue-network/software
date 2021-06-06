package parsers

import "errors"

type headValueBuilder struct {
	name    string
	version string
	imports []ImportSingle
	loads   []LoadSingle
}

func createHeadValueBuilder() HeadValueBuilder {
	out := headValueBuilder{
		name:    "",
		version: "",
		imports: nil,
		loads:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *headValueBuilder) Create() HeadValueBuilder {
	return createHeadValueBuilder()
}

// WithName adds a name to the builder
func (app *headValueBuilder) WithName(name string) HeadValueBuilder {
	app.name = name
	return app
}

// WithVersion adds a version to the builder
func (app *headValueBuilder) WithVersion(version string) HeadValueBuilder {
	app.version = version
	return app
}

// WithImport adds a import to the builder
func (app *headValueBuilder) WithImport(imp []ImportSingle) HeadValueBuilder {
	app.imports = imp
	return app
}

// WithLoad adds a load to the builder
func (app *headValueBuilder) WithLoad(load []LoadSingle) HeadValueBuilder {
	app.loads = load
	return app
}

// Now builds a new HeadValue instance
func (app *headValueBuilder) Now() (HeadValue, error) {
	if app.name != "" {
		return createHeadValueWithName(app.name), nil
	}

	if app.version != "" {
		return createHeadValueWithVersion(app.version), nil
	}

	if app.imports == nil {
		app.imports = []ImportSingle{}
	}

	if len(app.imports) > 0 {
		return createHeadValueWithImport(app.imports), nil
	}

	if app.loads == nil {
		app.loads = []LoadSingle{}
	}

	if len(app.loads) > 0 {
		return createHeadValueWithLoad(app.loads), nil
	}

	return nil, errors.New("the HeadValue is invalid")
}
