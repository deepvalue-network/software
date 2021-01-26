package parsers

import "errors"

type variableDeclarationBuilder struct {
	typ      Type
	variable string
	dir      VariableDirection
}

func createVariableDeclarationBuilder() VariableDeclarationBuilder {
	out := variableDeclarationBuilder{
		typ:      nil,
		variable: "",
		dir:      nil,
	}

	return &out
}

// Create initializes the builder
func (app *variableDeclarationBuilder) Create() VariableDeclarationBuilder {
	return createVariableDeclarationBuilder()
}

// WithType adds type to the builder
func (app *variableDeclarationBuilder) WithType(typ Type) VariableDeclarationBuilder {
	app.typ = typ
	return app
}

// WithVariable adds a variable to the builder
func (app *variableDeclarationBuilder) WithVariable(variable string) VariableDeclarationBuilder {
	app.variable = variable
	return app
}

// WithDirection adds a VariableDirection instance to the builder
func (app *variableDeclarationBuilder) WithDirection(dir VariableDirection) VariableDeclarationBuilder {
	app.dir = dir
	return app
}

// Now builds a new VariableDeclaration instance
func (app *variableDeclarationBuilder) Now() (VariableDeclaration, error) {
	if app.typ == nil {
		return nil, errors.New("the Type is mandatory in order to build a VariableDeclaration instance")
	}

	if app.variable == "" {
		return nil, errors.New("the variable is mandatory in order to build a VariableDeclaration instance")
	}

	if app.dir != nil {
		return createVariableDeclarationWithDirection(app.typ, app.variable, app.dir), nil
	}

	return createVariableDeclaration(app.typ, app.variable), nil
}
