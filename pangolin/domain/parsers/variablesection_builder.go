package parsers

import "errors"

type variableSectionBuilder struct {
	declarations []VariableDeclaration
}

func createVariableSectionBuilder() VariableSectionBuilder {
	out := variableSectionBuilder{
		declarations: nil,
	}

	return &out
}

// Create initializes the builder
func (app *variableSectionBuilder) Create() VariableSectionBuilder {
	return createVariableSectionBuilder()
}

// WithDeclarations add declarations to the builder
func (app *variableSectionBuilder) WithDeclarations(declarations []VariableDeclaration) VariableSectionBuilder {
	app.declarations = declarations
	return app
}

// Now builds a new VariableSection instance
func (app *variableSectionBuilder) Now() (VariableSection, error) {
	if app.declarations == nil {
		return nil, errors.New("the []VariableDeclaration is mandatory in order to build a VariableSection instance")
	}

	return createVariableSection(app.declarations), nil
}
