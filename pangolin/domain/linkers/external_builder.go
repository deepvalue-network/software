package linkers

import "errors"

type externalBuilder struct {
	name       string
	executable Executable
}

func createExternalBuilder() ExternalBuilder {
	out := externalBuilder{
		name:       "",
		executable: nil,
	}

	return &out
}

// Create initializes the builder
func (app *externalBuilder) Create() ExternalBuilder {
	return createExternalBuilder()
}

// WithName adds a name to the builder
func (app *externalBuilder) WithName(name string) ExternalBuilder {
	app.name = name
	return app
}

// WithExecutable adds an executable to the builder
func (app *externalBuilder) WithExecutable(executable Executable) ExternalBuilder {
	app.executable = executable
	return app
}

// Now builds a new External instance
func (app *externalBuilder) Now() (External, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build an External instance")
	}

	if app.executable == nil {
		return nil, errors.New("the executable is mandatory in order to build an External instance")
	}

	return createExternal(app.name, app.executable), nil
}
