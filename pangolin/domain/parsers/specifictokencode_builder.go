package parsers

import "errors"

type specificTokenCodeBuilder struct {
	variableName    VariableName
	patternVariable string
	amount          VariableName
}

func createSpecificTokenCodeBuilder() SpecificTokenCodeBuilder {
	out := specificTokenCodeBuilder{
		variableName:    nil,
		patternVariable: "",
		amount:          nil,
	}

	return &out
}

// Create initializes the builder
func (app *specificTokenCodeBuilder) Create() SpecificTokenCodeBuilder {
	return createSpecificTokenCodeBuilder()
}

// WithVariableName adds a content to the builder
func (app *specificTokenCodeBuilder) WithVariableName(variableName VariableName) SpecificTokenCodeBuilder {
	app.variableName = variableName
	return app
}

// WithAmount adds an amount to the builder
func (app *specificTokenCodeBuilder) WithAmount(amount VariableName) SpecificTokenCodeBuilder {
	app.amount = amount
	return app
}

// WithPatternVariable adds a pattern variable to the builder
func (app *specificTokenCodeBuilder) WithPatternVariable(pattern string) SpecificTokenCodeBuilder {
	app.patternVariable = pattern
	return app
}

// Now builds a new SpecificTokenCode
func (app *specificTokenCodeBuilder) Now() (SpecificTokenCode, error) {
	if app.variableName == nil {
		return nil, errors.New("the variableName is mandatory in order to build a SpecificTokenCode instance")
	}

	if app.patternVariable == "" {
		return nil, errors.New("the pattern variable is mandatory in order to build a SpecificTokenCode instance")
	}

	if app.amount != nil {
		return createSpecificTokenCodeWithAmount(app.variableName, app.patternVariable, app.amount), nil
	}

	return createSpecificTokenCode(app.variableName, app.patternVariable), nil
}
