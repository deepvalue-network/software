package linkers

import "errors"

type executableBuilder struct {
	app    Application
	script Script
}

func createExecutableBuilder() ExecutableBuilder {
	out := executableBuilder{
		app:    nil,
		script: nil,
	}

	return &out
}

// Create initializes the builder
func (app *executableBuilder) Create() ExecutableBuilder {
	return createExecutableBuilder()
}

// WithApplication adds an application to the builder
func (app *executableBuilder) WithApplication(appli Application) ExecutableBuilder {
	app.app = appli
	return app
}

// WithScript adds a script to the builder
func (app *executableBuilder) WithScript(script Script) ExecutableBuilder {
	app.script = script
	return app
}

// Now builds a new Executable instance
func (app *executableBuilder) Now() (Executable, error) {
	if app.app != nil {
		return createExecutableWithApplication(app.app), nil
	}

	if app.script != nil {
		return createExecutableWithScript(app.script), nil
	}

	return nil, errors.New("the Executable is invalid")
}
