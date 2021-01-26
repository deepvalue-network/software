package token

import "errors"

type codeBuilder struct {
	ret     string
	token   string
	pattern string
	amount  string
}

func createCodeBuilder() CodeBuilder {
	out := codeBuilder{
		ret:     "",
		token:   "",
		pattern: "",
		amount:  "",
	}

	return &out
}

// Create initializes the builder
func (app *codeBuilder) Create() CodeBuilder {
	return createCodeBuilder()
}

// WithReturn adds a return to the builder
func (app *codeBuilder) WithReturn(ret string) CodeBuilder {
	app.ret = ret
	return app
}

// WithToken adds a token to the builder
func (app *codeBuilder) WithToken(token string) CodeBuilder {
	app.token = token
	return app
}

// WithPattern adds a pattern to the builder
func (app *codeBuilder) WithPattern(pattern string) CodeBuilder {
	app.pattern = pattern
	return app
}

// WithAmount adds an amount to the builder
func (app *codeBuilder) WithAmount(amount string) CodeBuilder {
	app.amount = amount
	return app
}

// Now builds a new Code instance
func (app *codeBuilder) Now() (Code, error) {
	if app.ret == "" {
		return nil, errors.New("the return variable is mandatory in order to build a Code instance")
	}

	if app.token == "" {
		return nil, errors.New("the token variable is mandatory in order to build a Code instance")
	}

	if app.pattern != "" && app.amount != "" {
		return createCodeWithPatternAndAmount(app.ret, app.token, app.pattern, app.amount), nil
	}

	if app.pattern != "" {
		return createCodeWithPattern(app.ret, app.token, app.pattern), nil
	}

	if app.amount != "" {
		return createCodeWithAmount(app.ret, app.token, app.amount), nil
	}

	return createCode(app.ret, app.token), nil
}
