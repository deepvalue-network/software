package linkers

import "errors"

type programBuilder struct {
	testable Testable
	language LanguageApplication
}

func createProgramBuilder() ProgramBuilder {
	out := programBuilder{
		testable: nil,
		language: nil,
	}

	return &out
}

// Create initializes the builder
func (app *programBuilder) Create() ProgramBuilder {
	return createProgramBuilder()
}

// WithTestable adds a testable instance to the builder
func (app *programBuilder) WithTestable(testable Testable) ProgramBuilder {
	app.testable = testable
	return app
}

// WithLanguage adds a language instance to the builder
func (app *programBuilder) WithLanguage(language LanguageApplication) ProgramBuilder {
	app.language = language
	return app
}

// Now builds a new Program instance
func (app *programBuilder) Now() (Program, error) {
	if app.testable != nil {
		return createProgramWithTestable(app.testable), nil
	}

	if app.language != nil {
		return createProgramWithLanguage(app.language), nil
	}

	return nil, errors.New("the Program is invalid")
}
