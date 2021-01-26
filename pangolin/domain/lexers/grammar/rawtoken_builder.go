package grammar

import (
	"errors"
	"fmt"
)

type rawTokenBuilder struct {
	value string
	code  string
	index int
	gr    string
}

func createRawTokenBuilder() RawTokenBuilder {
	out := rawTokenBuilder{
		value: "",
		code:  "",
		index: -1,
		gr:    "",
	}

	return &out
}

// Create initializes the builder
func (app *rawTokenBuilder) Create() RawTokenBuilder {
	return createRawTokenBuilder()
}

// WithValue adds value to the builder
func (app *rawTokenBuilder) WithValue(value string) RawTokenBuilder {
	app.value = value
	return app
}

// WithCode adds code to the builder
func (app *rawTokenBuilder) WithCode(code string) RawTokenBuilder {
	app.code = code
	return app
}

// WithIndex adds the index to the builder
func (app *rawTokenBuilder) WithIndex(index int) RawTokenBuilder {
	app.index = index
	return app
}

// WithGrammar adds a grammar name to the builder
func (app *rawTokenBuilder) WithGrammar(grammar string) RawTokenBuilder {
	app.gr = grammar
	return app
}

// Now builds a new RawToken instance
func (app *rawTokenBuilder) Now() (RawToken, error) {
	if app.value == "" {
		return nil, errors.New("the value is mandatory in order to build a RawToken instance")
	}

	if app.code == "" {
		return nil, errors.New("the code is mandatory in order to build a RawToken instance")
	}

	if app.index == -1 {
		return nil, errors.New("the index is mandatory in order to build a RawToken instance")
	}

	if app.index < 0 {
		str := fmt.Sprintf("the index (%d) must be greater or equal to zero", app.index)
		return nil, errors.New(str)
	}

	if app.gr == "" {
		return nil, errors.New("the grammar is mandatory in order to build a RawToken instance")
	}

	return createRawToken(app.value, app.code, app.index, app.gr), nil
}
