package parsers

import "errors"

type codeMatchBuilder struct {
	content          VariableName
	section          VariableName
	patternVariables []string
}

func createCodeMatchBuilder() CodeMatchBuilder {
	out := codeMatchBuilder{
		content:          nil,
		section:          nil,
		patternVariables: nil,
	}

	return &out
}

// Create initializes the builder
func (app *codeMatchBuilder) Create() CodeMatchBuilder {
	return createCodeMatchBuilder()
}

// WithContent adds a content to the builder
func (app *codeMatchBuilder) WithContent(content VariableName) CodeMatchBuilder {
	app.content = content
	return app
}

// WithSection adds a section to the builder
func (app *codeMatchBuilder) WithSection(section VariableName) CodeMatchBuilder {
	app.section = section
	return app
}

// WithPatternVariables adds a patternVariables to the builder
func (app *codeMatchBuilder) WithPatternVariables(patterns []string) CodeMatchBuilder {
	app.patternVariables = patterns
	return app
}

// Now builds a new CodeMatch instance
func (app *codeMatchBuilder) Now() (CodeMatch, error) {
	if app.content == nil {
		return nil, errors.New("the content VariableName is mandatory in order to build a CodeMatch instance")
	}

	if app.section == nil {
		return nil, errors.New("the section VariableName is mandatory in order to build a CodeMatch instance")
	}

	if app.patternVariables == nil {
		app.patternVariables = []string{}
	}

	if len(app.patternVariables) <= 0 {
		return nil, errors.New("the patternVariables are mandatory in order to build a CodeMatch instance")
	}

	return createCodeMatch(app.content, app.section, app.patternVariables), nil
}
