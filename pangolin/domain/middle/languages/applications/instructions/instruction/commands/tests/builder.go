package tests

import "errors"

type builder struct {
	name         string
	variable     string
	instructions []Instruction
}

func createBuilder() Builder {
	out := builder{
		name:         "",
		variable:     "",
		instructions: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithName adds a name to the builder
func (app *builder) WithName(name string) Builder {
	app.name = name
	return app
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

// Now builds a new Test instance
func (app *builder) Now() (Test, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Test instance")
	}

	if app.variable == "" {
		return nil, errors.New("the variable is mandatory in order to build a Test instance")
	}

	if app.instructions != nil && len(app.instructions) <= 0 {
		app.instructions = nil
	}

	if app.instructions == nil {
		return nil, errors.New("there must be at least 1 Instruction in order to build a Test instance")
	}

	return createTest(app.name, app.variable, app.instructions), nil
}
