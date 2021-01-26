package linkers

import "errors"

type externalBuilder struct {
	name        string
	application Application
	script      Script
}

func createExternalBuilder() ExternalBuilder {
	out := externalBuilder{
		name:        "",
		application: nil,
		script:      nil,
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

// WithApplication adds an application to the builder
func (app *externalBuilder) WithApplication(application Application) ExternalBuilder {
	app.application = application
	return app
}

// WithScript adds a script to the builder
func (app *externalBuilder) WithScript(script Script) ExternalBuilder {
	app.script = script
	return app
}

// Now builds a new External instance
func (app *externalBuilder) Now() (External, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build an External instance")
	}

	if app.application != nil {
		return createExternalWithApplication(app.name, app.application), nil
	}

	if app.script != nil {
		return createExternalWithScript(app.name, app.script), nil
	}

	return nil, errors.New("the External instance is invalid")
}
