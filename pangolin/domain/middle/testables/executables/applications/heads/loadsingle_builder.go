package heads

import "errors"

type loadSingleBuilder struct {
	internal string
	external string
}

func createLoadSingleBuilder() LoadSingleBuilder {
	out := loadSingleBuilder{
		internal: "",
		external: "",
	}

	return &out
}

// Create initializes the builder
func (app *loadSingleBuilder) Create() LoadSingleBuilder {
	return createLoadSingleBuilder()
}

// WithInternal adds an internal to the builder
func (app *loadSingleBuilder) WithInternal(internal string) LoadSingleBuilder {
	app.internal = internal
	return app
}

// WithExternal adds an external to the builder
func (app *loadSingleBuilder) WithExternal(external string) LoadSingleBuilder {
	app.external = external
	return app
}

// Now builds a new LoadSingle instance
func (app *loadSingleBuilder) Now() (LoadSingle, error) {
	if app.internal == "" {
		return nil, errors.New("the internal is mandatory in order to build a LoadSingle instance")
	}

	if app.external == "" {
		return nil, errors.New("the external is mandatory in order to build a LoadSingle instance")
	}

	return createLoadSingle(app.internal, app.external), nil
}
