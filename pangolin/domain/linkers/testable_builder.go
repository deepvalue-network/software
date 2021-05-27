package linkers

import "errors"

type testableBuilder struct {
	executable Executable
	language   LanguageReference
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

// WithExecutable adds an executable to the builder
func (app *testableBuilder) WithExecutable(executable Executable) TestableBuilder {
	app.executable = executable
	return app
}

// WithLanguage adds a language to the builder
func (app *testableBuilder) WithLanguage(language LanguageReference) TestableBuilder {
	app.language = language
	return app
}

// Now builds a new Testable instance
func (app *testableBuilder) Now() (Testable, error) {
	if app.executable != nil {
		return createTestableWithExecutable(app.executable), nil
	}

	if app.language != nil {
		return createTestableWithLanguage(app.language), nil
	}

	return nil, errors.New("the Testable instance is invalid")
}
