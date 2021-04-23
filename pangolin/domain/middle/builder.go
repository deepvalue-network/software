package middle

import (
	"errors"

	"github.com/deepvalue-network/software/pangolin/domain/middle/applications"
	"github.com/deepvalue-network/software/pangolin/domain/middle/languages"
	"github.com/deepvalue-network/software/pangolin/domain/middle/scripts"
)

type builder struct {
	app    applications.Application
	lang   languages.Language
	script scripts.Script
}

func createBuilder() Builder {
	out := builder{
		app:    nil,
		lang:   nil,
		script: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithApplication adds an application to the builder
func (app *builder) WithApplication(appli applications.Application) Builder {
	app.app = appli
	return app
}

// WithLanguage adds a language to the builder
func (app *builder) WithLanguage(lang languages.Language) Builder {
	app.lang = lang
	return app
}

// WithScript adds a script to the builder
func (app *builder) WithScript(script scripts.Script) Builder {
	app.script = script
	return app
}

// Now builds a new Program instance
func (app *builder) Now() (Program, error) {
	if app.app != nil {
		return createProgramWithApplication(app.app), nil
	}

	if app.lang != nil {
		return createProgramWithLanguage(app.lang), nil
	}

	if app.script != nil {
		return createProgramWithScript(app.script), nil
	}

	return nil, errors.New("the Program is invalid")
}
