package testables

import (
	"errors"

	"github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables"
	"github.com/deepvalue-network/software/pangolin/domain/middle/testables/languages/definitions"
)

type builder struct {
	executable executables.Executable
	language   definitions.Definition
}

func createBuilder() Builder {
	out := builder{
		executable: nil,
		language:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithExecutable adds an executable to the builder
func (app *builder) WithExecutable(executable executables.Executable) Builder {
	app.executable = executable
	return app
}

// WithLanguage adds a language to the builder
func (app *builder) WithLanguage(language definitions.Definition) Builder {
	app.language = language
	return app
}

// Now builds a new Testable instance
func (app *builder) Now() (Testable, error) {
	if app.executable != nil {
		return createTestableWithExecutable(app.executable), nil
	}

	if app.language != nil {
		return createTestableWithLanguage(app.language), nil
	}

	return nil, errors.New("the Testable instance is invalid")
}
