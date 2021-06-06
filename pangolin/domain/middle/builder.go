package middle

import (
	"errors"

	"github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications"
	"github.com/deepvalue-network/software/pangolin/domain/middle/testables"
)

type builder struct {
	testable testables.Testable
	language applications.Application
}

func createBuilder() Builder {
	out := builder{
		testable: nil,
		language: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithTestable adds a testable to the builder
func (app *builder) WithTestable(testable testables.Testable) Builder {
	app.testable = testable
	return app
}

// WithLanguage adds a language to the builder
func (app *builder) WithLanguage(language applications.Application) Builder {
	app.language = language
	return app
}

// Now builds a new Program instance
func (app *builder) Now() (Program, error) {
	if app.testable != nil {
		return createProgramWithTestable(app.testable), nil
	}

	if app.language != nil {
		return createProgramWithLanguage(app.language), nil
	}

	return nil, errors.New("the Program is invalid")
}
