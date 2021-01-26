package parsers

import "errors"

type declarationBuilder struct {
	name string
	typ  Type
}

func createDeclarationBuilder() DeclarationBuilder {
	out := declarationBuilder{
		name: "",
		typ:  nil,
	}

	return &out
}

// Create initializes the builder
func (app *declarationBuilder) Create() DeclarationBuilder {
	return createDeclarationBuilder()
}

// WithVariable adds a variable name to the builder
func (app *declarationBuilder) WithVariable(name string) DeclarationBuilder {
	app.name = name
	return app
}

// WithType adds a type to the builder
func (app *declarationBuilder) WithType(typ Type) DeclarationBuilder {
	app.typ = typ
	return app
}

// Now builds a new Declaration instance
func (app *declarationBuilder) Now() (Declaration, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Declaration instance")
	}

	if app.typ == nil {
		return nil, errors.New("the type is mandatory in order to build a Declaration instance")
	}

	return createDeclaration(app.name, app.typ), nil
}
