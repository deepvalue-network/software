package parsers

import "errors"

type mainSectionBuilder struct {
	instructions []Instruction
}

func createMainSectionBuilder() MainSectionBuilder {
	out := mainSectionBuilder{
		instructions: nil,
	}

	return &out
}

// Create initializes the builder
func (app *mainSectionBuilder) Create() MainSectionBuilder {
	return createMainSectionBuilder()
}

// WithInstructions add instructions to the builder
func (app *mainSectionBuilder) WithInstructions(ins []Instruction) MainSectionBuilder {
	app.instructions = ins
	return app
}

// Now builds a new MainSection instance
func (app *mainSectionBuilder) Now() (MainSection, error) {
	if app.instructions == nil {
		return nil, errors.New("the []Instruction are mandatory in order to build a MainSection instance")
	}

	return createMainSection(app.instructions), nil
}
