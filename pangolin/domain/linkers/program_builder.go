package linkers

import (
	"errors"
)

type programBuilder struct {
	app    Application
	lang   Language
	script Script
}

func createProgramBuilder() ProgramBuilder {
	out := programBuilder{
		app:    nil,
		lang:   nil,
		script: nil,
	}

	return &out
}

// Create initializes the builder
func (app *programBuilder) Create() ProgramBuilder {
	return createProgramBuilder()
}

// WithApplication adds an application to the builder
func (app *programBuilder) WithApplication(appli Application) ProgramBuilder {
	app.app = appli
	return app
}

// WithLanguage adds a language to the builder
func (app *programBuilder) WithLanguage(lang Language) ProgramBuilder {
	app.lang = lang
	return app
}

// WithScript adds a script to the builder
func (app *programBuilder) WithScript(script Script) ProgramBuilder {
	app.script = script
	return app
}

// Now builds a new Program instance
func (app *programBuilder) Now() (Program, error) {
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
