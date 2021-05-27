package executables

import (
	"errors"

	"github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables/applications"
	"github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables/scripts"
)

type builder struct {
	application applications.Application
	script      scripts.Script
}

func createBuilder() Builder {
	out := builder{
		application: nil,
		script:      nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithApplication adds an application to the builder
func (app *builder) WithApplication(application applications.Application) Builder {
	app.application = application
	return app
}

// WithScript adds a script to the builder
func (app *builder) WithScript(script scripts.Script) Builder {
	app.script = script
	return app
}

// Now builds a new Executable instance
func (app *builder) Now() (Executable, error) {
	if app.application != nil {
		return createExecutableWithApplication(app.application), nil
	}

	if app.script != nil {
		return createExecutableWithScript(app.script), nil
	}

	return nil, errors.New("the Executable is invalid")
}
