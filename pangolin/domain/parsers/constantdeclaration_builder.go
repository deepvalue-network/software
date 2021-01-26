package parsers

import "errors"

type constantDeclarationBuilder struct {
	constant string
	typ      Type
	value    Value
}

func createConstantDeclarationBuilder() ConstantDeclarationBuilder {
	out := constantDeclarationBuilder{
		constant: "",
		typ:      nil,
		value:    nil,
	}

	return &out
}

// Create initializes the builder
func (app *constantDeclarationBuilder) Create() ConstantDeclarationBuilder {
	return createConstantDeclarationBuilder()
}

// WithConstant adds a constant to the builder
func (app *constantDeclarationBuilder) WithConstant(constant string) ConstantDeclarationBuilder {
	app.constant = constant
	return app
}

// WithType adds type to the builder
func (app *constantDeclarationBuilder) WithType(typ Type) ConstantDeclarationBuilder {
	app.typ = typ
	return app
}

// WithValue adds a value to the builder
func (app *constantDeclarationBuilder) WithValue(value Value) ConstantDeclarationBuilder {
	app.value = value
	return app
}

// Now builds a new ConstantDeclaration instance
func (app *constantDeclarationBuilder) Now() (ConstantDeclaration, error) {
	if app.constant == "" {
		return nil, errors.New("the constant is mandatory in order to build a ConstantDeclaration instance")
	}

	if app.typ == nil {
		return nil, errors.New("the Type is mandatory in order to build a ConstantDeclaration instance")
	}

	if app.value == nil {
		return nil, errors.New("the Value is mandatory in order to build a ConstantDeclaration instance")
	}

	return createConstantDeclaration(app.constant, app.typ, app.value), nil
}
