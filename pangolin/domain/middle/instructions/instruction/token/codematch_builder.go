package token

import "errors"

type codeMatchBuilder struct {
	ret         string
	sectionName string
	token       string
	patterns    []string
}

func createCodeMatchBuilder() CodeMatchBuilder {
	out := codeMatchBuilder{
		ret:         "",
		sectionName: "",
		token:       "",
		patterns:    nil,
	}

	return &out
}

// Create initializes the builder
func (app *codeMatchBuilder) Create() CodeMatchBuilder {
	return createCodeMatchBuilder()
}

// WithReturn adds a return variable to the builder
func (app *codeMatchBuilder) WithReturn(ret string) CodeMatchBuilder {
	app.ret = ret
	return app
}

// WithSectionName adds a sectionName to the builder
func (app *codeMatchBuilder) WithSectionName(sectionName string) CodeMatchBuilder {
	app.sectionName = sectionName
	return app
}

// WithToken adds a token to the builder
func (app *codeMatchBuilder) WithToken(token string) CodeMatchBuilder {
	app.token = token
	return app
}

// WithPatterns add patterns to the builder
func (app *codeMatchBuilder) WithPatterns(patterns []string) CodeMatchBuilder {
	app.patterns = patterns
	return app
}

// Now builds a new CodeMatch instance
func (app *codeMatchBuilder) Now() (CodeMatch, error) {
	if app.ret == "" {
		return nil, errors.New("the return variable is mandatory in order to build a CodeMatch instance")
	}

	if app.sectionName == "" {
		return nil, errors.New("the sectionName is mandatory in order to build a CodeMatch instance")
	}

	if app.token == "" {
		return nil, errors.New("the token is mandatory in order to build a CodeMatch instance")
	}

	if app.patterns == nil {
		app.patterns = []string{}
	}

	if len(app.patterns) <= 0 {
		return nil, errors.New("there must be at least 1 pattern in order to build a CodeMatch instance")
	}

	return createCodeMatch(app.ret, app.sectionName, app.token, app.patterns), nil
}
