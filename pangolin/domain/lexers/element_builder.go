package lexers

import (
	"errors"

	"github.com/steve-care-software/products/pangolin/domain/lexers/grammar"
)

type elementBuilder struct {
	rule grammar.Rule
	code string
}

func createElementBuilder() ElementBuilder {
	out := elementBuilder{
		rule: nil,
		code: "",
	}

	return &out
}

// Create initializes the builder
func (app *elementBuilder) Create() ElementBuilder {
	return createElementBuilder()
}

// WithRule adds a rule to the builder
func (app *elementBuilder) WithRule(rule grammar.Rule) ElementBuilder {
	app.rule = rule
	return app
}

// WithCode adds a code to the builder
func (app *elementBuilder) WithCode(code string) ElementBuilder {
	app.code = code
	return app
}

// Now builds a new Element instance
func (app *elementBuilder) Now() (Element, error) {
	if app.rule == nil {
		return nil, errors.New("the Rule is mandatory in order to build a Element instance")
	}

	if app.code == "" {
		return nil, errors.New("the code is mandatory in order to build a Element instance")
	}

	return createElement(app.rule, app.code), nil
}
