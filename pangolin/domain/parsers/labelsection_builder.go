package parsers

import "errors"

type labelSectionBuilder struct {
	declarations []LabelDeclaration
}

func createLabelSectionBuilder() LabelSectionBuilder {
	out := labelSectionBuilder{
		declarations: nil,
	}

	return &out
}

// Create initializes the builder
func (app *labelSectionBuilder) Create() LabelSectionBuilder {
	return createLabelSectionBuilder()
}

// WithDeclarations add declarations to the builder
func (app *labelSectionBuilder) WithDeclarations(decl []LabelDeclaration) LabelSectionBuilder {
	app.declarations = decl
	return app
}

// Now builds a new LabelSection instance
func (app *labelSectionBuilder) Now() (LabelSection, error) {
	if app.declarations == nil {
		return nil, errors.New("the []LabelDeclaration are mandatory in order to build a LabelSection instance")
	}

	return createLabelSection(app.declarations), nil
}
