package parsers

import "errors"

type scopeBuilder struct {
	isInternal bool
	isExternal bool
}

func createScopeBuilder() ScopeBuilder {
	out := scopeBuilder{
		isInternal: false,
		isExternal: false,
	}

	return &out
}

// Create initializes the builder
func (app *scopeBuilder) Create() ScopeBuilder {
	return createScopeBuilder()
}

// IsInternal flags the builder as internal
func (app *scopeBuilder) IsInternal() ScopeBuilder {
	app.isInternal = true
	return app
}

// IsExternal flags the builder as external
func (app *scopeBuilder) IsExternal() ScopeBuilder {
	app.isExternal = true
	return app
}

// Now builds a new Scope instance
func (app *scopeBuilder) Now() (Scope, error) {
	if app.isInternal {
		return createScopeWithInternal(), nil
	}

	if app.isExternal {
		return createScopeWithExternal(), nil
	}

	return nil, errors.New("the Scope is invalid")
}
