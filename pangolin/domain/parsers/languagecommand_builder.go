package parsers

import "errors"

type languageCommandBuilder struct {
	variable string
	values   []LanguageValue
}

func createLanguageCommandBuilder() LanguageCommandBuilder {
	out := languageCommandBuilder{
		variable: "",
		values:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *languageCommandBuilder) Create() LanguageCommandBuilder {
	return createLanguageCommandBuilder()
}

// WithVariable adds a variable to the builder
func (app *languageCommandBuilder) WithVariable(variable string) LanguageCommandBuilder {
	app.variable = variable
	return app
}

// WithValues add values to the builder
func (app *languageCommandBuilder) WithValues(values []LanguageValue) LanguageCommandBuilder {
	app.values = values
	return app
}

// Now builds a new LanguageValue instance
func (app *languageCommandBuilder) Now() (LanguageCommand, error) {
	if app.variable == "" {
		return nil, errors.New("the variable is mandatory in order to build a LanguageValue instance")
	}

	if app.values != nil && len(app.values) <= 0 {
		app.values = nil
	}

	if app.values == nil {
		return nil, errors.New("there must be at least 1 LanguageValue instance in order to build a LanguageCommand instance")
	}

	return createLanguageCommand(app.variable, app.values), nil
}
