package parsers

import "errors"

type constantSectionBuilder struct {
	declarations []ConstantDeclaration
}

func createConstantSectionBuilder() ConstantSectionBuilder {
	out := constantSectionBuilder{
		declarations: nil,
	}

	return &out
}

// Create initializes the builder
func (app *constantSectionBuilder) Create() ConstantSectionBuilder {
	return createConstantSectionBuilder()
}

// WithDeclarations add declarations to the builder
func (app *constantSectionBuilder) WithDeclarations(declarations []ConstantDeclaration) ConstantSectionBuilder {
	app.declarations = declarations
	return app
}

// Now builds a new ConstantSection instance
func (app *constantSectionBuilder) Now() (ConstantSection, error) {
	if app.declarations == nil {
		return nil, errors.New("the []ConstantDeclaration is mandatory in order to build a ConstantSection instance")
	}

	return createConstantSection(app.declarations), nil
}
