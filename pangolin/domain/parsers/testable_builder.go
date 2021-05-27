package parsers

import "errors"

type testableBuilder struct {
	executable Executable
	language   LanguageDefinition
}

func createTestableBuilder() TestableBuilder {
	out := testableBuilder{
		executable: nil,
		language:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *testableBuilder) Create() TestableBuilder {
	return createTestableBuilder()
}

// Create initializes the builder
func (app *testableBuilder) WithExecutable(executable Executable) TestableBuilder {
	app.executable = executable
	return app
}

// WithLanguage adds a language to the builder
func (app *testableBuilder) WithLanguage(language LanguageDefinition) TestableBuilder {
	app.language = language
	return app
}

// Now builds a new testable instance
func (app *testableBuilder) Now() (Testable, error) {
	if app.executable != nil {
		return createTestableWithExecutable(app.executable), nil
	}

	if app.language != nil {
		return createTestableWithLanguage(app.language), nil
	}

	return nil, errors.New("the Testable is invalid")
}
