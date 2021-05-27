package parsers

import "errors"

type executableBuilder struct {
	application Application
	script      Script
}

func createExecutableBuilder() ExecutableBuilder {
	out := executableBuilder{
		application: nil,
		script:      nil,
	}

	return &out
}

// Create initializes the builder
func (app *executableBuilder) Create() ExecutableBuilder {
	return createExecutableBuilder()
}

// WithApplication adds an application to the builder
func (app *executableBuilder) WithApplication(application Application) ExecutableBuilder {
	app.application = application
	return app
}

// WithScript adds a script to the builder
func (app *executableBuilder) WithScript(script Script) ExecutableBuilder {
	app.script = script
	return app
}

// Now builds a new Executable instance
func (app *executableBuilder) Now() (Executable, error) {
	if app.application != nil {
		return createExecutableWithApplication(app.application), nil
	}

	if app.script != nil {
		return createExecutableWithScript(app.script), nil
	}

	return nil, errors.New("the Executable is invalid")
}
