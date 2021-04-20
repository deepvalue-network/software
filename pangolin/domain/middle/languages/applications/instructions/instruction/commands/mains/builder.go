package mains

import "errors"

type builder struct {
	variable     string
	instructions []Instruction
}

func createBuilder() Builder {
	out := builder{
		variable:     "",
		instructions: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithVariable adds a variable to the builder
func (app *builder) WithVariable(variable string) Builder {
	app.variable = variable
	return app
}

// WithInstructions add instructions to the builder
func (app *builder) WithInstructions(ins []Instruction) Builder {
	app.instructions = ins
	return app
}

// Now builds a new Main instance
func (app *builder) Now() (Main, error) {
	if app.variable == "" {
		return nil, errors.New("the variable is mandatory in order to build a Main instance")
	}

	if app.instructions != nil && len(app.instructions) <= 0 {
		app.instructions = nil
	}

	if app.instructions == nil {
		return nil, errors.New("there must be at least 1 Instruction in order to build a Main instance")
	}

	return createMain(app.variable, app.instructions), nil
}
