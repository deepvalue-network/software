package tokens

import (
	"errors"

	"github.com/deepvalue-network/software/adrien/domain/rules"
)

type contentBuilder struct {
	token string
	rule  rules.Rule
}

func createContentBuilder() ContentBuilder {
	out := contentBuilder{
		token: "",
		rule:  nil,
	}

	return &out
}

// Create initializes the builder
func (app *contentBuilder) Create() ContentBuilder {
	return createContentBuilder()
}

// WithToken adds a token to the builder
func (app *contentBuilder) WithToken(token string) ContentBuilder {
	app.token = token
	return app
}

// WithRule adds a rule to the builder
func (app *contentBuilder) WithRule(rule rules.Rule) ContentBuilder {
	app.rule = rule
	return app
}

// Now builds a new Content instance
func (app *contentBuilder) Now() (Content, error) {
	if app.token != "" {
		return createContentWithToken(app.token), nil
	}

	if app.rule != nil {
		return createContentWithRule(app.rule), nil
	}

	return nil, errors.New("the content is invalid")
}
