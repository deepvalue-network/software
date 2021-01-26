package variable

import (
	"errors"

	var_value "github.com/steve-care-software/products/pangolin/domain/middle/variables/variable/value"
)

type builder struct {
	isGlobal    bool
	isImmutable bool
	isMandatory bool
	isIncoming  bool
	isOutgoing  bool
	name        string
	value       var_value.Value
}

func createBuilder() Builder {
	out := builder{
		isGlobal:    false,
		isImmutable: false,
		isMandatory: false,
		isIncoming:  false,
		isOutgoing:  false,
		name:        "",
		value:       nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// IsGlobal sets the builder as global
func (app *builder) IsGlobal() Builder {
	app.isGlobal = true
	return app
}

// IsImmutable sets the builder as immutable
func (app *builder) IsImmutable() Builder {
	app.isImmutable = true
	return app
}

// WithName adds a name to the builder
func (app *builder) WithName(name string) Builder {
	app.name = name
	return app
}

// WithValue adds a value to the builder
func (app *builder) WithValue(val var_value.Value) Builder {
	app.value = val
	return app
}

// IsIncoming adds an incoming
func (app *builder) IsIncoming() Builder {
	app.isIncoming = true
	return app
}

// IsOutgoing adds an outgoing to the builder
func (app *builder) IsOutgoing() Builder {
	app.isOutgoing = true
	return app
}

// IsMandatory flags the builder as  to the builder
func (app *builder) IsMandatory() Builder {
	app.isMandatory = true
	return app
}

// Now builds a new Variable instance
func (app *builder) Now() (Variable, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Variable instance")
	}

	if app.value == nil {
		return nil, errors.New("the value is mandatory in order to build a Variable instance")
	}

	return createVariable(
		app.isGlobal,
		app.isImmutable,
		app.isMandatory,
		app.isIncoming,
		app.isOutgoing,
		app.name,
		app.value,
	), nil
}
