package parsers

import "errors"

type scopesBuilder struct {
	list []Scope
}

func createScopesBuilder() ScopesBuilder {
	out := scopesBuilder{
		list: nil,
	}

	return &out
}

// Create initializes the builder
func (app *scopesBuilder) Create() ScopesBuilder {
	return createScopesBuilder()
}

// WithScopes add scopes to the builder
func (app *scopesBuilder) WithScopes(scopes []Scope) ScopesBuilder {
	app.list = scopes
	return app
}

// Now builds a new Scopes instance
func (app *scopesBuilder) Now() (Scopes, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Scope in order to build a Scopes instance")
	}

	return createScopes(app.list), nil
}
