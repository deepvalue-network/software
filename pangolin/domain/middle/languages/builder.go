package languages

import (
	"errors"

	"github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications"
	"github.com/deepvalue-network/software/pangolin/domain/middle/languages/definitions"
)

type builder struct {
	def definitions.Definition
	app applications.Application
}

func createBuilder() Builder {
	out := builder{
		def: nil,
		app: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithDefinition adds a definition to the builder
func (app *builder) WithDefinition(def definitions.Definition) Builder {
	app.def = def
	return app
}

// WithApplication adds an application to the builder
func (app *builder) WithApplication(ins applications.Application) Builder {
	app.app = ins
	return app
}

// Now builds a new Language instance
func (app *builder) Now() (Language, error) {
	if app.def != nil {
		return createLanguageWithDefinition(app.def), nil
	}

	if app.app != nil {
		return createLanguageWithApplication(app.app), nil
	}

	return nil, errors.New("the Language is invalid")
}
