package parsers

import "errors"

type specificTokenCodeBuilder struct {
	content         VariableName
	tokenVariable   string
	patternVariable string
	amount          VariableName
}

func createSpecificTokenCodeBuilder() SpecificTokenCodeBuilder {
	out := specificTokenCodeBuilder{
		content:         nil,
		tokenVariable:   "",
		patternVariable: "",
		amount:          nil,
	}

	return &out
}

// Create initializes the builder
func (app *specificTokenCodeBuilder) Create() SpecificTokenCodeBuilder {
	return createSpecificTokenCodeBuilder()
}

// WithContent adds a content to the builder
func (app *specificTokenCodeBuilder) WithContent(content VariableName) SpecificTokenCodeBuilder {
	app.content = content
	return app
}

// WithAmount adds an amount to the builder
func (app *specificTokenCodeBuilder) WithAmount(amount VariableName) SpecificTokenCodeBuilder {
	app.amount = amount
	return app
}

// WithTokenVariable adds a tokenVariable to the builder
func (app *specificTokenCodeBuilder) WithTokenVariable(tokenVariable string) SpecificTokenCodeBuilder {
	app.tokenVariable = tokenVariable
	return app
}

// WithPatternVariable adds a pattern variable to the builder
func (app *specificTokenCodeBuilder) WithPatternVariable(pattern string) SpecificTokenCodeBuilder {
	app.patternVariable = pattern
	return app
}

// Now builds a new SpecificTokenCode
func (app *specificTokenCodeBuilder) Now() (SpecificTokenCode, error) {
	if app.content == nil {
		return nil, errors.New("the content variable is mandatory in order to build a SpecificTokenCode instance")
	}

	if app.tokenVariable == "" {
		return nil, errors.New("the tokenVariable is mandatory in order to build a SpecificTokenCode instance")
	}

	if app.patternVariable == "" {
		return nil, errors.New("the pattern variable is mandatory in order to build a SpecificTokenCode instance")
	}

	if app.amount != nil {
		return createSpecificTokenCodeWithAmount(app.content, app.tokenVariable, app.patternVariable, app.amount), nil
	}

	return createSpecificTokenCode(app.content, app.tokenVariable, app.patternVariable), nil
}
